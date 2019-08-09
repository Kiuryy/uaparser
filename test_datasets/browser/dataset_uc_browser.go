package test_datasets_browser

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_UCBrowser = []test_datasets.TestDataset{
	{"Mozilla/5.0 (Linux; U; Android 2.3.4; en-US; MT11i Build/4.0.2.A.0.62) AppleWebKit/534.31 (KHTML, like Gecko) UCBrowser/9.0.1.275 U3/0.8.0 Mobile Safari/534.31",
		vars.UserAgent{
			vars.Browser{vars.BrowserUCBrowser, version.Version{9, 0, 1}},
			vars.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{2, 3, 4}, "Gingerbread"},
			vars.DevicePhone,
		},
		"Android 2.3 Gingerbread",
	},
	{"Mozilla/5.0 (Linux; U; Android 4.0.4; en-US; Micromax P255 Build/IMM76D) AppleWebKit/534.31 (KHTML, like Gecko) UCBrowser/9.2.0.308 U3/0.8.0 Mobile Safari/534.31",
		vars.UserAgent{
			vars.Browser{vars.BrowserUCBrowser, version.Version{9, 2, 0}},
			vars.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{4, 0, 4}, "Ice Cream Sandwich"},
			vars.DevicePhone,
		},
		"Android 4 Ice Cream Sandwich",
	},
	{"UCWEB/2.0 (Java; U; MIDP-2.0; en-US; MicromaxQ5) U2/1.0.0 UCBrowser/9.4.0.342 U2/1.0.0 Mobile",
		vars.UserAgent{
			vars.Browser{vars.BrowserUCBrowser, version.Version{9, 4, 0}},
			vars.OS{vars.PlatformUnknown, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DevicePhone,
		},
		"",
	},
}
