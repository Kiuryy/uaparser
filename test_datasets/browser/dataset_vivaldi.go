package test_datasets_browser

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/useragent"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Vivaldi = []test_datasets.TestDataset{
	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36 Vivaldi/2.1.1337.36",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserVivaldi, version.Version{2, 1, 1337}},
			useragent.OS{vars.PlatformWindows, vars.OSWindows, version.Version{10, 0, 0}, ""},
			vars.DeviceComputer,
		},
	},
	{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36 Vivaldi/1.2.490.43",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserVivaldi, version.Version{1, 2, 490}},
			useragent.OS{vars.PlatformLinux, vars.OSLinux, version.Version{0, 0, 0}, ""},
			vars.DeviceComputer,
		},
	},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.88 Safari/537.36 Vivaldi/2.4.1488.35",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserVivaldi, version.Version{2, 4, 1488}},
			useragent.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 14, 1}, "Mojave"},
			vars.DeviceComputer,
		},
	},
}
