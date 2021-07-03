package facio

import (
	"reflect"
	"testing"
)

type testFacio struct {
	name string
	want Facio
	got  *Facio
}

func getTestFacio() []testFacio {
	var tts []testFacio

	{
		want := Facio{
			client: &Client{
				BaseURL: "example.com",
			},
		}
		got := NewDefaultFacio("example.com")

		name := "Default Facio"

		tts = append(tts, testFacio{
			name: name,
			want: want,
			got:  got,
		})
	}

	{
		want := Facio{
			client: &Client{
				BaseURL: "example.com",
			},
		}
		client := NewClient("example.com")

		name := "Default Facio"

		got := NewFacio(&client)

		tts = append(tts, testFacio{
			name: name,
			want: want,
			got:  got,
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
