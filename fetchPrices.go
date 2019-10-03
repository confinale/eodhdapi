package eodhdapi

import (
	"context"
	"github.com/gitu/eodhdapi/exchanges"
	"strings"
	"time"
)

// EODPrice is the price information for a single Asset
type EODPrice struct {
	Code          string `json:"code,omitempty" bson:"code"`
	Ex            string `json:"exchange_short_name,omitempty" bson:"exchange_short_name"`
	Name          string
	Date          string  `json:"date,omitempty" bson:"date"`
	Open          float64 `json:"open,omitempty" bson:"open"`
	High          float64 `json:"high,omitempty" bson:"high"`
	Low           float64 `json:"low,omitempty" bson:"low"`
	Close         float64 `json:"close,omitempty" bson:"close"`
	AdjustedClose float64 `json:"adjusted_close,omitempty" bson:"adjusted_close"`
	Volume        float64 `json:"volume,omitempty" bson:"volume"`

	MarketCapitalization *float64
	EMA_50               *float64
	EMA_200              *float64
	High_250             *float64
	Low_250              *float64
	Prev_close           *float64
	Change               *float64
	ChangePercent        *float64

	Ticker string `json:"tickers,omitempty" bson:"ticker"`
}

const dateFormat = "2006-01-02"

// FetchPrices Fetches End of day for the exchange only date part of time will be used
func (d *EODhd) FetchPrices(ctx context.Context, info chan EODPrice, exchange *exchanges.Exchange, date time.Time, symbols ...string) error {

	urlParams := []urlParam{{"fmt", "csv"}, {"filter", "extended"}, {"date", date.Format(dateFormat)}}
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
	reader.skipMissingFields = 7

	for reader.Next() {
		i, err := buildPrice(reader)
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

func buildPrice(r *csvReaderMap) (EODPrice, error) {
	g := EODPrice{}
	var err error

	if g.Code, err = r.asString("Code"); err != nil {
		return EODPrice{}, err
	}
	if g.Ex, err = r.asString("Ex"); err != nil {
		return EODPrice{}, err
	}
	if g.Name, err = r.asString("Name"); err != nil {
		return EODPrice{}, err
	}
	if g.Date, err = r.asString("Date"); err != nil {
		return EODPrice{}, err
	}
	if g.Open, err = r.asFloat64("Open"); err != nil {
		return EODPrice{}, err
	}
	if g.High, err = r.asFloat64("High"); err != nil {
		return EODPrice{}, err
	}
	if g.Low, err = r.asFloat64("Low"); err != nil {
		return EODPrice{}, err
	}
	if g.Close, err = r.asFloat64("Close"); err != nil {
		return EODPrice{}, err
	}
	if g.AdjustedClose, err = r.asFloat64("Adjusted_close"); err != nil {
		return EODPrice{}, err
	}
	if g.Volume, err = r.asFloat64("Volume"); err != nil {
		return EODPrice{}, err
	}

	if g.MarketCapitalization, err = r.asOptionalFloat64("MarketCapitalization"); err != nil {
		return EODPrice{}, err
	}
	if g.EMA_50, err = r.asOptionalFloat64("EMA_50"); err != nil {
		return EODPrice{}, err
	}
	if g.EMA_200, err = r.asOptionalFloat64("EMA_200"); err != nil {
		return EODPrice{}, err
	}
	if g.High_250, err = r.asOptionalFloat64("High_250"); err != nil {
		return EODPrice{}, err
	}
	if g.Low_250, err = r.asOptionalFloat64("Low_250"); err != nil {
		return EODPrice{}, err
	}
	if g.Prev_close, err = r.asOptionalFloat64("Prev_close"); err != nil {
		return EODPrice{}, err
	}
	if g.Change, err = r.asOptionalFloat64("Change"); err != nil {
		return EODPrice{}, err
	}
	if g.ChangePercent, err = r.asOptionalFloat64("Change_%"); err != nil {
		return EODPrice{}, err
	}

	g.Ticker = g.Code + "." + g.Ex
	return g, nil
}
