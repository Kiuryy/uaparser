// Package uaparser provides fast and reliable abstraction
// of HTTP User-Agent strings. The philosophy is to identify
// technologies that holds >1% market share, and to avoid
// expending resources and accuracy on guessing at esoteric UA
// strings.
package uaparser

import "strings"

//go:generate stringer -type=DeviceType,BrowserName,OSName,Platform -output=const_string.go
//go:generate gofmt -s -w .

// DeviceType (int) returns a constant.
type DeviceType int

// A complete list of supported devices in the
// form of constants.
const (
	DeviceUnknown DeviceType = iota
	DeviceComputer
	DeviceTablet
	DevicePhone
	DeviceConsole
	DeviceWearable
	DeviceTV
)

// StringTrimPrefix is like String() but trims the "Device" prefix
func (d DeviceType) StringTrimPrefix() string {
	return strings.TrimPrefix(d.String(), "Device")
}

// BrowserName (int) returns a constant.
type BrowserName int

// A complete list of supported web browsers in the
// form of constants.
const (
	BrowserUnknown BrowserName = iota
	BrowserChrome
	BrowserChromium
	BrowserIE
	BrowserEdge
	BrowserSafari
	BrowserFirefox
	BrowserAndroid
	BrowserOpera
	BrowserBrave
	BrowserVivaldi
	BrowserUCBrowser
	BrowserQQ
	BrowserSamsung
	BrowserYandex
	BrowserCocCoc
	BrowserBot // Bot list begins here
	BrowserAppleBot
	BrowserBaiduBot
	BrowserBingBot
	BrowserDuckDuckGoBot
	BrowserFacebookBot
	BrowserGoogleBot
	BrowserLinkedInBot
	BrowserMsnBot
	BrowserPingdomBot
	BrowserTwitterBot
	BrowserYandexBot
	BrowserCocCocBot
	BrowserYahooBot // Bot list ends here
)

// StringTrimPrefix is like String() but trims the "Browser" prefix
func (b BrowserName) StringTrimPrefix() string {
	return strings.TrimPrefix(b.String(), "Browser")
}

// OSName (int) returns a constant.
type OSName int

// A complete list of supported OSes in the
// form of constants. For handling particular versions
// of operating systems (e.g. Windows 2000), see
// the README.md file.
const (
	OSUnknown OSName = iota
	OSWindows
	OSMacOS
	OSiOS
	OSAndroid
	OSChromeOS
	OSLinux
	OSBot
)

// StringTrimPrefix is like String() but trims the "OS" prefix
func (o OSName) StringTrimPrefix() string {
	return strings.TrimPrefix(o.String(), "OS")
}

// Platform (int) returns a constant.
type Platform int

// A complete list of supported platforms in the
// form of constants. Many OSes report their
// true platform, such as Android OS being Linux
// platform.
const (
	PlatformUnknown Platform = iota
	PlatformWindows
	PlatformMac
	PlatformLinux
	PlatformiPad
	PlatformiPhone
	PlatformiPod
	PlatformBot
)

// StringTrimPrefix is like String() but trims the "Platform" prefix
func (p Platform) StringTrimPrefix() string {
	return strings.TrimPrefix(p.String(), "Platform")
}

// Version contains three segments: 'Major', 'Minor' and 'Patch'
type Version struct {
	Major int
	Minor int
	Patch int
}

// Less compares the given version with the current version object: `if ver1.Less(ver2) {}`
func (v Version) Less(c Version) bool {
	if v.Major < c.Major {
		return true
	}

	if v.Major > c.Major {
		return false
	}

	if v.Minor < c.Minor {
		return true
	}

	if v.Minor > c.Minor {
		return false
	}

	return v.Patch < c.Patch
}

// UserAgent stores the info about the browser, OS and device type
type UserAgent struct {
	Browser    Browser
	OS         OS
	DeviceType DeviceType
}

// Browser contains the name and version of the browser
type Browser struct {
	Name    BrowserName
	Version Version
}

// OS contains the name, version and platform of the OS
type OS struct {
	Platform Platform
	Name     OSName
	Version  Version
}

// Reset resets the UserAgent to it's zero value
func (ua *UserAgent) Reset() {
	ua.Browser = Browser{}
	ua.OS = OS{}
	ua.DeviceType = DeviceUnknown
}

// IsBot returns true if the UserAgent represent a bot
func (ua *UserAgent) IsBot() bool {
	if ua.Browser.Name >= BrowserBot && ua.Browser.Name <= BrowserYahooBot {
		return true
	}
	if ua.OS.Name == OSBot {
		return true
	}
	if ua.OS.Platform == PlatformBot {
		return true
	}
	return false
}

// Parse accepts a raw user agent (string) and returns the UserAgent.
func Parse(ua string) *UserAgent {
	dest := new(UserAgent)
	parse(ua, dest)
	return dest
}

// ParseUserAgent is the same as Parse, but populates the supplied UserAgent.
// It is the caller's responsibility to call Reset() on the UserAgent before
// passing it to this function.
func ParseUserAgent(ua string, dest *UserAgent) {
	parse(ua, dest)
}

func parse(ua string, dest *UserAgent) {
	ua = normalise(ua)
	switch {
	case len(ua) == 0:
		dest.OS.Platform = PlatformUnknown
		dest.OS.Name = OSUnknown
		dest.Browser.Name = BrowserUnknown
		dest.DeviceType = DeviceUnknown

		// stop on on first case returning true
	case dest.evalOS(ua):
	case dest.evalBrowserName(ua):
	default:
		dest.evalBrowserVersion(ua)
		dest.evalDevice(ua)
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
