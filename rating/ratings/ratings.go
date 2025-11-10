package ratings

import (
	"fmt"
	"time"
)

// Rating struct
type Rating struct {
	ID        string
	AvgRating float64
	Ratings   []UserRating
}

type UserRating struct {
	UID      int
	Rating   float64
	Comments Comment
}

// String method for Comment
func (c Comment) String() string {
	return fmt.Sprintf("Comment: %s, Date: %s", c.Comment, c.Date)
}

// String method for UserRating
func (u UserRating) String() string {
	return fmt.Sprintf("UID: %d, Rating: %.2f/5, Comments: %s", u.UID, u.Rating, u.Comments)
}

// GetRatingInfo returns all rating information as a formatted string
func (r Rating) String() string {
	return fmt.Sprintf("ID: %s, Avg Rating: %.2f, Ratings: %v", r.ID, r.AvgRating, r.Ratings)
}

// AddRating adds a new UserRating to the Rating
func (r *Rating) AddRating(uid int, rating float64, comments string) error {
	// Validate rating is between 0 and 5
	if rating < 0 || rating > 5 {
		return fmt.Errorf("invalid rating %.2f: must be between 0 and 5", rating)
	}

	newRating := UserRating{
		UID:    uid,
		Rating: rating,
		Comments: Comment{
			Comment: comments,
			Date:    time.Now(),
		},
	}
	r.Ratings = append(r.Ratings, newRating)

	// Calculate new average
	total := 0.0
	for _, ur := range r.Ratings {
		total += ur.Rating
	}
	r.AvgRating = total / float64(len(r.Ratings))

	return nil
}
