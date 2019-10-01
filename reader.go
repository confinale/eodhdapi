package eodhdapi

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
)

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
	// only skips missing fields if at least X fields found (some generators remove duplicate delimiters at the end
	skipMissingFields int
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
			return "", fmt.Errorf("field: %s not found", value)
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
			if len(r.current) <= r.skipMissingFields {
				// occurs constantly
				continue
			}
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
