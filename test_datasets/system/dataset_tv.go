package test_datasets_system

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_TV = []test_datasets.TestParserDataset{
	{"Mozilla/5.0 (CrKey armv7l 1.4.15250) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.0 Safari/537.36", // Chromecast
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{31, 0, 1650}},
			vars.OS{vars.PlatformUnknown, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DeviceTV,
		},
		"",
	},
	{"Mozilla/5.0 (Linux; GoogleTV 3.2; VAP430 Build/MASTER) AppleWebKit/534.24 (KHTML, like Gecko) Chrome/11.0.696.77 Safari/534.24", // Google TV
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{11, 0, 696}},
			vars.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{0, 0, 0}, ""},
			vars.DeviceTV,
		},
		"Android",
	},
	{"Mozilla/5.0 (Linux; Android 5.0; ADT-1 Build/LPX13D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.89 Mobile Safari/537.36", // Android TV
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{40, 0, 2214}},
			vars.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{5, 0, 0}, "Lollipop"},
			vars.DeviceTV,
		},
		"Android 5 Lollipop",
	},
	{"Mozilla/5.0 (Linux; Android 4.2.2; AFTB Build/JDQ39) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.173 Mobile Safari/537.22", // Amazon Fire
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{25, 0, 1364}},
			vars.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{4, 2, 2}, "Jelly Bean"},
			vars.DeviceTV,
		},
		"Android 4.2 Jelly Bean",
	},
	{"Mozilla/5.0 (Unknown; Linux armv7l) AppleWebKit/537.1+ (KHTML, like Gecko) Safari/537.1+ LG def.Browser/6.00.00(+mouse+3D+SCREEN+TUNER; LGE; GLOBAL-PLAT5; 03.07.01; 0x00000001;); LG NetCast.TV-2013/03.17.01 (LG, GLOBAL-PLAT4, wired)", // LG TV
		vars.UserAgent{
			vars.Browser{vars.BrowserUnknown, version.Version{0, 0, 0}},
			vars.OS{vars.PlatformLinux, vars.OSLinux, version.Version{0, 0, 0}, ""},
			vars.DeviceTV,
		},
		"Linux",
	},
	{"Mozilla/5.0 (X11; FreeBSD; U; Viera; de-DE) AppleWebKit/537.11 (KHTML, like Gecko) Viera/3.10.0 Chrome/23.0.1271.97 Safari/537.11", // Panasonic Viera
		vars.UserAgent{
			vars.Browser{vars.BrowserChrome, version.Version{23, 0, 1271}},
			vars.OS{vars.PlatformLinux, vars.OSLinux, version.Version{0, 0, 0}, ""},
			vars.DeviceTV,
		},
		"Linux",
	},
	{"Roku/DVP-5.2 (025.02E03197A)", // Roku
		vars.UserAgent{
			vars.Browser{vars.BrowserUnknown, version.Version{0, 0, 0}},
			vars.OS{vars.PlatformUnknown, vars.OSUnknown, version.Version{0, 0, 0}, ""},
			vars.DeviceTV,
		},
		"",
	},
	{"mozilla/5.0 (smart-tv; linux; tizen 2.3) applewebkit/538.1 (khtml, like gecko) samsungbrowser/1.0 tv safari/538.1", // Samsung SmartTV
		vars.UserAgent{
			vars.Browser{vars.BrowserSamsung, version.Version{1, 0, 0}},
			vars.OS{vars.PlatformLinux, vars.OSLinux, version.Version{0, 0, 0}, ""},
			vars.DeviceTV,
		},
		"Linux",
	},
	{"mozilla/5.0 (linux; u) applewebkit/537.36 (khtml, like gecko) version/4.0 mobile safari/537.36 smarttv/6.0 (netcast)",
		vars.UserAgent{
			vars.Browser{vars.BrowserUnknown, version.Version{4, 0, 0}},
			vars.OS{vars.PlatformLinux, vars.OSLinux, version.Version{0, 0, 0}, ""},
			vars.DeviceTV,
		},
		"Linux",
	},
}
