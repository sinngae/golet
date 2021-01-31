package gland

var (
	RetCodeSuccess int = 0
	MsgSuccess         = "success"
	CodeMsgSuccess     = NewRetCodeMsg(RetCodeSuccess, MsgSuccess)

	RetCodeFailure int = 1
	MsgFailure         = "failure"
	CodeMsgFailure     = NewRetCodeMsg(RetCodeFailure, MsgFailure)

	RetCodeUnknown int = 2
	MsgUnknown         = "unknown"
	CodeMsgUnknown     = NewRetCodeMsg(RetCodeUnknown, MsgUnknown)
)
