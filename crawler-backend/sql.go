package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/renstrom/fuzzysearch/fuzzy" // partial string matching
)

// SQL Schema
/*
CREATE DATABASE teltechcrawler;
USE teltechcrawler;
CREATE TABLE people (
	number VARCHAR(10) NOT NULL,
	name VARCHAR(50) NOT NULL,
	address VARCHAR(250) NOT NULL
);
*/

// initializeSQL initializes the SQL database connection, authenticating this session
func initializeSQL(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	var err error
	dB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}

// queryDB takes the number from the inputted person struct and searches the database.
func findMatchFromDB(data *person) queriedPerson {
	// Pass slice as a function argument and store database query in it
	personList := []queriedPerson{}
	getPerson(&personList, data.Number)

	// If there are no results, return status -2
	if len(personList) == 0 {
		return queriedPerson{" ", " ", -2}
	}

	// If there is no name provided and more than one search result, return status -3
	if len(personList) > 1 && data.Name == " " {
		return queriedPerson{" ", " ", -3}
	}

	// Store the index of the closest match
	// Compare the scraped name with user inputted string, and generate a rank
	// if the rank (higher the better) is greater than the current stored rank, replace the stored rank
	closestMatchIndex := 0
	for i, personDetail := range personList {
		personList[i].Ranking = fuzzy.RankMatch(strings.ToLower(data.Name), strings.ToLower(personDetail.Name))
		if personList[i].Ranking > personList[closestMatchIndex].Ranking {
			closestMatchIndex = i
		}
	}
	// Return the details of who's information was the closest
	return personList[closestMatchIndex]
}
