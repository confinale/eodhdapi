package eodhdapi

import (
	"golang.org/x/text/encoding/charmap"
	"io"
	"log"
	"net/http"
)

// The EODMapping does map a ticker to an isin
type EODMapping struct {
	Code     string `json:"code" bson:"code"`
	Exchange string `json:"exchange" bson:"exchange"`
	Country  string `json:"country" bson:"country"`
	Isin     string `json:"isin" bson:"isin"`
	Name     string `json:"name" bson:"name"`

	Ticker string `json:"ticker" bson:"ticker"`
}

func (d *EODhd) GetEtfs() ([]EODMapping, error) {
	mapping := make(chan EODMapping)
	done := make(chan []EODMapping, 1)
	go func(f chan EODMapping, d chan []EODMapping) {
		mappings := make([]EODMapping, 0)
		for fu := range f {
			mappings = append(mappings, fu)
		}
		d <- mappings
	}(mapping, done)
	err := d.LoadEtfs(mapping)
	if err != nil {
		return nil, err
	}
	close(mapping)
	mappings := <-done

	return mappings, nil
}

// LoadEtfs Loads the ETFS into a EODMappings
func (d *EODhd) LoadEtfs(info chan EODMapping) error {
	//ETF Code	Exchange	Country	ISIN	ETF Name

	etfListRequest, err := http.DefaultClient.Get("https://eodhistoricaldata.com/download/List_Of_Supported_ETFs.csv")
	if err != nil {
		return err
	}
	defer etfListRequest.Body.Close()
	reader, err := newCsvReader(charmap.ISO8859_1.NewDecoder().Reader(etfListRequest.Body))

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

		v := EODMapping{
			Code:     line[0],
			Exchange: line[1],
			Country:  line[2],
			Isin:     line[3],
			Name:     line[4],
		}

		v.Ticker = v.Code + "." + v.Exchange

		info <- v
	}

	return nil
}
