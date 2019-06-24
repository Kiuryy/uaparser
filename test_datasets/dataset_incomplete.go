package test_datasets

import (
	"github.com/Kiuryy/uaparser/useragent"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Incomplete = []TestDataset{
	{"Mozilla/5.0 (iPad; CPU OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) GSA/6.0.51363 Mobile/12F69 Safari/600.1.4", // Google search app (GSA) for iOS -- it's Safari in disguise as of v6
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformiPad, vars.OSiOS, version.Version{8, 3, 0}, ""},
			vars.DeviceTablet,
		},
	},
	{"Microsoft-CryptoAPI/10.0", // OCSP fetcher
		useragent.UserAgent{
			useragent.Browser{vars.BrowserUnknown, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformWindows, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DeviceComputer,
		},
	}, {"trustd (unknown version) CFNetwork/811.7.2 Darwin/16.7.0 (x86_64)", // OCSP fetcher
		useragent.UserAgent{
			useragent.Browser{vars.BrowserUnknown, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformMac, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DeviceComputer,
		},
	}, {"ocspd (unknown version) CFNetwork/520.5.3 Darwin/11.4.2 (x86_64)(MacBookAir5%2C2)", // OCSP fetcher
		useragent.UserAgent{
			useragent.Browser{vars.BrowserUnknown, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformMac, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DeviceComputer,
		},
	},
	{"Mozilla/5.0 (Mobile; rv:26.0) Gecko/26.0 Firefox/26.0", //firefox OS
		useragent.UserAgent{
			useragent.Browser{vars.BrowserFirefox, version.Version{26, 0, 0}},
			useragent.OS{vars.PlatformUnknown, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DevicePhone,
		},
	},
	{"Opera/9.80 (S60; SymbOS; Opera Mobi/352; U; de) Presto/2.4.15 Version/10.00",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserOpera, version.Version{10, 0, 0}},
			useragent.OS{vars.PlatformUnknown, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DevicePhone,
		},
	},

	// Can't parse Browser due to limitation of user agent library

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.7; rv:9.0) Gecko/20111222 Thunderbird/9.0.1",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserUnknown, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 7, 0}, "Lion"},
			vars.DeviceComputer,
		},
	},
	{"Opera/9.80 (S60; SymbOS; Opera Mobi/352; U; de) Presto/2.4.15 Version/10.00",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserOpera, version.Version{10, 0, 0}},
			useragent.OS{vars.PlatformUnknown, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DevicePhone,
		},
	},
	{"Opera/9.80 (S60; SymbOS; Opera Mobi/352; U; de) Presto/2.4.15 Version/10.00",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserOpera, version.Version{10, 0, 0}},
			useragent.OS{vars.PlatformUnknown, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DevicePhone,
		},
	},
}
