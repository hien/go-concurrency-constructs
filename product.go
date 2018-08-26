package main

import (
	"os"

	"github.com/anaskhan96/soup"
)

// A Product represents a product on amazon
type Product struct {
	Name    string
	Link    string
	Image   string
	Price   string
	Reviews []Review
}

// GetReviews gets the product's top reviews from amazon product page
func (product *Product) GetReviews() {
	// now := time.Now().UTC()
	resp, err := soup.Get(product.Link)
	// fmt.Println("Fetching time: ", time.Since(now))

	// now = time.Now().UTC()

	if err != nil {
		os.Exit(1)
	}

	doc := soup.HTMLParse(resp)

	reviewsContainer := doc.Find("div", "class", "reviews-content")

	if reviewsContainer.Error != nil {
		return
	}

	rawReviews := reviewsContainer.FindAll("div", "class", "review")
	reviews := []Review{}

	for _, rawReview := range rawReviews {
		review := Review{}
		err := review.ParseReviews(rawReview)

		if err == nil {
			reviews = append(reviews, review)
		}
	}

	product.Reviews = reviews
	// fmt.Println("Review parsing time: ", time.Since(now))
}
