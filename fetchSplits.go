package eodhdapi

import (
	"context"
	"github.com/gitu/eodhdapi/exchanges"
	"strings"
	"time"
)

type EODSplit struct {
	Code   string `json:"code,omitempty" bson:"code"`
	Ex     string `json:"exchange_short_name,omitempty" bson:"exchange_short_name"`
	Ticker string `json:"tickers,omitempty" bson:"ticker"`
	Date   string `json:"date,omitempty" bson:"date"`
	Split  string `json:"split,omitempty" bson:"split"`
}

// FetchPrices Fetches End of day for the exchange only date part of time will be used
func (d *EODhd) FetchSplits(ctx context.Context, info chan EODSplit, exchange *exchanges.Exchange, date time.Time, symbols ...string) error {

	urlParams := []urlParam{{"fmt", "csv"}, {"type", "splits"}, {"filter", "extended"}, {"date", date.Format(dateFormat)}}
	if len(symbols) > 0 {
		urlParams = append(urlParams, urlParam{"symbols", strings.Join(symbols, ",")})
	}

	res, err := d.readPath("/eod-bulk-last-day/"+exchange.Code, urlParams...)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	reader, err := newCsvReaderMap(res.Body, false, true)
	if err != nil {
		return err
	}
	reader.skipMissingFields = 4

	for reader.Next() {
		i, err := buildSplit(reader)
		if err != nil {
			return err
		}

		info <- i
		if reader.trackVisits {
			// skip tracking after first visit
			reader.trackVisits = false
		}
	}
	if !reader.trackVisits {
		err = reader.checkAllVisited()
		if err != nil {
			return err
		}
	}

	return nil
}

func buildSplit(r *csvReaderMap) (EODSplit, error) {
	g := EODSplit{}
	var err error

	if g.Code, err = r.asString("Code"); err != nil {
		return EODSplit{}, err
	}
	if g.Ex, err = r.asString("Ex"); err != nil {
		return EODSplit{}, err
	}
	if g.Date, err = r.asString("Date"); err != nil {
		return EODSplit{}, err
	}
	if g.Split, err = r.asString("Split"); err != nil {
		return EODSplit{}, err
	}

	g.Ticker = g.Code + "." + g.Ex
	return g, nil
}
