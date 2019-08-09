package test_datasets_browser

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_IE = []test_datasets.TestDataset{
	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)",
		vars.UserAgent{
			vars.Browser{vars.BrowserIE, version.Version{10, 0, 0}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{6, 2, 0}, "8"},
			vars.DeviceComputer,
		},
		"Windows 8",
	},
	{"Mozilla/5.0 (Windows NT 6.3; Trident/7.0; .NET4.0E; .NET4.0C; rv:11.0) like Gecko",
		vars.UserAgent{
			vars.Browser{vars.BrowserIE, version.Version{11, 0, 0}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{6, 3, 0}, "8.1"},
			vars.DeviceComputer,
		},
		"Windows 8.1",
	},
	{"Mozilla/4.0 (compatible; MSIE 5.01; Windows NT 5.0; SV1; .NET CLR 1.1.4322; .NET CLR 1.0.3705; .NET CLR 2.0.50727)",
		vars.UserAgent{
			vars.Browser{vars.BrowserIE, version.Version{5, 0, 1}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{5, 0, 0}, "2000"},
			vars.DeviceComputer,
		},
		"Windows 2000",
	},
	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/4.0; GTB6.4; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; OfficeLiveConnector.1.3; OfficeLivePatch.0.0; .NET CLR 1.1.4322)",
		vars.UserAgent{
			vars.Browser{vars.BrowserIE, version.Version{7, 0, 0}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{6, 1, 0}, "7"},
			vars.DeviceComputer,
		},
		"Windows 7",
	},
	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0; Touch)", //Windows Surface RT tablet
		vars.UserAgent{
			vars.Browser{vars.BrowserIE, version.Version{10, 0, 0}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{6, 2, 0}, "8"},
			vars.DeviceTablet,
		},
		"Windows 8",
	},
}
