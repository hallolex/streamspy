package containers

type ChannelInfo struct {
	IsMature            bool         `json:"mature"`
	IsPartner           bool         `json:"partner"`
	Status              string       `json:"status"`
	BroadcasterLanguage string       `json:"broadcaster_language"`
	DisplayName         string       `json:"display_name"`
	Game                string       `json:"game"`
	Language            string       `json:"language"`
	Id                  int64        `json:"_id"`
	Name                string       `json:"name"`
	CreatedAt           string       `json:"created_at"`
	UpdatedAt           string       `json:"updated_at"`
	Logo                string       `json:"logo"`
	VideoBanner         string       `json:"video_banner"`
	ProfileBanner       string       `json:"profile_banner"`
	Url                 string       `json:"url"`
	Views               int64        `json:"views"`
	Followers           int64        `json:"followers"`
	Links               ChannelLinks `json:"_links"`
}
