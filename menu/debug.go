package menu

// start a debugging session inside the terminal with a timer and an animated gif

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"

	_ "github.com/mattn/go-sqlite3"

	"time"
)

func Debug(duration int, programName string) {

	data, err := ioutil.ReadFile("./random_facts.json")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	var facts []string
	// convert json to array
	json.Unmarshal(data, &facts)

	// asign file content to data

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
			// clean terminal each second
			fmt.Print("\033[H\033[2J")
			fmt.Println("Debugging session in progress...")
			fmt.Println("Remaining time: ", duration)

			if duration%2 == 0 {
				fmt.Println("  __")
				fmt.Println("<(o )___")
				fmt.Println(" ( ._> /")
				fmt.Println("  `---'")
				fmt.Println("Random fact: ", facts[getRandomNumber(0, len(facts))])

			} else {

				fmt.Println("    __")
				fmt.Println("___( o)>")
				fmt.Println("\\ <_. )")
				fmt.Println(" `---'")

			}

			duration--
			time.Sleep(1 * time.Second)
		}
	}

}

func getRandomNumber(min int, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
