package device

import (
	"github.com/Kiuryy/uaparser/const"
	"github.com/Kiuryy/uaparser/userAgent"
	"strings"
)

func Eval(u *userAgent.UserAgent, ua string) {
	switch {

	case u.OS.Platform == _const.PlatformWindows || u.OS.Platform == _const.PlatformMac || u.OS.Name == _const.OSChromeOS:
		if strings.Contains(ua, "mobile") || strings.Contains(ua, "touch") {
			u.DeviceType = _const.DeviceTablet // windows rt, linux haxor tablets
			return
		}
		u.DeviceType = _const.DeviceComputer

	case u.OS.Platform == _const.PlatformiPad || u.OS.Platform == _const.PlatformiPod || strings.Contains(ua, "tablet") || strings.Contains(ua, "kindle/") || strings.Contains(ua, "playbook"):
		u.DeviceType = _const.DeviceTablet

	case u.OS.Platform == _const.PlatformiPhone || strings.Contains(ua, "phone"):
		u.DeviceType = _const.DevicePhone

		// long list of smarttv and tv dongle identifiers
	case strings.Contains(ua, "tv") || strings.Contains(ua, "crkey") || strings.Contains(ua, "googletv") || strings.Contains(ua, "aftb") || strings.Contains(ua, "aftt") || strings.Contains(ua, "aftm") || strings.Contains(ua, "adt-") || strings.Contains(ua, "roku") || strings.Contains(ua, "viera") || strings.Contains(ua, "aquos") || strings.Contains(ua, "dtv") || strings.Contains(ua, "appletv") || strings.Contains(ua, "smarttv") || strings.Contains(ua, "tuner") || strings.Contains(ua, "smart-tv") || strings.Contains(ua, "hbbtv") || strings.Contains(ua, "netcast") || strings.Contains(ua, "vizio"):
		u.DeviceType = _const.DeviceTV

	case u.OS.Name == _const.OSAndroid:
		// android phones report as "mobile", android tablets should not but often do -- http://android-developers.blogspot.com/2010/12/android-browser-user-agent-issues.html
		if strings.Contains(ua, "mobile") {
			u.DeviceType = _const.DevicePhone
			return
		}

		if strings.Contains(ua, "tablet") || strings.Contains(ua, "nexus 7") || strings.Contains(ua, "nexus 9") || strings.Contains(ua, "nexus 10") || strings.Contains(ua, "xoom") {
			u.DeviceType = _const.DeviceTablet
			return
		}

		u.DeviceType = _const.DevicePhone // default to phone

	case strings.Contains(ua, "nintendo") || strings.Contains(ua, "xbox") || strings.Contains(ua, "playstation"):
		u.DeviceType = _const.DeviceConsole

	case strings.Contains(ua, "glass") || strings.Contains(ua, "watch") || strings.Contains(ua, "sm-v"):
		u.DeviceType = _const.DeviceWearable

		// specifically above "mobile" string check as Kindle Fire tablets report as "mobile"
	case strings.Contains(ua, "kindle/") || strings.Contains(ua, "silk/") && !strings.Contains(ua, "sd4930ur"):
		u.DeviceType = _const.DeviceTablet

	case strings.Contains(ua, "mobile") || strings.Contains(ua, "touch") || strings.Contains(ua, " mobi") || strings.Contains(ua, "webos"): //anything "mobile"/"touch" that didn't get captured as tablet, console or wearable is presumed a phone
		u.DeviceType = _const.DevicePhone

	case u.OS.Name == _const.OSLinux: // linux goes last since it's in so many other device types (tvs, wearables, android-based stuff)
		u.DeviceType = _const.DeviceComputer

	default:
		u.DeviceType = _const.DeviceUnknown
	}
}
