package test_datasets_browser

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Safari = []test_datasets.TestDataset{
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/600.7.12 (KHTML, like Gecko) Version/8.0.7 Safari/600.7.12",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{8, 0, 7}},
			vars.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 10, 4}, "Yosemite"},
			vars.DeviceComputer,
		},
		"OS X Yosemite 10.10",
	},
	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_5; en-us) AppleWebKit/525.26.2 (KHTML, like Gecko) Version/3.2 Safari/525.26.12",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{3, 2, 0}},
			vars.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 5, 5}, "Leopard"},
			vars.DeviceComputer,
		},
		"OS X Leopard 10.5",
	},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12) AppleWebKit/602.1.32 (KHTML, like Gecko) Version/10.0 Safari/602.1.32", // macOS Sierra dev beta
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{10, 0, 0}},
			vars.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 12, 0}, "Sierra"},
			vars.DeviceComputer,
		},
		"macOS Sierra 10.12",
	},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 Safari/605.1.15",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{12, 0, 0}},
			vars.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 14, 0}, "Mojave"},
			vars.DeviceComputer,
		},
		"macOS Mojave 10.14",
	},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0.2 Safari/605.1.15",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{12, 0, 2}},
			vars.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 15, 1}, "Catalina"},
			vars.DeviceComputer,
		},
		"macOS Catalina 10.15",
	},
}
