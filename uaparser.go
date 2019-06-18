// Package uaparser provides fast and reliable abstraction
// of HTTP User-Agent strings. The philosophy is to identify
// browsers and OSes that are currently popular, and to avoid
// expending resources and accuracy on guessing at barely used
// or legacy technologies.
package uaparser

import (
	"github.com/Kiuryy/uaparser/browser"
	"github.com/Kiuryy/uaparser/const"
	"github.com/Kiuryy/uaparser/device"
	"github.com/Kiuryy/uaparser/system"
	"github.com/Kiuryy/uaparser/userAgent"
	"strings"
)

// Parse accepts a raw user agent (string) and returns the UserAgent.
func Parse(ua string) *userAgent.UserAgent {
	dest := new(userAgent.UserAgent)
	parse(ua, dest)
	return dest
}

// ParseUserAgent is the same as Parse, but populates the supplied UserAgent.
// It is the caller's responsibility to call Reset() on the UserAgent before
// passing it to this function.
func ParseUserAgent(ua string, dest *userAgent.UserAgent) {
	parse(ua, dest)
}

func parse(ua string, dest *userAgent.UserAgent) {
	ua = normalise(ua)
	switch {
	case len(ua) == 0:
		dest.OS.Platform = _const.PlatformUnknown
		dest.OS.Name = _const.OSUnknown
		dest.Browser.Name = _const.BrowserUnknown
		dest.DeviceType = _const.DeviceUnknown

		// stop on on first case returning true
	case system.Eval(dest, ua):
	case browser.EvalName(dest, ua):
	default:
		browser.EvalVersion(dest, ua)
		device.Eval(dest, ua)
	}
}

// normalise normalises the user supplied agent string so that
// we can more easily parse it.
func normalise(ua string) string {
	if len(ua) <= 1024 {
		var buf [1024]byte
		ascii := copyLower(buf[:len(ua)], ua)
		if !ascii {
			// Fall back for non ascii characters
			return strings.ToLower(ua)
		}
		return string(buf[:len(ua)])
	}
	// Fallback for unusually long strings
	return strings.ToLower(ua)
}

// copyLower copies a lowercase version of s to b. It assumes s contains only single byte characters
// and will panic if b is nil or is not long enough to contain all the bytes from s.
// It returns early with false if any characters were non ascii.
func copyLower(b []byte, s string) bool {
	for j := 0; j < len(s); j++ {
		c := s[j]
		if c > 127 {
			return false
		}

		if 'A' <= c && c <= 'Z' {
			c += 'a' - 'A'
		}

		b[j] = c
	}
	return true
}