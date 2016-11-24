package containers

type TwitchReply struct {
	TotalNumStreams int64    `json:"_total"`
	Links           Links    `json:"_links"`
	Streams         []Stream `json:"streams"`
}
