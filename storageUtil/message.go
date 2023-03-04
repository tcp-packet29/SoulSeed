package storageUtil

type Message struct {
	Msg       string `json:"msg"` //json because redis
	Timestamp int64  `json:"timestamp"`
	UID_Hex   string `json:"uid_hex"`
}
