package facio

type Req struct {
	BaseURL string
	Method  string
	Headers map[string]string
}

func NewRequest(method, url string) (*Req, error) {
	var err error
	var r Req

	method, err = checkMethod(method)
	if err != nil {
		return &r, err
	}

	url, err = parseURL(url)

	if err != nil {
		return &r, err
	}

	r.BaseURL = url
	r.Method = method

	return &r, err
}
