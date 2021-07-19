package domain

type RoomId uint32

type Rooms interface {
	CreateNewRoom() Room
	Find(id RoomId) Room
	GC()
}

type Room interface {
	IsThisRoom(id RoomId) bool
	ConnectNewPlayer() (PlayerId, error)
	GetInactiveSeconds() int
	GetRoomStatus() RoomStatus
}

type RoomStatus struct {
	RoomId         RoomId         `json:"roomId,omitempty"`
	SelectablePack []string       `json:"selectablePack,omitempty"`
	Pack           []string       `json:"pack,omitempty"`
	Players        []PlayerStatus `json:"players,omitempty"`
	IsShown        bool           `json:"isShown,omitempty"`
}
