package eodhdapi

import (
	"context"
	"fmt"
	"github.com/confinale/eodhdapi/exchanges"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

// FetchFundamentals Fetches Fundamentals for the exchange
func (d *EODhd) FetchFundamentals(ctx context.Context, fundamentals chan Fundamentals, exchange *exchanges.Exchange, pagesize int, lenient bool) error {

	if exchange.ForceLenient {
		lenient = true
	}
	for _, e := range exchange.ExchangeCodeComponents {

		offset := 0

		newElements := pagesize
		for newElements == pagesize {
			newElements = 0
			res, err := d.readPath("/bulk-fundamentals/"+e,
				urlParam{"fmt", "csv"},
				urlParam{"offset", strconv.Itoa(offset)},
				urlParam{"limit", strconv.Itoa(pagesize)})

			if err != nil {
				return err
			}

			defer res.Body.Close()
			if res.StatusCode != 200 {
				log.Printf("body for url: %s - code %d: %v\n", strings.ReplaceAll(res.Request.URL.String(), d.token, "******"), res.StatusCode, res.Body)
				return fmt.Errorf("received non 200 error code: %d", res.StatusCode)
			}

			reader, err := newCsvReaderMap(res.Body, lenient, !lenient)
			if err != nil {
				return err
			}
			for reader.Next() {
				f, err := buildFundamental(reader, exchange.Code)
				if err != nil {
					if !lenient {
						return errors.Wrap(err, fmt.Sprintf("while parsing line: %.50s", strings.Join(reader.current, ",")))
					}
					log.Println(err, strings.Join(reader.current, ","))
					continue
				}

				fundamentals <- f

				if reader.trackVisits {
					// skip tracking after first visit
					reader.trackVisits = false
				}

				newElements++
			}
			offset += newElements
		}
	}

	return nil
}

// FetchFundamentalsTicker gets multiple symbols (currently a wrapper for single fetches - does multiple network calls
func (d *EODhd) FetchFundamentalsTicker(ctx context.Context, fundamentals chan Fundamentals, exchange string, symbol ...string) error {
	for _, s := range symbol {
		fu, err := d.FetchFundamentalsSymbol(ctx, exchange, s)
		if err != nil {
			return err
		}
		fundamentals <- fu
	}
	return nil
}

// FetchFundamentalsSymbol Fetches Fundamentals for a single symbol
func (d *EODhd) FetchFundamentalsSymbol(ctx context.Context, exchange, symbol string) (Fundamentals, error) {

	fu := Fundamentals{}
	urlParams := []urlParam{}

	path := "/fundamentals/" + symbol + "." + exchange
	if exchange == "BOND" {
		path = "/bond-fundamentals/" + symbol
	}
	res, err := d.readPath(path, urlParams...)
	if err != nil {
		return fu, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fu, err
	}

	err = fu.UnmarshalJSON(body)
	if err != nil {
		return fu, err
	}
	return fu, nil
}

func buildFundamental(reader *csvReaderMap, exchange string) (Fundamentals, error) {
	var err error
	f := Fundamentals{
		LastUpdate: time.Now(),
		General:    General{},
	}
	err = f.General.fill(reader, "General_")
	if err != nil {
		return Fundamentals{}, err
	}
	f.Ticker = f.General.Code + "." + exchange
	return f, err
}

func (g *General) fill(reader *csvReaderMap, prefix string) error {
	var err error
	if g.Code, err = reader.asString(prefix + "Code"); err != nil {
		return err
	}
	if g.Type, err = reader.asString(prefix + "Type"); err != nil {
		return err
	}
	if g.Name, err = reader.asString(prefix + "Name"); err != nil {
		return err
	}
	if g.Exchange, err = reader.asString(prefix + "Exchange"); err != nil {
		return err
	}
	if g.CurrencyCode, err = reader.asString(prefix + "CurrencyCode"); err != nil {
		return err
	}
	if g.CurrencyName, err = reader.asString(prefix + "CurrencyName"); err != nil {
		return err
	}
	if g.CurrencySymbol, err = reader.asString(prefix + "CurrencySymbol"); err != nil {
		return err
	}
	if g.CountryName, err = reader.asString(prefix + "CountryName"); err != nil {
		return err
	}
	if g.CountryISO, err = reader.asString(prefix + "CountryISO"); err != nil {
		return err
	}
	if g.ISIN, err = reader.asOptionalString(prefix + "ISIN"); err != nil {
		return err
	}
	if g.Sector, err = reader.asString(prefix + "Sector"); err != nil {
		return err
	}
	if g.Industry, err = reader.asString(prefix + "Industry"); err != nil {
		return err
	}
	if g.Description, err = reader.asString(prefix + "Description"); err != nil {
		return err
	}
	if g.FullTimeEmployees, err = reader.asOptionalInt(prefix + "FullTimeEmployees"); err != nil {
		return err
	}
	if g.UpdatedAt, err = reader.asOptionalStringLenient(prefix + "UpdatedAt"); err != nil {
		return err
	}

	if g.Cusip, err = reader.asOptionalStringLenient(prefix + "CUSIP"); err != nil {
		return err
	}
	return nil
}
