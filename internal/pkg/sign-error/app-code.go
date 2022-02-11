package sign_error

//go:generate codegen

type code = int

// common
const (
	// ErrUnknown 500 unknown error
	ErrUnknown code = iota + 100000
	// ErrBindRequest 400 parse request failed
	ErrBindRequest
	// ErrRPCCall 500 rpc call failed
	ErrRPCCall
)

// user
const (
	_ code = iota + 100100
	// ErrLogin 400 login failed
	ErrLogin
	// ErrAuth 400 auth failed
	ErrAuth
)

// task
const (
	_ code = iota + 100200
	// ErrEncodeParam 400 encode param failed
	ErrEncodeParam
	// ErrDecodeParam 500 decode param failed
	ErrDecodeParam
)
