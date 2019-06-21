package browser

import (
	"github.com/Kiuryy/uaparser/useragent"
	"github.com/Kiuryy/uaparser/vars"
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
// 			Major int
// 			Minor int
// 			Patch int
// 		}
// }

// EvalVersion retrieves the browser version from the given UA string
// Methods used in order:
// 1st: look for generic version/#
// 2nd: look for browser-specific instructions (e.g. chrome/34)
// 3rd: infer from OS (iOS only)
func EvalVersion(u *useragent.UserAgent, ua string) {
	// if there is a 'version/#' attribute with numeric version, use it -- except for Chrome since Android vendors sometimes hijack version/#
	if u.Browser.Name != vars.BrowserChrome && u.Browser.Version.FindVersionNumber(ua, "version/") {
		return
	}

	switch u.Browser.Name {

	case vars.BrowserChrome:
		// match both chrome and crios
		_ = u.Browser.Version.FindVersionNumber(ua, "chrome/") || u.Browser.Version.FindVersionNumber(ua, "crios/") || u.Browser.Version.FindVersionNumber(ua, "crmo/")
	case vars.BrowserChromium:
		_ = u.Browser.Version.FindVersionNumber(ua, "chromium/") || u.Browser.Version.FindVersionNumber(ua, "chrome/")

	case vars.BrowserYandex:
		_ = u.Browser.Version.FindVersionNumber(ua, "yabrowser/")

	case vars.BrowserVivaldi:
		_ = u.Browser.Version.FindVersionNumber(ua, "vivaldi/")

	case vars.BrowserQQ:
		_ = u.Browser.Version.FindVersionNumber(ua, "qq/") || u.Browser.Version.FindVersionNumber(ua, "qqbrowser/")

	case vars.BrowserEdge:
		_ = u.Browser.Version.FindVersionNumber(ua, "edge/") || u.Browser.Version.FindVersionNumber(ua, "edgios/") || u.Browser.Version.FindVersionNumber(ua, "edga/") || u.Browser.Version.FindVersionNumber(ua, "edg/")

	case vars.BrowserIE:
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

	case vars.BrowserFirefox:
		_ = u.Browser.Version.FindVersionNumber(ua, "firefox/") || u.Browser.Version.FindVersionNumber(ua, "fxios/")

	case vars.BrowserUCBrowser:
		_ = u.Browser.Version.FindVersionNumber(ua, "ucbrowser/")

	case vars.BrowserOpera:
		_ = u.Browser.Version.FindVersionNumber(ua, "opr/") || u.Browser.Version.FindVersionNumber(ua, "opios/") || u.Browser.Version.FindVersionNumber(ua, "opera/")

	case vars.BrowserBrave:
		_ = u.Browser.Version.FindVersionNumber(ua, "brave/") || u.Browser.Version.FindVersionNumber(ua, "brave chrome/")

	case vars.BrowserCocCoc:
		_ = u.Browser.Version.FindVersionNumber(ua, "coc_coc_browser/")
	}
}

// EvalName retrieves the browser name from the given UA string
func EvalName(u *useragent.UserAgent, ua string) {

	if strings.Contains(ua, "applewebkit") {
		evalWebkitBrowserName(u, ua)
	} else {
		evalNonWebkitBrowserName(u, ua)
	}
}

// Retrieve browser name from UA strings containing 'applewebkit'
func evalWebkitBrowserName(u *useragent.UserAgent, ua string) {

	switch {
	case strings.Contains(ua, "googlebot"):
		u.Browser.Name = vars.BrowserGoogleBot

	case strings.Contains(ua, "qq/") || strings.Contains(ua, "qqbrowser/"):
		u.Browser.Name = vars.BrowserQQ

	case strings.Contains(ua, "opr/") || strings.Contains(ua, "opios/"):
		u.Browser.Name = vars.BrowserOpera

	case strings.Contains(ua, "brave/") || strings.Contains(ua, "brave chrome/"):
		u.Browser.Name = vars.BrowserBrave

	case strings.Contains(ua, "vivaldi/"):
		u.Browser.Name = vars.BrowserVivaldi

	case strings.Contains(ua, "edge/") || strings.Contains(ua, "edg/") || strings.Contains(ua, "edgios/") || strings.Contains(ua, "edga/"):
		u.Browser.Name = vars.BrowserEdge

	case strings.Contains(ua, "iemobile/") || strings.Contains(ua, "msie "):
		u.Browser.Name = vars.BrowserIE

	case strings.Contains(ua, "ucbrowser/") || strings.Contains(ua, "ucweb/"):
		u.Browser.Name = vars.BrowserUCBrowser

	case strings.Contains(ua, "samsungbrowser/"):
		u.Browser.Name = vars.BrowserSamsung

	case strings.Contains(ua, "coc_coc_browser/"):
		u.Browser.Name = vars.BrowserCocCoc

	case strings.Contains(ua, "yabrowser/"):
		u.Browser.Name = vars.BrowserYandex

	case strings.Contains(ua, "chromium/"):
		u.Browser.Name = vars.BrowserChromium

		// Edge, Vivaldi and other chrome-identifying browsers must evaluate before chrome, unless we want to add more overhead
	case strings.Contains(ua, "chrome/") || strings.Contains(ua, "crios/") || strings.Contains(ua, "crmo/"):
		u.Browser.Name = vars.BrowserChrome

	case strings.Contains(ua, "android") && !strings.Contains(ua, "chrome/") && strings.Contains(ua, "version/") && !strings.Contains(ua, "like android"):
		// Android WebView on Android >= 4.4 is purposefully being identified as Chrome above -- https://developer.chrome.com/multidevice/webview/overview
		u.Browser.Name = vars.BrowserAndroid

	case strings.Contains(ua, "fxios"):
		u.Browser.Name = vars.BrowserFirefox

		// AppleBot uses webkit signature as well
	case strings.Contains(ua, "applebot"):
		u.Browser.Name = vars.BrowserAppleBot

		// presume it's safari unless an esoteric browser is being specified (webOSBrowser, SamsungBrowser, etc.)
	case strings.Contains(ua, "like gecko") && strings.Contains(ua, "mozilla/") && strings.Contains(ua, "safari/") && !strings.Contains(ua, "linux") && !strings.Contains(ua, "android") && !strings.Contains(ua, "browser/") && !strings.Contains(ua, "os/") && !strings.Contains(ua, "yabrowser/"):
		u.Browser.Name = vars.BrowserSafari

		// if we got this far and the device is iPhone or iPad, assume safari. Some agents don't actually contain the word "safari"
	case strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad"):
		u.Browser.Name = vars.BrowserSafari

		// Google's search app on iPhone, leverages native Safari rather than Chrome
	case strings.Contains(ua, " gsa/"):
		u.Browser.Name = vars.BrowserSafari

	default:
		evalNonWebkitBrowserName(u, ua)

	}
}

// Retrieve browser name from UA strings not containing 'applewebkit'
func evalNonWebkitBrowserName(u *useragent.UserAgent, ua string) {

	switch {
	case strings.Contains(ua, "qq/") || strings.Contains(ua, "qqbrowser/"):
		u.Browser.Name = vars.BrowserQQ

	case strings.Contains(ua, "msie") || strings.Contains(ua, "trident"):
		u.Browser.Name = vars.BrowserIE

	case strings.Contains(ua, "gecko") && strings.Contains(ua, "firefox"):
		u.Browser.Name = vars.BrowserFirefox

	case strings.Contains(ua, "presto") || strings.Contains(ua, "opera"):
		u.Browser.Name = vars.BrowserOpera

	case strings.Contains(ua, "ucbrowser"):
		u.Browser.Name = vars.BrowserUCBrowser

	case strings.Contains(ua, "applebot"):
		u.Browser.Name = vars.BrowserAppleBot

	case strings.Contains(ua, "baiduspider"):
		u.Browser.Name = vars.BrowserBaiduBot

	case strings.Contains(ua, "adidxbot") || strings.Contains(ua, "bingbot") || strings.Contains(ua, "bingpreview"):
		u.Browser.Name = vars.BrowserBingBot

	case strings.Contains(ua, "duckduckbot"):
		u.Browser.Name = vars.BrowserDuckDuckGoBot

	case strings.Contains(ua, "facebot") || strings.Contains(ua, "facebookexternalhit"):
		u.Browser.Name = vars.BrowserFacebookBot

	case strings.Contains(ua, "googlebot"):
		u.Browser.Name = vars.BrowserGoogleBot

	case strings.Contains(ua, "linkedinbot"):
		u.Browser.Name = vars.BrowserLinkedInBot

	case strings.Contains(ua, "msnbot"):
		u.Browser.Name = vars.BrowserMsnBot

	case strings.Contains(ua, "pingdom.com_bot"):
		u.Browser.Name = vars.BrowserPingdomBot

	case strings.Contains(ua, "twitterbot"):
		u.Browser.Name = vars.BrowserTwitterBot

	case strings.Contains(ua, "yandex") || strings.Contains(ua, "yadirectfetcher"):
		u.Browser.Name = vars.BrowserYandexBot

	case strings.Contains(ua, "yahoo"):
		u.Browser.Name = vars.BrowserYahooBot

	case strings.Contains(ua, "coccocbot"):
		u.Browser.Name = vars.BrowserCocCocBot

	case strings.Contains(ua, "phantomjs"):
		u.Browser.Name = vars.BrowserBot

	default:
		u.Browser.Name = vars.BrowserUnknown

	}
}
