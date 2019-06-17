package uaparser

import (
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

// Retrieve browser name from UA strings
func (u *UserAgent) evalBrowserName(ua string) bool {

	if strings.Contains(ua, "applewebkit") {
		return u.evalWebkitBrowserName(ua)
	}

	return u.evalNonWebkitBrowserName(ua)
}

// Retrieve browser name from UA strings containing 'applewebkit'
func (u *UserAgent) evalWebkitBrowserName(ua string) bool {

	switch {
	case strings.Contains(ua, "googlebot"):
		u.Browser.Name = BrowserGoogleBot

	case strings.Contains(ua, "qq/") || strings.Contains(ua, "qqbrowser/"):
		u.Browser.Name = BrowserQQ

	case strings.Contains(ua, "opr/") || strings.Contains(ua, "opios/"):
		u.Browser.Name = BrowserOpera

	case strings.Contains(ua, "brave/") || strings.Contains(ua, "brave chrome/"):
		u.Browser.Name = BrowserBrave

	case strings.Contains(ua, "vivaldi/"):
		u.Browser.Name = BrowserVivaldi

	case strings.Contains(ua, "edge/") || strings.Contains(ua, "edg/") || strings.Contains(ua, "edgios/") || strings.Contains(ua, "edga/"):
		u.Browser.Name = BrowserEdge

	case strings.Contains(ua, "iemobile/") || strings.Contains(ua, "msie "):
		u.Browser.Name = BrowserIE

	case strings.Contains(ua, "ucbrowser/") || strings.Contains(ua, "ucweb/"):
		u.Browser.Name = BrowserUCBrowser

	case strings.Contains(ua, "samsungbrowser/"):
		u.Browser.Name = BrowserSamsung

	case strings.Contains(ua, "coc_coc_browser/"):
		u.Browser.Name = BrowserCocCoc

	case strings.Contains(ua, "yabrowser/"):
		u.Browser.Name = BrowserYandex

	case strings.Contains(ua, "chromium/"):
		u.Browser.Name = BrowserChromium

		// Edge, Silk and other chrome-identifying browsers must evaluate before chrome, unless we want to add more overhead
	case strings.Contains(ua, "chrome/") || strings.Contains(ua, "crios/") || strings.Contains(ua, "crmo/"):
		u.Browser.Name = BrowserChrome

	case strings.Contains(ua, "android") && !strings.Contains(ua, "chrome/") && strings.Contains(ua, "version/") && !strings.Contains(ua, "like android"):
		// Android WebView on Android >= 4.4 is purposefully being identified as Chrome above -- https://developer.chrome.com/multidevice/webview/overview
		u.Browser.Name = BrowserAndroid

	case strings.Contains(ua, "fxios"):
		u.Browser.Name = BrowserFirefox

		// AppleBot uses webkit signature as well
	case strings.Contains(ua, "applebot"):
		u.Browser.Name = BrowserAppleBot

		// presume it's safari unless an esoteric browser is being specified (webOSBrowser, SamsungBrowser, etc.)
	case strings.Contains(ua, "like gecko") && strings.Contains(ua, "mozilla/") && strings.Contains(ua, "safari/") && !strings.Contains(ua, "linux") && !strings.Contains(ua, "android") && !strings.Contains(ua, "browser/") && !strings.Contains(ua, "os/") && !strings.Contains(ua, "yabrowser/"):
		u.Browser.Name = BrowserSafari

		// if we got this far and the device is iPhone or iPad, assume safari. Some agents don't actually contain the word "safari"
	case strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad"):
		u.Browser.Name = BrowserSafari

		// Google's search app on iPhone, leverages native Safari rather than Chrome
	case strings.Contains(ua, " gsa/"):
		u.Browser.Name = BrowserSafari

	default:
		return u.evalNonWebkitBrowserName(ua)

	}

	return u.maybeBot(ua)
}

// Retrieve browser name from UA strings not containing 'applewebkit'
func (u *UserAgent) evalNonWebkitBrowserName(ua string) bool {

	switch {
	case strings.Contains(ua, "qq/") || strings.Contains(ua, "qqbrowser/"):
		u.Browser.Name = BrowserQQ

	case strings.Contains(ua, "msie") || strings.Contains(ua, "trident"):
		u.Browser.Name = BrowserIE

	case strings.Contains(ua, "gecko") && (strings.Contains(ua, "firefox") || strings.Contains(ua, "iceweasel") || strings.Contains(ua, "seamonkey") || strings.Contains(ua, "icecat")):
		u.Browser.Name = BrowserFirefox

	case strings.Contains(ua, "presto") || strings.Contains(ua, "opera"):
		u.Browser.Name = BrowserOpera

	case strings.Contains(ua, "ucbrowser"):
		u.Browser.Name = BrowserUCBrowser

	case strings.Contains(ua, "applebot"):
		u.Browser.Name = BrowserAppleBot

	case strings.Contains(ua, "baiduspider"):
		u.Browser.Name = BrowserBaiduBot

	case strings.Contains(ua, "adidxbot") || strings.Contains(ua, "bingbot") || strings.Contains(ua, "bingpreview"):
		u.Browser.Name = BrowserBingBot

	case strings.Contains(ua, "duckduckbot"):
		u.Browser.Name = BrowserDuckDuckGoBot

	case strings.Contains(ua, "facebot") || strings.Contains(ua, "facebookexternalhit"):
		u.Browser.Name = BrowserFacebookBot

	case strings.Contains(ua, "googlebot"):
		u.Browser.Name = BrowserGoogleBot

	case strings.Contains(ua, "linkedinbot"):
		u.Browser.Name = BrowserLinkedInBot

	case strings.Contains(ua, "msnbot"):
		u.Browser.Name = BrowserMsnBot

	case strings.Contains(ua, "pingdom.com_bot"):
		u.Browser.Name = BrowserPingdomBot

	case strings.Contains(ua, "twitterbot"):
		u.Browser.Name = BrowserTwitterBot

	case strings.Contains(ua, "yandex") || strings.Contains(ua, "yadirectfetcher"):
		u.Browser.Name = BrowserYandexBot

	case strings.Contains(ua, "yahoo"):
		u.Browser.Name = BrowserYahooBot

	case strings.Contains(ua, "coccocbot"):
		u.Browser.Name = BrowserCocCocBot

	case strings.Contains(ua, "phantomjs"):
		u.Browser.Name = BrowserBot

	default:
		u.Browser.Name = BrowserUnknown

	}

	return u.maybeBot(ua)
}

// Retrieve browser version
// Methods used in order:
// 1st: look for generic version/#
// 2nd: look for browser-specific instructions (e.g. chrome/34)
// 3rd: infer from OS (iOS only)
func (u *UserAgent) evalBrowserVersion(ua string) {
	// if there is a 'version/#' attribute with numeric version, use it -- except for Chrome since Android vendors sometimes hijack version/#
	if u.Browser.Name != BrowserChrome && u.Browser.Version.findVersionNumber(ua, "version/") {
		return
	}

	switch u.Browser.Name {

	case BrowserChrome:
		// match both chrome and crios
		_ = u.Browser.Version.findVersionNumber(ua, "chrome/") || u.Browser.Version.findVersionNumber(ua, "crios/") || u.Browser.Version.findVersionNumber(ua, "crmo/")
	case BrowserChromium:
		_ = u.Browser.Version.findVersionNumber(ua, "chromium/") || u.Browser.Version.findVersionNumber(ua, "chrome/")

	case BrowserYandex:
		_ = u.Browser.Version.findVersionNumber(ua, "yabrowser/")

	case BrowserVivaldi:
		_ = u.Browser.Version.findVersionNumber(ua, "vivaldi/")

	case BrowserQQ:
		_ = u.Browser.Version.findVersionNumber(ua, "qq/") || u.Browser.Version.findVersionNumber(ua, "qqbrowser/")

	case BrowserEdge:
		_ = u.Browser.Version.findVersionNumber(ua, "edge/") || u.Browser.Version.findVersionNumber(ua, "edgios/") || u.Browser.Version.findVersionNumber(ua, "edga/") || u.Browser.Version.findVersionNumber(ua, "edg/")

	case BrowserIE:
		if u.Browser.Version.findVersionNumber(ua, "msie ") {
			return
		}

		// get MSIE version from trident version https://en.wikipedia.org/wiki/Trident_(layout_engine)
		if u.Browser.Version.findVersionNumber(ua, "trident/") {
			// convert trident versions 3-7 to MSIE version
			if (u.Browser.Version.Major >= 3) && (u.Browser.Version.Major <= 7) {
				u.Browser.Version.Major += 4
			}
		}

	case BrowserFirefox:
		_ = u.Browser.Version.findVersionNumber(ua, "firefox/") || u.Browser.Version.findVersionNumber(ua, "fxios/")

	case BrowserSafari: // executes typically if we're on iOS and not using a familiar browser
		u.Browser.Version = u.OS.Version
		// early Safari used a version number +1 to OS version
		if (u.Browser.Version.Major <= 3) && (u.Browser.Version.Major >= 1) {
			u.Browser.Version.Major++
		}

	case BrowserUCBrowser:
		_ = u.Browser.Version.findVersionNumber(ua, "ucbrowser/")

	case BrowserOpera:
		_ = u.Browser.Version.findVersionNumber(ua, "opr/") || u.Browser.Version.findVersionNumber(ua, "opios/") || u.Browser.Version.findVersionNumber(ua, "opera/")

	case BrowserBrave:
		_ = u.Browser.Version.findVersionNumber(ua, "brave/") || u.Browser.Version.findVersionNumber(ua, "brave chrome/")

	case BrowserCocCoc:
		_ = u.Browser.Version.findVersionNumber(ua, "coc_coc_browser/")
	}
}
