package player

import "g2/domain"

type player struct {
	id             domain.PlayerId
	estimate       string
	isControllable bool
}

func (p *player) IsOwnId(id domain.PlayerId) bool {
	return p.id == id
}

func (p *player) SetEstimate(selected string) error {
	p.estimate = selected
	return nil
}

func (p *player) Reset() {
	p.estimate = ""
}

func (p *player) GetPlayerStatus() domain.PlayerStatus {
	return domain.PlayerStatus{
		Id:             p.id,
		IsSelected:     p.estimate != "",
		Selected:       p.estimate,
		IsControllable: p.isControllable,
	}
}

func (p *player) Promote() {
	p.isControllable = true
}

func (p *player) GetPlayerId() domain.PlayerId {
	return p.id
}

func NewPlayer(id domain.PlayerId) domain.Player {
	return &player{
		id:             id,
		isControllable: false,
	}
}
