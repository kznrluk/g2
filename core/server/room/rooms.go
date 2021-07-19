package room

import (
	"g2/domain"
	"math/rand"
)

var roomsSingleton map[domain.RoomId]domain.Room

func InitRooms() {
	roomsSingleton = make(map[domain.RoomId]domain.Room)
}

type rooms struct{}

func (r *rooms) CreateNewRoom() domain.Room {
	id := domain.RoomId(rand.Int31())
	if r.Find(id) != nil {
		return r.CreateNewRoom() // まさか埋まらないでしょ
	}

	room := NewRoom(id)
	roomsSingleton[id] = room
	return room
}

func (r *rooms) Find(id domain.RoomId) domain.Room {
	d, ok := roomsSingleton[id]
	if !ok {
		return nil
	}
	return d
}

func (r *rooms) GC() {
	panic("implement me")
}

func GetRooms() domain.Rooms {
	return &rooms{}
}
