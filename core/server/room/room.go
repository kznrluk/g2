package room

import (
	"g2/domain"
	"g2/player"
)

type room struct {
	id      domain.RoomId
	players domain.Players
}

func (r *room) SetEstimate(id domain.PlayerId, selected string) error {
	panic("implement me")
}

func (r *room) RevealAll(id domain.PlayerId) error {
	panic("implement me")
}

func (r *room) ResetAll(id domain.PlayerId) error {
	panic("implement me")
}

func (r *room) GetRoomId() domain.RoomId {
	panic("implement me")
}

func (r *room) IsThisRoom(id domain.RoomId) bool {
	panic("implement me")
}

func (r *room) ConnectNewPlayer() (domain.PlayerId, error) {
	p := r.players.AddPlayer()
	return p.GetPlayerId(), nil
}

func (r *room) GetInactiveSeconds() int {
	panic("implement me")
}

func (r *room) GetRoomStatus() domain.RoomStatus {
	return domain.RoomStatus{
		RoomId:         r.id,
		SelectablePack: []string{"Mountain Goat Pack"},
		Pack:           []string{"0", "1/2", "1", "2", "3", "5", "8", "13", "20", "40", "100", "?", "☕️"},
		Players:        r.players.GetPlayerStatusList().ToStructArray(),
		IsShown:        false,
	}
}

func NewRoom(id domain.RoomId) domain.Room {
	return &room{
		id:      id,
		players: player.NewPlayers(),
	}
}
