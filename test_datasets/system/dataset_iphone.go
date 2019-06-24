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
}
