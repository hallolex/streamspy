package main

import (
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hallolex/streamspy/database"
	"github.com/hallolex/streamspy/twitch"
)

func main() {
	// set up database connection
	database.SetupDatabase()
	defer database.CloseDatabase()

	// handle routing
	http.HandleFunc("/ParseGame/", twitch.ParseGame)
	http.HandleFunc("/Database/", database.Database)
	http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
}
