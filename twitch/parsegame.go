package twitch

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/hallolex/streamspy/containers"
	"github.com/hallolex/streamspy/utils"
)

func ParseGame(w http.ResponseWriter, r *http.Request) {
	var gameName string = html.EscapeString(r.URL.Path[11:]) // cut off /parsegame/

	fmt.Fprintf(w, "Parsing: "+gameName+"\n\n")

	var gameURL string = "https://api.twitch.tv/kraken/streams?game=" + url.QueryEscape(gameName) + "&limit=10&client_id=njuhnnm201z2bniflxdypamzbv127xl"

	res, err := http.Get(gameURL)
	utils.PanicError(err)

	defer res.Body.Close()
	utils.PanicError(err)

	body, err := ioutil.ReadAll(res.Body)
	utils.PanicError(err)

	s, err := GetStreams([]byte(body))

	var streamerNames string = ""
	i := 0
	for i < len(s.Streams) {
		streamerNames += s.Streams[i].ChannelInfo.DisplayName + "\n"
		i = i + 1
	}
	fmt.Fprintf(w, streamerNames)
}

func GetStreams(body []byte) (*containers.TwitchReply, error) {
	var s = new(containers.TwitchReply)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
