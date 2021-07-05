package facio

import (
	"net/http"
	"reflect"
	"testing"
)

type testFacio struct {
	name string
	want interface{}
	got  interface{}
}

func getTestFacio() []testFacio {
	var tts []testFacio

	{
		name := "Default Facio"

		cfg := &config{
			baseURL: "example.com",
		}

		facio := &facio{
			config:     cfg,
			httpClient: &http.Client{},
		}

		want := facio
		got := NewFacio(cfg)

		tts = append(tts, testFacio{
			name: name,
			want: want,
			got:  got,
		})
	}
	{
		// Using the public API: http://api.dataatwork.org/v1
		name := "Get /jobs"

		facio := NewFacio(NewConfig("http://api.dataatwork.org/v1"))
		res := facio.Get("jobs")

		tts = append(tts, testFacio{
			name: name,
			want: false,
			got:  res.Error.isNil,
		})
	}
	{
		// Using the public API: http://api.dataatwork.org/v1
		name := "Get expect status 200"

		facio := NewFacio(NewConfig("http://api.dataatwork.org/v1"))
		res := facio.Get("jobs")

		tts = append(tts, testFacio{
			name: name,
			want: http.StatusOK,
			got:  res.StatusCode(),
		})
	}
	return tts
}

func TestFacio(t *testing.T) {
	tts := getTestFacio()

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.want, tt.got) {
				t.Errorf("Want: %+v\n Got: %+v", tt.want, tt.got)
			}
		})
	}
}
