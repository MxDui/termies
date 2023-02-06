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

	// on the first run, the default menu is displayed

	fmt.Println("Welcome to the program")

	// make the file to store the data in the current directory

	// create file if it doesn't exist
	if _, err := os.Stat("./data.db"); os.IsNotExist(err) {
		file, err := os.Create("./data.db")
		checkErr(err)
		defer file.Close()
	}

	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)

	db.Exec("CREATE TABLE IF NOT EXISTS data (id INTEGER PRIMARY KEY, date TEXT, time TEXT, temperature REAL, humidity REAL, pressure REAL)")
	db.Exec("CREATE TABLE IF NOT EXISTS sessions (id INTEGER PRIMARY KEY, date TEXT, duration INTEGER)")

	var menuSwitch = flag.String("menu", "default", "menu name")
	var debugDuration = flag.Int("duration", 10, "debug duration")

	flag.Parse()

	switch *menuSwitch {
	case "default":
		// do something
	case "debug":
		// do something
		menu.Debug(*debugDuration)

	case "report":
		// do something
		menu.Report()

	case "help":
		// do something

	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
