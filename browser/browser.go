package browser

import (
	"github.com/Kiuryy/uaparser/const"
	"github.com/Kiuryy/uaparser/system"
	"github.com/Kiuryy/uaparser/userAgent"
	"strings"
)

// Browser struct contains the lowercase name of the browser, along
// with its browser version number. Browser are grouped together without
// consideration for device. For example, Chrome (Chrome/43.0) and Chrome for iOS
// (CriOS/43.0) would both return as "chrome" (name) and 43.0 (version). Similarly
// Internet Explorer 11 and Edge 12 would return as "ie" and "11" or "12", respectively.
// type Browser struct {
// 		Name    BrowserName
// 		Version struct {
// 				Major int
// 			Minor int
// 			Patch int
// 		}
// }

// Retrieve browser version
// Methods used in order:
// 1st: look for generic version/#
// 2nd: look for browser-specific instructions (e.g. chrome/34)
// 3rd: infer from OS (iOS only)
func EvalVersion(u *userAgent.UserAgent, ua string) {
	// if there is a 'version/#' attribute with numeric version, use it -- except for Chrome since Android vendors sometimes hijack version/#
	if u.Browser.Name != _const.BrowserChrome && u.Browser.Version.FindVersionNumber(ua, "version/") {
		return
	}

	switch u.Browser.Name {

	case _const.BrowserChrome:
		// match both chrome and crios
		_ = u.Browser.Version.FindVersionNumber(ua, "chrome/") || u.Browser.Version.FindVersionNumber(ua, "crios/") || u.Browser.Version.FindVersionNumber(ua, "crmo/")
	case _const.BrowserChromium:
		_ = u.Browser.Version.FindVersionNumber(ua, "chromium/") || u.Browser.Version.FindVersionNumber(ua, "chrome/")

	case _const.BrowserYandex:
		_ = u.Browser.Version.FindVersionNumber(ua, "yabrowser/")

	case _const.BrowserVivaldi:
		_ = u.Browser.Version.FindVersionNumber(ua, "vivaldi/")

	case _const.BrowserQQ:
		_ = u.Browser.Version.FindVersionNumber(ua, "qq/") || u.Browser.Version.FindVersionNumber(ua, "qqbrowser/")

	case _const.BrowserEdge:
		_ = u.Browser.Version.FindVersionNumber(ua, "edge/") || u.Browser.Version.FindVersionNumber(ua, "edgios/") || u.Browser.Version.FindVersionNumber(ua, "edga/") || u.Browser.Version.FindVersionNumber(ua, "edg/")

	case _const.BrowserIE:
		if u.Browser.Version.FindVersionNumber(ua, "msie ") {
			return
		}

		// get MSIE version from trident version https://en.wikipedia.org/wiki/Trident_(layout_engine)
		if u.Browser.Version.FindVersionNumber(ua, "trident/") {
			// convert trident versions 3-7 to MSIE version
			if (u.Browser.Version.Major >= 3) && (u.Browser.Version.Major <= 7) {
				u.Browser.Version.Major += 4
			}
		}

	case _const.BrowserFirefox:
		_ = u.Browser.Version.FindVersionNumber(ua, "firefox/") || u.Browser.Version.FindVersionNumber(ua, "fxios/")

	case _const.BrowserSafari: // executes typically if we're on iOS and not using a familiar browser
		u.Browser.Version = u.OS.Version
		// early Safari used a version number +1 to OS version
		if (u.Browser.Version.Major <= 3) && (u.Browser.Version.Major >= 1) {
			u.Browser.Version.Major++
		}

	case _const.BrowserUCBrowser:
		_ = u.Browser.Version.FindVersionNumber(ua, "ucbrowser/")

	case _const.BrowserOpera:
		_ = u.Browser.Version.FindVersionNumber(ua, "opr/") || u.Browser.Version.FindVersionNumber(ua, "opios/") || u.Browser.Version.FindVersionNumber(ua, "opera/")

	case _const.BrowserBrave:
		_ = u.Browser.Version.FindVersionNumber(ua, "brave/") || u.Browser.Version.FindVersionNumber(ua, "brave chrome/")

	case _const.BrowserCocCoc:
		_ = u.Browser.Version.FindVersionNumber(ua, "coc_coc_browser/")
	}
}

// Retrieve browser name from UA strings
func EvalName(u *userAgent.UserAgent, ua string) bool {

	if strings.Contains(ua, "applewebkit") {
		return evalWebkitBrowserName(u, ua)
	}

	return evalNonWebkitBrowserName(u, ua)
}

