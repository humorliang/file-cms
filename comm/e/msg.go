package e

//获取异常信息
func GetMsg(msgCode int) string {
	m, ok := msg[msgCode]
	if ok {
		return m
	}
	return msg[INTERSERVER_ERROR]
}
