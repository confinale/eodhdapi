package eodhdapi

import (
	"context"
	"fmt"
	"github.com/confinale/eodhdapi/exchanges"
	freshcache "github.com/confinale/eodhdapi/util/afr"
	"github.com/confinale/eodhdapi/util/afr/diskcache"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestEODhd_FetchDividendsTicker(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/api/div/AAAPL.US" {
			rw.WriteHeader(404)
			return
		}
		date := req.URL.Query().Get("date")
		format := req.URL.Query().Get("fmt")

		filename := fmt.Sprintf("test-data/div/AAPL.US_from_%s.%s", date, format)
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			t.Logf("file does not exist: %s", filename)
			rw.WriteHeader(404)
			return
		}

		bytes, err := ioutil.ReadFile(filename)
		require.NoError(t, err)
		_, err = rw.Write(bytes)
		require.NoError(t, err)
	}))

	type fields struct {
		token   string
		baseURL string
		clt     *http.Client
	}
	type args struct {
		ctx    context.Context
		ticker string
		from   time.Time
	}
	tests := []struct {
		name                 string
		fields               fields
		args                 args
		wantErr              bool
		wantEodDividendCount int
	}{
		{
			name: "AAAPL.US",
			fields: fields{
				token:   "TOKEN",
				baseURL: server.URL + "/api",
				clt:     server.Client(),
			},
			args: args{
				ctx:    context.Background(),
				ticker: "AAAPL.US",
				from:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr:              false,
			wantEodDividendCount: 29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &EODhd{
				token:   tt.fields.token,
				baseURL: tt.fields.baseURL,
				clt:     tt.fields.clt,
			}

			dividends := make(chan EODDividend)
			done := make(chan int, 1)

			go func(f chan EODDividend, d chan int) {
				count := 0
				for range f {
					count++
				}
				d <- count
			}(dividends, done)
			if err := d.FetchDividendsTicker(tt.args.ctx, dividends, tt.args.ticker, tt.args.from); (err != nil) != tt.wantErr {
				t.Errorf("FetchFundamentals() error = %v, wantErr %v", err, tt.wantErr)
			}
			close(dividends)

			count := <-done

			require.Equal(t, tt.wantEodDividendCount, count)
		})
	}
}

func TestEODhd_FetchDividends(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/api/eod-bulk-last-day/F" {
			rw.WriteHeader(404)
			return
		}
		date := req.URL.Query().Get("date")
		symbols := req.URL.Query().Get("symbols")
		format := req.URL.Query().Get("fmt")

		filename := fmt.Sprintf("test-data/eod-bulk-last-day/F_dividends_date_%s_%s.%s", date, symbols, format)
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			t.Logf("file does not exist: %s", filename)
			rw.WriteHeader(404)
			return
		}

		bytes, err := ioutil.ReadFile(filename)
		require.NoError(t, err)
		_, err = rw.Write(bytes)
		require.NoError(t, err)
	}))

	type fields struct {
		token   string
		baseURL string
		clt     *http.Client
	}
	type args struct {
		ctx      context.Context
		exchange *exchanges.Exchange
		date     time.Time
		symbols  []string
	}
	tests := []struct {
		name                 string
		fields               fields
		args                 args
		wantErr              bool
		wantEodDividendCount int
	}{
		{
			name: "F",
			fields: fields{
				token:   "TOKEN",
				baseURL: server.URL + "/api",
				clt:     server.Client(),
			},
			args: args{
				ctx:      context.Background(),
				exchange: exchanges.All().GetByCode("F"),
				date:     time.Date(2019, 10, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr:              false,
			wantEodDividendCount: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &EODhd{
				token:   tt.fields.token,
				baseURL: tt.fields.baseURL,
				clt:     tt.fields.clt,
			}

			dividends := make(chan EODDividend)
			done := make(chan int, 1)

			go func(f chan EODDividend, d chan int) {
				count := 0
				for range f {
					count++
				}
				d <- count
			}(dividends, done)
			if err := d.FetchDividends(tt.args.ctx, dividends, tt.args.exchange, tt.args.date, tt.args.symbols...); (err != nil) != tt.wantErr {
				t.Errorf("FetchFundamentals() error = %v, wantErr %v", err, tt.wantErr)
			}
			close(dividends)

			count := <-done

			require.Equal(t, tt.wantEodDividendCount, count)
		})
	}
}

func TestEODhd_FetchDividends_TestAll(t *testing.T) {
	if os.Getenv("EODHD_TOKEN") == "" {
		t.Skipf("no env variable EODHD_TOKEN set, will skip this test")
		t.SkipNow()
	}

	c := diskcache.New("cache/TestEODhd_FetchDividends_TestAll")
	tr := freshcache.NewTransport(c)

	d := NewEOD(DefaultProxyURL, os.Getenv("EODHD_TOKEN"), tr)

	for _, e := range exchanges.All() {

		t.Run(e.Code, func(t *testing.T) {

			dividends := make(chan EODDividend)
			done := make(chan int, 1)

			go func(f chan EODDividend, d chan int) {
				count := 0
				for range f {
					count++
				}
				d <- count
			}(dividends, done)

			if err := d.FetchDividends(context.Background(), dividends, e, time.Date(2019, 9, 25, 0, 0, 0, 0, time.UTC)); err != nil {
				t.Errorf("FetchDividends() error = %v", err)
			}
			close(dividends)

			count := <-done

			t.Logf("exchange %s had %d elements", e.Code, count)
		})
	}
}
