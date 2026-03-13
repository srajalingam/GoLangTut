package main

import (
	"fmt"
	"pricecalculator/filemanager"
	"pricecalculator/prices"
)

func main() {
	taxRates := []float64{0.05, 0.10, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPrice(fm, taxRate)
		fmt.Println(priceJob)
		priceJob.Process()
	}

}
