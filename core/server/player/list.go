package player

import (
	"fmt"
	"g2/domain"
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

type players struct {
	list []domain.Player
}

func (p *players) AddPlayer() domain.Player {
	// 多分ビミョい entropyは使いまわすべき？
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	s := NewPlayer(domain.PlayerId(id.String()))
	p.list = append(p.list, s)

	return s
}

func (p *players) Length() int {
	return len(p.list)
}

func (p *players) DeletePlayer(id domain.PlayerId) error {
	panic("implement me")
}

func (p *players) SetEstimate(id domain.PlayerId, selected string) error {
	s := p.find(id)
	if s == nil {
		return fmt.Errorf("player: プレイヤーが見つかりませんでした")
	}

	return s.SetEstimate(selected)
}

func (p *players) ResetAll() {
	for _, d := range p.list {
		d.Reset()
	}
}

func (p *players) GetPlayerStatusList() domain.PlayerStatusList {
	var li []domain.PlayerStatus
	for _, s := range p.list {
		li = append(li, s.GetPlayerStatus())
	}

	return NewPlayerList(li)
}

func (p *players) find(id domain.PlayerId) domain.Player {
	for _, s := range p.list {
		if s.IsOwnId(id) {
			return s
		}
	}

	return nil
}

func NewPlayers() domain.Players {
	var list []domain.Player
	return &players{
		list: list,
	}
}

type playerStatus struct {
	list []domain.PlayerStatus
}

func (p playerStatus) ToStructArray() []domain.PlayerStatus {
	return p.list
}

func NewPlayerList(list []domain.PlayerStatus) domain.PlayerStatusList {
	return &playerStatus{list: list}
}
