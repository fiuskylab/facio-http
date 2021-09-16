package facio

import (
	"fmt"
	"testing"
)

func TestUtils(t *testing.T) {
	var tts []testCases

	{
		name := "checkMethod Correct Method - Returned Method"
		want := "POST"

		got, _ := checkMethod("pOsT")

		tts = append(tts, testCases{
			name:     name,
			want:     want,
			got:      got,
			testType: Equal,
		})
	}

	{
		name := "checkMethod Correct Method - Returned Error"

		_, got := checkMethod("pOsT")

		tts = append(tts, testCases{
			name:     name,
			got:      got,
			testType: Nil,
		})
	}

	{
		name := "checkMethod Incorrect Method - Returned Error"
		method := "random_method"
		want := fmt.Errorf(msgInvalidMethod, method)

		_, got := checkMethod(method)

		tts = append(tts, testCases{
			name:     name,
			want:     want,
			got:      got,
			testType: Equal,
		})
	}

	{
		name := "parseEndpoint"

		tts = append(tts, testCases{
			name:     name,
			want:     "/endpoint",
			got:      parseEndpoint("/endpoint"),
			testType: Equal,
		})
		tts = append(tts, testCases{
			name:     name,
			want:     "/endpoint",
			got:      parseEndpoint("endpoint"),
			testType: Equal,
		})
	}

	runTests(t, tts)
}
