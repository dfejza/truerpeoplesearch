# TruePeopleSearch Scraper
Web scraper for the site http://truepeoplesearch.com
User can search a telephone number, and the resultant address will be displayed. If multiple names found, the user must specify the name.

# Stack
React communicates to the go backend via API. Go backend will scrape truepeoplesearch for the results. All results are stored in a SQL database, and future requests will use the SQL database rather than scrape again.

## Front-End
- React
- MaterialUI
## Back-End
- Go
- SQL