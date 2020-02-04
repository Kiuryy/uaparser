package test_datasets_browser

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Samsung = []test_datasets.TestParserDataset{
	{"Mozilla/5.0 (Linux; Android 7.0; SAMSUNG SM-G570M Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/7.4 Chrome/59.0.3071.125 Mobile Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserSamsung, version.Version{7, 4, 0}},
			vars.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{7, 0, 0}, "Nougat"},
			vars.DevicePhone,
		},
		"Android 7 Nougat",
	},
	{"Mozilla/5.0 (Linux; Android 9; SAMSUNG SM-G960F Build/PPR1.180610.011) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/9.2 Chrome/67.0.3396.87 Mobile Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserSamsung, version.Version{9, 2, 0}},
			vars.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{9, 0, 0}, "Pie"},
			vars.DevicePhone,
		},
		"Android 9 Pie",
	},
	{"Mozilla/5.0 (Linux; Android 5.1.1; SAMSUNG SM-J320M Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/6.4 Chrome/56.0.2924.87 Mobile Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserSamsung, version.Version{6, 4, 0}},
			vars.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{5, 1, 1}, "Lollipop"},
			vars.DevicePhone,
		},
		"Android 5.1 Lollipop",
	},
	{"Mozilla/5.0 (Linux; Android 8.1.0; SAMSUNG SM-N9600 Build/M1AJQ) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/8.0 Chrome/63.0.3239.111 Mobile Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserSamsung, version.Version{8, 0, 0}},
			vars.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{8, 1, 0}, "Oreo"},
			vars.DevicePhone,
		},
		"Android 8.1 Oreo",
	},
}
