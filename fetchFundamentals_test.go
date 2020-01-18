package eodhdapi

import (
	"context"
	"fmt"
	"github.com/confinale/eodhdapi/exchanges"
	freshcache "github.com/confinale/eodhdapi/util/afr"
	"github.com/confinale/eodhdapi/util/afr/diskcache"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

func TestEODhd_FetchFundamentalsTicker(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if path.Dir(req.URL.Path) != "/api/fundamentals" {
			rw.WriteHeader(404)
			return
		}
		name := path.Base(req.URL.Path)

		filename := fmt.Sprintf("test-data/fundamentals/%s.json", name)
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			t.Logf("file does not exist: %s", filename)
			rw.WriteHeader(404)
			return
		}

		b, err := ioutil.ReadFile(filename)
		require.NoError(t, err)
		_, err = rw.Write(b)
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
		symbols  []string
	}
	tests := []struct {
		name                  string
		fields                fields
		args                  args
		wantErr               bool
		wantFundamentalsCount int
	}{
		{
			name: "AAPL/ABB",
			fields: fields{
				token:   "TOKEN",
				baseURL: server.URL + "/api",
				clt:     server.Client(),
			},
			args: args{
				ctx:      context.Background(),
				exchange: exchanges.All().GetByCode("US"),
				symbols:  []string{"AAPL", "ABB"},
			},
			wantErr:               false,
			wantFundamentalsCount: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &EODhd{
				token:   tt.fields.token,
				baseURL: tt.fields.baseURL,
				clt:     tt.fields.clt,
			}

			fundamentals := make(chan Fundamentals)
			done := make(chan int, 1)

			go func(f chan Fundamentals, d chan int) {
				count := 0
				for range f {
					count++
				}
				d <- count
			}(fundamentals, done)
			if err := d.FetchFundamentalsTicker(tt.args.ctx, fundamentals, tt.args.exchange.Code, tt.args.symbols...); (err != nil) != tt.wantErr {
				t.Errorf("FetchFundamentals() error = %v, wantErr %v", err, tt.wantErr)
			}
			close(fundamentals)

			count := <-done

			require.Equal(t, tt.wantFundamentalsCount, count)
		})
	}
}

func TestEODhd_FetchFundamentals(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/api/bulk-fundamentals/F" {
			rw.WriteHeader(404)
			return
		}
		limit := req.URL.Query().Get("limit")
		offset := req.URL.Query().Get("offset")
		format := req.URL.Query().Get("fmt")

		filename := fmt.Sprintf("test-data/bulk-fundamentals/F_limit_%s_offset_%s.%s", limit, offset, format)
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			t.Logf("file does not exist: %s", filename)
			rw.WriteHeader(404)
			return
		}

		b, err := ioutil.ReadFile(filename)
		require.NoError(t, err)
		_, err = rw.Write(b)
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
		pagesize int
	}
	tests := []struct {
		name                  string
		fields                fields
		args                  args
		wantErr               bool
		wantFundamentalsCount int
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
				pagesize: 10,
			},
			wantErr:               false,
			wantFundamentalsCount: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &EODhd{
				token:   tt.fields.token,
				baseURL: tt.fields.baseURL,
				clt:     tt.fields.clt,
			}

			fundamentals := make(chan Fundamentals)
			done := make(chan int, 1)

			go func(f chan Fundamentals, d chan int) {
				count := 0
				for range f {
					count++
				}
				d <- count
			}(fundamentals, done)
			if err := d.FetchFundamentals(tt.args.ctx, fundamentals, tt.args.exchange, tt.args.pagesize, false); (err != nil) != tt.wantErr {
				t.Errorf("FetchFundamentals() error = %v, wantErr %v", err, tt.wantErr)
			}
			close(fundamentals)

			count := <-done

			require.Equal(t, tt.wantFundamentalsCount, count)
		})
	}
}

func TestEODhd_FetchFundamentals_TestAll(t *testing.T) {
	if os.Getenv("EODHD_TOKEN") == "" {
		t.Skipf("no env variable EODHD_TOKEN set, will skip this test")
		t.SkipNow()
	}

	c := diskcache.New("cache/TestEODhd_FetchFundamentals_TestAll")
	tr := freshcache.NewTransport(c)

	d := NewEOD(DefaultProxyURL, os.Getenv("EODHD_TOKEN"), tr)

	for _, e := range exchanges.All() {

		t.Run(e.Code, func(t *testing.T) {

			fundamentals := make(chan Fundamentals)
			done := make(chan int, 1)

			go func(f chan Fundamentals, d chan int) {
				count := 0
				for range f {
					count++
				}
				d <- count
			}(fundamentals, done)

			if err := d.FetchFundamentals(context.Background(), fundamentals, e, 1000, false); err != nil {
				t.Errorf("FetchFundamentals() error = %v", err)
			}
			close(fundamentals)

			count := <-done

			t.Logf("exchange %s had %d elements", e.Code, count)
		})
	}
}

func TestEODhd_FetchFundamentalsTicker_TestETFS(t *testing.T) {
	if os.Getenv("EODHD_TOKEN") == "" {
		t.Skipf("no env variable EODHD_TOKEN set, will skip this test")
		t.SkipNow()
	}

	c := diskcache.New("cache/TestEODhd_FetchFundamentalsTicker_TestETFS")
	tr := freshcache.NewTransport(c)

	d := NewEOD(DefaultProxyURL, os.Getenv("EODHD_TOKEN"), tr)

	r := rand.New(rand.NewSource(33))

	mappings, err := d.GetEtfs()
	if err != nil {
		t.Error(err)
	}

	for i := 1; i <= 30; i++ {
		intn := r.Intn(len(mappings))
		m := mappings[intn]

		t.Run(m.Ticker, func(t *testing.T) {
			if _, err := d.FetchFundamentalsSymbol(context.Background(), m.Exchange, m.Code); err != nil {
				t.Errorf("FetchFundamentals() error = %v", err)
			}
		})
	}
}

func TestEODhd_FetchFundamentalsSymbol_TestAll(t *testing.T) {
	if os.Getenv("EODHD_TOKEN") == "" {
		t.Skipf("no env variable EODHD_TOKEN set, will skip this test")
		t.SkipNow()
	}

	c := diskcache.New("cache/TestEODhd_FetchFundamentalsSymbol_TestAll")
	tr := freshcache.NewTransport(c)

	d := NewEOD(DefaultProxyURL, os.Getenv("EODHD_TOKEN"), tr)

	for _, e := range exchanges.All() {

		t.Run(e.Code, func(t *testing.T) {
			r := rand.New(rand.NewSource(42))
			symbols, err := d.GetSymbols(context.Background(), e)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}

			for i := 1; i <= Min(20, len(symbols)/2); i++ {
				intn := r.Intn(len(symbols))
				s := symbols[intn]

				t.Run(s.Ticker, func(t *testing.T) {
					if _, err := d.FetchFundamentalsSymbol(context.Background(), e.Code, s.Code); err != nil {
						t.Errorf("FetchFundamentals() error = %v", err)
					}
				})
			}

		})
	}
}
