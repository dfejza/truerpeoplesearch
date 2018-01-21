package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery" // web crawler
)

// CrawlSite will scrape https://www.truepeoplesearch.com/results?phoneno=(XXX)XXX-XXXX
//  where (XXX)XXX-XXXX is the user inputted phone number
// The resulting information is then stored in SQL database
func crawlSite(data *person) queriedPerson {
	// Build the URL
	var websiteURL = searchByPhoneURL + "(" + data.Number[:3] + ")" + data.Number[3:]
	// Build the doc explorer
	doc, err := goquery.NewDocument(websiteURL)
	if err != nil {
		log.Fatal(err)
	}

	// struct holding what to be returned after the crawl
	// if Ranking == -2, then no results
	// if Ranking == -3, multiple results with no name
	// otherwise read the content of the struct
	closestMatch := queriedPerson{" ", " ", -2}

	// Find and stop at the first instance of records class. This is where the number of hits is written
	text := doc.Find(".record-count:first-of-type").Text()

	// Search for the correct keyword
	pos := strings.Index(text, searchTerm)

	// If at least 1 record was found according to the keyword
	if pos > -1 {
		// Scrape the search results
		doc.Find(".card-summary").Each(func(i int, s *goquery.Selection) {
			// For each item found, parse the string found in a div with the classname h4
			crawledName := s.Find("div .h4").Text()
			// Remove trailing space, the result is the name
			crawledName = strings.TrimSpace(crawledName)
			// Find the link associated with the name
			correspondingLink, _ := s.Find("a").Attr("href")

			// Crawl into the stored link and parse for the address
			docLink, err := goquery.NewDocument(uRLBase + correspondingLink)
			if err != nil {
				log.Fatal(err)
			}

			crawledAddress := docLink.Find("body .shadow-form .content-value .link-to-more:first-of-type").Text()
			// Remove the all text after 'Map" substring
			crawledAddress = crawledAddress[:strings.Index(crawledAddress, "Map")]
			// Remove trailing space, the result is the name
			crawledAddress = strings.TrimSpace(crawledAddress)
			// Remove new lines
			crawledAddress = strings.Replace(crawledAddress, "\n", " ", -1)

			// Store in SQL
			err = createPerson(data.Number, crawledName, crawledAddress)
			if err != nil {
				log.Fatal(err)
			}

		})

		// Now that the data is entered, call the function that would have been used to query the information
		return findMatchFromDB(data)
	}

	return closestMatch
}
