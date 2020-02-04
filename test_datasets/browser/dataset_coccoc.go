package test_datasets_browser

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_CocCoc = []test_datasets.TestParserDataset{
	{"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) coc_coc_browser/42.0 CoRom/36.0.1985.144 Chrome/36.0.1985.144 Safari/537.36",
		vars.UserAgent{
			vars.Browser{vars.BrowserCocCoc, version.Version{42, 0, 0}},
			vars.OS{vars.PlatformWindows, vars.OSWindows, version.Version{6, 3, 0}, "8.1"},
			vars.DeviceComputer,
		},
		"Windows 8.1",
	},
}
