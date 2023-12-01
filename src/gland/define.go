package gland

var (
	RetCodeSuccess int = 0
	MsgSuccess         = "success"
	CodeMsgSuccess     = New(RetCodeSuccess, MsgSuccess, nil)

	RetCodeFailure int = 1
	MsgFailure         = "failure"
	CodeMsgFailure     = New(RetCodeFailure, MsgFailure, nil)

	RetCodeUnknown int = 2
	MsgUnknown         = "unknown"
	CodeMsgUnknown     = New(RetCodeUnknown, MsgUnknown, nil)
)
