// SiGG-GoLang-On-the-Fly //

package fabric

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
)

func getDNFromCertString(certString string) (string, error) {
	cert, err := getCertificateFromBytes(certString)
	if err != nil {
		return "", err
	}
	return getDN(&cert.Subject), nil
}

// borrowed from fabric-chaincode-go to guarantee the same
// resolution of "DN" string from x509 certs
func getDN(name *pkix.Name) string {
	r := name.ToRDNSequence()
	return r.String()
}

func getCertificateFromBytes(certString string) (*x509.Certificate, error) {
	idbytes, err := base64.StdEncoding.DecodeString(certString)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(idbytes)
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	return cert, nil
}
