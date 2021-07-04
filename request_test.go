package facio

import (
	"reflect"
	"testing"
)

type testRequest struct {
	name string
	want *request
	got  *request
}

func getRequestWithHeaders() []testRequest {
	var tts []testRequest

	{
		// Expects "Authorization"
		headerName := Authorization
		headerNameStr := "Authorization"
		headerValue := "Bearer randomtoken"

		want := &request{
			headers: map[string]string{
				headerNameStr: headerValue,
			},
			headerMap: HeaderMap{
				headerName: {headerValue},
			},
			endpoint: "/profile",
			method:   "GET",
		}

		got, _ := NewRequest("GET", "/profile")

		got.AddHeader(headerName, headerValue)

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

		want := &request{
			headers: map[string]string{
				headerNameStr: headerValue,
			},
			headerMap: HeaderMap{
				headerName: headerValueArr,
			},
			endpoint: "/profile",
			method:   "GET",
		}

		got, _ := NewRequest("GET", "/profile")

		got.AddHeader(headerName, headerValueArr...)

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

		want := &request{
			headers:   headerValue,
			headerMap: headerValueArr,
			endpoint:  "/profile",
			method:    "GET",
		}

		got, _ := NewRequest("GET", "/profile")

		got.AddHeaders(headerValueArr)

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
