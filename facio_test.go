package facio

import (
	"reflect"
	"testing"
)

type testFacio struct {
	name string
	want *Facio
	got  *Facio
}

func getTestFacio() []testFacio {
	var tts []testFacio

	{
		name := "Default Facio"

		facio := Facio{
			client: &client{
				baseURL: "example.com",
			},
		}
		want := &facio
		got, _ := NewDefaultFacio("example.com")

		tts = append(tts, testFacio{
			name: name,
			want: want,
			got:  got,
		})
	}

	{
		name := "New Facio"

		client, _ := NewClient("example.com")

		facio := Facio{
			client: client,
		}

		want := &facio

		got := NewFacio(client)

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
