package vars

// UserAgent stores the info about the browser, OS and device type
type UserAgent struct {
	Browser    Browser
	OS         OS
	DeviceType DeviceType
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
