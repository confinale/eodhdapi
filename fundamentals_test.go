package eodhdapi

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gitu/eodhdapi/exchanges"
	"github.com/stretchr/testify/require"
)

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

			fundamentals := make(chan Fundamentals, 20000)
			if err := d.FetchFundamentals(tt.args.ctx, fundamentals, tt.args.exchange, tt.args.pagesize, false); (err != nil) != tt.wantErr {
				t.Errorf("FetchFundamentals() error = %v, wantErr %v", err, tt.wantErr)
			}
			close(fundamentals)

			count := 0

			for range fundamentals {
				count++
			}

			require.Equal(t, tt.wantFundamentalsCount, count)
		})
	}
}
