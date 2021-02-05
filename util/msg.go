package util

const (
	SUCCESS = 200
	ERROR   = 500

	LackParameter = 111000 + iota //缺少参数
	QueryFails                    //查询失败
)

var MsgFlags = map[int]string{
	SUCCESS:       "ok",
	ERROR:         "fail",
	LackParameter: "缺少参数",
	QueryFails:    "查询失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
