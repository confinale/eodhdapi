package eodhdapi

import (
	"bytes"
	"encoding/json"
	"github.com/mailru/easyjson/jlexer"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
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
				t.Log(string(r.Data[Max(0, r.GetPos()-20):Min(len(r.Data), r.GetPos()+20)]))
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
