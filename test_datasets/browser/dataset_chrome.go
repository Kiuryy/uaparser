package test_datasets_browser

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Chrome = []test_datasets.TestParserDataset{
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.130 Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{43, 0, 2357}},
			vars.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 10, 4}, "Yosemite"},
			vars.DeviceComputer,
		},
		"OS X Yosemite 10.10",
	},
	{"Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; en) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/534.48.3",
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{19, 0, 1084}},
			vars.OS{vars.PlatformiPhone, vars.OSiOS, version.Version{5, 1, 1}, ""},
			vars.DevicePhone,
		},
		"iOS 5.1",
	},
	{"Mozilla/5.0 (Linux; Android 6.0; Nexus 5X Build/MDB08L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.76 Mobile Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{46, 0, 2490}},
			vars.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{6, 0, 0}, "Marshmallow"},
			vars.DevicePhone,
		},
		"Android 6 Marshmallow",
	},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{45, 0, 2454}},
			vars.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 11, 0}, "El Capitan"},
			vars.DeviceComputer,
		},
		"OS X El Capitan 10.11",
	},
	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{74, 0, 3729}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{10, 0, 0}, ""},
			vars.DeviceComputer,
		},
		"Windows 10",
	},
	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.80 Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{75, 0, 3770}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{10, 0, 0}, ""},
			vars.DeviceComputer,
		},
		"Windows 10",
	},
	{"Mozilla/5.0 (X11; CrOS x86_64 11895.95.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.125 Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{74, 0, 3729}},
			vars.OS{vars.PlatformLinux, vars.OSChromeOS, version.Version{74, 0, 3729}, ""},
			vars.DeviceComputer,
		},
		"ChromeOS 74",
	},
	{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{91, 0, 4472}},
			vars.OS{vars.PlatformLinux, vars.OSLinux, version.Version{0, 0, 0}, ""},
			vars.DeviceComputer,
		},
		"Linux",
	},
}
