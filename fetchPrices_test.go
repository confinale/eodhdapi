package eodhdapi

import (
	"context"
	"fmt"
	freshcache "github.com/confinale/eodhdapi/util/afr"
	"github.com/confinale/eodhdapi/util/afr/diskcache"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/confinale/eodhdapi/exchanges"
	"github.com/stretchr/testify/require"
)

func TestEODhd_FetchEOD(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/api/eod-bulk-last-day/F" {
			rw.WriteHeader(404)
			return
		}
		date := req.URL.Query().Get("date")
		symbols := req.URL.Query().Get("symbols")
		format := req.URL.Query().Get("fmt")

		filename := fmt.Sprintf("test-data/eod-bulk-last-day/F_date_%s_%s.%s", date, symbols, format)
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
		name            string
		fields          fields
		args            args
		wantErr         bool
		wantPricesCount int
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
				date:     time.Date(2019, 9, 25, 0, 0, 0, 0, time.UTC),
			},
			wantErr:         false,
			wantPricesCount: 20,
		},
		{
			name: "F - two tickers",
			fields: fields{
				token:   "TOKEN",
				baseURL: server.URL + "/api",
				clt:     server.Client(),
			},
			args: args{
				ctx:      context.Background(),
				exchange: exchanges.All().GetByCode("F"),
				date:     time.Date(2019, 9, 24, 0, 0, 0, 0, time.UTC),
				symbols:  []string{"CON", "BAYN"},
			},
			wantErr:         false,
			wantPricesCount: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &EODhd{
				token:   tt.fields.token,
				baseURL: tt.fields.baseURL,
				clt:     tt.fields.clt,
			}

			prices := make(chan EODPrice)
			done := make(chan int, 1)

			go func(f chan EODPrice, d chan int) {
				count := 0
				for range f {
					count++
				}
				d <- count
			}(prices, done)
			if err := d.FetchPrices(tt.args.ctx, prices, tt.args.exchange, tt.args.date, tt.args.symbols...); (err != nil) != tt.wantErr {
				t.Errorf("FetchPrices() error = %v, wantErr %v", err, tt.wantErr)
			}
			close(prices)

			count := <-done

			require.Equal(t, tt.wantPricesCount, count)
		})
	}
}

func TestEODhd_FetchEOD_Ticker(t *testing.T) {
	if os.Getenv("EODHD_TOKEN") == "" {
		t.Skipf("no env variable EODHD_TOKEN set, will skip this test")
		t.SkipNow()
	}

	c := diskcache.New("cache/TestEODhd_FetchEOD_Ticker")
	tr := freshcache.NewTransport(c)

	d := NewEOD(DefaultProxyURL, os.Getenv("EODHD_TOKEN"), tr)

	tickers := []string{"AAPL.US", "CL.COMM", "APM.HM", "US084664CR08.BOND", "SYN.BE", "CBCYB.US", "PSFYX.US"}

	for _, ti := range tickers {

		t.Run(ti, func(t *testing.T) {

			tis := strings.Split(ti, ".")

			prices := make(chan EODPrice)
			done := make(chan int, 1)

			go func(f chan EODPrice, d chan int) {
				count := 0
				for range f {
					count++
				}
				d <- count
			}(prices, done)

			if err := d.FetchTickerPrices(context.Background(), prices, tis[0], tis[1],
				time.Date(2000, 1, 25, 0, 0, 0, 0, time.UTC),
				time.Date(2019, 9, 25, 0, 0, 0, 0, time.UTC),
				"d"); err != nil {
				t.Errorf("FetchTickerPrices() error = %v", err)
			}
			close(prices)

			count := <-done

			t.Logf("ticker %s had %d prices", ti, count)
		})
	}
}

func TestEODhd_FetchEOD_TestAll(t *testing.T) {
	if os.Getenv("EODHD_TOKEN") == "" {
		t.Skipf("no env variable EODHD_TOKEN set, will skip this test")
		t.SkipNow()
	}

	c := diskcache.New("cache/TestEODhd_FetchEOD_TestAll")
	tr := freshcache.NewTransport(c)

	d := NewEOD(DefaultProxyURL, os.Getenv("EODHD_TOKEN"), tr)

	for _, e := range exchanges.All() {

		t.Run(e.Code, func(t *testing.T) {

			prices := make(chan EODPrice)
			done := make(chan int, 1)

			go func(f chan EODPrice, d chan int) {
				count := 0
				for range f {
					count++
				}
				d <- count
			}(prices, done)

			if err := d.FetchPrices(context.Background(), prices, e, time.Date(2019, 9, 25, 0, 0, 0, 0, time.UTC)); err != nil {
				t.Errorf("FetchPrices() error = %v", err)
			}
			close(prices)

			count := <-done

			t.Logf("exchange %s had %d elements", e.Code, count)
		})
	}
}
