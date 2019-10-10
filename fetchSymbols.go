package eodhdapi

import (
	"context"
	"github.com/gitu/eodhdapi/exchanges"
)

// The Symbol does map a ticker to an isin
type Symbol struct {
	Code     string `json:"code" bson:"code"`
	Type     string `json:"type" bson:"type"`
	Exchange string `json:"exchange" bson:"exchange"`
	Country  string `json:"country" bson:"country"`
	Name     string `json:"name" bson:"name"`
	Currency string `json:"name" bson:"name"`

	Ticker string `json:"ticker" bson:"ticker"`
}

// FetchSymbols Loads the Symbols for an exchange
func (d *EODhd) GetSymbols(ctx context.Context, exchange *exchanges.Exchange) ([]Symbol, error) {

	mapping := make(chan Symbol)
	done := make(chan []Symbol, 1)
	go func(f chan Symbol, d chan []Symbol) {
		mappings := make([]Symbol, 0)
		for fu := range f {
			mappings = append(mappings, fu)
		}
		d <- mappings
	}(mapping, done)
	err := d.FetchSymbols(ctx, mapping, exchange)
	if err != nil {
		return nil, err
	}
	close(mapping)
	mappings := <-done

	return mappings, nil
}

// FetchSymbols Loads the Symbols for an exchange
func (d *EODhd) FetchSymbols(ctx context.Context, info chan Symbol, exchange *exchanges.Exchange) error {

	urlParams := []urlParam{{"fmt", "csv"}}

	res, err := d.readPath("/exchanges/"+exchange.Code, urlParams...)

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
		i, err := buildSymbol(reader)
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

func buildSymbol(reader *csvReaderMap) (Symbol, error) {
	var err error
	g := Symbol{}
	if g.Code, err = reader.asString("Code"); err != nil {
		return g, err
	}
	if g.Name, err = reader.asString("Name"); err != nil {
		return g, err
	}
	if g.Country, err = reader.asString("Country"); err != nil {
		return g, err
	}
	if g.Exchange, err = reader.asString("Exchange"); err != nil {
		return g, err
	}
	if g.Currency, err = reader.asString("Currency"); err != nil {
		return g, err
	}
	if g.Type, err = reader.asString("Type"); err != nil {
		return g, err
	}
	g.Ticker = g.Code + "." + g.Exchange

	return g, err
}
