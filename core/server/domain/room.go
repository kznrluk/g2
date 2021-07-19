package domain

type RoomId uint16

type Rooms interface {
	Add(room Room)
	Find(id RoomId) Room
	GC()
}

type Room interface {
	IsThisRoom(id RoomId) bool
	ConnectPlayer() (error, PlayerId)
	GetInactiveTime() int
}
