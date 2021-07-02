package facio

type header uint

// All HTTP Headers
const (
	Authorization header = iota
	AcceptEncoding
	ContentType
)

func (h header) getHeader() string {
	switch h {
	case Authorization:
		return "Authorization"
	case AcceptEncoding:
		return "Accept-Encoding"
	case ContentType:
		return "Content-Type"
	default:
		return ""
	}
}