// Retrieve browser name from UA strings containing 'applewebkit'
func evalWebkitBrowserName(u *userAgent.UserAgent, ua string) bool {

	switch {
	case strings.Contains(ua, "googlebot"):
		u.Browser.Name = _const.BrowserGoogleBot

	case strings.Contains(ua, "qq/") || strings.Contains(ua, "qqbrowser/"):
		u.Browser.Name = _const.BrowserQQ

	case strings.Contains(ua, "opr/") || strings.Contains(ua, "opios/"):
		u.Browser.Name = _const.BrowserOpera

	case strings.Contains(ua, "brave/") || strings.Contains(ua, "brave chrome/"):
		u.Browser.Name = _const.BrowserBrave

	case strings.Contains(ua, "vivaldi/"):
		u.Browser.Name = _const.BrowserVivaldi

	case strings.Contains(ua, "edge/") || strings.Contains(ua, "edg/") || strings.Contains(ua, "edgios/") || strings.Contains(ua, "edga/"):
		u.Browser.Name = _const.BrowserEdge

	case strings.Contains(ua, "iemobile/") || strings.Contains(ua, "msie "):
		u.Browser.Name = _const.BrowserIE

	case strings.Contains(ua, "ucbrowser/") || strings.Contains(ua, "ucweb/"):
		u.Browser.Name = _const.BrowserUCBrowser

	case strings.Contains(ua, "samsungbrowser/"):
		u.Browser.Name = _const.BrowserSamsung

	case strings.Contains(ua, "coc_coc_browser/"):
		u.Browser.Name = _const.BrowserCocCoc

	case strings.Contains(ua, "yabrowser/"):
		u.Browser.Name = _const.BrowserYandex

	case strings.Contains(ua, "chromium/"):
		u.Browser.Name = _const.BrowserChromium

		// Edge, Silk and other chrome-identifying browsers must evaluate before chrome, unless we want to add more overhead
	case strings.Contains(ua, "chrome/") || strings.Contains(ua, "crios/") || strings.Contains(ua, "crmo/"):
		u.Browser.Name = _const.BrowserChrome

	case strings.Contains(ua, "android") && !strings.Contains(ua, "chrome/") && strings.Contains(ua, "version/") && !strings.Contains(ua, "like android"):
		// Android WebView on Android >= 4.4 is purposefully being identified as Chrome above -- https://developer.chrome.com/multidevice/webview/overview
		u.Browser.Name = _const.BrowserAndroid

	case strings.Contains(ua, "fxios"):
		u.Browser.Name = _const.BrowserFirefox

		// AppleBot uses webkit signature as well
	case strings.Contains(ua, "applebot"):
		u.Browser.Name = _const.BrowserAppleBot

		// presume it's safari unless an esoteric browser is being specified (webOSBrowser, SamsungBrowser, etc.)
	case strings.Contains(ua, "like gecko") && strings.Contains(ua, "mozilla/") && strings.Contains(ua, "safari/") && !strings.Contains(ua, "linux") && !strings.Contains(ua, "android") && !strings.Contains(ua, "browser/") && !strings.Contains(ua, "os/") && !strings.Contains(ua, "yabrowser/"):
		u.Browser.Name = _const.BrowserSafari

		// if we got this far and the device is iPhone or iPad, assume safari. Some agents don't actually contain the word "safari"
	case strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad"):
		u.Browser.Name = _const.BrowserSafari

		// Google's search app on iPhone, leverages native Safari rather than Chrome
	case strings.Contains(ua, " gsa/"):
		u.Browser.Name = _const.BrowserSafari

	default:
		return evalNonWebkitBrowserName(u, ua)

	}

	return system.MaybeBot(u, ua)
}

// Retrieve browser name from UA strings not containing 'applewebkit'
func evalNonWebkitBrowserName(u *userAgent.UserAgent, ua string) bool {

	switch {
	case strings.Contains(ua, "qq/") || strings.Contains(ua, "qqbrowser/"):
		u.Browser.Name = _const.BrowserQQ

	case strings.Contains(ua, "msie") || strings.Contains(ua, "trident"):
		u.Browser.Name = _const.BrowserIE

	case strings.Contains(ua, "gecko") && (strings.Contains(ua, "firefox") || strings.Contains(ua, "iceweasel") || strings.Contains(ua, "seamonkey") || strings.Contains(ua, "icecat")):
		u.Browser.Name = _const.BrowserFirefox

	case strings.Contains(ua, "presto") || strings.Contains(ua, "opera"):
		u.Browser.Name = _const.BrowserOpera

	case strings.Contains(ua, "ucbrowser"):
		u.Browser.Name = _const.BrowserUCBrowser

	case strings.Contains(ua, "applebot"):
		u.Browser.Name = _const.BrowserAppleBot

	case strings.Contains(ua, "baiduspider"):
		u.Browser.Name = _const.BrowserBaiduBot

	case strings.Contains(ua, "adidxbot") || strings.Contains(ua, "bingbot") || strings.Contains(ua, "bingpreview"):
		u.Browser.Name = _const.BrowserBingBot

	case strings.Contains(ua, "duckduckbot"):
		u.Browser.Name = _const.BrowserDuckDuckGoBot

	case strings.Contains(ua, "facebot") || strings.Contains(ua, "facebookexternalhit"):
		u.Browser.Name = _const.BrowserFacebookBot

	case strings.Contains(ua, "googlebot"):
		u.Browser.Name = _const.BrowserGoogleBot

	case strings.Contains(ua, "linkedinbot"):
		u.Browser.Name = _const.BrowserLinkedInBot

	case strings.Contains(ua, "msnbot"):
		u.Browser.Name = _const.BrowserMsnBot

	case strings.Contains(ua, "pingdom.com_bot"):
		u.Browser.Name = _const.BrowserPingdomBot

	case strings.Contains(ua, "twitterbot"):
		u.Browser.Name = _const.BrowserTwitterBot

	case strings.Contains(ua, "yandex") || strings.Contains(ua, "yadirectfetcher"):
		u.Browser.Name = _const.BrowserYandexBot

	case strings.Contains(ua, "yahoo"):
		u.Browser.Name = _const.BrowserYahooBot

	case strings.Contains(ua, "coccocbot"):
		u.Browser.Name = _const.BrowserCocCocBot

	case strings.Contains(ua, "phantomjs"):
		u.Browser.Name = _const.BrowserBot

	default:
		u.Browser.Name = _const.BrowserUnknown

	}

	return system.MaybeBot(u, ua)
}
