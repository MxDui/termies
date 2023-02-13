package menu

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Journal is a struct that contains the data for a journal entry
type Journal struct {
	ID          int
	Date        string
	Time        string
	Description string
	Mood        string
}

// AddJournalEntry adds a journal entry to the database
func AddJournalEntry(description string, mood string) {
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO journal(date, time, description, mood) VALUES(?, ?, ?, ?)")
	checkErr(err)

	res, err := stmt.Exec(time.Now(), time.Now(), description, mood)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}

// GetJournalEntries returns all journal entries
func GetJournalEntries() []Journal {
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM journal")
	checkErr(err)

	var entries []Journal

	for rows.Next() {
		var entry Journal
		err = rows.Scan(&entry.ID, &entry.Date, &entry.Time, &entry.Description, &entry.Mood)
		checkErr(err)
		entries = append(entries, entry)
	}

	return entries
}

// GetJournalEntry returns a journal entry by ID
func GetJournalEntry(id int) Journal {
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)

	var entry Journal

	err = db.QueryRow("SELECT * FROM journal WHERE id=?", id).Scan(&entry.ID, &entry.Date, &entry.Time, &entry.Description, &entry.Mood)
	checkErr(err)

	return entry
}

// DeleteJournalEntry deletes a journal entry by ID
func DeleteJournalEntry(id int) {

	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)

	stmt, err := db.Prepare("DELETE FROM journal WHERE id=?")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

}
