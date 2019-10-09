package eodhdapi

import (
	"context"
	"github.com/gitu/eodhdapi/exchanges"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

type EODDividend struct {
	Code     string          `json:"code,omitempty" bson:"code"`
	Ex       string          `json:"exchange_short_name,omitempty" bson:"exchange_short_name"`
	Ticker   string          `json:"tickers,omitempty" bson:"ticker"`
	Date     string          `json:"date,omitempty" bson:"date"`
	Dividend decimal.Decimal `json:"dividend,omitempty" bson:"dividend"`
}

// FetchDividendsTicker fetches the dividends of a single ticker
func (d *EODhd) FetchDividendsTicker(context context.Context, dividends chan EODDividend, ticker string, from time.Time) interface{} {

	urlParams := []urlParam{{"fmt", "csv"}, {"type", "dividends"}, {"filter", "extended"}, {"date", from.Format(dateFormat)}}

	res, err := d.readPath("/div/"+ticker, urlParams...)

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
		i, err := buildDividendSingle(reader, splitTicker[0], splitTicker[1])
		if err != nil {
			return err
		}

		dividends <- i
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
func (d *EODhd) FetchDividends(ctx context.Context, info chan EODDividend, exchange *exchanges.Exchange, date time.Time, symbols ...string) error {

	urlParams := []urlParam{{"fmt", "csv"}, {"type", "dividends"}, {"filter", "extended"}, {"date", date.Format(dateFormat)}}
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
		i, err := buildDividend(reader)
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

func buildDividend(r *csvReaderMap) (EODDividend, error) {
	g := EODDividend{}
	var err error

	if g.Code, err = r.asString("Code"); err != nil {
		return EODDividend{}, err
	}
	if g.Ex, err = r.asString("Ex"); err != nil {
		return EODDividend{}, err
	}
	if g.Date, err = r.asString("Date"); err != nil {
		return EODDividend{}, err
	}
	if g.Dividend, err = r.asDecimal("Dividend"); err != nil {
		return EODDividend{}, err
	}

	g.Ticker = g.Code + "." + g.Ex
	return g, nil
}

func buildDividendSingle(r *csvReaderMap, code, exchange string) (EODDividend, error) {
	g := EODDividend{}
	var err error

	if g.Date, err = r.asString("Date"); err != nil {
		return EODDividend{}, err
	}
	if g.Dividend, err = r.asDecimal("Dividends"); err != nil {
		return EODDividend{}, err
	}

	g.Code = code
	g.Ex = exchange
	g.Ticker = g.Code + "." + g.Ex
	return g, nil
}
