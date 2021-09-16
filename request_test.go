package facio

import (
	"fmt"
	"testing"
)

func TestRequest(t *testing.T) {
	tts := []testCases{}

	{
		gotR, gotErr := NewRequest("GET", "https://example.com/")

		tts = append(tts, testCases{
			name:     "NewRequest - Correct",
			got:      gotR.BaseURL,
			want:     "https://example.com",
			testType: Equal,
		})

		tts = append(tts, testCases{
			name:     "NewRequest - Correct",
			got:      gotErr,
			testType: Nil,
		})
	}
	{
		method := "random_method"
		gotR, gotErr := NewRequest(method, "https://example.com/")

		tts = append(tts, testCases{
			name:     "NewRequest - Correct",
			got:      gotR,
			want:     &Req{},
			testType: Equal,
		})

		tts = append(tts, testCases{
			name:     "NewRequest - Correct",
			got:      gotErr,
			want:     fmt.Errorf(msgInvalidMethod, method),
			testType: Equal,
		})
	}
	{
		method := "GET"
		gotR, gotErr := NewRequest(method, "")

		tts = append(tts, testCases{
			name:     "NewRequest - Correct",
			got:      gotR,
			want:     &Req{},
			testType: Equal,
		})

		tts = append(tts, testCases{
			name:     "NewRequest - Correct",
			got:      gotErr,
			want:     fmt.Errorf(msgInvalidURL, ""),
			testType: Equal,
		})
	}

	{
		gotR, _ := NewRequest("GET", "https://example.com/")
		gotR.AddHeader(AcceptEncoding, "application/json")

		tts = append(tts, testCases{
			name:     "NewRequest - Correct",
			got:      gotR.Headers[AcceptEncoding.getHeader()],
			want:     "application/json",
			testType: Equal,
		})
	}

	{
		gotR, _ := NewRequest("GET", "https://example.com/")
		headerName := "bar"
		headerValue := "foo"
		gotR.AddCustomHeader(headerName, headerValue)

		tts = append(tts, testCases{
			name:     "NewRequest - Correct",
			got:      gotR.Headers[headerName],
			want:     headerValue,
			testType: Equal,
		})
	}

	runTests(t, tts)
}
