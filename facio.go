package facio

// Facio is the main struct to execute all actions and config the client
type Facio struct {
	// Client stores the main config
	client *Client
}

// NewFacio returns Facio with given Client
func NewFacio(client *Client) *Facio {
	return &Facio{client: client}
}

// NewDefaultFacio returns a Facio
func NewDefaultFacio(baseURL string) *Facio {
	client := NewClient(baseURL)
	return &Facio{
		client: &client,
	}
}