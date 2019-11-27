package eodhdapi

import (
	"bytes"
	"encoding/json"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func asDec(s string) *Decimal {
	v, err := decimal.NewFromString(s)
	if err != nil {
		log.Fatal(v)
	}
	i := Decimal(v)
	return &i
}

func TestJsonParsingDecimal(t *testing.T) {
	values := []struct {
		input    string
		expected *Decimal
	}{
		{"\"32.3\"", asDec("32.3")},
		{"\"32\"", asDec("32")},
		{"\"-\"", &Decimal{}},
		{"\"null\"", &Decimal{}},
		{"null", &Decimal{}},
		{"32", asDec("32")},
		{"32.3", asDec("32.3")},
		{"\"99,5\"", asDec("99.5")},
	}

	for _, v := range values {
		t.Run(v.input, func(t *testing.T) {
			var d Decimal
			if err := easyjson.Unmarshal([]byte(v.input), &d); err != nil {
				t.Fatalf("failed reading json: %s", err)
			}
			assert.Equal(t, v.expected, &d)
		})
	}
}

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
				t.Log(string(r.Data[Max(0, r.GetPos()-30):Min(len(r.Data), r.GetPos()+30)]))
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

			fu2 := Fundamentals{}
			r2 := jlexer.Lexer{
				Data: g,
			}
			fu2.UnmarshalEasyJSON(&r2)
			if r.Error() != nil {
				t.Log(r.Error())
				t.FailNow()
			}

			assert.Equal(t, fu, fu2)
		})
	}
}
