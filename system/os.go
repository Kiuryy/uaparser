package system

import (
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
	"strings"
)

// Eval parses the user agent and sets the system information (OS, Platform)
func Eval(u *vars.UserAgent, ua string) {
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
		evalLinux(u, ua, agentPlatform)

	case "x11", "linux":
		evalLinux(u, ua, agentPlatform)

	case "ipad", "iphone", "ipod touch", "ipod":
		evaliOS(u, specs, agentPlatform)

	case "macintosh":
		evalMac(u, ua)

	default:
		switch {

		// Windows, Xbox
		case strings.Contains(ua, "windows ") || strings.Contains(ua, "microsoft-cryptoapi"):
			evalWindows(u, ua)

			// Linux, Android
		case strings.Contains(ua, "linux") || strings.Contains(ua, "android"):
			evalLinux(u, ua, agentPlatform)

			// Apple CFNetwork
		case strings.Contains(ua, "cfnetwork") && strings.Contains(ua, "darwin"):
			evalMac(u, ua)

		default:
			u.OS.Platform = vars.PlatformUnknown
			u.OS.Name = vars.OSUnknown
		}
	}
}

// evalLinux returns the `Platform`, `OSName` and Version of UAs with
// 'linux' listed as their platform.
func evalLinux(u *vars.UserAgent, ua string, agentPlatform string) {

	switch {

	// Android, Kindle Fire
	case strings.Contains(ua, "android") || strings.Contains(ua, "googletv"):
		// Android
		u.OS.Platform = vars.PlatformLinux
		u.OS.Name = vars.OSAndroid
		u.OS.Version.FindVersionNumber(agentPlatform, "android ")

		if versionAlias, ok := androidVersionAlias[version.Version{u.OS.Version.Major, u.OS.Version.Minor, 0}]; ok {
			u.OS.VersionAlias = versionAlias
		}

		// ChromeOS
	case strings.Contains(ua, "cros"):
		u.OS.Platform = vars.PlatformLinux
		u.OS.Name = vars.OSChromeOS
		u.OS.Version.FindVersionNumber(ua, "chrome/")

		// Linux, "Linux-like"
	case strings.Contains(ua, "x11") || strings.Contains(ua, "bsd") || strings.Contains(ua, "suse") || strings.Contains(ua, "debian") || strings.Contains(ua, "ubuntu"):
		u.OS.Platform = vars.PlatformLinux
		u.OS.Name = vars.OSLinux

	default:
		u.OS.Platform = vars.PlatformLinux
		u.OS.Name = vars.OSLinux
	}
}

// evaliOS returns the `Platform`, `OSName` and Version of UAs with
// 'ipad' or 'iphone' listed as their platform.
func evaliOS(u *vars.UserAgent, uaPlatform string, agentPlatform string) {

	switch uaPlatform {
	// iPhone
	case "iphone":
		u.OS.Platform = vars.PlatformiPhone
		u.OS.Name = vars.OSiOS
		evaliOSVersion(u, agentPlatform)

		// iPad
	case "ipad":
		u.OS.Platform = vars.PlatformiPad
		u.OS.Name = vars.OSiOS
		evaliOSVersion(u, agentPlatform)

		// iPod
	case "ipod touch", "ipod":
		u.OS.Platform = vars.PlatformiPod
		u.OS.Name = vars.OSiOS
		evaliOSVersion(u, agentPlatform)

	default:
		u.OS.Platform = vars.PlatformiPad
		u.OS.Name = vars.OSUnknown
	}
}

// getiOSVersion accepts the platform portion of a UA string and returns a Version.
func evaliOSVersion(u *vars.UserAgent, uaPlatformGroup string) {
	if i := strings.Index(uaPlatformGroup, "cpu iphone os "); i != -1 {
		u.OS.Version.Parse(uaPlatformGroup[i+14:])
		return
	}

	if i := strings.Index(uaPlatformGroup, "cpu os "); i != -1 {
		u.OS.Version.Parse(uaPlatformGroup[i+7:])
		return
	}

	u.OS.Version.Parse(uaPlatformGroup)
}

func evalWindows(u *vars.UserAgent, ua string) {

	switch {

	// No windows version
	case !strings.Contains(ua, "windows "):
		u.OS.Platform = vars.PlatformWindows
		u.OS.Name = vars.OSUnknown

	case strings.Contains(ua, "windows nt ") && u.OS.Version.FindVersionNumber(ua, "windows nt "):
		u.OS.Platform = vars.PlatformWindows
		u.OS.Name = vars.OSWindows

	case strings.Contains(ua, "windows xp"):
		u.OS.Platform = vars.PlatformWindows
		u.OS.Name = vars.OSWindows
		u.OS.Version.Major = 5
		u.OS.Version.Minor = 1
		u.OS.Version.Patch = 0

	default:
		u.OS.Platform = vars.PlatformWindows
		u.OS.Name = vars.OSUnknown

	}

	if versionAlias, ok := windowsVersionAlias[version.Version{u.OS.Version.Major, u.OS.Version.Minor, 0}]; ok {
		u.OS.VersionAlias = versionAlias
	}
}

func evalMac(u *vars.UserAgent, uaPlatformGroup string) {
	u.OS.Platform = vars.PlatformMac
	u.OS.Name = vars.OSUnknown

	if i := strings.Index(uaPlatformGroup, "os x 10"); i != -1 {
		u.OS.Name = vars.OSMacOS
		u.OS.Version.Parse(uaPlatformGroup[i+5:])

		if versionAlias, ok := macVersionAlias[version.Version{u.OS.Version.Major, u.OS.Version.Minor, 0}]; ok {
			u.OS.VersionAlias = versionAlias
		}
	}
}
