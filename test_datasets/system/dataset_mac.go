package test_datasets_system

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Mac = []test_datasets.TestDataset{
	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_5; en-us) AppleWebKit/525.26.2 (KHTML, like Gecko) Version/3.2 Safari/525.26.12",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{3, 2, 0}},
			vars.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 5, 5}, "Leopard"},
			vars.DeviceComputer,
		},
		"OS X Leopard 10.5",
	},
	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_2; en-US) AppleWebKit/533.1 (KHTML, like Gecko) Chrome/5.0.329.0 Safari/533.1",
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{5, 0, 329}},
			vars.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 6, 2}, "Snow Leopard"},
			vars.DeviceComputer,
		},
		"OS X Snow Leopard 10.6",
	},
	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.1.6) Gecko/20091201 Firefox/3.5.6 (.NET CLR 3.5.30729)",
		vars.UserAgent{
			vars.Browser{vars.BrowserFirefox, version.Version{3, 5, 6}},
			vars.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 6, 0}, "Snow Leopard"},
			vars.DeviceComputer,
		},
		"OS X Snow Leopard 10.6",
	},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/534.52.7 (KHTML, like Gecko) Version/5.1.2 Safari/534.52.7",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{5, 1, 2}},
			vars.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 7, 2}, "Lion"},
			vars.DeviceComputer,
		},
		"OS X Lion 10.7",
	},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/535.7 (KHTML, like Gecko) Chrome/16.0.912.75 Safari/535.7",
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{16, 0, 912}},
			vars.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 7, 2}, "Lion"},
			vars.DeviceComputer,
		},
		"OS X Lion 10.7",
	},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8) AppleWebKit/535.18.5 (KHTML, like Gecko) Version/5.2 Safari/535.18.5",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{5, 2, 0}},
			vars.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 8, 0}, "Mountain Lion"},
			vars.DeviceComputer,
		},
		"OS X Mountain Lion 10.8",
	},
	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_8; en-US) AppleWebKit/532.5 (KHTML, like Gecko) Chrome/4.0.249.0 Safari/532.5",
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{4, 0, 249}},
			vars.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 8, 0}, "Mountain Lion"},
			vars.DeviceComputer,
		},
		"OS X Mountain Lion 10.8",
	},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9) AppleWebKit/537.35.1 (KHTML, like Gecko) Version/6.1 Safari/537.35.1",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{6, 1, 0}},
			vars.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 9, 0}, "Mavericks"},
			vars.DeviceComputer,
		},
		"OS X Mavericks 10.9",
	},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10) AppleWebKit/538.34.48 (KHTML, like Gecko) Version/8.0 Safari/538.35.8",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{8, 0, 0}},
			vars.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 10, 0}, "Yosemite"},
			vars.DeviceComputer,
		},
		"OS X Yosemite 10.10",
	},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10) AppleWebKit/538.32 (KHTML, like Gecko) Version/7.1 Safari/538.4",
		vars.UserAgent{
			vars.Browser{vars.BrowserSafari, version.Version{7, 1, 0}},
			vars.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 10, 0}, "Yosemite"},
			vars.DeviceComputer,
		},
		"OS X Yosemite 10.10",
	},
}
