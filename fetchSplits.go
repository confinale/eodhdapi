package eodhdapi

import (
	"context"
	"errors"
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

// FetchSplitsTicker fetches the splits of a single ticker
func (d *EODhd) FetchSplitsTicker(context context.Context, splits chan EODSplit, ticker string, from time.Time) error {

	urlParams := []urlParam{{"fmt", "csv"}, {"filter", "extended"}, {"date", from.Format(dateFormat)}}

	res, err := d.readPath("/splits/"+ticker, urlParams...)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	reader, err := newCsvReaderMap(res.Body, false, true)
	if err != nil {
		return err
	}
	reader.skipMissingFields = 2

	splitTicker := strings.Split(ticker, ".")

	if len(splitTicker) != 2 {
		return errors.New("expected ticker to be in format CODE.EX")
	}

	for reader.Next() {
		i, err := buildSplitSingle(reader, splitTicker[0], splitTicker[1])
		if err != nil {
			return err
		}

		splits <- i
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

func buildSplitSingle(r *csvReaderMap, code, exchange string) (EODSplit, error) {
	g := EODSplit{}
	var err error

	if g.Date, err = r.asString("Date"); err != nil {
		return EODSplit{}, err
	}
	if g.Split, err = r.asString("Stock Splits"); err != nil {
		return EODSplit{}, err
	}

	g.Code = code
	g.Ex = exchange
	g.Ticker = g.Code + "." + g.Ex
	return g, nil
}
