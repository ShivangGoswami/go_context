package google

import (
	"context"
	"net/http"
)

//Results is an ordered list of search results.
type Results []Results

// A Result contains the title and URL of a search result.
type Result struct {
	Title, URL string
}

//Search sends query to Google search and returns the results.
func Search(ctx context.Context, query string) (Results, error) {
	//Prepare the Google Search API request.
	req, err := http.NewRequest("GET", "https://ajax.googleapis.com/ajax/services/search/web?v=1.0", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Set("q", query)

	//If ctx is carrying
}
