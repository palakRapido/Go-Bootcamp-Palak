package main

import (
	"fmt"
	"goBootCamp/rating/ratings"
)

func main() {
	fmt.Println("Welcome to Rating Project!")

	// Create a Rating instance
	productRating := &ratings.Rating{
		ID:        "product123",
		AvgRating: 0,
		Ratings:   []ratings.UserRating{},
	}

	e := productRating.AddRating(1, 5.0, "Excellent product!")
	if e != nil {
		fmt.Println(e)

	}

	e = productRating.AddRating(2, -5.0, "Very good quality")
	if e != nil {
		fmt.Println(e)

	}

	e = productRating.AddRating(3, 2.0, "Highly recommended")
	if e != nil {
		fmt.Println(e)

	}
	fmt.Println(productRating)

}
