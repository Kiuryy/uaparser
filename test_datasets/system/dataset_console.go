package test_datasets_system

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Console = []test_datasets.TestParserDataset{
	{"Mozilla/5.0 (Nintendo Switch; WifiWebAuthApplet) AppleWebKit/601.6 (KHTML, like Gecko) NF/4.0.0.5.9 NintendoBrowser/5.1.0.13341",
		vars.UserAgent{
			vars.Browser{vars.BrowserUnknown, version.Version{0, 0, 0}},
			vars.OS{vars.PlatformUnknown, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DeviceConsole,
		},
		"",
	},
	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0; Xbox; Xbox One)",
		vars.UserAgent{
			vars.Browser{vars.BrowserIE, version.Version{10, 0, 0}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{6, 2, 0}, "8"},
			vars.DeviceConsole,
		},
		"Windows 8",
	},
	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36 Edge/15.15063",
		vars.UserAgent{
			vars.Browser{vars.BrowserEdge, version.Version{15, 15063, 0}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{10, 0, 0}, ""},
			vars.DeviceConsole,
		},
		"Windows 10",
	},
	{"Mozilla/5.0 (PlayStation 4 5.01) AppleWebKit/601.2 (KHTML, like Gecko)",
		vars.UserAgent{
			vars.Browser{vars.BrowserUnknown, version.Version{0, 0, 0}},
			vars.OS{vars.PlatformUnknown, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DeviceConsole,
		},
		"",
	},
}
