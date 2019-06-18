package useragent

import (
	"github.com/Kiuryy/uaparser/const"
	"github.com/Kiuryy/uaparser/version"
)

// Browser contains the name and version of the browser
type Browser struct {
	Name    _const.BrowserName
	Version version.Version
}

// OS contains the name, version and platform of the OS
type OS struct {
	Platform     _const.Platform
	Name         _const.OSName
	Version      version.Version
	VersionAlias string
}

// UserAgent stores the info about the browser, OS and device type
type UserAgent struct {
	Browser    Browser
	OS         OS
	DeviceType _const.DeviceType
}

// Reset resets the UserAgent to it's zero value
func (ua *UserAgent) Reset() {
	ua.Browser = Browser{}
	ua.OS = OS{}
	ua.DeviceType = _const.DeviceUnknown
}

// IsBot returns true if the UserAgent represent a bot
func (ua *UserAgent) IsBot() bool {
	if ua.Browser.Name >= _const.BrowserBot && ua.Browser.Name <= _const.BrowserYahooBot {
		return true
	}
	if ua.OS.Name == _const.OSBot {
		return true
	}
	if ua.OS.Platform == _const.PlatformBot {
		return true
	}
	return false
}
