package facio

type header uint

// All HTTP Headers
const (
	// Standard Headers

	AIM                         header = iota // A-IM Header
	Accept                                    // Accept Header
	AcceptCharset                             // Accept-Charset Header
	AcceptEncoding                            // Accept-Encoding Header
	AcceptLanguage                            // Accept-Language Header
	AcceptDatetime                            // Accept-Datetime Header
	AccessControlRequestMethod                // Access-Control-Request-Method Header
	AccessControlRequestHeaders               // Access-Control-Request-Headers Header
	Authorization                             // Authorization Header
	CacheControl                              // Cache-Control Header
	Connection                                // Connection Header
	ContentLength                             // Content-Length Header
	ContentType                               // Content-Type Header
	Cookie                                    // Cookie Header
	Date                                      // Date Header
	Expect                                    // Expect Header
	Forwarded                                 // Forwarded Header
	From                                      // From Header
	Host                                      // Host Header
	IfMatch                                   // If-Match Header
	IfModifiedSince                           // If-Modified-Since Header
	IfNoneMatch                               // If-None-Match Header
	IfRange                                   // If-Range Header
	IfUnmodifiedSince                         // If-Unmodified-Since Header
	MaxForwards                               // Max-Forwards Header
	Origin                                    // Origin Header
	Pragma                                    // Pragma Header
	ProxyAuthorization                        // Proxy-Authorization Header
	Range                                     // Range Header
	Referer                                   // Referer Header
	TE                                        // TE Header
	UserAgent                                 // User-Agent Header
	Upgrade                                   // Upgrade Header
	Via                                       // Via Header
	Warning                                   // Warning Header

	// Non-standard Headers
	Dnt            // Dnt Header
	XRequestedWith // X-Requested-With Header
	XCSRFToken     // X-CSRF-Token Header
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
