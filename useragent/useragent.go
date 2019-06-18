package useragent

import (
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

// Browser contains the name and version of the browser
type Browser struct {
	Name    vars.BrowserName
	Version version.Version
}

// OS contains the name, version and platform of the OS
type OS struct {
	Platform     vars.Platform
	Name         vars.OSName
	Version      version.Version
	VersionAlias string
}

// UserAgent stores the info about the browser, OS and device type
type UserAgent struct {
	Browser    Browser
	OS         OS
	DeviceType vars.DeviceType
}

// Reset resets the UserAgent to it's zero value
func (ua *UserAgent) Reset() {
	ua.Browser = Browser{}
	ua.OS = OS{}
	ua.DeviceType = vars.DeviceUnknown
}

// IsBot returns true if the UserAgent represent a bot
func (ua *UserAgent) IsBot() bool {
	if ua.Browser.Name >= vars.BrowserBot && ua.Browser.Name <= vars.BrowserYahooBot {
		return true
	}
	if ua.OS.Name == vars.OSBot {
		return true
	}
	if ua.OS.Platform == vars.PlatformBot {
		return true
	}
	return false
}
