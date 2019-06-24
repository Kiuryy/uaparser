package test_datasets_browser

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/useragent"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Edge = []test_datasets.TestDataset{
	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.123",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserEdge, version.Version{12, 123, 0}},
			useragent.OS{vars.PlatformWindows, vars.OSWindows, version.Version{10, 0, 0}, ""},
			vars.DeviceComputer,
		},
	},
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 12_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 EdgiOS/44.3.5 Mobile/15E148 Safari/605.1.15",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserEdge, version.Version{12, 0, 0}},
			useragent.OS{vars.PlatformiPhone, vars.OSiOS, version.Version{12, 3, 1}, ""},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (iPad; CPU OS 12_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 EdgiOS/44.3.2 Mobile/15E148 Safari/605.1.15",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserEdge, version.Version{12, 0, 0}},
			useragent.OS{vars.PlatformiPad, vars.OSiOS, version.Version{12, 3, 1}, ""},
			vars.DeviceTablet,
		},
	},
	{"Mozilla/5.0 (Linux; Android 9; motorola one) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.90 Mobile Safari/537.36 EdgA/42.0.2.3728",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserEdge, version.Version{42, 0, 2}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{9, 0, 0}, "Pie"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3800.0 Safari/537.36 Edg/76.0.172.0",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserEdge, version.Version{76, 0, 172}},
			useragent.OS{vars.PlatformWindows, vars.OSWindows, version.Version{10, 0, 0}, ""},
			vars.DeviceComputer,
		},
	},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3803.0 Safari/537.36 Edg/76.0.176.0",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserEdge, version.Version{76, 0, 176}},
			useragent.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 14, 5}, "Mojave"},
			vars.DeviceComputer,
		},
	},
}
