package domain

type Players interface {
	ConnectPlayer() Player
	DisconnectPlayer(Player)

	SetEstimate(id PlayerId, selected string) error
	RevealAll()
	ResetAll()

	GetPlayerStatusList() PlayerStatusList
	find(PlayerId) Player
}

type PlayerStatusList interface {
	ToJSON() string
}

type PlayerStatus struct {
	Id             PlayerId `json:"id,omitempty"`
	IsSelected     bool     `json:"isSelected,omitempty"`
	Selected       string   `json:"selected,omitempty"`
	IsControllable bool     `json:"isControllable,omitempty"`
}

type PlayerId string

type Player interface {
	IsOwnId(PlayerId) bool

	GetPlayerStatus() PlayerStatus
}
