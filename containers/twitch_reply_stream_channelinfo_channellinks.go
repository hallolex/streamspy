package containers

type ChannelLinks struct {
	Self          string `json:"self"`
	Follows       string `json:"follows"`
	Commercial    string `json:"commercial"`
	StreamKey     string `json:"stream_key"`
	Chat          string `json:"chat"`
	Features      string `json:"features"`
	Subscriptions string `json:"subscriptions"`
	Editors       string `json:"editors"`
	Teams         string `json:"teams"`
	Videos        string `json:"videos"`
}
