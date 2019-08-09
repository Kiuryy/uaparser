package test_datasets_system

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Windows = []test_datasets.TestDataset{
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
	{"Mozilla/5.0 (Windows; U; Windows NT 6.1; sk; rv:1.9.1.7) Gecko/20091221 Firefox/3.5.7",
		vars.UserAgent{
			vars.Browser{vars.BrowserFirefox, version.Version{3, 5, 7}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{6, 1, 0}, "7"},
			vars.DeviceComputer,
		},
		"Windows 7",
	},
	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)",
		vars.UserAgent{
			vars.Browser{vars.BrowserIE, version.Version{10, 0, 0}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{6, 2, 0}, "8"},
			vars.DeviceComputer,
		},
		"Windows 8",
	},
	{"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.15 (KHTML, like Gecko) Chrome/24.0.1295.0 Safari/537.15",
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{24, 0, 1295}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{6, 2, 0}, "8"},
			vars.DeviceComputer,
		},
		"Windows 8",
	},
	{"Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko",
		vars.UserAgent{
			vars.Browser{vars.BrowserIE, version.Version{11, 0, 0}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{6, 3, 0}, "8.1"},
			vars.DeviceTablet,
		},
		"Windows 8.1",
	},
	{"Mozilla/5.0 (IE 11.0; Windows NT 6.3; Trident/7.0; .NET4.0E; .NET4.0C; rv:11.0) like Gecko",
		vars.UserAgent{
			vars.Browser{vars.BrowserIE, version.Version{11, 0, 0}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{6, 3, 0}, "8.1"},
			vars.DeviceComputer,
		},
		"Windows 8.1",
	},
	{"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; SLCC1; .NET CLR 2.0.50727; .NET CLR 1.1.4322; InfoPath.2; .NET CLR 3.5.21022; .NET CLR 3.5.30729; MS-RTC LM 8; OfficeLiveConnector.1.4; OfficeLivePatch.1.3; .NET CLR 3.0.30729)",
		vars.UserAgent{
			vars.Browser{vars.BrowserIE, version.Version{8, 0, 0}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{6, 0, 0}, "Vista"},
			vars.DeviceComputer,
		},
		"Windows Vista",
	},
	{"Mozilla/5.0 (Windows; U; Windows NT 5.1; cs; rv:1.9.1.8) Gecko/20100202 Firefox/3.5.8",
		vars.UserAgent{
			vars.Browser{vars.BrowserFirefox, version.Version{3, 5, 8}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{5, 1, 0}, "XP"},
			vars.DeviceComputer,
		},
		"Windows XP",
	},
	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; )",
		vars.UserAgent{
			vars.Browser{vars.BrowserIE, version.Version{7, 0, 0}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{5, 1, 0}, "XP"},
			vars.DeviceComputer,
		},
		"Windows XP",
	},
}
