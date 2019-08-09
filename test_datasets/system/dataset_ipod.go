package test_datasets_system

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_iPod = []test_datasets.TestDataset{
	{"mozilla/5.0 (ipod touch; cpu iphone os 9_3_3 like mac os x) applewebkit/601.1.46 (khtml, like gecko) version/9.0 mobile/13g34 safari/601.1",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{9, 0, 0}},
			vars.OS{vars.PlatformiPod, vars.OSiOS, version.Version{9, 3, 3}, ""},
			vars.DeviceTablet,
		},
		"iOS 9.3",
	},
	{"mozilla/5.0 (ipod; cpu iphone os 6_1_6 like mac os x) applewebkit/536.26 (khtml, like gecko) version/6.0 mobile/10b500 safari/8536.25",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{6, 0, 0}},
			vars.OS{vars.PlatformiPod, vars.OSiOS, version.Version{6, 1, 6}, ""},
			vars.DeviceTablet,
		},
		"iOS 6.1",
	},
}
