<<<<<<< HEAD
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
=======
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
>>>>>>> 4aad3e1b64427d5ebafb07f037b140f7eb3a6511
