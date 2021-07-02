package facio_test

import (
	"reflect"
	"testing"

	facio "github.com/fiuskylab/facio-http"
)

type testClient struct {
	name string
	want facio.Client
	got  facio.Client
}

func getNewClients() []testClient {
	return []testClient{
		{
			name: "Correct Client",
			want: facio.Client{
				BaseURL: "example.com",
				Request: facio.Request{
					Headers: make(facio.HeaderResult),
				},
				HeaderMap: make(facio.HeaderMap),
			},
			got: facio.NewClient("example.com"),
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
