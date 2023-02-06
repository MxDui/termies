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
	rows, err := db.Query("SELECT * FROM data")
	checkErr(err)

	// create a csv file
	file, err := os.Create("./report.csv")
	checkErr(err)
	defer file.Close()

	// create a csv writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// write the header
	header := []string{"ID", "Date", "Time", "Temperature", "Humidity", "Pressure"}
	err = writer.Write(header)
	checkErr(err)

	// write the data
	for rows.Next() {
		var id int
		var date string
		var time string
		var temperature float64
		var humidity float64
		var pressure float64

		err = rows.Scan(&id, &date, &time, &temperature, &humidity, &pressure)
		checkErr(err)

		// convert the data to strings
		idStr := strconv.Itoa(id)
		temperatureStr := strconv.FormatFloat(temperature, 'f', 2, 64)
		humidityStr := strconv.FormatFloat(humidity, 'f', 2, 64)
		pressureStr := strconv.FormatFloat(pressure, 'f', 2, 64)

		// create a slice of strings
		data := []string{idStr, date, time, temperatureStr, humidityStr, pressureStr}

		// write the data to the csv file
		err = writer.Write(data)
		checkErr(err)
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
