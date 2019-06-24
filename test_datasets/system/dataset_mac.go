package test_datasets_system

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/useragent"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Mac = []test_datasets.TestDataset{
	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_5; en-us) AppleWebKit/525.26.2 (KHTML, like Gecko) Version/3.2 Safari/525.26.12",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{3, 2, 0}},
			useragent.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 5, 5}, "Leopard"},
			vars.DeviceComputer,
		},
	},
	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_2; en-US) AppleWebKit/533.1 (KHTML, like Gecko) Chrome/5.0.329.0 Safari/533.1",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserChrome, version.Version{5, 0, 329}},
			useragent.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 6, 2}, "Snow Leopard"},
			vars.DeviceComputer,
		},
	},
	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.1.6) Gecko/20091201 Firefox/3.5.6 (.NET CLR 3.5.30729)",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserFirefox, version.Version{3, 5, 6}},
			useragent.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 6, 0}, "Snow Leopard"},
			vars.DeviceComputer,
		},
	},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/534.52.7 (KHTML, like Gecko) Version/5.1.2 Safari/534.52.7",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{5, 1, 2}},
			useragent.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 7, 2}, "Lion"},
			vars.DeviceComputer,
		},
	},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/535.7 (KHTML, like Gecko) Chrome/16.0.912.75 Safari/535.7",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserChrome, version.Version{16, 0, 912}},
			useragent.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 7, 2}, "Lion"},
			vars.DeviceComputer,
		},
	},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8) AppleWebKit/535.18.5 (KHTML, like Gecko) Version/5.2 Safari/535.18.5",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{5, 2, 0}},
			useragent.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 8, 0}, "Mountain Lion"},
			vars.DeviceComputer,
		},
	},
	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_8; en-US) AppleWebKit/532.5 (KHTML, like Gecko) Chrome/4.0.249.0 Safari/532.5",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserChrome, version.Version{4, 0, 249}},
			useragent.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 8, 0}, "Mountain Lion"},
			vars.DeviceComputer,
		},
	},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9) AppleWebKit/537.35.1 (KHTML, like Gecko) Version/6.1 Safari/537.35.1",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{6, 1, 0}},
			useragent.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 9, 0}, "Mavericks"},
			vars.DeviceComputer,
		},
	},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10) AppleWebKit/538.34.48 (KHTML, like Gecko) Version/8.0 Safari/538.35.8",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{8, 0, 0}},
			useragent.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 10, 0}, "Yosemite"},
			vars.DeviceComputer,
		},
	},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10) AppleWebKit/538.32 (KHTML, like Gecko) Version/7.1 Safari/538.4",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{7, 1, 0}},
			useragent.OS{vars.PlatformMac, vars.OSMacOS, version.Version{10, 10, 0}, "Yosemite"},
			vars.DeviceComputer,
		},
	},
}
