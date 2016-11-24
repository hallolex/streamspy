package database

import (
	"database/sql"
	"fmt"
	"html"
	"net/http"

	"github.com/hallolex/streamspy/containers"
	"github.com/hallolex/streamspy/utils"
)

var Db *sql.DB
var err error

const (
	DB_HOST = "tcp(eu-cdbr-azure-west-a.cloudapp.net:3306)"
	DB_NAME = "streamspydb"
	DB_USER = "bf9c32a84b81fc"
	DB_PASS = "fb6aa96a"
)

func SetupDatabase() {
	Db, err = sql.Open("mysql", DB_USER+":"+DB_PASS+"@"+DB_HOST+"/"+DB_NAME)
	utils.PanicError(err)
}

func CloseDatabase() {
	Db.Close()
}

func Database(w http.ResponseWriter, r *http.Request) {

	user := new(containers.StreamSpyUser)
	user.Token = html.EscapeString(r.URL.Path[10:]) // get token from url
	fmt.Fprintf(w, "Token: "+user.Token+"\n\n")

	err = Db.QueryRow("SELECT id, username, slack_url, slack_threshold FROM users WHERE token = ?", user.Token).Scan(&user.Id, &user.Username, &user.SlackURL, &user.SlackThreshold)
	utils.PanicError(err)

	fmt.Fprintf(w, "Token belongs to user: "+user.Username)
}
