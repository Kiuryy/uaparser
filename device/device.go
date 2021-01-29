package device

import (
	"github.com/Kiuryy/uaparser/vars"
	"strings"
)

// Eval parses the user agent and sets the device type
func Eval(u *vars.UserAgent, ua string) {
	switch {

	case strings.Contains(ua, "nintendo") || strings.Contains(ua, "xbox") || strings.Contains(ua, "playstation"):
		u.DeviceType = vars.DeviceConsole

	case u.OS.Platform == vars.PlatformWindows || u.OS.Platform == vars.PlatformMac || u.OS.Name == vars.OSChromeOS:
		if strings.Contains(ua, "mobile") || strings.Contains(ua, "touch") {
			u.DeviceType = vars.DeviceTablet // windows rt, linux haxor tablets
			return
		}
		u.DeviceType = vars.DeviceComputer

	case u.OS.Platform == vars.PlatformiPad || u.OS.Platform == vars.PlatformiPod || strings.Contains(ua, "tablet") || strings.Contains(ua, "kindle/") || strings.Contains(ua, "playbook"):
		u.DeviceType = vars.DeviceTablet

	case u.OS.Platform == vars.PlatformiPhone || strings.Contains(ua, "phone"):
		u.DeviceType = vars.DevicePhone

		// long list of smarttv and tv dongle identifiers
	case strings.Contains(ua, "tv") || strings.Contains(ua, "crkey") || strings.Contains(ua, "googletv") || strings.Contains(ua, "aftb") || strings.Contains(ua, "aftt") || strings.Contains(ua, "aftm") || strings.Contains(ua, "adt-") || strings.Contains(ua, "roku") || strings.Contains(ua, "viera") || strings.Contains(ua, "aquos") || strings.Contains(ua, "dtv") || strings.Contains(ua, "appletv") || strings.Contains(ua, "smarttv") || strings.Contains(ua, "tuner") || strings.Contains(ua, "smart-tv") || strings.Contains(ua, "hbbtv") || strings.Contains(ua, "netcast") || strings.Contains(ua, "vizio"):
		u.DeviceType = vars.DeviceTV

	case u.OS.Name == vars.OSAndroid:
		// android phones report as "mobile", android tablets should not but often do -- http://android-developers.blogspot.com/2010/12/android-browser-user-agent-issues.html
		if strings.Contains(ua, "mobile") {
			u.DeviceType = vars.DevicePhone
			return
		}

		if strings.Contains(ua, "tablet") || strings.Contains(ua, "nexus 7") || strings.Contains(ua, "nexus 9") || strings.Contains(ua, "nexus 10") || strings.Contains(ua, "xoom") || strings.Contains(ua, "sm-t") || strings.Contains(ua, "; kf") || strings.Contains(ua, "; t1") || strings.Contains(ua, "lenovo tab") {
			u.DeviceType = vars.DeviceTablet
			return
		}

		u.DeviceType = vars.DevicePhone // default to phone

	case strings.Contains(ua, "mobile") || strings.Contains(ua, "touch") || strings.Contains(ua, " mobi") || strings.Contains(ua, "webos"): //anything "mobile"/"touch" that didn't get captured as tablet, console or wearable is presumed a phone
		u.DeviceType = vars.DevicePhone

	case u.OS.Name == vars.OSLinux: // linux goes last since it's in so many other device types (tvs, android-based stuff)
		u.DeviceType = vars.DeviceComputer

	default:
		u.DeviceType = vars.DeviceUnknown
	}
}
