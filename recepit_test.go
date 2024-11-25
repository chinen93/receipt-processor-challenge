package main

import (
	"fmt"
	"testing"
)

func TestPointsForRetailer(t *testing.T) {
	var tests = []struct {
		retailer string
		expected int
	}{
		{"Target", 6},
	}

	for _, tt := range tests {

		testname := tt.retailer
		t.Run(testname, func(t *testing.T) {
			ans := pointsForRetailer(tt.retailer)
			if ans != tt.expected {
				t.Errorf("got %d, want %d", ans, tt.expected)
			}
		})
	}
}

func TestPointsForTotal(t *testing.T) {
	var tests = []struct {
		total    string
		expected int
	}{
		{"1.25", 25},
		{"5.00", 50},
		{"10.30", 0},
	}

	for _, tt := range tests {

		testname := tt.total
		t.Run(testname, func(t *testing.T) {
			ans := pointsForTotal(tt.total)
			if ans != tt.expected {
				t.Errorf("got %d, want %d", ans, tt.expected)
			}
		})
	}
}

func TestPointsForItems(t *testing.T) {
	items0 := []item{}
	items1 := []item{
		{Description: "012", Price: "10.25"},
	}
	items2 := []item{
		{Description: "012", Price: "10.25"},
		{Description: "01234", Price: "15.40"},
	}
	items3 := []item{
		{Description: "012", Price: "10.25"},
		{Description: "01234", Price: "15.40"},
		{Description: "Emils Cheese Pizza", Price: "12.25"},
	}

	var tests = []struct {
		id       int
		items    []item
		expected int
	}{
		{0, items0, 0},
		{1, items1, 3},
		{2, items2, 8},
		{3, items3, 11},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.id)
		t.Run(testname, func(t *testing.T) {
			ans := pointsForItems(tt.items)
			if ans != tt.expected {
				t.Errorf("got %d, want %d", ans, tt.expected)
			}
		})
	}
}

func TestPointsForPurchaseDate(t *testing.T) {
	var tests = []struct {
		purchaseDate string
		expected     int
	}{
		{"2022-01-02", 0},
		{"2022-01-01", 6},
	}

	for _, tt := range tests {

		testname := tt.purchaseDate
		t.Run(testname, func(t *testing.T) {
			ans := pointsForPurchaseDate(tt.purchaseDate)
			if ans != tt.expected {
				t.Errorf("got %d, want %d", ans, tt.expected)
			}
		})
	}
}

func TestPointsForPurchaseTime(t *testing.T) {
	var tests = []struct {
		purchaseTime string
		expected     int
	}{
		{"13:59", 0},
		{"14:00", 10},
		{"15:59", 10},
		{"16:01", 0},
	}

	for _, tt := range tests {

		testname := tt.purchaseTime
		t.Run(testname, func(t *testing.T) {
			ans := pointsForPurchaseTime(tt.purchaseTime)
			if ans != tt.expected {
				t.Errorf("got %d, want %d", ans, tt.expected)
			}
		})
	}
}
