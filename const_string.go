// Code generated by "stringer -type=DeviceType,BrowserName,OSName,Platform -output=const_string.go"; DO NOT EDIT.

package uasurfer

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DeviceUnknown-0]
	_ = x[DeviceComputer-1]
	_ = x[DeviceTablet-2]
	_ = x[DevicePhone-3]
	_ = x[DeviceConsole-4]
	_ = x[DeviceWearable-5]
	_ = x[DeviceTV-6]
}

const _DeviceType_name = "DeviceUnknownDeviceComputerDeviceTabletDevicePhoneDeviceConsoleDeviceWearableDeviceTV"

var _DeviceType_index = [...]uint8{0, 13, 27, 39, 50, 63, 77, 85}

func (i DeviceType) String() string {
	if i < 0 || i >= DeviceType(len(_DeviceType_index)-1) {
		return "DeviceType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DeviceType_name[_DeviceType_index[i]:_DeviceType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BrowserUnknown-0]
	_ = x[BrowserChrome-1]
	_ = x[BrowserChromium-2]
	_ = x[BrowserIE-3]
	_ = x[BrowserEdge-4]
	_ = x[BrowserSafari-5]
	_ = x[BrowserFirefox-6]
	_ = x[BrowserAndroid-7]
	_ = x[BrowserOpera-8]
	_ = x[BrowserBrave-9]
	_ = x[BrowserVivaldi-10]
	_ = x[BrowserUCBrowser-11]
	_ = x[BrowserQQ-12]
	_ = x[BrowserSamsung-13]
	_ = x[BrowserYandex-14]
	_ = x[BrowserCocCoc-15]
	_ = x[BrowserBot-16]
	_ = x[BrowserAppleBot-17]
	_ = x[BrowserBaiduBot-18]
	_ = x[BrowserBingBot-19]
	_ = x[BrowserDuckDuckGoBot-20]
	_ = x[BrowserFacebookBot-21]
	_ = x[BrowserGoogleBot-22]
	_ = x[BrowserLinkedInBot-23]
	_ = x[BrowserMsnBot-24]
	_ = x[BrowserPingdomBot-25]
	_ = x[BrowserTwitterBot-26]
	_ = x[BrowserYandexBot-27]
	_ = x[BrowserCocCocBot-28]
	_ = x[BrowserYahooBot-29]
}

const _BrowserName_name = "BrowserUnknownBrowserChromeBrowserChromiumBrowserIEBrowserEdgeBrowserSafariBrowserFirefoxBrowserAndroidBrowserOperaBrowserBraveBrowserVivaldiBrowserUCBrowserBrowserQQBrowserSamsungBrowserYandexBrowserCocCocBrowserBotBrowserAppleBotBrowserBaiduBotBrowserBingBotBrowserDuckDuckGoBotBrowserFacebookBotBrowserGoogleBotBrowserLinkedInBotBrowserMsnBotBrowserPingdomBotBrowserTwitterBotBrowserYandexBotBrowserCocCocBotBrowserYahooBot"

var _BrowserName_index = [...]uint16{0, 14, 27, 42, 51, 62, 75, 89, 103, 115, 127, 141, 157, 166, 180, 193, 206, 216, 231, 246, 260, 280, 298, 314, 332, 345, 362, 379, 395, 411, 426}

func (i BrowserName) String() string {
	if i < 0 || i >= BrowserName(len(_BrowserName_index)-1) {
		return "BrowserName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BrowserName_name[_BrowserName_index[i]:_BrowserName_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OSUnknown-0]
	_ = x[OSWindows-1]
	_ = x[OSMacOS-2]
	_ = x[OSiOS-3]
	_ = x[OSAndroid-4]
	_ = x[OSChromeOS-5]
	_ = x[OSLinux-6]
	_ = x[OSBot-7]
}

const _OSName_name = "OSUnknownOSWindowsOSMacOSOSiOSOSAndroidOSChromeOSOSLinuxOSBot"

var _OSName_index = [...]uint8{0, 9, 18, 25, 30, 39, 49, 56, 61}

func (i OSName) String() string {
	if i < 0 || i >= OSName(len(_OSName_index)-1) {
		return "OSName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OSName_name[_OSName_index[i]:_OSName_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PlatformUnknown-0]
	_ = x[PlatformWindows-1]
	_ = x[PlatformMac-2]
	_ = x[PlatformLinux-3]
	_ = x[PlatformiPad-4]
	_ = x[PlatformiPhone-5]
	_ = x[PlatformiPod-6]
	_ = x[PlatformBot-7]
}

const _Platform_name = "PlatformUnknownPlatformWindowsPlatformMacPlatformLinuxPlatformiPadPlatformiPhonePlatformiPodPlatformBot"

var _Platform_index = [...]uint8{0, 15, 30, 41, 54, 66, 80, 92, 103}

func (i Platform) String() string {
	if i < 0 || i >= Platform(len(_Platform_index)-1) {
		return "Platform(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Platform_name[_Platform_index[i]:_Platform_index[i+1]]
}
