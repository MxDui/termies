package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

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

	fmt.Println(db)

	var menu = flag.String("menu", "default", "menu name")

	flag.Parse()

	switch *menu {
	case "default":
		// do something
	case "debug":
		// do something

	case "report":
		// do something
	case "help":
		// do something

	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
