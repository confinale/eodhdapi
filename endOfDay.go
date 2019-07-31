package eodhdapi

import (
	"context"
	"encoding/csv"
	"github.com/gitu/eodhdapi/exchanges"
	"io"
	"strconv"
	"time"
)

// EODInfo is the price information for a single Asset
type EODInfo struct {
	Code              string    `json:"code,omitempty" bson:"code"`
	ExchangeShortName string    `json:"exchange_short_name,omitempty" bson:"exchange_short_name"`
	Date              time.Time `json:"date,omitempty" bson:"date"`
	Open              float64   `json:"open,omitempty" bson:"open"`
	High              float64   `json:"hig,omitemptyh" bson:"high"`
	Low               float64   `json:"low,omitempty" bson:"low"`
	Close             float64   `json:"close,omitempty" bson:"close"`
	AdjustedClose     float64   `json:"adjusted_close,omitempty" bson:"adjusted_close"`
	Volume            float64   `json:"volume,omitempty" bson:"volume"`
	Ticker            string    `json:"tickers,omitempty" bson:"ticker"`
}

const dateFormat = "2006-01-02"

// FetchEOD Fetches End of day for the exchange
func (d *EODhd) FetchEOD(ctx context.Context, info chan EODInfo, exchange *exchanges.Exchange, date time.Time) error {
	res, err := d.readPath("/eod-bulk-last-day/"+exchange.Code, urlParam{"fmt", "csv"}, urlParam{"date", date.Format(dateFormat)})

	if err != nil {
		return err
	}

	defer res.Body.Close()

	reader, err := newCsvReaderWithFirstLine(res.Body, "Code,Ex,Date,Open,High,Low,Close,Adjusted_close,Volume,Prev_close,Change,Change_%")

	if err != nil {
		return err
	}

	for {

		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil && !(err == csv.ErrFieldCount && len(line) > 8) {
			//log.Println(err)
			continue
		}

		date, err := time.Parse(dateFormat, line[2])
		if err != nil {
			continue
		}
		open, err := strconv.ParseFloat(line[3], 64)
		if err != nil {
			continue
		}
		high, err := strconv.ParseFloat(line[4], 64)
		if err != nil {
			continue
		}
		low, err := strconv.ParseFloat(line[5], 64)
		if err != nil {
			continue
		}
		origClose, err := strconv.ParseFloat(line[6], 64)
		if err != nil {
			continue
		}
		adjustedClose, err := strconv.ParseFloat(line[7], 64)
		if err != nil {
			continue
		}
		volume, err := strconv.ParseFloat(line[8], 64)
		if err != nil {
			continue
		}

		v := EODInfo{
			Code:              line[0],
			ExchangeShortName: line[1],
			Date:              date,
			Open:              open,
			High:              high,
			Low:               low,
			Close:             origClose,
			AdjustedClose:     adjustedClose,
			Volume:            volume,
			Ticker:            line[0] + "." + line[1],
		}

		info <- v
	}

	return nil
}

// FetchEODForTicker Fetches End of day for a single ticker
func (d *EODhd) FetchEODForTicker(ctx context.Context, info chan EODInfo, code, exchange string) error {
	ticker := code + "." + exchange

	res, err := d.readPath("/eod/"+ticker, urlParam{"fmt", "csv"}, urlParam{"period", "d"}, urlParam{"from", "2000-01-01"})

	if err != nil {
		return err
	}

	defer res.Body.Close()

	reader, err := newCsvReaderWithFirstLine(res.Body, "Date,Open,High,Low,Close,Adjusted_close,Volume")

	if err != nil {
		return err
	}

	for {

		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			continue
		}

		open, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			continue
		}
		high, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			continue
		}
		low, err := strconv.ParseFloat(line[3], 64)
		if err != nil {
			continue
		}
		origClose, err := strconv.ParseFloat(line[4], 64)
		if err != nil {
			continue
		}
		adjustedClose, err := strconv.ParseFloat(line[5], 64)
		if err != nil {
			continue
		}
		volume, err := strconv.ParseFloat(line[5], 64)
		if err != nil {
			continue
		}
		date, err := time.Parse(dateFormat, line[0])
		if err != nil {
			continue
		}

		v := EODInfo{
			Code:              code,
			ExchangeShortName: exchange,
			Date:              date,
			Open:              open,
			High:              high,
			Low:               low,
			Close:             origClose,
			AdjustedClose:     adjustedClose,
			Volume:            volume,
			Ticker:            ticker,
		}

		info <- v
	}

	return nil
}
