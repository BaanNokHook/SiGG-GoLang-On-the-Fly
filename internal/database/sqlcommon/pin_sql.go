// SiGG-GoLang-On-the-Fly //
package sqlcommon

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/hyperledger/firefly-common/pkg/dbsql"
	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly-common/pkg/log"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var (
	pinColumns = []string{
		"namespace",
		"masked",
		"hash",
		"batch_id",
		"batch_hash",
		"idx",
		"signer",
		"dispatched",
		"created",
	}
	pinFilterFieldMap = map[string]string{
		"batch":     "batch_id",
		"batchhash": "batch_hash",
		"index":     "idx",
	}
)

const pinsTable = "pins"

func (s *SQLCommon) UpsertPin(ctx context.Context, pin *core.Pin) (err error) {
	ctx, tx, autoCommit, err := s.BeginOrUseTx(ctx)
	if err != nil {
		return err
	}
	defer s.RollbackTx(ctx, tx, autoCommit)

	// Do a select within the transaction to detemine if the UUID already exists
	pinRows, tx, err := s.QueryTx(ctx, pinsTable, tx,
		sq.Select(s.SequenceColumn(), "masked", "dispatched").
			From(pinsTable).
			Where(sq.Eq{
				"hash":      pin.Hash,
				"batch_id":  pin.Batch,
				"idx":       pin.Index,
				"namespace": pin.Namespace,
			}))
	if err != nil {
		return err
	}
	existing := pinRows.Next()

	if existing {
		err := pinRows.Scan(&pin.Sequence, &pin.Masked, &pin.Dispatched)
		pinRows.Close()
		if err != nil {
			return i18n.WrapError(ctx, err, coremsgs.MsgDBReadErr, pinsTable)
		}
		// Pin's can only go from undispatched, to dispatched - so no update here.
		log.L(ctx).Debugf("Existing pin returned at sequence %d", pin.Sequence)
	} else {
		pinRows.Close()
		if err = s.attemptPinInsert(ctx, tx, pin); err != nil {
			return err
		}

	}

	return s.CommitTx(ctx, tx, autoCommit)
}

func (s *SQLCommon) attemptPinInsert(ctx context.Context, tx *dbsql.TXWrapper, pin *core.Pin) (err error) {
	pin.Sequence, err = s.InsertTx(ctx, pinsTable, tx,
		s.setPinInsertValues(sq.Insert(pinsTable).Columns(pinColumns...), pin),
		func() {
			log.L(ctx).Debugf("Triggering creation event for pin %d", pin.Sequence)
			s.callbacks.OrderedCollectionNSEvent(database.CollectionPins, core.ChangeEventTypeCreated, pin.Namespace, pin.Sequence)
		},
	)
	return err
}

func (s *SQLCommon) setPinInsertValues(query sq.InsertBuilder, pin *core.Pin) sq.InsertBuilder {
	return query.Values(
		pin.Namespace,
		pin.Masked,
		pin.Hash,
		pin.Batch,
		pin.BatchHash,
		pin.Index,
		pin.Signer,
		pin.Dispatched,
		pin.Created,
	)
}

func (s *SQLCommon) InsertPins(ctx context.Context, pins []*core.Pin) error {
	ctx, tx, autoCommit, err := s.BeginOrUseTx(ctx)
	if err != nil {
		return err
	}
	defer s.RollbackTx(ctx, tx, autoCommit)

	if s.features.MultiRowInsert {
		query := sq.Insert(pinsTable).Columns(pinColumns...)
		for _, pin := range pins {
			query = s.setPinInsertValues(query, pin)
		}
		sequences := make([]int64, len(pins))
		err := s.InsertTxRows(ctx, pinsTable, tx, query, func() {
			for i, pin := range pins {
				pin.Sequence = sequences[i]
				s.callbacks.OrderedCollectionNSEvent(database.CollectionPins, core.ChangeEventTypeCreated, pin.Namespace, pin.Sequence)
			}
		}, sequences, true /* we want the caller to be able to retry with individual upserts */)
		if err != nil {
			return err
		}
	} else {
		// Fall back to individual inserts grouped in a TX
		for _, pin := range pins {
			if err := s.attemptPinInsert(ctx, tx, pin); err != nil {
				return err
			}
		}
	}

	return s.CommitTx(ctx, tx, autoCommit)
}

func (s *SQLCommon) pinResult(ctx context.Context, row *sql.Rows) (*core.Pin, error) {
	pin := core.Pin{}
	err := row.Scan(
		&pin.Namespace,
		&pin.Masked,
		&pin.Hash,
		&pin.Batch,
		&pin.BatchHash,
		&pin.Index,
		&pin.Signer,
		&pin.Dispatched,
		&pin.Created,
		&pin.Sequence,
	)
	if err != nil {
		return nil, i18n.WrapError(ctx, err, coremsgs.MsgDBReadErr, pinsTable)
	}
	return &pin, nil
}

func (s *SQLCommon) GetPins(ctx context.Context, namespace string, filter ffapi.Filter) (message []*core.Pin, fr *ffapi.FilterResult, err error) {

	cols := append([]string{}, pinColumns...)
	cols = append(cols, s.SequenceColumn())
	query, fop, fi, err := s.FilterSelect(
		ctx, "", sq.Select(cols...).From(pinsTable),
		filter, pinFilterFieldMap, []interface{}{"sequence"}, sq.Eq{"namespace": namespace})
	if err != nil {
		return nil, nil, err
	}

	rows, tx, err := s.Query(ctx, pinsTable, query)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	pin := []*core.Pin{}
	for rows.Next() {
		d, err := s.pinResult(ctx, rows)
		if err != nil {
			return nil, nil, err
		}
		pin = append(pin, d)
	}

	return pin, s.QueryRes(ctx, pinsTable, tx, fop, fi), err

}

func (s *SQLCommon) UpdatePins(ctx context.Context, namespace string, filter ffapi.Filter, update ffapi.Update) (err error) {

	ctx, tx, autoCommit, err := s.BeginOrUseTx(ctx)
	if err != nil {
		return err
	}
	defer s.RollbackTx(ctx, tx, autoCommit)

	query, err := s.BuildUpdate(sq.Update(pinsTable).Where(sq.Eq{"namespace": namespace}), update, pinFilterFieldMap)
	if err != nil {
		return err
	}

	query, err = s.FilterUpdate(ctx, query, filter, pinFilterFieldMap)
	if err != nil {
		return err
	}

	_, err = s.UpdateTx(ctx, pinsTable, tx, query, nil /* no change events filter based update */)
	if err != nil {
		return err
	}

	return s.CommitTx(ctx, tx, autoCommit)
}
