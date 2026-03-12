package prices

import (
	"bufio"
	"fmt"
	"os"
	"pricecalculator/conversion"
)

type TaxIncludedPrice struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices []float64
}

func (job *TaxIncludedPrice) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Println(result)
}

func (job *TaxIncludedPrice) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		file.Close()
		return
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println("Error converting strings to floats:", err)
		file.Close()
		return
	}
	job.InputPrices = prices
	file.Close()
}

func NewTaxIncludedPrice(taxRate float64) *TaxIncludedPrice {
	return &TaxIncludedPrice{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
