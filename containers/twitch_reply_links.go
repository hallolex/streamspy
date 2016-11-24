package containers

type Links struct {
	Self     string `json:"self"`
	Next     string `json:"next"`
	Featured string `json:"featured"`
	Summary  string `json:"summary"`
	Followed string `json:"followed"`
}
