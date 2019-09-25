package eodhdapi

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/texttheater/golang-levenshtein/levenshtein"
	"io"
	"log"
	"net/http"
	"strconv"
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

func newCsvReaderMap(r io.Reader, lenient, trackVisits bool) (*csvReaderMap, error) {
	colMap := make(map[string]int)
	reader := csv.NewReader(r)
	reader.Comma = ','
	reader.ReuseRecord = true
	// skip first line
	firstLine, err := reader.Read()
	if err != nil {
		return nil, err
	}
	for k, v := range firstLine {
		colMap[v] = k
	}
	return &csvReaderMap{
		reader:      reader,
		fields:      colMap,
		visits:      make(map[string]bool),
		trackVisits: trackVisits,
		lenient:     lenient,
	}, nil
}

type csvReaderMap struct {
	reader      *csv.Reader
	fields      map[string]int
	trackVisits bool
	visits      map[string]bool
	current     []string
	lenient     bool
}

func (r *csvReaderMap) asOptionalStringLenient(value string) (*string, error) {
	val, err := r.asString(value)
	if err != nil {
		return nil, nil
	}
	if len(val) == 0 {
		return nil, nil
	}
	return &val, nil
}

func (r *csvReaderMap) asOptionalString(value string) (*string, error) {
	val, err := r.asString(value)
	if err != nil {
		return nil, err
	}
	if len(val) == 0 {
		return nil, nil
	}
	return &val, nil
}
func (r *csvReaderMap) asString(value string) (string, error) {
	if r.trackVisits {
		r.markVisited(value)
	}
	i, ok := r.fields[value]
	if !ok {
		if !r.lenient {
			closestMatch := "N/A"
			minDistance := 99999
			for v := range r.fields {
				distance := levenshtein.DistanceForStrings([]rune(v), []rune(value), levenshtein.DefaultOptions)
				if distance < minDistance {
					minDistance = distance
					closestMatch = v
				}
			}

			return "", fmt.Errorf("field: %s not found - closest match: %s", value, closestMatch)
		}
		return "", nil
	}

	return r.current[i], nil
}

func (r *csvReaderMap) asFloat64(value string) (float64, error) {
	s, err := r.asString(value)
	if err != nil {
		return 0, err
	}
	if r.lenient && len(s) == 0 {
		return 0, nil
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return f, fmt.Errorf("error while parsing %s: %v", value, err)
	}
	return f, err
}

func (r *csvReaderMap) asOptionalFloat64(value string) (*float64, error) {
	s, err := r.asOptionalString(value)
	if err != nil || s == nil {
		return nil, err
	}
	if r.lenient && len(*s) == 0 {
		return nil, nil
	}
	val, err := strconv.ParseFloat(*s, 64)
	if err != nil {
		return nil, fmt.Errorf("error while parsing %s: %v", value, err)
	}
	return &val, err
}

func (r *csvReaderMap) asOptionalInt(value string) (*int, error) {
	s, err := r.asOptionalString(value)
	if err != nil || s == nil {
		return nil, err
	}
	if r.lenient && len(*s) == 0 {
		return nil, nil
	}
	i, err := strconv.Atoi(*s)
	if err != nil {
		return nil, fmt.Errorf("error while parsing %s: %v", value, err)
	}
	return &i, err
}

func (r *csvReaderMap) asInt(value string) (int, error) {
	s, err := r.asString(value)
	if err != nil {
		return 0, err
	}
	if r.lenient && len(s) == 0 {
		return 0, nil
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return i, fmt.Errorf("error while parsing %s: %v", value, err)
	}
	return i, err
}

func (r *csvReaderMap) Next() bool {
	for {
		var err error
		r.current, err = r.reader.Read()

		if err == io.EOF {
			return false
		} else if errors.Is(err, csv.ErrFieldCount) {
			// occurs constantly
			continue
		} else if err != nil {
			// should not occur
			return false
		}

		return true
	}
}

func (r *csvReaderMap) markVisited(s string) {
	r.visits[s] = true
}

func (r *csvReaderMap) checkAllVisited(ignored ...string) error {
	notVisited := make([]string, 0)
CORE:
	for k := range r.fields {
		if !r.visits[k] {
			for _, i := range ignored {
				if i == k {
					continue CORE
				}
			}
			if len(notVisited) > 10 {
				notVisited = append(notVisited, "...")
				break
			}
			notVisited = append(notVisited, k)
		}
	}

	if len(notVisited) > 0 {
		return fmt.Errorf("fields not visited: %+v", notVisited)
	}
	return nil
}
