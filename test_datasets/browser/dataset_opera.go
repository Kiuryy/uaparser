package test_datasets_browser

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Opera = []test_datasets.TestDataset{
	{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36 OPR/18.0.1284.68",
		vars.UserAgent{
			vars.Browser{vars.BrowserOpera, version.Version{18, 0, 1284}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{6, 1, 0}, "7"},
			vars.DeviceComputer,
		},
		"Windows 7",
	},
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 8_4 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) OPiOS/10.2.0.93022 Mobile/12H143 Safari/9537.53",
		vars.UserAgent{
			vars.Browser{vars.BrowserOpera, version.Version{10, 2, 0}},
			vars.OS{vars.PlatformiPhone, vars.OSiOS, version.Version{8, 4, 0}, ""},
			vars.DevicePhone,
		},
		"iOS 8.4",
	},
}
