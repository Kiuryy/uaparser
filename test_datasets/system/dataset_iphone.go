package test_datasets_system

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/useragent"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_iPhone = []test_datasets.TestDataset{
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/546.10 (KHTML, like Gecko) Version/6.0 Mobile/7E18WD Safari/8536.25",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{6, 0, 0}},
			useragent.OS{vars.PlatformiPhone, vars.OSiOS, version.Version{7, 0, 0}, ""},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 8_0_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12A405 Safari/600.1.4",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{8, 0, 0}},
			useragent.OS{vars.PlatformiPhone, vars.OSiOS, version.Version{8, 0, 2}, ""},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_0 like Mac OS X; en-us) AppleWebKit/532.9 (KHTML, like Gecko) Version/4.0.5 Mobile/8A293 Safari/6531.22.7",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{4, 0, 5}},
			useragent.OS{vars.PlatformiPhone, vars.OSiOS, version.Version{4, 0, 0}, ""},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{5, 1, 0}},
			useragent.OS{vars.PlatformiPhone, vars.OSiOS, version.Version{5, 0, 0}, ""},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/546.10 (KHTML, like Gecko) Version/6.0 Mobile/7E18WD Safari/8536.25",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{6, 0, 0}},
			useragent.OS{vars.PlatformiPhone, vars.OSiOS, version.Version{7, 0, 0}, ""},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 10_2_1 like Mac OS X) AppleWebKit/602.4.6 (KHTML, like Gecko) Mobile/14D27 [FBAN/FBIOS;FBAV/86.0.0.48.52;FBBV/53842252;FBDV/iPhone9,1;FBMD/iPhone;FBSN/iOS;FBSV/10.2.1;FBSS/2;FBCR/Verizon;FBID/phone;FBLC/en_US;FBOP/5;FBRV/0]",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformiPhone, vars.OSiOS, version.Version{10, 2, 1}, ""},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 8_0_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12A405 Safari/600.1.4",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{8, 0, 0}},
			useragent.OS{vars.PlatformiPhone, vars.OSiOS, version.Version{8, 0, 2}, ""},
			vars.DevicePhone,
		},
	},
}
