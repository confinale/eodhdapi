package eodhdapi

import (
	"context"
	"fmt"
	freshcache "github.com/gitu/eodhdapi/util/afr"
	"github.com/gitu/eodhdapi/util/afr/diskcache"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gitu/eodhdapi/exchanges"
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
		name             string
		fields           fields
		args             args
		wantErr          bool
		wantEodInfoCount int
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
			wantErr:          false,
			wantEodInfoCount: 20,
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
			wantErr:          false,
			wantEodInfoCount: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &EODhd{
				token:   tt.fields.token,
				baseURL: tt.fields.baseURL,
				clt:     tt.fields.clt,
			}

			infos := make(chan EODInfo)
			done := make(chan int, 1)

			go func(f chan EODInfo, d chan int) {
				count := 0
				for range f {
					count++
				}
				d <- count
			}(infos, done)
			if err := d.FetchEOD(tt.args.ctx, infos, tt.args.exchange, tt.args.date, tt.args.symbols...); (err != nil) != tt.wantErr {
				t.Errorf("FetchFundamentals() error = %v, wantErr %v", err, tt.wantErr)
			}
			close(infos)

			count := <-done

			require.Equal(t, tt.wantEodInfoCount, count)
		})
	}
}

func TestEODhd_FetchEOD_TestAll(t *testing.T) {
	if os.Getenv("EODHD_TOKEN") == "" {
		t.Skipf("no env variable EODHD_TOKEN set, will skip this test")
		t.SkipNow()
	}

	c := diskcache.New("cache")
	tr := freshcache.NewTransport(c)

	d := NewEOD(DefaultURL, os.Getenv("EODHD_TOKEN"), tr)

	for _, e := range exchanges.All() {

		t.Run(e.Code, func(t *testing.T) {

			infos := make(chan EODInfo)
			done := make(chan int, 1)

			go func(f chan EODInfo, d chan int) {
				count := 0
				for range f {
					count++
				}
				d <- count
			}(infos, done)

			if err := d.FetchEOD(context.Background(), infos, e, time.Date(2019, 9, 25, 0, 0, 0, 0, time.UTC)); err != nil {
				t.Errorf("FetchEOD() error = %v", err)
			}
			close(infos)

			count := <-done

			t.Logf("exchange %s had %d elements", e.Code, count)
		})
	}
}

func TestEODhd_FetchDividends_TestAll(t *testing.T) {
	if os.Getenv("EODHD_TOKEN") == "" {
		t.Skipf("no env variable EODHD_TOKEN set, will skip this test")
		t.SkipNow()
	}

	c := diskcache.New("cache")
	tr := freshcache.NewTransport(c)

	d := NewEOD(DefaultURL, os.Getenv("EODHD_TOKEN"), tr)

	for _, e := range exchanges.All() {

		t.Run(e.Code, func(t *testing.T) {

			infos := make(chan EODDividend)
			done := make(chan int, 1)

			go func(f chan EODDividend, d chan int) {
				count := 0
				for range f {
					count++
				}
				d <- count
			}(infos, done)

			if err := d.FetchDividends(context.Background(), infos, e, time.Date(2019, 9, 25, 0, 0, 0, 0, time.UTC)); err != nil {
				t.Errorf("FetchDividends() error = %v", err)
			}
			close(infos)

			count := <-done

			t.Logf("exchange %s had %d elements", e.Code, count)
		})
	}
}

func TestEODhd_FetchSplits_TestAll(t *testing.T) {
	if os.Getenv("EODHD_TOKEN") == "" {
		t.Skipf("no env variable EODHD_TOKEN set, will skip this test")
		t.SkipNow()
	}

	c := diskcache.New("cache")
	tr := freshcache.NewTransport(c)

	d := NewEOD(DefaultURL, os.Getenv("EODHD_TOKEN"), tr)

	for _, e := range exchanges.All() {

		t.Run(e.Code, func(t *testing.T) {

			infos := make(chan EODSplit)
			done := make(chan int, 1)

			go func(f chan EODSplit, d chan int) {
				count := 0
				for range f {
					count++
				}
				d <- count
			}(infos, done)

			if err := d.FetchSplits(context.Background(), infos, e, time.Date(2019, 9, 25, 0, 0, 0, 0, time.UTC)); err != nil {
				t.Errorf("FetchSplits() error = %v", err)
			}
			close(infos)

			count := <-done

			t.Logf("exchange %s had %d elements", e.Code, count)
		})
	}
}
