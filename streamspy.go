package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hallolex/streamspy/containers"
	"github.com/hallolex/streamspy/utils"
)

var db *sql.DB

func GetStreams(body []byte) (*containers.TwitchReply, error) {
	var s = new(containers.TwitchReply)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

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

const (
	DB_HOST = "tcp(eu-cdbr-azure-west-a.cloudapp.net:3306)"
	DB_NAME = "streamspydb"
	DB_USER = "bf9c32a84b81fc"
	DB_PASS = "fb6aa96a"
)

func SetupDatabase() {
	var err error
	db, err = sql.Open("mysql", DB_USER+":"+DB_PASS+"@"+DB_HOST+"/"+DB_NAME)
	utils.PanicError(err)
	defer db.Close()

	err = db.Ping()
	utils.PanicError(err)
}

func CloseDatabase() {
	db.Close()
}

func Database(w http.ResponseWriter, r *http.Request) {

	var err error
	db, err = sql.Open("mysql", DB_USER+":"+DB_PASS+"@"+DB_HOST+"/"+DB_NAME)
	utils.PanicError(err)
	defer db.Close()

	err = db.Ping()
	utils.PanicError(err)

	user := new(containers.StreamSpyUser)
	user.Token = html.EscapeString(r.URL.Path[10:]) // get token from url
	fmt.Fprintf(w, "Token: "+user.Token+"\n\n")

	err = db.QueryRow("SELECT id, username, slack_url, slack_threshold FROM users WHERE token = ?", user.Token).Scan(&user.Id, &user.Username, &user.SlackURL, &user.SlackThreshold)
	utils.PanicError(err)

	fmt.Fprintf(w, "Token belongs to user: "+user.Username)
}

func main() {
	//SetupDatabase()
	http.HandleFunc("/ParseGame/", ParseGame)
	http.HandleFunc("/Database/", Database)
	http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
}
