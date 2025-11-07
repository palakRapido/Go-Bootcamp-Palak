package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {

	color.Red("Hello, World!")

	type ProductReview struct {
		ProductName  string
		CustomerName string
		Rating       int
		Comments     []string
	}
	
	// Create a ProductReview struct instance
	review := ProductReview{
		ProductName:  "Apple iPhone",
		CustomerName: "Palak",
		Rating:       4,
		Comments:     []string{"Great product!"},
	}
	
	// Append "nice" to comments
	review.Comments = append(review.Comments, "nice")
	
	// Print the review
	fmt.Println("\n========== Product Review ==========")
	fmt.Printf("Product: %s\n", review.ProductName)
	fmt.Printf("Customer: %s\n", review.CustomerName)
	fmt.Printf("Rating: %d/5\n", review.Rating)
	
	// Print stars based on rating
	fmt.Print("Stars: ")
	for i := 0; i < review.Rating; i++ {
		fmt.Print("⭐")
	}
	fmt.Println()
	
	fmt.Println("\nComments:")
	for i, comment := range review.Comments {
		fmt.Printf("%d. %s\n", i+1, comment)
	}
	fmt.Println("====================================")
	
	// Switch case for rating feedback
	switch review.Rating {
	case 5, 4:
		color.Green("\nThanks for your positive feedback!")
	case 3, 2, 1:
		color.Red("\nWe regret and will try to improve!")
	default:
		fmt.Println("\nInvalid rating")
	}
	
	fmt.Println("\nLength of comments:", len(review.Comments))
	
		
	// Create a map of fruits with prices
		fruits := map[string]int{
			"Apple":      50,
			"Banana":     30,
			"Strawberry": 120,
		}
					
		// Check if "Banana" exists in map
		price, exists := fruits["Orange"]
		if exists {
			fmt.Printf("Orange exists in map with price: $%d\n", price)
		} else {
			fmt.Println("Orange does not exist in map")
		}
		
		fmt.Println("\nFruits Map:")
		for fruit, price := range fruits {
			fmt.Printf("%s: $%d\n", fruit, price)
		}

}
