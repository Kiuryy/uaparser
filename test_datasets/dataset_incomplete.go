package test_datasets

import (
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Incomplete = []TestParserDataset{
	{"Mozilla/5.0 (iPad; CPU OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) GSA/6.0.51363 Mobile/12F69 Safari/600.1.4", // Google search app (GSA) for iOS -- it's Safari in disguise as of v6
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{0, 0, 0}},
			vars.OS{vars.PlatformiPad, vars.OSiOS, version.Version{8, 3, 0}, ""},
			vars.DeviceTablet,
		},
		"iOS 8.3",
	},
	{"Microsoft-CryptoAPI/10.0", // OCSP fetcher
		vars.UserAgent{
			vars.Browser{vars.BrowserUnknown, version.Version{0, 0, 0}},
			vars.OS{vars.PlatformWindows, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DeviceComputer,
		},
		"",
	},
	{"trustd (unknown version) CFNetwork/811.7.2 Darwin/16.7.0 (x86_64)", // OCSP fetcher
		vars.UserAgent{
			vars.Browser{vars.BrowserUnknown, version.Version{0, 0, 0}},
			vars.OS{vars.PlatformMac, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DeviceComputer,
		},
		"",
	},
	{"ocspd (unknown version) CFNetwork/520.5.3 Darwin/11.4.2 (x86_64)(MacBookAir5%2C2)", // OCSP fetcher
		vars.UserAgent{
			vars.Browser{vars.BrowserUnknown, version.Version{0, 0, 0}},
			vars.OS{vars.PlatformMac, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DeviceComputer,
		},
		"",
	},
	{"Mozilla/5.0 (Mobile; rv:26.0) Gecko/26.0 Firefox/26.0", //firefox OS
		vars.UserAgent{
			vars.Browser{vars.BrowserFirefox, version.Version{26, 0, 0}},
			vars.OS{vars.PlatformUnknown, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DevicePhone,
		},
		"",
	},
	{"Opera/9.80 (S60; SymbOS; Opera Mobi/352; U; de) Presto/2.4.15 Version/10.00",
		vars.UserAgent{
			vars.Browser{vars.BrowserOpera, version.Version{10, 0, 0}},
			vars.OS{vars.PlatformUnknown, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DevicePhone,
		},
		"",
	},

	// Can't parse Browser due to limitation of user agent library

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.7; rv:9.0) Gecko/20111222 Thunderbird/9.0.1",
		vars.UserAgent{
			vars.Browser{vars.BrowserUnknown, version.Version{0, 0, 0}},
			vars.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 7, 0}, "Lion"},
			vars.DeviceComputer,
		},
		"OS X Lion 10.7",
	},
	{"Opera/9.80 (S60; SymbOS; Opera Mobi/352; U; de) Presto/2.4.15 Version/10.00",
		vars.UserAgent{
			vars.Browser{vars.BrowserOpera, version.Version{10, 0, 0}},
			vars.OS{vars.PlatformUnknown, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DevicePhone,
		},
		"",
	},
	{"Opera/9.80 (S60; SymbOS; Opera Mobi/352; U; de) Presto/2.4.15 Version/10.00",
		vars.UserAgent{
			vars.Browser{vars.BrowserOpera, version.Version{10, 0, 0}},
			vars.OS{vars.PlatformUnknown, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DevicePhone,
		},
		"",
	},

	// Incomplete user agents

	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/",
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{0, 0, 0}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{10, 0, 0}, ""},
			vars.DeviceComputer,
		},
		"Windows 10",
	},
	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko)",
		vars.UserAgent{
			vars.Browser{vars.BrowserUnknown, version.Version{0, 0, 0}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{10, 0, 0}, ""},
			vars.DeviceComputer,
		},
		"Windows 10",
	},
	{"Mozilla/5.0 (Windows NT 6.1; Win64; x64)",
		vars.UserAgent{
			vars.Browser{vars.BrowserUnknown, version.Version{0, 0, 0}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{6, 1, 0}, "7"},
			vars.DeviceComputer,
		},
		"Windows 7",
	},
	{"Mozilla/5.0 (X11; Linux x86_64)",
		vars.UserAgent{
			vars.Browser{vars.BrowserUnknown, version.Version{0, 0, 0}},
			vars.OS{vars.PlatformLinux, vars.OSLinux, version.Version{0, 0, 0}, ""},
			vars.DeviceComputer,
		},
		"Linux",
	},
	{"Mozilla/5.0 (X11;",
		vars.UserAgent{
			vars.Browser{vars.BrowserUnknown, version.Version{0, 0, 0}},
			vars.OS{vars.PlatformUnknown, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DeviceUnknown,
		},
		"",
	},
	{"Mozilla/5.0",
		vars.UserAgent{
			vars.Browser{vars.BrowserUnknown, version.Version{0, 0, 0}},
			vars.OS{vars.PlatformUnknown, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DeviceUnknown,
		},
		"",
	},
	{"Mozilla",
		vars.UserAgent{
			vars.Browser{vars.BrowserUnknown, version.Version{0, 0, 0}},
			vars.OS{vars.PlatformUnknown, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DeviceUnknown,
		},
		"",
	},
	{"Chrome",
		vars.UserAgent{
			vars.Browser{vars.BrowserUnknown, version.Version{0, 0, 0}},
			vars.OS{vars.PlatformUnknown, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DeviceUnknown,
		},
		"",
	},
}
