package domain

type WSResponseWriter interface {
	Write(s []byte)
	WriteAll(s []byte)
}

type WSRequest struct {
	Api  string `json:"api"`
	Path string `json:"path"`
	Data []byte `json:"data"`
}
