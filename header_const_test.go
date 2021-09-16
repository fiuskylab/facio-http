package facio

import (
	"testing"
)

func TestHeader(t *testing.T) {
	tts := []testCases{}

	tts = append(tts, testCases{
		name:     "Header Test",
		got:      AIM.getHeader(),
		want:     "A-IM",
		testType: Equal,
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		got:      Accept.getHeader(),
		want:     "Accept",
		testType: Equal,
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		got:      AcceptEncoding.getHeader(),
		want:     "Accept-Encoding",
		testType: Equal,
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		got:      AcceptLanguage.getHeader(),
		want:     "Accept-Language",
		testType: Equal,
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		got:      AcceptDatetime.getHeader(),
		want:     "Accept-Datetime",
		testType: Equal,
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		got:      AccessControlAllowCredentials.getHeader(),
		want:     "Access-Control-Allow-Credentials",
		testType: Equal,
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		got:      AccessControlAllowHeaders.getHeader(),
		want:     "Access-Control-Allow-Headers",
		testType: Equal,
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		want:     AccessControlAllowMethods.getHeader(),
		got:      "Access-Control-Allow-Methods",
		testType: Equal,
	})

	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     AccessControlAllowOrigin.getHeader(),
		got:      "Access-Control-Allow-Origin",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     AccessControlExposeHeaders.getHeader(),
		got:      "Access-Control-Expose-Headers",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     AccessControlMaxAge.getHeader(),
		got:      "Access-Control-Max-Age",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     AccessControlRequestMethod.getHeader(),
		got:      "Access-Control-Request-Method",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     AccessControlRequestHeaders.getHeader(),
		got:      "Access-Control-Request-Headers",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     Authorization.getHeader(),
		got:      "Authorization",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     CacheControl.getHeader(),
		got:      "Cache-Control",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     Connection.getHeader(),
		got:      "Connection",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     ContentLength.getHeader(),
		got:      "Content-Length",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     ContentType.getHeader(),
		got:      "Content-Type",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     Cookie.getHeader(),
		got:      "Cookie",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     Date.getHeader(),
		got:      "Date",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     Expect.getHeader(),
		got:      "Expect",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     Forwarded.getHeader(),
		got:      "Forwarded",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     From.getHeader(),
		got:      "From",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     Host.getHeader(),
		got:      "Host",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     IfMatch.getHeader(),
		got:      "If-Match",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     IfModifiedSince.getHeader(),
		got:      "If-Modified-Since",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     IfNoneMatch.getHeader(),
		got:      "If-None-Match",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     IfRange.getHeader(),
		got:      "If-Range",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     IfUnmodifiedSince.getHeader(),
		got:      "If-Unmodified-Since",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     MaxForwards.getHeader(),
		got:      "Max-Forwards",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     Origin.getHeader(),
		got:      "Origin",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     Pragma.getHeader(),
		got:      "Pragma",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     ProxyAuthorization.getHeader(),
		got:      "Proxy-Authorization",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     Range.getHeader(),
		got:      "Range",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     Referer.getHeader(),
		got:      "Referer",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     TE.getHeader(),
		got:      "TE",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     UserAgent.getHeader(),
		got:      "User-Agent",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     Upgrade.getHeader(),
		got:      "Upgrade",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     Via.getHeader(),
		got:      "Via",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     Warning.getHeader(),
		got:      "Warning",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     Dnt.getHeader(),
		got:      "Dnt",
	})

	// Non-standard Headers
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     XRequestedWith.getHeader(),
		got:      "X-Requested-With",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     XCSRFToken.getHeader(),
		got:      "X-CSRF-Token",
	})
	tts = append(tts, testCases{
		name:     "Header Test",
		testType: Equal,
		want:     rand_header.getHeader(),
		got:      "",
	})

	runTests(t, tts)
}
