package test_datasets_browser

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Edge = []test_datasets.TestParserDataset{
	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.123",
		vars.UserAgent{
			vars.Browser{vars.BrowserEdge, version.Version{12, 123, 0}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{10, 0, 0}, ""},
			vars.DeviceComputer,
		},
		"Windows 10",
	},
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 12_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 EdgiOS/44.3.5 Mobile/15E148 Safari/605.1.15",
		vars.UserAgent{
			vars.Browser{vars.BrowserEdge, version.Version{12, 0, 0}},
			vars.OS{vars.PlatformiPhone, vars.OSiOS, version.Version{12, 3, 1}, ""},
			vars.DevicePhone,
		},
		"iOS 12.3",
	},
	{"Mozilla/5.0 (iPad; CPU OS 12_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 EdgiOS/44.3.2 Mobile/15E148 Safari/605.1.15",
		vars.UserAgent{
			vars.Browser{vars.BrowserEdge, version.Version{12, 0, 0}},
			vars.OS{vars.PlatformiPad, vars.OSiOS, version.Version{12, 3, 1}, ""},
			vars.DeviceTablet,
		},
		"iOS 12.3",
	},
	{"Mozilla/5.0 (Linux; Android 9; motorola one) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.90 Mobile Safari/537.36 EdgA/42.0.2.3728",
		vars.UserAgent{
			vars.Browser{vars.BrowserEdge, version.Version{42, 0, 2}},
			vars.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{9, 0, 0}, "Pie"},
			vars.DevicePhone,
		},
		"Android 9 Pie",
	},
	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3800.0 Safari/537.36 Edg/76.0.172.0",
		vars.UserAgent{
			vars.Browser{vars.BrowserEdge, version.Version{76, 0, 172}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{10, 0, 0}, ""},
			vars.DeviceComputer,
		},
		"Windows 10",
	},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3803.0 Safari/537.36 Edg/76.0.176.0",
		vars.UserAgent{
			vars.Browser{vars.BrowserEdge, version.Version{76, 0, 176}},
			vars.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 14, 5}, "Mojave"},
			vars.DeviceComputer,
		},
		"macOS Mojave 10.14",
	},
}
