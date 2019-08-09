package test_datasets_browser

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_QQ = []test_datasets.TestDataset{
	{"Mozilla/5.0 (Windows NT 6.2; WOW64; Trident/7.0; Touch; .NET4.0E; .NET4.0C; .NET CLR 3.5.30729; .NET CLR 2.0.50727; .NET CLR 3.0.30729; InfoPath.3; Tablet PC 2.0; QQBrowser/7.6.21433.400; rv:11.0) like Gecko",
		vars.UserAgent{
			vars.Browser{vars.BrowserQQ, version.Version{7, 6, 21433}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{6, 2, 0}, "8"},
			vars.DeviceTablet,
		},
		"Windows 8",
	},
	{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.124 Safari/537.36 QQBrowser/9.0.2191.400",
		vars.UserAgent{
			vars.Browser{vars.BrowserQQ, version.Version{9, 0, 2191}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{6, 1, 0}, "7"},
			vars.DeviceComputer,
		},
		"Windows 7",
	},
	{"mozilla/5.0 (iphone; cpu iphone os 8_1_2 like mac os x) applewebkit/600.1.4 (khtml, like gecko) mobile/12b440 qq/5.3.0.319 nettype/wifi mem/205",
		vars.UserAgent{
			vars.Browser{vars.BrowserQQ, version.Version{5, 3, 0}},
			vars.OS{vars.PlatformiPhone, vars.OSiOS, version.Version{8, 1, 2}, ""},
			vars.DevicePhone,
		},
		"iOS 8.1",
	},
}
