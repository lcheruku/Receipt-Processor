package services


import (
	"math"
	"receipt-processor/models"
	"strconv"
	"strings"
)

/*	Rules:

	These rules collectively define how many points should be awarded to a receipt.

	1. One point for every alphanumeric character in the retailer name.
	2. 50 points if the total is a round dollar amount with no cents.
	3. 25 points if the total is a multiple of 0.25.
	4. 5 points for every two items on the receipt.
	5. If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	6. 6 points if the day in the purchase date is odd.
	7. 10 points if the time of purchase is after 2:00pm and before 4:00pm.
*/

// funciton for calculating points
func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// 1. One point for every alphanumeric character in the retailer name.
	for _, char := range receipt.Retailer {
		if isAlphanumeric(char) {
			points++
		}
	}

	// 2. 50 points if the total is a round dollar amount with no cents.
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == math.Trunc(total) {
		points += 50
	}
	// 3. 25 points if the total is a multiple of 0.25.
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}
	
	// 4. 5 points for every two items on the receipt.
	points += (len(receipt.Items) / 2) * 5
	
	/* 5. If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. 
	The result is the number of points earned. */
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// 6. 6 points if the day in the purchase date is odd.
	dateParts := strings.Split(receipt.PurchaseDate, "-")
	day, _ := strconv.Atoi(dateParts[2])
	if day%2 != 0 {
		points += 6
	}

	// 7. 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	timeParts := strings.Split(receipt.PurchaseTime, ":")
	hour, _ := strconv.Atoi(timeParts[0])
	if hour == 14 {
		points += 10
	}
	
	// return the calculated points
	return points
}

func isAlphanumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}
