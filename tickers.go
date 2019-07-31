package eodhdapi

import (
	"context"
	"github.com/gitu/eodhdapi/exchanges"
	"io"
	"log"
)

//-- deactiv go:generate easytags $GOFILE json,bson

// EODTicker Ticker Info
type EODTicker struct {
	Code     string `json:"code" bson:"code"`
	Name     string `json:"name" bson:"name"`
	Country  string `json:"country" bson:"country"`
	Exchange string `json:"exchange" bson:"exchange"`
	Currency string `json:"currency" bson:"currency"`
	Type     string `json:"type" bson:"type"`

	Ticker            string   `json:"ticker" bson:"ticker"`
	ExchangeShortName string   `json:"exchange_short_name" bson:"exchange_short_name"`
	OpenfigiTickers   []string `json:"openfigi_tickers" bson:"openfigi_tickers"`
}

// FetchTickers Fetches End of day for the exchange
func (d *EODhd) FetchTickers(ctx context.Context, info chan EODTicker, exchange *exchanges.Exchange) error {

	res, err := d.readPath("/exchanges/"+exchange.Code, urlParam{"fmt", "csv"})

	if err != nil {
		return err
	}

	defer res.Body.Close()

	reader, err := newCsvReaderWithFirstLine(res.Body, "Code,Name,Country,Exchange,Currency,Type")

	if err != nil {
		return err
	}

	for {

		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
			continue
		}

		v := EODTicker{
			Code:     line[0],
			Name:     line[1],
			Country:  line[2],
			Exchange: line[3],
			Currency: line[4],
			Type:     line[5],
		}

		v.ExchangeShortName = exchange.Code
		v.Ticker = v.Code + "." + v.ExchangeShortName

		for _, figi := range exchange.FigiExchangeCodes {
			v.OpenfigiTickers = append(v.OpenfigiTickers, v.Code+"."+figi)
		}
		info <- v
	}

	return nil
}
