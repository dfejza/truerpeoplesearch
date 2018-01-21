package main

import (
	"database/sql"
	"fmt"
)

// person is the expected data structure for API input and SQL
type person struct {
	Number  string `json:"number"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// queriedPerson is the restful API response object
type queriedPerson struct {
	Name    string
	Address string
	Ranking int
}

//
//SQL related database transactions
//

// getPerson lookups the specified number in the SQL database, and stores all hits
// into the queriedPerson slice
func getPerson(pD *[]queriedPerson, lookupNumber string) error {
	statement := fmt.Sprintf("SELECT name, address FROM people WHERE number=%s", lookupNumber)
	rows, err := dB.Query(statement)
	defer rows.Close()
	for rows.Next() {
		var name, address string
		err = rows.Scan(&name, &address)
		*pD = append(*pD, queriedPerson{name, address, -2})
	}
	if err != nil {
		return err
	}
	return nil
}

// checkIfStored checks if the given number has been quried and stored into the database in the past
func checkIfStored(lookupNumber string) bool {
	var exists bool
	statement := fmt.Sprintf("SELECT 1 FROM people WHERE number=%s", lookupNumber)
	err := dB.QueryRow(statement).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return false
	}
	return exists
}

// createPerson will create a new entry into the people table
func createPerson(number string, name string, address string) error {
	statement := fmt.Sprintf("INSERT INTO people(number, name, address) VALUES('%s', '%s', '%s')", number, name, address)
	_, err := dB.Exec(statement)
	if err != nil {
		return err
	}
	return nil
}
