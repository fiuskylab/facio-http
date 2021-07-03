package facio_test

import (
	"reflect"
	"testing"

	"github.com/fiuskylab/facio-http"
)

type testHeader struct {
	name string
	want interface{}
	got  interface{}
}

func getTestHeaders() []testHeader {
	var tts []testHeader

	{
		token := "random_token"
		encs := []string{"application/json", "text/plain"}
		encsStr := "application/json,text/plain"

		hm := facio.HeaderMap{
			facio.Authorization:  {token},
			facio.AcceptEncoding: encs,
		}

		want := facio.HeaderResult{
			"Authorization":   token,
			"Accept-Encoding": encsStr,
		}

		client := facio.
			NewClient("example.com")
		client.AddHeaders(hm)

		req, _ := client.GetRequest()

		got := req.Headers

		tt := testHeader{
			name: "Correct Header",
			want: want,
			got:  got,
		}
		tts = append(tts, tt)
	}

	return tts
}

func TestHeaderBuild(t *testing.T) {
	tts := getTestHeaders()

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.want, tt.got) {
				t.Errorf("Want: %+v\n Got: %+v", tt.want, tt.got)
			}
		})
	}
}
