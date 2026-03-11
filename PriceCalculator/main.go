package main

import (
	"fmt"
	"pricecalculator/prices"
)

func main() {
	taxRates := []float64{0.05, 0.10, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPrice(taxRate)
		fmt.Println(priceJob)
		priceJob.Process()
	}

}
