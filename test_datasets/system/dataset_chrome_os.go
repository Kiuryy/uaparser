package test_datasets_system

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_ChromeOS = []test_datasets.TestParserDataset{
	{"Mozilla/5.0 (X11; CrOS x86_64 11151.113.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.127 Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{71, 0, 3578}},
			vars.OS{vars.PlatformLinux, vars.OSChromeOS, version.Version{71, 0, 3578}, ""},
			vars.DeviceComputer,
		},
		"ChromeOS 71",
	},
	{"Mozilla/5.0 (X11; CrOS x86_64 11647.154.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.114 Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{73, 0, 3683}},
			vars.OS{vars.PlatformLinux, vars.OSChromeOS, version.Version{73, 0, 3683}, ""},
			vars.DeviceComputer,
		},
		"ChromeOS 73",
	},
	{"Mozilla/5.0 (X11; CrOS armv7l 5500.100.6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.120 Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{34, 0, 1847}},
			vars.OS{vars.PlatformLinux, vars.OSChromeOS, version.Version{34, 0, 1847}, ""},
			vars.DeviceComputer,
		},
		"ChromeOS 34",
	},
	{"Mozilla/5.0 (X11; CrOS x86_64 11316.98.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.74 Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{72, 0, 3626}},
			vars.OS{vars.PlatformLinux, vars.OSChromeOS, version.Version{72, 0, 3626}, ""},
			vars.DeviceComputer,
		},
		"ChromeOS 72",
	},
	{"Mozilla/5.0 (X11; CrOS x86_64 12105.42.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.42 Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{75, 0, 3770}},
			vars.OS{vars.PlatformLinux, vars.OSChromeOS, version.Version{75, 0, 3770}, ""},
			vars.DeviceComputer,
		},
		"ChromeOS 75",
	},
}
