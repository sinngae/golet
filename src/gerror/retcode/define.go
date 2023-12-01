package retcode

/*
0 是返回成功，不用带源服务编码
*/

const (
	Success = uint64(0)

	MsgSuc     = "success"
	MsgUnknown = "unknown"
	MsgNoData  = "no-data"
)
