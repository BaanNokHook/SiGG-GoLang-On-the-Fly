// SiGG-GoLang-On-the-Fly //
package ethereum

import "github.com/hyperledger/firefly-signer/pkg/abi"

var batchPinMethodABIV1 = &abi.Entry{
	Name: "pinBatch",
	Type: "function",
	Inputs: abi.ParameterArray{
		{
			InternalType: "string",
			Name:         "namespace",
			Type:         "string",
		},
		{
			InternalType: "bytes32",
			Name:         "uuids",
			Type:         "bytes32",
		},
		{
			InternalType: "bytes32",
			Name:         "batchHash",
			Type:         "bytes32",
		},
		{
			InternalType: "string",
			Name:         "payloadRef",
			Type:         "string",
		},
		{
			InternalType: "bytes32[]",
			Name:         "contexts",
			Type:         "bytes32[]",
		},
	},
}

var batchPinMethodABI = &abi.Entry{
	Name: "pinBatch",
	Type: "function",
	Inputs: abi.ParameterArray{
		{
			InternalType: "bytes32",
			Name:         "uuids",
			Type:         "bytes32",
		},
		{
			InternalType: "bytes32",
			Name:         "batchHash",
			Type:         "bytes32",
		},
		{
			InternalType: "string",
			Name:         "payloadRef",
			Type:         "string",
		},
		{
			InternalType: "bytes32[]",
			Name:         "contexts",
			Type:         "bytes32[]",
		},
	},
}

var networkActionMethodABI = &abi.Entry{
	Name: "networkAction",
	Type: "function",
	Inputs: abi.ParameterArray{
		{
			InternalType: "string",
			Name:         "action",
			Type:         "string",
		},
		{
			InternalType: "string",
			Name:         "payload",
			Type:         "string",
		},
	},
}

var batchPinEventABI = &abi.Entry{
	Name: "BatchPin",
	Type: "event",
	Inputs: abi.ParameterArray{
		{
			Indexed:      false,
			InternalType: "address",
			Name:         "author",
			Type:         "address",
		},
		{
			Indexed:      false,
			InternalType: "uint256",
			Name:         "timestamp",
			Type:         "uint256",
		},
		{
			Indexed:      false,
			InternalType: "string",
			Name:         "namespace",
			Type:         "string",
		},
		{
			Indexed:      false,
			InternalType: "bytes32",
			Name:         "uuids",
			Type:         "bytes32",
		},
		{
			Indexed:      false,
			InternalType: "bytes32",
			Name:         "batchHash",
			Type:         "bytes32",
		},
		{
			Indexed:      false,
			InternalType: "string",
			Name:         "payloadRef",
			Type:         "string",
		},
		{
			Indexed:      false,
			InternalType: "bytes32[]",
			Name:         "contexts",
			Type:         "bytes32[]",
		},
	},
}

var networkVersionMethodABI = &abi.Entry{
	Name:            "networkVersion",
	Type:            "function",
	StateMutability: "pure",
	Inputs:          abi.ParameterArray{},
	Outputs: abi.ParameterArray{
		{
			InternalType: "uint8",
			Type:         "uint8",
		},
	},
}
