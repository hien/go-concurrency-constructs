package main

import (
	"github.com/anaskhan96/soup"
)

// A Review represents a review on amazon
type Review struct {
	Name    string
	Rating  string
	Content string
}

// ParseReviews parses the HTML content
func (review *Review) ParseReviews(raw soup.Root) error {
	contentHolder := raw.Find("div", "class", "a-expander-content")

	if contentHolder.Error != nil {
		return contentHolder.Error
	}

	review.Name = raw.Find("span", "class", "a-profile-name").Text()
	review.Rating = raw.Find("span", "class", "a-icon-alt").Text()
	review.Content = contentHolder.Text()

	return nil
}
