package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"termies.mxdui.com/menu"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// create file if it doesn't exist
	if _, err := os.Stat("./data.db"); os.IsNotExist(err) {
		file, err := os.Create("./data.db")
		checkErr(err)
		defer file.Close()
	}

	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)

	db.Exec("CREATE TABLE IF NOT EXISTS data (id INTEGER PRIMARY KEY, date TEXT, time TEXT, temperature REAL, humidity REAL, pressure REAL)")
	db.Exec("CREATE TABLE IF NOT EXISTS sessions (id INTEGER PRIMARY KEY, date TEXT, duration INTEGER, program TEXT)")

	var menuSwitch = flag.String("m", "default", "menu name")
	var debugDuration = flag.Int("d", 10, "debug duration")
	var programName = flag.String("n", "default", "program name")

	flag.Parse()

	switch *menuSwitch {
	case "default":
		// do something
		menu.Help()

	case "debug":
		// do something
		menu.Debug(*debugDuration, *programName)

	case "report":
		// do something
		menu.Report()

	case "help":
		// do something
		menu.Help()

	default:
		// do something
		fmt.Println("Invalid menu name")
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
