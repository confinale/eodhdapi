package eodhdapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	freshcache "github.com/gitu/eodhdapi/util/afr"
	"github.com/gitu/eodhdapi/util/afr/diskcache"
	"github.com/mailru/easyjson/jlexer"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gitu/eodhdapi/exchanges"
	"github.com/stretchr/testify/require"
)

func TestJsonParsing(t *testing.T) {
	root := "test-data/fundamentals"
	files, err := ioutil.ReadDir(root)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	for _, f := range files {
		t.Run(f.Name(), func(t *testing.T) {
			if f.IsDir() {
				t.Log("is dir", f.Name())
				t.Skipped()
				return
			}

			data, err := ioutil.ReadFile(root + "/" + f.Name())
			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			fu := Fundamentals{}
			r := jlexer.Lexer{
				Data: data,
			}
			fu.UnmarshalEasyJSON(&r)

			if r.Error() != nil {
				t.Log(r.Error())
				t.FailNow()
			}

			b, err := json.MarshalIndent(fu, "", "  ")
			if err != nil {
				t.Fatalf("failed writing json: %s", err)
			}

			gp := filepath.Join("test-data/fundamentals_golden", f.Name())

			if _, err := os.Stat(gp); os.IsNotExist(err) {
				t.Log("create golden file")
				if err := ioutil.WriteFile(gp, b, 0644); err != nil {
					t.Fatalf("failed to create golden file: %s", err)
				}
			}
			g, err := ioutil.ReadFile(gp)
			if err != nil {
				t.Fatalf("failed reading .golden: %s", err)
			}
			t.Log(string(b))
			if !bytes.Equal(b, g) {
				t.Errorf("writtein json does not match .golden file")
			}

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

	c := diskcache.New("cache")
	tr := freshcache.NewTransport(c)

	d := NewEOD(DefaultURL, os.Getenv("EODHD_TOKEN"), tr)

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
