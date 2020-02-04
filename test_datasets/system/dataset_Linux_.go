package test_datasets_system

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Linux = []test_datasets.TestParserDataset{
	{"Mozilla/5.0 (X11; U; Linux x86_64; zh-TW; rv:1.9.0.8) Gecko/2009032712 Ubuntu/8.04 (hardy) Firefox/3.0.8 GTB5",
		vars.UserAgent{
			vars.Browser{vars.BrowserFirefox, version.Version{3, 0, 8}},
			vars.OS{vars.PlatformLinux, vars.OSLinux, version.Version{0, 0, 0}, ""},
			vars.DeviceComputer,
		},
		"Linux",
	},
	{"Mozilla/5.0 (X11; U; Linux i686; de; rv:1.9.1.5) Gecko/20091112 Iceweasel/3.5.5 (like Firefox/3.5.5; Debian-3.5.5-1)",
		vars.UserAgent{
			vars.Browser{vars.BrowserFirefox, version.Version{3, 5, 5}},
			vars.OS{vars.PlatformLinux, vars.OSLinux, version.Version{0, 0, 0}, ""},
			vars.DeviceComputer,
		},
		"Linux",
	},
	{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.45 Safari/535.19",
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{18, 0, 1025}},
			vars.OS{vars.PlatformLinux, vars.OSLinux, version.Version{0, 0, 0}, ""},
			vars.DeviceComputer,
		},
		"Linux",
	},
	{"Mozilla/5.0 (X11; U; Linux x86_64; en; rv:1.9.0.14) Gecko/20080528 Ubuntu/9.10 (karmic) Epiphany/2.22 Firefox/3.0",
		vars.UserAgent{
			vars.Browser{vars.BrowserFirefox, version.Version{3, 0, 0}},
			vars.OS{vars.PlatformLinux, vars.OSLinux, version.Version{0, 0, 0}, ""},
			vars.DeviceComputer,
		},
		"Linux",
	},
}
