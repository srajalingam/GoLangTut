package prices

import (
	"fmt"
	"pricecalculator/conversion"
	"pricecalculator/iomanager"
)

type TaxIncludedPrice struct {
	IOManager         iomanager.IOManager `json:"-"` //ignore
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPrice) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrices = result
	job.IOManager.WriteJson(job)
}

func (job *TaxIncludedPrice) LoadData() {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println("Error converting strings to floats:", err)
		return
	}
	job.InputPrices = prices
}

func NewTaxIncludedPrice(iom iomanager.IOManager, taxRate float64) *TaxIncludedPrice {
	return &TaxIncludedPrice{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
