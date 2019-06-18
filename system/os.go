package system

import (
	"github.com/Kiuryy/uaparser/const"
	"github.com/Kiuryy/uaparser/device"
	"github.com/Kiuryy/uaparser/userAgent"
	"github.com/Kiuryy/uaparser/version"
	"strings"
)

var macVersionAlias = map[version.Version]string{
	{10, 1, 0}:  "Puma",
	{10, 2, 0}:  "Jaguar",
	{10, 3, 0}:  "Panther",
	{10, 4, 0}:  "Tiger",
	{10, 5, 0}:  "Leopard",
	{10, 6, 0}:  "Snow Leopard",
	{10, 7, 0}:  "Lion",
	{10, 8, 0}:  "Mountain Lion",
	{10, 9, 0}:  "Mavericks",
	{10, 10, 0}: "Yosemite",
	{10, 11, 0}: "El Capitan",
	{10, 12, 0}: "Sierra",
	{10, 13, 0}: "High Sierra",
	{10, 14, 0}: "Mojave",
	{10, 15, 0}: "Catalina",
}

var windowsVersionAlias = map[version.Version]string{
	{6, 3, 0}: "8.1",
	{6, 2, 0}: "8",
	{6, 1, 0}: "7",
	{6, 0, 0}: "Vista",
	{5, 2, 0}: "XP",
	{5, 1, 0}: "XP",
	{5, 0, 0}: "2000",
}

func Eval(u *userAgent.UserAgent, ua string) bool {
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
			u.OS.Platform = _const.PlatformUnknown
			u.OS.Name = _const.OSUnknown
		}
	}

	return MaybeBot(u, ua)
}

// maybeBot checks if the UserAgent is a bot and sets
// all bot related fields if it is
func MaybeBot(u *userAgent.UserAgent, ua string) bool {
	if u.IsBot() {
		u.OS.Platform = _const.PlatformBot
		u.OS.Name = _const.OSBot
		device.Eval(u, ua)
		return true
	}
	return false
}

// evalLinux returns the `Platform`, `OSName` and Version of UAs with
// 'linux' listed as their platform.
func evalLinux(u *userAgent.UserAgent, ua string, agentPlatform string) {

	switch {

	// Android, Kindle Fire
	case strings.Contains(ua, "android") || strings.Contains(ua, "googletv"):
		// Android
		u.OS.Platform = _const.PlatformLinux
		u.OS.Name = _const.OSAndroid
		u.OS.Version.FindVersionNumber(agentPlatform, "android ")

		// ChromeOS
	case strings.Contains(ua, "cros"):
		u.OS.Platform = _const.PlatformLinux
		u.OS.Name = _const.OSChromeOS

		// Linux, "Linux-like"
	case strings.Contains(ua, "x11") || strings.Contains(ua, "bsd") || strings.Contains(ua, "suse") || strings.Contains(ua, "debian") || strings.Contains(ua, "ubuntu"):
		u.OS.Platform = _const.PlatformLinux
		u.OS.Name = _const.OSLinux

	default:
		u.OS.Platform = _const.PlatformLinux
		u.OS.Name = _const.OSLinux
	}
}

// evaliOS returns the `Platform`, `OSName` and Version of UAs with
// 'ipad' or 'iphone' listed as their platform.
func evaliOS(u *userAgent.UserAgent, uaPlatform string, agentPlatform string) {

	switch uaPlatform {
	// iPhone
	case "iphone":
		u.OS.Platform = _const.PlatformiPhone
		u.OS.Name = _const.OSiOS
		evaliOSVersion(u, agentPlatform)

		// iPad
	case "ipad":
		u.OS.Platform = _const.PlatformiPad
		u.OS.Name = _const.OSiOS
		evaliOSVersion(u, agentPlatform)

		// iPod
	case "ipod touch", "ipod":
		u.OS.Platform = _const.PlatformiPod
		u.OS.Name = _const.OSiOS
		evaliOSVersion(u, agentPlatform)

	default:
		u.OS.Platform = _const.PlatformiPad
		u.OS.Name = _const.OSUnknown
	}
}

// getiOSVersion accepts the platform portion of a UA string and returns a Version.
func evaliOSVersion(u *userAgent.UserAgent, uaPlatformGroup string) {
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

func evalWindows(u *userAgent.UserAgent, ua string) {

	switch {

	// No windows version
	case !strings.Contains(ua, "windows "):
		u.OS.Platform = _const.PlatformWindows
		u.OS.Name = _const.OSUnknown

	case strings.Contains(ua, "windows nt ") && u.OS.Version.FindVersionNumber(ua, "windows nt "):
		u.OS.Platform = _const.PlatformWindows
		u.OS.Name = _const.OSWindows

	case strings.Contains(ua, "windows xp"):
		u.OS.Platform = _const.PlatformWindows
		u.OS.Name = _const.OSWindows
		u.OS.Version.Major = 5
		u.OS.Version.Minor = 1
		u.OS.Version.Patch = 0

	default:
		u.OS.Platform = _const.PlatformWindows
		u.OS.Name = _const.OSUnknown

	}

	if versionAlias, ok := windowsVersionAlias[version.Version{u.OS.Version.Major, u.OS.Version.Minor, 0}]; ok {
		u.OS.VersionAlias = versionAlias
	}
}

func evalMac(u *userAgent.UserAgent, uaPlatformGroup string) {
	u.OS.Platform = _const.PlatformMac
	u.OS.Name = _const.OSUnknown

	if i := strings.Index(uaPlatformGroup, "os x 10"); i != -1 {
		u.OS.Name = _const.OSMacOS
		u.OS.Version.Parse(uaPlatformGroup[i+5:])

		if versionAlias, ok := macVersionAlias[version.Version{u.OS.Version.Major, u.OS.Version.Minor, 0}]; ok {
			u.OS.VersionAlias = versionAlias
		}
	}
}
