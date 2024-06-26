package pricecalculator

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

// first basic example
type priceAlias []float64
type taxRatesAlias []float64
type resultAlias map[float64][]float64

func CalculatePrice(prices priceAlias, taxRates taxRatesAlias) {
	result := make(resultAlias) // Initialize the result map
	for _, taxRate := range taxRates {
		taxIncludedPrices := make([]float64, len(prices))
		for indexPrice, price := range prices {
			taxIncludedPrices[indexPrice] = price * (1 + taxRate)
		}
		result[taxRate] = taxIncludedPrices
	}
	fmt.Println(result)
}

// *********************************************
// second intermediate example
type TaxIncludedPriceJob struct {
	TaxRate           float64            `json:"tax_rate"` //for omit use `json:"-"`
	InputPrices       []float64          `json:"input_prices"`
	TaxIncludedPrices map[string]float64 `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) Calculate() {
	err := job.LoadData()
	if err != nil {
		log.Fatal(err)
	}

	result := make(map[string]float64)
	for _, price := range job.InputPrices {

		taxIncludedPrice := price * (1 + job.TaxRate)
		floatPrice, err := strconv.ParseFloat(fmt.Sprintf("%.2f", taxIncludedPrice), 64)
		if err != nil {
			log.Fatal(err)
		}

		result[fmt.Sprintf("%.2f", price)] = floatPrice
	}
	job.TaxIncludedPrices = result
	nameFile := fmt.Sprintf("/home/jhonny/go/src/github/rigel-developer/golang-course/projects/price-calculator/prices_%v.json", job.TaxRate)

	saveDatainFile(nameFile, job)

}

func (job *TaxIncludedPriceJob) LoadData() error {
	prices, err := readPricesFromFile("/home/jhonny/go/src/github/rigel-developer/golang-course/projects/price-calculator/prices.txt")
	if err != nil {
		log.Fatal(err)
		return err
	}
	job.InputPrices = prices
	return nil

}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{},
		TaxRate:     taxRate,
	}
}

func CalculatePriceV2() {
	taxRates := []float64{0.1, 0.2, 0.3}
	//read prices from a file prices.txt
	job := &TaxIncludedPriceJob{}
	for _, taxRate := range taxRates {
		job = NewTaxIncludedPriceJob(taxRate)
		job.Calculate()

	}
	fmt.Println(job.TaxIncludedPrices)

}

func readPricesFromFile(filename string) ([]float64, error) {
	// read prices from the file
	//read file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var prices []float64
	for scanner.Scan() {
		price, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		prices = append(prices, price)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return prices, nil
}

func saveDatainFile(filename string, data interface{}) error {
	// write data to the file
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	encoder, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = file.Write(encoder)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil

}
