package test_datasets_browser

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Firefox = []test_datasets.TestParserDataset{
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) FxiOS/1.0 Mobile/12F69 Safari/600.1.4",
		vars.UserAgent{
			vars.Browser{vars.BrowserFirefox, version.Version{1, 0, 0}},
			vars.OS{vars.PlatformiPhone, vars.OSiOS, version.Version{8, 3, 0}, ""},
			vars.DevicePhone,
		},
		"iOS 8.3",
	},
	{"Mozilla/5.0 (Android 4.4; Tablet; rv:41.0) Gecko/41.0 Firefox/41.0",
		vars.UserAgent{
			vars.Browser{vars.BrowserFirefox, version.Version{41, 0, 0}},
			vars.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{4, 4, 0}, "KitKat"},
			vars.DeviceTablet,
		},
		"Android 4.4 KitKat",
	},
	{"Mozilla/5.0 (Android; Mobile; rv:40.0) Gecko/40.0 Firefox/40.0",
		vars.UserAgent{
			vars.Browser{vars.BrowserFirefox, version.Version{40, 0, 0}},
			vars.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{0, 0, 0}, ""},
			vars.DevicePhone,
		},
		"Android",
	},
	{"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:38.0) Gecko/20100101 Firefox/38.0",
		vars.UserAgent{
			vars.Browser{vars.BrowserFirefox, version.Version{38, 0, 0}},
			vars.OS{vars.PlatformLinux, vars.OSLinux, version.Version{0, 0, 0}, ""},
			vars.DeviceComputer,
		},
		"Linux",
	},
}
