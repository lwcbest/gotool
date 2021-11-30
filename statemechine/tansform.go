package statemechine

type Transform struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Event  string `json:"event"`
	Action string `json:"action"`
}