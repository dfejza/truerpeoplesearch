package main

import (
	"database/sql"

	"github.com/gin-gonic/gin" // restful API library
	_ "github.com/go-sql-driver/mysql"
)

var router *gin.Engine
var dB *sql.DB

// Program entry
func main() {
	// init sql
	// Credentials should be stored as enviornment variables...
	initializeSQL("teltech", "teltech", "teltechcrawler")

	// Routing stuff
	// Create the router
	router = gin.Default()

	// Define the routes
	initializeRoutes()

	// Listen
	router.Run("127.0.0.1:3001")
}

// Build go backend with
// 'go build .\main.go .\crawler.go .\constants.go .\model.go .\routes.go .\sql.go'
// Build react frontend using provided script, 'npm run build'
// Build the sql schema using the the commented code found on line 12 of sql.go
// frontend runs on port 3000, backend on 3001
//
// Caveats;
// - No detection of crawling denied due to truepeople.com bot detection
// - Front end handles sanitizing of input, however backend should check too
// - Front end CSS assumes desktop
// - Didn't look too much into the CORS error, I feel like I swept the problem under the rug with this approach
// - No loading indicator on frontend i.e. meaningful loading messages (Crawling Site -> Checking DB -> ect)
// - SQL errors are not properly reported and dealt with
// - Not to happy on how I handled error communication with the backend to the frontend. Errors are stored in the 'ranking' member of the return struct.
//
