package test_datasets_system

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/useragent"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_ChromeOS = []test_datasets.TestDataset{
	{"Mozilla/5.0 (X11; U; CrOS i686 9.10.0; en-US) AppleWebKit/532.5 (KHTML, like Gecko) Chrome/4.0.253.0 Safari/532.5",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserChrome, version.Version{4, 0, 253}},
			useragent.OS{vars.PlatformLinux, vars.OSChromeOS, version.Version{0, 0, 0}, ""},
			vars.DeviceComputer,
		},
	},
}
