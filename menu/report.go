package menu

// export all the data from the database to a csv file

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// Report is the main function for the report menu
func Report() {
	// create file if it doesn't exist
	if _, err := os.Stat("./data.db"); os.IsNotExist(err) {
		file, err := os.Create("./data.db")
		checkErr(err)
		defer file.Close()
	}

	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)

	// get all the data from the database
	rows, err := db.Query("SELECT * FROM sessions")
	checkErr(err)

	// create a csv file
	file, err := os.Create("./report.csv")
	checkErr(err)
	defer file.Close()

	// create a csv writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// write the header
	writer.Write([]string{"id", "date", "duration"})
	for rows.Next() {
		var id int
		var date string
		var duration int
		err = rows.Scan(&id, &date, &duration)
		checkErr(err)

		// write the data to the csv file
		writer.Write([]string{strconv.Itoa(id), date, strconv.Itoa(duration)})
	}

	// close the database connection
	db.Close()

	// print the report
	fmt.Println("Report created successfully")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
