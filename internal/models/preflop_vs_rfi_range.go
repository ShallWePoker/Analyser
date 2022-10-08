package models

type PreflopVsRFIRange struct {
	PlayerName      string             `json:"playerName"`
	HeroPosition    string             `json:"heroPosition"`
	VillainPosition string             `json:"villainPosition"`
	HolecardMap     map[string]Actions `json:"holecardMap"`
}
