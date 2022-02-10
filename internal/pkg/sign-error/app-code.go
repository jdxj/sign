package sign_error

// common
const (
	// ErrUnknown 500 unknown error
	ErrUnknown int = iota + 100000
	// ErrBindRequest 400 parse request failed
	ErrBindRequest
	// ErrRPCCall 500 rpc call failed
	ErrRPCCall
)

// user
const (
	errReserveUser int = iota + 100100
	// ErrLogin 400 login failed
	ErrLogin
	// ErrAuth 400 auth failed
	ErrAuth
)

// task
const (
	errReserveTask int = iota + 100200
	// ErrEncodeParam 400 encode param failed
	ErrEncodeParam
	// ErrDecodeParam 500 decode param failed
	ErrDecodeParam
)
