package vars

// UserAgent stores the info about the browser, OS and device type
type UserAgent struct {
	Browser    Browser
	OS         OS
	DeviceType DeviceType
}

// IsBot returns true if the Browser, OS or Platform is reported as "Bot"
func (ua *UserAgent) IsBot() bool {
	if (ua.Browser.Name >= BrowserBot && ua.Browser.Name <= BrowserYahooBot) ||
		ua.OS.Name == OSBot ||
		ua.OS.Platform == PlatformBot {
		return true
	}
	return false
}
