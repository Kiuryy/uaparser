package uaparser

import (
	"strconv"
	"strings"
)

func (u *UserAgent) evalOS(ua string) bool {
	s := strings.IndexRune(ua, '(')
	e := strings.IndexRune(ua, ')')
	if s > e {
		s = 0
		e = len(ua)
	}
	if e == -1 {
		e = len(ua)
	}

	agentPlatform := ua[s+1 : e]
	specsEnd := strings.Index(agentPlatform, ";")
	var specs string
	if specsEnd != -1 {
		specs = agentPlatform[:specsEnd]
	} else {
		specs = agentPlatform
	}

	//strict OS & version identification
	switch specs {
	case "android":
		u.evalLinux(ua, agentPlatform)

	case "x11", "linux":
		u.evalLinux(ua, agentPlatform)

	case "ipad", "iphone", "ipod touch", "ipod":
		u.evaliOS(specs, agentPlatform)

	case "macintosh":
		u.evalMac(ua)

	default:
		switch {

		// Windows, Xbox
		case strings.Contains(ua, "windows ") || strings.Contains(ua, "microsoft-cryptoapi"):
			u.evalWindows(ua)

			// Linux, Android
		case strings.Contains(ua, "linux") || strings.Contains(ua, "android"):
			u.evalLinux(ua, agentPlatform)

			// Apple CFNetwork
		case strings.Contains(ua, "cfnetwork") && strings.Contains(ua, "darwin"):
			u.evalMac(ua)

		default:
			u.OS.Platform = PlatformUnknown
			u.OS.Name = OSUnknown
		}
	}

	return u.maybeBot(ua)
}

// maybeBot checks if the UserAgent is a bot and sets
// all bot related fields if it is
func (u *UserAgent) maybeBot(ua string) bool {
	if u.IsBot() {
		u.OS.Platform = PlatformBot
		u.OS.Name = OSBot
		u.evalDevice(ua)
		return true
	}
	return false
}

// evalLinux returns the `Platform`, `OSName` and Version of UAs with
// 'linux' listed as their platform.
func (u *UserAgent) evalLinux(ua string, agentPlatform string) {

	switch {

	// Android, Kindle Fire
	case strings.Contains(ua, "android") || strings.Contains(ua, "googletv"):
		// Android
		u.OS.Platform = PlatformLinux
		u.OS.Name = OSAndroid
		u.OS.Version.findVersionNumber(agentPlatform, "android ")

		// ChromeOS
	case strings.Contains(ua, "cros"):
		u.OS.Platform = PlatformLinux
		u.OS.Name = OSChromeOS

		// Linux, "Linux-like"
	case strings.Contains(ua, "x11") || strings.Contains(ua, "bsd") || strings.Contains(ua, "suse") || strings.Contains(ua, "debian") || strings.Contains(ua, "ubuntu"):
		u.OS.Platform = PlatformLinux
		u.OS.Name = OSLinux

	default:
		u.OS.Platform = PlatformLinux
		u.OS.Name = OSLinux
	}
}

// evaliOS returns the `Platform`, `OSName` and Version of UAs with
// 'ipad' or 'iphone' listed as their platform.
func (u *UserAgent) evaliOS(uaPlatform string, agentPlatform string) {

	switch uaPlatform {
	// iPhone
	case "iphone":
		u.OS.Platform = PlatformiPhone
		u.OS.Name = OSiOS
		u.OS.getiOSVersion(agentPlatform)

		// iPad
	case "ipad":
		u.OS.Platform = PlatformiPad
		u.OS.Name = OSiOS
		u.OS.getiOSVersion(agentPlatform)

		// iPod
	case "ipod touch", "ipod":
		u.OS.Platform = PlatformiPod
		u.OS.Name = OSiOS
		u.OS.getiOSVersion(agentPlatform)

	default:
		u.OS.Platform = PlatformiPad
		u.OS.Name = OSUnknown
	}
}

func (u *UserAgent) evalWindows(ua string) {

	switch {

	// No windows version
	case !strings.Contains(ua, "windows "):
		u.OS.Platform = PlatformWindows
		u.OS.Name = OSUnknown

	case strings.Contains(ua, "windows nt ") && u.OS.Version.findVersionNumber(ua, "windows nt "):
		u.OS.Platform = PlatformWindows
		u.OS.Name = OSWindows

	case strings.Contains(ua, "windows xp"):
		u.OS.Platform = PlatformWindows
		u.OS.Name = OSWindows
		u.OS.Version.Major = 5
		u.OS.Version.Minor = 1
		u.OS.Version.Patch = 0

	default:
		u.OS.Platform = PlatformWindows
		u.OS.Name = OSUnknown

	}
}

func (u *UserAgent) evalMac(uaPlatformGroup string) {
	u.OS.Platform = PlatformMac
	if i := strings.Index(uaPlatformGroup, "os x 10"); i != -1 {
		u.OS.Name = OSMacOS
		u.OS.Version.parse(uaPlatformGroup[i+5:])

		return
	}
	u.OS.Name = OSUnknown
}

func (v *Version) findVersionNumber(s string, m string) bool {
	if ind := strings.Index(s, m); ind != -1 {
		return v.parse(s[ind+len(m):])
	}
	return false
}

// getiOSVersion accepts the platform portion of a UA string and returns
// a Version.
func (o *OS) getiOSVersion(uaPlatformGroup string) {
	if i := strings.Index(uaPlatformGroup, "cpu iphone os "); i != -1 {
		o.Version.parse(uaPlatformGroup[i+14:])
		return
	}

	if i := strings.Index(uaPlatformGroup, "cpu os "); i != -1 {
		o.Version.parse(uaPlatformGroup[i+7:])
		return
	}

	o.Version.parse(uaPlatformGroup)
}

// strToInt simply accepts a string and returns a `int`,
// with '0' being default.
func strToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

// strToVer accepts a string and returns a Version,
// with {0, 0, 0} being default.
func (v *Version) parse(str string) bool {
	if len(str) == 0 || str[0] < '0' || str[0] > '9' {
		return false
	}
	for i := 0; i < 3; i++ {
		empty := true
		val := 0
		l := len(str) - 1

		for k, c := range str {
			if c >= '0' && c <= '9' {
				if empty {
					val = int(c) - 48
					empty = false
					if k == l {
						str = str[:0]
					}
					continue
				}

				if val == 0 {
					if c == '0' {
						if k == l {
							str = str[:0]
						}
						continue
					}
					str = str[k:]
					break
				}

				val = 10*val + int(c) - 48
				if k == l {
					str = str[:0]
				}
				continue
			}
			str = str[k+1:]
			break
		}

		switch i {
		case 0:
			v.Major = val

		case 1:
			v.Minor = val

		case 2:
			v.Patch = val
		}
	}
	return true
}
