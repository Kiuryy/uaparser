package test_datasets_system

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_iPad = []test_datasets.TestDataset{
	{"Mozilla/5.0(iPad; U; CPU iPhone OS 3_2 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Version/4.0.4 Mobile/7B314 Safari/531.21.10",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{4, 0, 4}},
			vars.OS{vars.PlatformiPad, vars.OSiOS, version.Version{3, 2, 0}, ""},
			vars.DeviceTablet,
		},
		"iOS 3.2",
	},
	{"Mozilla/5.0 (iPad; CPU OS 9_0 like Mac OS X) AppleWebKit/601.1.17 (KHTML, like Gecko) Version/8.0 Mobile/13A175 Safari/600.1.4",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{8, 0, 0}},
			vars.OS{vars.PlatformiPad, vars.OSiOS, version.Version{9, 0, 0}, ""},
			vars.DeviceTablet,
		},
		"iOS 9",
	},
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 10_0 like Mac OS X) AppleWebKit/602.1.32 (KHTML, like Gecko) Version/10.0 Mobile/14A5261v Safari/602.1",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{10, 0, 0}},
			vars.OS{vars.PlatformiPhone, vars.OSiOS, version.Version{10, 0, 0}, ""},
			vars.DevicePhone,
		},
		"iOS 10",
	},
	{"Mozilla/5.0(iPad; U; CPU iPhone OS 3_2 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Version/4.0.4 Mobile/7B314 Safari/531.21.10",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{4, 0, 4}},
			vars.OS{vars.PlatformiPad, vars.OSiOS, version.Version{3, 2, 0}, ""},
			vars.DeviceTablet,
		},
		"iOS 3.2",
	},
	{"Mozilla/5.0 (iPad; CPU OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{5, 1, 0}},
			vars.OS{vars.PlatformiPad, vars.OSiOS, version.Version{5, 0, 0}, ""},
			vars.DeviceTablet,
		},
		"iOS 5",
	},
	{"Mozilla/5.0 (iPad; CPU OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5355d Safari/8536.25",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{6, 0, 0}},
			vars.OS{vars.PlatformiPad, vars.OSiOS, version.Version{6, 0, 0}, ""},
			vars.DeviceTablet,
		},
		"iOS 6",
	},
	{"Mozilla/5.0 (iPad; CPU OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{7, 0, 0}},
			vars.OS{vars.PlatformiPad, vars.OSiOS, version.Version{7, 0, 0}, ""},
			vars.DeviceTablet,
		},
		"iOS 7",
	},
	{"Mozilla/5.0 (iPad; CPU OS 7_0_2 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A501 Safari/9537.53",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{7, 0, 0}},
			vars.OS{vars.PlatformiPad, vars.OSiOS, version.Version{7, 0, 2}, ""},
			vars.DeviceTablet,
		},
		"iOS 7",
	},
}
