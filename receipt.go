package main

import (
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type item struct {
	Description string `json:"shortDescription"`
	Price       string `json:"price"`
}

type receipt struct {
	ID           int    `json:"id"`
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total        string `json:"total"`
	Points       int    `json:"points"`
	Items        []item `json:"items"`
}

var receipts = []receipt{}

func getReceiptPoints(ginContext *gin.Context) {
	l.Println("Getting Receipts Points")
	id, err := strconv.Atoi(ginContext.Param("id"))

	if err != nil {
		l.Printf("Invalid receipt: %s - %s", ginContext.Param("id"), err)
	}

	l.Printf("Requesting #%d Receipt Points", id)

	for _, receipt := range receipts {
		if receipt.ID == id {
			ginContext.IndentedJSON(http.StatusOK, receipt.Points)
			return
		}
	}

	l.Printf("#%d Receipt not found", id)
	ginContext.IndentedJSON(http.StatusNotFound, gin.H{"message": "receipt not found"})
}

func processReceipt(ginContext *gin.Context) {
	l.Println("Processing Receipts")

	var newReceipt receipt

	if err := ginContext.BindJSON(&newReceipt); err != nil {
		l.Printf("Bad Receipt - %s", err)
		l.Println(newReceipt)
		return
	}

	calculateReceipt(&newReceipt)

	receipts = append(receipts, newReceipt)
	ginContext.IndentedJSON(http.StatusCreated, newReceipt)
}

func calculateReceipt(receipt *receipt) {
	receipt.ID = len(receipts)

	l.Printf("Calculate Info for Receipt #%d", receipt.ID)

	points := 0
	points += pointsForRetailer(receipt.Retailer)
	points += pointsForTotal(receipt.Total)
	points += pointsForItems(receipt.Items)
	points += pointsForPurchaseDate(receipt.PurchaseDate)
	points += pointsForPurchaseTime(receipt.PurchaseTime)

	receipt.Points = points

	l.Println(receipt)
}

func pointsForRetailer(retailer string) int {
	// * One point for every alphanumeric character in the retailer name.
	points := 0

	for _, character := range retailer {
		if character >= 'a' && character <= 'z' {
			points += 1
		}
		if character >= 'A' && character <= 'Z' {
			points += 1
		}
		if character >= '0' && character <= '9' {
			points += 1
		}
	}

	l.Printf("- retailer '%s' gave %d points", retailer, points)

	return points
}

func pointsForTotal(total string) int {
	// * 50 points if the total is a round dollar amount with no cents.
	// * 25 points if the total is a multiple of `0.25`.
	points := 0

	parsedTotal, err := strconv.ParseFloat(total, 64)
	if err != nil {
		l.Printf("- %s gave %d points - wrong format - %s", total, points, err)
		return points
	}

	_, cents := math.Modf(parsedTotal)
	if cents == 0.0 {
		points += 50
	} else {
		if math.Mod(parsedTotal, 0.25) == 0 {
			points += 25
		}
	}

	l.Printf("- total %f gave %d points", parsedTotal, points)

	return points
}

func pointsForItems(items []item) int {
	// * If the trimmed length of the item description is a multiple of 3,
	// multiply the price by `0.2` and round up to the nearest integer.
	// The result is the number of points earned.
	// * 5 points for every two items on the receipt.
	points := 0

	pairPoints := (len(items) / 2) * 5
	points += pairPoints
	l.Printf("- #itens has %d pairs gave %d points", len(items)/2, pairPoints)

	for _, item := range items {
		if (len(strings.Trim(item.Description, " ")) % 3) == 0 {

			parsedPrice, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				l.Printf("- %s gave %d points - wrong format - %s", item.Price, points, err)
				return points
			}

			result := int(math.Ceil(parsedPrice * 0.2))
			l.Printf("- description '%s' gave %d points", item.Description, result)
			points += result
		}
	}

	return points
}

func pointsForPurchaseDate(purchaseDate string) int {
	// * 6 points if the day in the purchase date is odd.
	points := 0
	dateStr := "2006-01-02"

	parsedDate, err := time.Parse(dateStr, purchaseDate)
	if err != nil {
		l.Printf("%s gave %d points - wrong format - %s", purchaseDate, points, err)
		return points
	}

	day := parsedDate.Day()

	if (day % 2) == 1 {
		l.Println(day)
		points += 6
	}

	l.Printf("- date '%s' gave %d points", purchaseDate, points)
	return points
}

func pointsForPurchaseTime(purchaseTime string) int {
	// * 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	points := 0
	dateStr := "15:04"

	parsedTime, err := time.Parse(dateStr, purchaseTime)
	if err != nil {
		l.Printf("- %s gave %d points - wrong format - %s", purchaseTime, points, err)
		return points
	}

	if parsedTime.Hour() >= 14 && parsedTime.Hour() < 16 {
		points += 10
	}

	l.Printf("- time '%s' gave %d points", purchaseTime, points)
	return points
}
