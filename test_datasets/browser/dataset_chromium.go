package test_datasets_browser

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Chromium = []test_datasets.TestParserDataset{
	{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/69.0.3497.81 Chrome/69.0.3497.81 Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserChromium, version.Version{69, 0, 3497}},
			vars.OS{vars.PlatformLinux, vars.OSLinux, version.Version{0, 0, 0}, ""},
			vars.DeviceComputer,
		},
		"Linux",
	},
	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/75.0.3731.0 Chrome/75.0.3731.0 Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserChromium, version.Version{75, 0, 3731}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{10, 0, 0}, ""},
			vars.DeviceComputer,
		},
		"Windows 10",
	},
}
