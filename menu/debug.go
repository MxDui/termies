package menu

// start a debugging session inside the terminal with a timer and an animated gif

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"time"
)

func Debug(duration int, programName string) {

	fmt.Println("Starting debug session...")
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)

	// start a timer
	timer := time.NewTimer(time.Duration(duration) * time.Second)

	for {
		select {
		case <-timer.C:
			// save session to database
			stmt, err := db.Prepare("INSERT INTO sessions(date, duration, program) VALUES(?, ?, ?)")
			checkErr(err)

			res, err := stmt.Exec(time.Now(), duration, programName)
			checkErr(err)

			id, err := res.LastInsertId()
			checkErr(err)

			fmt.Println(id)
			fmt.Println("Debug session ended.")
			return
		default:
			fmt.Println("Debugging...")
			time.Sleep(1 * time.Second)
		}
	}

}
