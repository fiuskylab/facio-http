package facio

import (
	"reflect"
	"testing"
)

type testUtils struct {
	name string
	want interface{}
	got  interface{}
}

func getUtilsTests() []testUtils {
	var tts []testUtils

	{
		name := "checkMethod Correct Method - Returned Method"
		want := "POST"

		got, _ := checkMethod("pOsT")

		tts = append(tts, testUtils{
			name: name,
			want: want,
			got:  got,
		})
	}

	{
		name := "checkMethod Correct Method - Returned Error"
		want := newNilError()

		_, got := checkMethod("pOsT")

		tts = append(tts, testUtils{
			name: name,
			want: want,
			got:  got,
		})
	}

	{

		name := "checkMethod Incorrect Method - Returned Error"
		method := "random_method"
		want := newError(msgInvalidMethod, method)

		_, got := checkMethod(method)

		tts = append(tts, testUtils{
			name: name,
			want: want,
			got:  got,
		})
	}

	return tts
}

func TestUtils(t *testing.T) {
	tts := getUtilsTests()

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.want, tt.got) {
				t.Errorf("Want: %+v\n Got: %+v", tt.want, tt.got)
			}
		})
	}
}
