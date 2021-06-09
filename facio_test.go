package facio_test

import (
	"reflect"
	"testing"

	facio "github.com/fiuskylab/facio-http"
)

type testClient struct {
	name string
	want facio.Client
	got  facio.Client
}

func getNewClients() []testClient {
	return []testClient{
		{
			name: "Correct Client",
			want: facio.Client{
				BaseURL: "example.com",
				Request: facio.Request{
					Headers: make(map[string]string),
				},
			},
			got: facio.NewClient("example.com"),
		},
	}
}

func TestNewClient(t *testing.T) {
	tts := getNewClients()

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.want, tt.got) {
				t.Errorf("Want: %+v\n Got: %+v", tt.want, tt.got)
			}
		})
	}
}

type testRequest struct {
	name string
	want facio.Request
	got  facio.Request
}

func getRequestWithHeaders() []testRequest {
	var tts []testRequest

	{
		headerName := "Authorization"
		headerValue := "Bearer randomtoken"

		want := facio.Request{
			Headers: map[string]string{
				headerName: headerValue,
			},
		}

		got, _ := facio.
			NewClient("example.com").
			AddHeader(headerName, "Bearer randomtoken")

		tt := testRequest{
			name: "Correct Header",
			want: want,
			got:  got,
		}
		tts = append(tts, tt)
	}

	{
		headerName := "Accept-Encoding"
		headerValueArr := []string{"gzip", "deflate"}
		headerValue := "gzip, deflate"

		want := facio.Request{
			Headers: map[string]string{
				headerName: headerValue,
			},
		}

		got, _ := facio.
			NewClient("example.com").
			AddHeader(headerName, headerValueArr...)

		tt := testRequest{
			name: "Correct Header 2",
			want: want,
			got:  got,
		}
		tts = append(tts, tt)
	}

	{
		//headerName := "Accept-Encoding"
		headerValueArr := facio.HeaderMap{
			"Authorization":   {"Bearer Eyyyy.TOKEN_.AZZDD"},
			"Accept-Encoding": {"gzip", "deflate"},
		}
		headerValue := map[string]string{
			"Authorization":   "Bearer Eyyyy.TOKEN_.AZZDD",
			"Accept-Encoding": "gzip, deflate",
		}

		want := facio.Request{
			Headers: headerValue,
		}

		got, _ := facio.
			NewClient("example.com").
			AddHeaders(headerValueArr)

		tt := testRequest{
			name: "Correct Header 2",
			want: want,
			got:  got,
		}
		tts = append(tts, tt)
	}

	return tts
}

func TestNewRequest(t *testing.T) {
	tts := getRequestWithHeaders()

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.want, tt.got) {
				t.Errorf("Want: %+v\n Got: %+v", tt.want, tt.got)
			}
		})
	}
}
