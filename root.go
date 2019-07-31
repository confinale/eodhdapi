package eodhdapi

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// EODhd is an EOD Historical Data Client Info
type EODhd struct {
	token   string
	baseURL string
	clt     *http.Client
}

type urlParam struct {
	Key   string
	Value string
}

// NewEOD initializes a new eod historical data client
func NewEOD(eodHdURL, eodHdToken string, transport http.RoundTripper) *EODhd {
	client := http.Client{Transport: transport}
	return &EODhd{
		token:   eodHdToken,
		baseURL: eodHdURL,
		clt:     &client,
	}
}

func (d *EODhd) readPath(path string, params ...urlParam) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", d.baseURL, path), nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("api_token", d.token)

	for _, v := range params {
		q.Add(v.Key, v.Value)
	}

	req.URL.RawQuery = q.Encode()

	return d.clt.Do(req)
}

func newCsvReaderWithFirstLine(r io.Reader, expectedFirstLine string) (*csv.Reader, error) {
	reader := csv.NewReader(r)
	reader.Comma = ','
	reader.ReuseRecord = true
	// skip first line
	first, err := reader.Read()
	if err != nil {
		return nil, err
	}
	joinedFirst := strings.Join(first, ",")
	if joinedFirst != expectedFirstLine {
		log.Printf("First Line was vs expected:\n was:      [%s]\nexpected: [%s]\n", joinedFirst, expectedFirstLine)
		return nil, errors.New("failed first line check")
	}
	return reader, nil
}

func newCsvReader(r io.Reader) (*csv.Reader, error) {
	reader := csv.NewReader(r)
	reader.Comma = ','
	reader.ReuseRecord = true
	// skip first line
	_, err := reader.Read()
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func newCsvReaderMap(r io.Reader) (*csv.Reader, map[string]int, error) {
	colMap := make(map[string]int)
	reader := csv.NewReader(r)
	reader.Comma = ','
	reader.ReuseRecord = true
	// skip first line
	firstLine, err := reader.Read()
	if err != nil {
		return nil, nil, err
	}
	for k, v := range firstLine {
		colMap[v] = k
	}
	return reader, colMap, nil
}
