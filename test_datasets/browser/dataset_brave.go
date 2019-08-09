package test_datasets_browser

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Brave = []test_datasets.TestDataset{
	{"Mozilla/5.0 (Linux; Android 9; ONEPLUS A5010) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.157 Mobile Safari/537.36 Brave/74",
		vars.UserAgent{
			vars.Browser{vars.BrowserBrave, version.Version{74, 0, 0}},
			vars.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{9, 0, 0}, "Pie"},
			vars.DevicePhone,
		},
		"Android 9 Pie",
	},
	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.38 Safari/537.36 Brave/75",
		vars.UserAgent{
			vars.Browser{vars.BrowserBrave, version.Version{75, 0, 0}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{10, 0, 0}, ""},
			vars.DeviceComputer,
		},
		"Windows 10",
	},
	{"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) brave/0.7.10 Chrome/47.0.2526.110 Brave/0.36.5 Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserBrave, version.Version{0, 7, 10}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{6, 1, 0}, "7"},
			vars.DeviceComputer,
		},
		"Windows 7",
	},
}
