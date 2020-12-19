package vars

// DeviceType (int) returns a constant.
type DeviceType int

// A complete list of supported device types in the
// form of constants.
const (
	DeviceUnknown DeviceType = iota
	DeviceComputer
	DeviceTablet
	DevicePhone
	DeviceConsole
	DeviceTV
)

// OSName (int) returns a constant.
type OSName int

// A complete list of supported OS in the
// form of constants.
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

// Platform (int) returns a constant.
type Platform int

// A complete list of supported platforms in the
// form of constants. Many OS report their true
// platform, such as Android OS being Linux
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
