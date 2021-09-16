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
			got:      gotR.baseURL,
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
			got:      gotR.headers[AcceptEncoding.getHeader()],
			want:     "application/json",
			testType: Equal,
		})
	}
	{
		gotR, _ := NewRequest("GET", "https://example.com/")
		gotR.AddCustomHeader(AcceptEncoding.getHeader(), "application/json")

		tts = append(tts, testCases{
			name:     "NewRequest - Correct",
			got:      gotR.headers[AcceptEncoding.getHeader()],
			want:     "application/json",
			testType: Equal,
		})
	}

	{
		gotR, _ := NewRequest("GET", "https://example.com/")
		gotR.SetURLForm(map[string][]string{
			"val": {"foo", "bar"},
		})

		tts = append(tts, testCases{
			name:     "NewRequest - Correct",
			got:      gotR.formURL["val"],
			want:     []string{"foo", "bar"},
			testType: Equal,
		})
	}

	runTests(t, tts)
}
