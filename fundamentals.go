package eodhdapi

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/gitu/eodhdapi/exchanges"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

// FetchSingleFundamental Fetches Fundamentals for the exchange
func (d *EODhd) FetchSingleFundamental(ctx context.Context, ticker string) ([]byte, error) {
	res, err := d.readPath("/fundamentals/" + ticker)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return bodyBytes, nil

}

// FetchFundamentals Fetches Fundamentals for the exchange
func (d *EODhd) FetchFundamentals(ctx context.Context, fundamentals chan Fundamentals, exchange *exchanges.Exchange) error {
	for _, e := range exchange.ExchangeCodeComponents {

		res, err := d.readPath("/bulk-fundamentals/" + e)

		if err != nil {
			return err
		}

		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Println(res.Body)
		}

		reader, valMap, values, err := newCsvReaderMap(res.Body)

		fmt.Printf(" available fundamentals for exchange [%s]: %v", exchange.Code, valMap)

		if err != nil {
			return err
		}
		for {

			line, err := reader.Read()

			if err == io.EOF {
				break
			} else if errors.Is(err, csv.ErrFieldCount) {
				// occurs constantly
				continue
			} else if err != nil {
				// should not occur
				return err
			}

			fullTimeEmployees, _ := strconv.ParseInt(line[valMap["General_FullTimeEmployees"]], 10, 32)

			data := make(map[string]string)

			for k, v := range values {
				data[v] = line[k]
			}

			ticker := line[valMap["General_Code"]] + "." + exchange.Code
			f := Fundamentals{
				Ticker:     ticker,
				LastUpdate: time.Now(),
				General: General{
					Code:              line[valMap["General_Code"]],
					Type:              line[valMap["General_Type"]],
					Name:              line[valMap["General_Name"]],
					Exchange:          exchange.Code,
					CurrencyCode:      line[valMap["General_CurrencyCode"]],
					CurrencyName:      line[valMap["General_CurrencyName"]],
					CurrencySymbol:    line[valMap["General_CurrencySymbol"]],
					CountryName:       line[valMap["General_CountryName"]],
					CountryISO:        line[valMap["General_CountryISO"]],
					ISIN:              line[valMap["General_ISIN"]],
					Sector:            line[valMap["General_Sector"]],
					Industry:          line[valMap["General_Industry"]],
					Description:       line[valMap["General_Description"]],
					FullTimeEmployees: int(fullTimeEmployees),
					UpdatedAt:         line[valMap["General_UpdatedAt"]],
				},
				Data: data,
			}

			fundamentals <- f

		}
	}

	return nil
}

// Fundamentals for a ticker
type Fundamentals struct {
	Ticker     string            `bson:"ticker"`
	LastUpdate time.Time         `bson:"last_update"`
	General    General           `bson:"general"`
	Data       map[string]string `bson:"data"`
}

// General information about an asset
type General struct {
	Code              string `bson:"code"`
	Type              string `bson:"type"`
	Name              string `bson:"name"`
	Exchange          string `bson:"exchange"`
	CurrencyCode      string `bson:"currency_code"`
	CurrencyName      string `bson:"currency_name"`
	CurrencySymbol    string `bson:"currency_symbol"`
	CountryName       string `bson:"country_name"`
	CountryISO        string `bson:"country_iso"`
	ISIN              string `bson:"isin"`
	CUSIP             string `bson:"cusip"`
	Sector            string `bson:"sector"`
	Industry          string `bson:"industry"`
	Description       string `bson:"description"`
	FullTimeEmployees int    `bson:"full_time_employees"`
	UpdatedAt         string `bson:"updated_at"`
}
