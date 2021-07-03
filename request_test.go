package facio

import (
	"reflect"
	"testing"
)

type testRequest struct {
	name string
	want Request
	got  Request
}

func getRequestWithHeaders() []testRequest {
	var tts []testRequest

	{
		// Expects "Authorization"
		headerName := Authorization
		headerNameStr := "Authorization"
		headerValue := "Bearer randomtoken"

		want := Request{
			Headers: map[string]string{
				headerNameStr: headerValue,
			},
			Endpoint: "/profile",
			Method:   "GET",
		}

		client :=
			NewClient("example.com")
		client.AddHeader(headerName, headerValue)

		got, _ := client.NewRequest("GET", "/profile")

		tt := testRequest{
			name: "Correct Header",
			want: want,
			got:  got,
		}
		tts = append(tts, tt)
	}

	{
		headerName := AcceptEncoding
		headerNameStr := "Accept-Encoding"
		headerValueArr := []string{"gzip", "deflate"}
		headerValue := "gzip,deflate"

		want := Request{
			Headers: map[string]string{
				headerNameStr: headerValue,
			},
			Endpoint: "/profile",
			Method:   "GET",
		}

		client :=
			NewClient("example.com")

		_ = client.
			AddHeader(headerName, headerValueArr...)

		got, _ := client.NewRequest("GET", "/profile")

		tt := testRequest{
			name: "Correct Header 2",
			want: want,
			got:  got,
		}
		tts = append(tts, tt)
	}

	{
		//headerName := "Accept-Encoding"
		headerValueArr := HeaderMap{
			Authorization:  {"Bearer Eyyyy.TOKEN_.AZZDD"},
			AcceptEncoding: {"gzip", "deflate"},
		}
		headerValue := map[string]string{
			"Authorization":   "Bearer Eyyyy.TOKEN_.AZZDD",
			"Accept-Encoding": "gzip,deflate",
		}

		want := Request{
			Headers:  headerValue,
			Endpoint: "/profile",
			Method:   "GET",
		}

		client :=
			NewClient("example.com")

		client.
			AddHeaders(headerValueArr)

		got, _ := client.NewRequest("GET", "/profile")

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
