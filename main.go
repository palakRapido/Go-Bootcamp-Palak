package main

import (
	"bufio"
	"fmt"
	"goBootCamp/rating/ratings"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Welcome to Rating Project!")
	fmt.Println("=========================================")

	scanner := bufio.NewScanner(os.Stdin)

	// Get product ID
	fmt.Print("Enter Product ID: ")
	scanner.Scan()
	productID := scanner.Text()

	// Create a Rating instance
	productRating := &ratings.Rating{
		ID:        productID,
		AvgRating: 0,
		Ratings:   []ratings.UserRating{},
	}

	// Loop to get ratings from user
	for {
		fmt.Println("\n--- Add New Rating ---")
		fmt.Print("Enter User ID (or 'quit' to exit): ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		// Check if user wants to quit
		if strings.ToLower(input) == "quit" || strings.ToLower(input) == "q" {
			break
		}

		var uid int
		_, err := fmt.Sscanf(input, "%d", &uid)
		if err != nil {
			fmt.Println("Error: Invalid User ID. Please enter a number.")
			continue
		}

		// Get rating value
		fmt.Print("Enter rating (0-5): ")
		var ratingValue float64
		_, err = fmt.Scanf("%f", &ratingValue)
		if err != nil {
			fmt.Println("Error: Invalid rating. Please enter a number.")
			scanner.Scan() // Clear the scanner buffer
			continue
		}
		//scanner.Scan() // Clear the newline

		// Get comments
		fmt.Print("Enter comments: ")
		scanner.Scan()
		comments := scanner.Text()

		// Add the rating
		e := productRating.AddRating(uid, ratingValue, comments)
		if e != nil {
			fmt.Println("Error:", e)
		} else {
			fmt.Println("✓ Rating added successfully!")

		}
	}

	// Display final results
	fmt.Println("\n========================================")
	fmt.Println("         PRODUCT RATINGS SUMMARY")
	fmt.Println("========================================")
	fmt.Printf("Product ID: %s\n", productRating.ID)
	//fmt.Printf("Average Rating: %.2f/5\n", productRating.AvgRating)
	if productRating.AvgRating >= 3 {
		color.Green("Great Current Average Rating: %.2f/5 (Total: %d ratings)", productRating.AvgRating, len(productRating.Ratings))
	} else {
		color.Red("Sorry Current Average Rating: %.2f/5 (Total: %d ratings)\n", productRating.AvgRating, len(productRating.Ratings))
	}
	fmt.Printf("Total Ratings: %d\n", len(productRating.Ratings))

	if len(productRating.Ratings) > 0 {
		fmt.Println("\nIndividual Ratings:")
		fmt.Println("-------------------")
		for i, rating := range productRating.Ratings {
			fmt.Printf("%d. %s\n", i+1, rating)
		}
	} else {
		fmt.Println("\nNo ratings added.")
	}

	fmt.Println("========================================")
	fmt.Println("Thank you for using Rating Project!")
}
