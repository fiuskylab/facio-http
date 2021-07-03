package facio

import (
	"reflect"
	"testing"
)

type testClient struct {
	name string
	want Client
	got  Client
}

func getNewClients() []testClient {
	return []testClient{
		{
			name: "Correct Client",
			want: Client{
				BaseURL: "example.com",
			},
			got: NewClient("example.com"),
		},
		{
			name: "Trimmed URL",
			want: Client{
				BaseURL: "example.com",
			},
			got: NewClient("example.com/"),
		},
	}
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
