package main

import (
	"errors"
	"fmt"
	"math"
)

func avg(x float64, y float64) float64 {
	return (x + y) / 2
}

// func getStockPriceChange(prevPrice, currentPrice float64) (float64, float64) {
// 	change := currentPrice - prevPrice
// 	percentChange := (change / prevPrice) * 100
// 	return change, percentChange
// }

func getStockPriceChangeWithError(prevPrice, currentPrice float64) (float64, float64, error) {
	if prevPrice == 0 {
		err := errors.New("Previous price cannot be zero")
		return 0, 0, err
	}
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange, nil
}

func getNamedStockPriceChange(prevPrice, currentPrice float64) (change, percentChange float64) {
	change = currentPrice - prevPrice
	percentChange = (change / prevPrice) * 100
	return change, percentChange
}

func getStockPriceChange(prevPrice, currentPrice float64) (float64, float64) {
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange
}

func main() {

	///====================1
	// x := 5.75
	// y := 6.25

	// result := avg(x, y)

	// fmt.Printf("Average of %.2f and %.2f = %.2f\n", x, y, result)

	///====================2
	// prevStockPrice := 0.0
	// currentStockPrice := 100000.0

	// change, percentChange := getStockPriceChange(prevStockPrice, currentStockPrice)

	// if change < 0 {
	// 	fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(change), math.Abs(percentChange))
	// } else {
	// 	fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", change, percentChange)
	// }

	///====================3
	// prevStockPrice := 0.0
	// currentStockPrice := 100000.0

	// change, percentChange, err := getStockPriceChangeWithError(prevStockPrice, currentStockPrice)

	// if err != nil {
	// 	fmt.Println("Sorry! There was an error: ", err)
	// } else {
	// 	if change < 0 {
	// 		fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(change), math.Abs(percentChange))
	// 	} else {
	// 		fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", change, percentChange)
	// 	}
	// }
	///===================== 4
	// prevStockPrice := 100000.0
	// currentStockPrice := 90000.0

	// change, percentChange := getNamedStockPriceChange(prevStockPrice, currentStockPrice)

	// if change < 0 {
	// 	fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(change), math.Abs(percentChange))
	// } else {
	// 	fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", change, percentChange)
	// }

	///========================= 5
	prevStockPrice := 80000.0
	currentStockPrice := 120000.0

	change, _ := getStockPriceChange(prevStockPrice, currentStockPrice)

	if change < 0 {
		fmt.Printf("The Stock Price decreased by $%.2f\n", math.Abs(change))
	} else {
		fmt.Printf("The Stock Price increased by $%.2f\n", change)
	}
}