package containers

type Stream struct {
	Id            int64         `json:"_id"`
	Game          string        `json:"game"`
	Viewers       int64         `json:"viewers"`
	VideoHeight   int64         `json:"video_height"`
	AverageFps    float64       `json:"average_fps"`
	Delay         int64         `json:"delay"`
	CreatedAt     string        `json:"created_at"`
	IsPlaylist    bool          `json:"is_playlist"`
	PreviewImages PreviewImages `json:"preview"`
	ChannelInfo   ChannelInfo   `json:"channel"`
	Links         StreamLinks   `json:"_links"`
}
