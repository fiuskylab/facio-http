package facio

// Store every info about a request to be made
type Request struct {
	// Headers sent to request
	// ["Header Name": ["Header Value 1", "Header Value 2"]]
	Headers HeaderResult

	// Endpoint to fetch based on Client's BaseURL
	Endpoint string

	// Method type POST, GET, PATCH, PUT, etc
	// Not case sensitive
	Method string
}
