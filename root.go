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

const DefaultURL = "https://eodhistoricaldata.com/api"

// NewEOD initializes a new eod historical data client
func NewDefaultEOD(eodHdToken string) *EODhd {
	return &EODhd{
		token:   eodHdToken,
		baseURL: DefaultURL,
		clt:     http.DefaultClient,
	}
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

type ErrUnknownStatus struct {
	message string
}

func (e *ErrUnknownStatus) Error() string {
	return e.message
}

type ErrNotFound struct {
	message string
}

func (e *ErrNotFound) Error() string {
	return e.message
}

type ErrTooManyRequests struct {
	message string
}

func (e *ErrTooManyRequests) Error() string {
	return e.message
}

func StatusError(response *http.Response) error {
	switch response.StatusCode {
	case 429:
		return &ErrTooManyRequests{message: fmt.Sprintf("not expected status %d - %s", response.StatusCode, response.Status)}
	case 404:
		return &ErrNotFound{message: fmt.Sprintf("not expected status %d - %s", response.StatusCode, response.Status)}
	default:
		return &ErrUnknownStatus{message: fmt.Sprintf("not expected status %d - %s", response.StatusCode, response.Status)}
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

	response, err := d.clt.Do(req)
	if err == nil && response.StatusCode != 200 {
		return response, StatusError(response)
	}
	return response, err
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
