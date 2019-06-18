package _const

// DeviceType (int) returns a constant.
type DeviceType int

const (
	DeviceUnknown DeviceType = iota
	DeviceComputer
	DeviceTablet
	DevicePhone
	DeviceConsole
	DeviceWearable
	DeviceTV
)

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
