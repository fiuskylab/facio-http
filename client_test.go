package facio

import (
	"reflect"
	"testing"
)

type testClient struct {
	name string
	want *client
	got  *client
}

func getNewClients() []testClient {
	var tts []testClient
	{
		name := "Correct Client"
		want := &client{
			baseURL: "example.com",
		}
		got, _ := NewClient("example.com")
		tts = append(tts, testClient{
			name: name,
			want: want,
			got:  got,
		})
	}
	{
		name := "Trimmed URL"
		want := &client{
			baseURL: "example.com",
		}
		got, _ := NewClient("example.com/")

		tts = append(tts, testClient{
			name: name,
			want: want,
			got:  got,
		})
	}

	return tts
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
