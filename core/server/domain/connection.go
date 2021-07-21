package domain

type WSResponseWriter interface {
	Write(s []byte)
	WriteAll(s []byte)
}

type WSRequest struct {
	Api    string `json:"api"`
	Id     string `json:"id"`
	RoomId int32  `json:"roomId"`
	Data   []byte `json:"data"`
}
