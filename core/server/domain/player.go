package domain

type Players interface {
	AddPlayer() Player
	DeletePlayer(id PlayerId) error
	Length() int

	SetEstimate(id PlayerId, selected string) error
	ResetAll()

	GetPlayerStatusList() PlayerStatusList
}

type PlayerStatusList interface {
	ToStructArray() []PlayerStatus
}

type PlayerStatus struct {
	Id             PlayerId `json:"id,omitempty"`
	IsSelected     bool     `json:"isSelected,omitempty"`
	Selected       string   `json:"selected,omitempty"`
	IsControllable bool     `json:"isControllable,omitempty"`
}

type PlayerId string

type Player interface {
	GetPlayerId() PlayerId
	IsOwnId(PlayerId) bool

	Reset()

	SetEstimate(selected string) error
	Promote()

	GetPlayerStatus() PlayerStatus
}
