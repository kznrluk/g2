package handler

import (
	"encoding/json"
	"fmt"
	"g2/domain"
	"g2/room"
	"github.com/gobwas/ws"
	"net"

	"github.com/gobwas/ws/wsutil"
)

type wsResponseWriter struct {
	currentConn net.Conn
	member      []net.Conn
}

func (w *wsResponseWriter) WriteAll(b []byte) {
	for _, m := range w.member {
		err := wsutil.WriteServerMessage(m, ws.OpText, b)
		if err != nil {
			// パイプ破損してる
			_ = w.currentConn.Close()
		}
	}
}

func (w *wsResponseWriter) Write(b []byte) {
	err := wsutil.WriteServerMessage(w.currentConn, ws.OpText, b)
	if err != nil {
		// パイプ破損してる
		_ = w.currentConn.Close()
	}
}

func NewResponseWriter(w net.Conn, member []net.Conn) domain.WSResponseWriter {
	return &wsResponseWriter{
		w,
		member,
	}
}

var ConnectionPool map[domain.RoomId][]net.Conn

func HandleWebSocket(conn net.Conn) {
	for {
		msg, _, err := wsutil.ReadClientData(conn)
		if err != nil {
			continue // EOFも入ってくる
		}

		var data domain.WSRequest
		err = json.Unmarshal(msg, &data)
		if err != nil {
			fmt.Printf("WebSocket: 未知のデータを着信しました %s", err.Error())
			continue
		}

		ConnectionPool[0] = append(ConnectionPool[0], conn)

		writer := NewResponseWriter(conn, ConnectionPool[0])
		switch data.Api {
		case "CreateRoom":
			WSCreateRoomHandler(writer, data.Data)
		case "Join":

		case "Ping":
			WSPingHandler(writer, data.Data)
		default:
			fmt.Printf("WebSocket: 未知のAPIです %s", data.Api)
		}
	}
}

//TODO: usecaseに移したい

func WSPingHandler(writer domain.WSResponseWriter, data []byte) {
	writer.WriteAll([]byte("Pong"))
}

type WSCreateRoomData struct {
}

func WSCreateRoomHandler(writer domain.WSResponseWriter, data []byte) {
	rooms := room.GetRooms()
	r := rooms.CreateNewRoom()
	id, _ := r.ConnectNewPlayer()
	writer.Write([]byte("{ \"id\": \"" + id + "\" }")) // なんとかしたい

	b, err := json.Marshal(r.GetRoomStatus())
	if err != nil {
		writer.WriteAll([]byte(err.Error()))
	}

	writer.WriteAll(b)
}

func WSJoinHandler(writer domain.WSResponseWriter, data []byte) {

}
