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

// Return the header name as string
func (h header) getHeader() string {
	switch h {
	// Standard Headers
	case AIM:
		return "A-IM"
	case Accept:
		return "Accept"
	case AcceptCharset:
		return "Accept-Charset"
	case AcceptEncoding:
		return "Accept-Encoding"
	case AcceptLanguage:
		return "Accept-Language"
	case AcceptDatetime:
		return "Accept-Datetime"
	case AccessControlRequestMethod:
		return "Access-Control-Request-Method"
	case AccessControlRequestHeaders:
		return "Access-Control-Request-Headers"
	case Authorization:
		return "Authorization"
	case CacheControl:
		return "Cache-Control"
	case Connection:
		return "Connection"
	case ContentLength:
		return "Content-Length"
	case ContentType:
		return "Content-Type"
	case Cookie:
		return "Cookie"
	case Date:
		return "Date"
	case Expect:
		return "Expect"
	case Forwarded:
		return "Forwarded"
	case From:
		return "From"
	case Host:
		return "Host"
	case IfMatch:
		return "If-Match"
	case IfModifiedSince:
		return "If-Modified-Since"
	case IfNoneMatch:
		return "If-None-Match"
	case IfRange:
		return "If-Range"
	case IfUnmodifiedSince:
		return "If-Unmodified-Since"
	case MaxForwards:
		return "Max-Forwards"
	case Origin:
		return "Origin"
	case Pragma:
		return "Pragma"
	case ProxyAuthorization:
		return "Proxy-Authorization"
	case Range:
		return "Range"
	case Referer:
		return "Referer"
	case TE:
		return "TE"
	case UserAgent:
		return "User-Agent"
	case Upgrade:
		return "Upgrade"
	case Via:
		return "Via"
	case Warning:
		return "Warning"

	// Non-standard Headers
	case Dnt:
		return "Dnt"
	case XRequestedWith:
		return "X-Requested-With"
	case XCSRFToken:
		return "X-CSRF-Token"
	default:
		return ""
	}
}

