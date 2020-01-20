package eodhdapi

//go:generate go run github.com/mailru/easyjson/easyjson -omit_empty -all fundamentals_unknown_allowed.go

type SectorWeightsGroup struct {
	Cyclical  SectorWeights `json:"Cyclical"`
	Defensive SectorWeights `json:"Defensive"`
	Sensitive SectorWeights `json:"Sensitive"`
}
