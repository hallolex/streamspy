package main

//HTTP_PLATFORM_PORT env variable

import (
    "fmt"
    "net/http"
	"net/url"
	"html"
	"io/ioutil"
	"encoding/json"
)

type TwitchReply struct {
  TotalNumStreams int64 `json:"_total"`
	Links Links `json:"_links"`
	Streams []Streams `json:"streams"`
}

type Streams struct {
	Id int64 `json:"_id"`
	Game string `json:"game"`
	Viewers int64 `json:"viewers"`
	VideoHeight int64 `json:"video_height"`
	AverageFps float64 `json:"average_fps"`
	Delay int64 `json:"delay"`
	CreatedAt string `json:"created_at"`
	IsPlaylist bool `json:"is_playlist"`
	PreviewImages PreviewImages `json:"preview"`
	ChannelInfo ChannelInfo `json:"channel"`
	Links StreamLinks `json:"_links"`
}

type Links struct {
	Self string `json:"self"`
	Next string `json:"next"`
	Featured string `json:"featured"`
	Summary string `json:"summary"`
	Followed string `json:"followed"`
}

type StreamLinks struct {
	Self string `json:"self"`
}

type ChannelLinks struct {
	Self string `json:"self"`
	Follows string `json:"follows"`
	Commercial string `json:"commercial"`
	StreamKey string `json:"stream_key"`
	Chat string `json:"chat"`
	Features string `json:"features"`
	Subscriptions string `json:"subscriptions"`
	Editors string `json:"editors"`
	Teams string `json:"teams"`
	Videos string `json:"videos"`
}

type ChannelInfo struct {
	IsMature bool `json:"mature"`
	IsPartner bool `json:"partner"`
	Status string `json:"status"`
	BroadcasterLanguage string `json:"broadcaster_language"`
	DisplayName string `json:"display_name"`
	Game string `json:"game"`
	Language string `json:"language"`
	Id int64 `json:"_id"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Logo string `json:"logo"`
	VideoBanner string `json:"video_banner"`
	ProfileBanner string `json:"profile_banner"`
	Url string `json:"url"`
	Views int64 `json:"views"`
	Followers int64 `json:"followers"`
	Links ChannelLinks `json:"_links"`
}

type PreviewImages struct {
	Small string `json:"small"`
	Medium string `json:"medium"`
	Large string `json:"large"`
	Template string `json:"template"`
}

func perror(err error) {
    if err != nil {
        panic(err)
    }
}

func GetStreams(body []byte) (*TwitchReply, error) {
    var s = new(TwitchReply)
    err := json.Unmarshal(body, &s)
    if(err != nil){
        fmt.Println("whoops:", err)
    }
    return s, err
}

func ParseGame(w http.ResponseWriter, r *http.Request) {
	var gameName string = html.EscapeString(r.URL.Path[11:])

	fmt.Fprintf(w, "Parsing: " + gameName + "\n\n") // cut off /parsegame/

	var gameURL string = "https://api.twitch.tv/kraken/streams?game=" + url.QueryEscape(gameName) + "&limit=25&client_id=njuhnnm201z2bniflxdypamzbv127xl"

	res, err := http.Get(gameURL)
	perror(err)

	defer res.Body.Close()
	perror(err)

	body, err := ioutil.ReadAll(res.Body)
	perror(err)


	s, err := GetStreams([]byte(body))

	var streamerNames string = ""
	i := 0
    for i < len(s.Streams) {
        streamerNames += s.Streams[i].ChannelInfo.DisplayName + "\n"
        i = i + 1
    }
	fmt.Fprintf(w, streamerNames)
}

func main() {
	http.HandleFunc("/ParseGame/", ParseGame)
    http.ListenAndServe(":8080", nil)
}
