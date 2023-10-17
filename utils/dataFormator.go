package utils

import (
	"fmt"
	"time"
)

func ParseDateInput(input string) (time.Time, error) {
	// Define the expected date layout (YYYY-MM-DD)
	layout := "2006-01-02"

	// Parse the user input using the defined layout
	parsedTime, err := time.Parse(layout, input)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		// return
	}

	return parsedTime, nil
}
