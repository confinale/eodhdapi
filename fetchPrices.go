package eodhdapi

import (
	"context"
	"github.com/confinale/eodhdapi/exchanges"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

// EODPrice is the price information for a single Asset
type EODPrice struct {
	Code          string `json:"code,omitempty" bson:"code"`
	Ex            string `json:"exchange_short_name,omitempty" bson:"exchange_short_name"`
	Name          string
	Date          string           `json:"date,omitempty" bson:"date"`
	Open          *decimal.Decimal `json:"open,omitempty" bson:"open"`
	High          *decimal.Decimal `json:"high,omitempty" bson:"high"`
	Low           *decimal.Decimal `json:"low,omitempty" bson:"low"`
	Close         *decimal.Decimal `json:"close,omitempty" bson:"close"`
	AdjustedClose *decimal.Decimal `json:"adjusted_close,omitempty" bson:"adjusted_close"`
	Volume        decimal.Decimal  `json:"volume,omitempty" bson:"volume"`

	Price *decimal.Decimal `json:"price,omitempty" bson:"price"`
	Yield *decimal.Decimal `json:"yield,omitempty" bson:"yield"`

	MarketCapitalization *decimal.Decimal `json:"market_capitalization,omitempty" bson:"market_capitalization"`
	EMA_50               *decimal.Decimal `json:"ema_50,omitempty" bson:"ema_50"`
	EMA_200              *decimal.Decimal `json:"ema_200,omitempty" bson:"ema_200"`
	High_250             *decimal.Decimal `json:"high_250,omitempty" bson:"high_250"`
	Low_250              *decimal.Decimal `json:"low_250,omitempty" bson:"low_250"`
	Prev_close           *decimal.Decimal `json:"prev_close,omitempty" bson:"prev_close"`
	Change               *decimal.Decimal `json:"change,omitempty" bson:"change"`
	ChangePercent        *decimal.Decimal `json:"change_percent,omitempty" bson:"change_percent"`

	Avgvol50d  *decimal.Decimal `json:"avg_vol_50d,omitempty" bson:"avg_vol_50d"`
	Avgvol200d *decimal.Decimal `json:"avg_vol_200d,omitempty" bson:"avg_vol_200d"`
	Avgvol14d  *decimal.Decimal `json:"avg_vol_14d,omitempty" bson:"avg_vol_14d"`

	Ticker string `json:"ticker,omitempty" bson:"ticker"`
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
	reader.skipMissingFields = 4

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

// FetchTickerPrices Fetches End of day for the ticker only date part of time will be used - use ‘d’ for daily, ‘w’ for weekly and ‘m’ for monthly prices
func (d *EODhd) FetchTickerPrices(ctx context.Context, info chan EODPrice, symbol, exchange string, from, to time.Time, period string) error {

	urlParams := []urlParam{{"fmt", "csv"},
		//{"filter", "extended"},
		{"from", from.Format(dateFormat)},
		{"to", to.Format(dateFormat)},
		{"period", period},
	}

	res, err := d.readPath("/eod/"+symbol+"."+exchange, urlParams...)

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
		if s, err := reader.asString("Volume"); err == nil && s == "" {
			continue
		}
		i, err := buildPriceTicker(reader, symbol, exchange)
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

func buildPriceTicker(r *csvReaderMap, code, exchange string) (EODPrice, error) {
	g := EODPrice{
		Code: code,
		Ex:   exchange,
	}
	var err error

	if g.Date, err = r.asString("Date"); err != nil {
		return EODPrice{}, err
	}
	if g.Open, err = r.asOptionalDecimal("Open"); err != nil {
		return EODPrice{}, err
	}
	if g.High, err = r.asOptionalDecimal("High"); err != nil {
		return EODPrice{}, err
	}
	if g.Low, err = r.asOptionalDecimal("Low"); err != nil {
		return EODPrice{}, err
	}
	if g.Close, err = r.asOptionalDecimal("Close"); err != nil {
		return EODPrice{}, err
	}
	if g.AdjustedClose, err = r.asOptionalDecimal("Adjusted_close"); err != nil {
		return EODPrice{}, err
	}
	if g.Volume, err = r.asDecimal("Volume"); err != nil {
		return EODPrice{}, err
	}

	if g.Price, err = r.asOptionalDecimal("Price"); err != nil {
		return EODPrice{}, err
	}
	if g.Yield, err = r.asOptionalDecimal("Yield"); err != nil {
		return EODPrice{}, err
	}

	if g.MarketCapitalization, err = r.asOptionalDecimal("MarketCapitalization"); err != nil {
		return EODPrice{}, err
	}
	if g.EMA_50, err = r.asOptionalDecimal("EMA_50"); err != nil {
		return EODPrice{}, err
	}
	if g.EMA_200, err = r.asOptionalDecimal("EMA_200"); err != nil {
		return EODPrice{}, err
	}
	if g.High_250, err = r.asOptionalDecimal("High_250"); err != nil {
		return EODPrice{}, err
	}
	if g.Low_250, err = r.asOptionalDecimal("Low_250"); err != nil {
		return EODPrice{}, err
	}
	if g.Prev_close, err = r.asOptionalDecimal("Prev_close"); err != nil {
		return EODPrice{}, err
	}
	if g.Change, err = r.asOptionalDecimal("Change"); err != nil {
		return EODPrice{}, err
	}
	if g.ChangePercent, err = r.asOptionalDecimal("Change_%"); err != nil {
		return EODPrice{}, err
	}

	if g.Avgvol50d, err = r.asOptionalDecimal("Avgvol_50d"); err != nil {
		return EODPrice{}, err
	}
	if g.Avgvol200d, err = r.asOptionalDecimal("Avgvol_200d"); err != nil {
		return EODPrice{}, err
	}
	if g.Avgvol14d, err = r.asOptionalDecimal("Avgvol_14d"); err != nil {
		return EODPrice{}, err
	}

	g.Ticker = g.Code + "." + g.Ex
	return g, nil
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
	if g.Open, err = r.asOptionalDecimal("Open"); err != nil {
		return EODPrice{}, err
	}
	if g.High, err = r.asOptionalDecimal("High"); err != nil {
		return EODPrice{}, err
	}
	if g.Low, err = r.asOptionalDecimal("Low"); err != nil {
		return EODPrice{}, err
	}
	if g.Close, err = r.asOptionalDecimal("Close"); err != nil {
		return EODPrice{}, err
	}
	if g.AdjustedClose, err = r.asOptionalDecimal("Adjusted_close"); err != nil {
		return EODPrice{}, err
	}
	if g.Volume, err = r.asDecimal("Volume"); err != nil {
		return EODPrice{}, err
	}

	if g.Price, err = r.asOptionalDecimal("Price"); err != nil {
		return EODPrice{}, err
	}
	if g.Yield, err = r.asOptionalDecimal("Yield"); err != nil {
		return EODPrice{}, err
	}

	if g.MarketCapitalization, err = r.asOptionalDecimal("MarketCapitalization"); err != nil {
		return EODPrice{}, err
	}
	if g.EMA_50, err = r.asOptionalDecimal("EMA_50"); err != nil {
		return EODPrice{}, err
	}
	if g.EMA_200, err = r.asOptionalDecimal("EMA_200"); err != nil {
		return EODPrice{}, err
	}
	if g.High_250, err = r.asOptionalDecimal("High_250"); err != nil {
		return EODPrice{}, err
	}
	if g.Low_250, err = r.asOptionalDecimal("Low_250"); err != nil {
		return EODPrice{}, err
	}
	if g.Prev_close, err = r.asOptionalDecimal("Prev_close"); err != nil {
		return EODPrice{}, err
	}
	if g.Change, err = r.asOptionalDecimal("Change"); err != nil {
		return EODPrice{}, err
	}
	if g.ChangePercent, err = r.asOptionalDecimal("Change_%"); err != nil {
		return EODPrice{}, err
	}

	if g.Avgvol50d, err = r.asOptionalDecimal("Avgvol_50d"); err != nil {
		return EODPrice{}, err
	}
	if g.Avgvol200d, err = r.asOptionalDecimal("Avgvol_200d"); err != nil {
		return EODPrice{}, err
	}
	if g.Avgvol14d, err = r.asOptionalDecimal("Avgvol_14d"); err != nil {
		return EODPrice{}, err
	}

	g.Ticker = g.Code + "." + g.Ex
	return g, nil
}
