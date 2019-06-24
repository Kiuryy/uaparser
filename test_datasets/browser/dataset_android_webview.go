package test_datasets_browser

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/useragent"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_AndroidWebview = []test_datasets.TestDataset{
	{"Mozilla/5.0 (Linux; U; Android 2.2; en-us; DROID2 GLOBAL Build/S273) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{4, 0, 0}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{2, 2, 0}, "Froyo"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 4.0.3; de-ch; HTC Sensation Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari53/4.30",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{4, 0, 0}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{4, 0, 3}, "Ice Cream Sandwich"},
			vars.DevicePhone,
		},
	},
}
