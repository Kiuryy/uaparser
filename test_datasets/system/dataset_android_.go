package test_datasets_system

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/useragent"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Android = []test_datasets.TestDataset{
	{"Mozilla/5.0 (Linux; U; Android 1.0; en-us; dream) AppleWebKit/525.10+ (KHTML,like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{3, 0, 4}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{1, 0, 0}, ""},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 1.0; en-us; generic) AppleWebKit/525.10 (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{3, 0, 4}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{1, 0, 0}, ""},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 1.0.3; de-de; A80KSC Build/ECLAIR) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{4, 0, 0}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{1, 0, 3}, ""},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 1.5; en-gb; T-Mobile G1 Build/CRC1) AppleWebKit/528.5+ (KHTML, like Gecko) Version/3.1.2 Mobile Safari/525.20.1",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{3, 1, 2}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{1, 5, 0}, "Cupcake"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 1.5; es-; FBW1_4 Build/MASTER) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{3, 0, 4}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{1, 5, 0}, "Cupcake"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux U; Android 1.5 en-us hero) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{3, 0, 4}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{1, 5, 0}, "Cupcake"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 1.5; en-us; Opus One Build/RBE.00.00) AppleWebKit/528.18.1 (KHTML, like Gecko) Version/3.1.1 Mobile Safari/525.20.1",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{3, 1, 1}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{1, 5, 0}, "Cupcake"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 1.6; ar-us; SonyEricssonX10i Build/R2BA026) AppleWebKit/528.5+ (KHTML, like Gecko) Version/3.1.2 Mobile Safari/525.20.1",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{3, 1, 2}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{1, 6, 0}, "Donut"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 1.6; en-gb; HTC Tattoo Build/DRC79) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{3, 0, 4}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{1, 6, 0}, "Donut"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 1.6; ja-jp; Docomo HT-03A Build/DRD08) AppleWebKit/525.10 (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{3, 0, 4}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{1, 6, 0}, "Donut"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 2.1; en-us; Nexus One Build/ERD62) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{4, 0, 0}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{2, 1, 0}, "Éclair"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 2.1-update1; en-au; HTC_Desire_A8183 V1.16.841.1 Build/ERE27) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{4, 0, 0}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{2, 1, 0}, "Éclair"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 2.1; en-us; generic) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{3, 0, 4}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{2, 1, 0}, "Éclair"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 2.2; en-sa; HTC_DesireHD_A9191 Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{4, 0, 0}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{2, 2, 0}, "Froyo"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 2.2.1; en-gb; HTC_DesireZ_A7272 Build/FRG83D) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{4, 0, 0}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{2, 2, 1}, "Froyo"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 2.3.3; en-us; Sensation_4G Build/GRI40) AppleWebKit/533.1 (KHTML, like Gecko) Version/5.0 Safari/533.16",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{5, 0, 0}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{2, 3, 3}, "Gingerbread"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 2.3.5; ko-kr; SHW-M250S Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{4, 0, 0}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{2, 3, 5}, "Gingerbread"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 2.3.7; ja-jp; L-02D Build/GWK74) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{4, 0, 0}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{2, 3, 7}, "Gingerbread"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 3.0; en-us; Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{4, 0, 0}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{3, 0, 0}, "Honeycomb"},
			vars.DeviceTablet,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 4.0.1; en-us; sdk Build/ICS_MR0) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{4, 0, 0}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{4, 0, 1}, "Ice Cream Sandwich"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 4.1.1; en-us; Nexus S Build/JRO03E) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{4, 0, 0}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{4, 1, 1}, "Jelly Bean"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 4.1; en-gb; Build/JRN84D) AppleWebKit/534.30 (KHTML like Gecko) Version/4.0 Mobile Safari/534.30",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{4, 0, 0}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{4, 1, 0}, "Jelly Bean"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 4.1.1; el-gr; MB525 Build/JRO03H; CyanogenMod-10) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{4, 0, 0}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{4, 1, 1}, "Jelly Bean"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 4.1.1; fr-fr; MB525 Build/JRO03H; CyanogenMod-10) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{4, 0, 0}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{4, 1, 1}, "Jelly Bean"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 4.2; en-us; Nexus 10 Build/JVP15I) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{4, 0, 0}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{4, 2, 0}, "Jelly Bean"},
			vars.DeviceTablet,
		},
	},
	{"Mozilla/5.0 (Linux; U; Android 4.2; ro-ro; LT18i Build/4.1.B.0.431) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAndroid, version.Version{4, 0, 0}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{4, 2, 0}, "Jelly Bean"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; Android 4.3; Nexus 7 Build/JWR66D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.111 Safari/537.36",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserChrome, version.Version{27, 0, 1453}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{4, 3, 0}, "Jelly Bean"},
			vars.DeviceTablet,
		},
	},
	{"Mozilla/5.0 (Linux; Android 4.4; Nexus 7 Build/KOT24) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.105 Safari/537.36",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserChrome, version.Version{30, 0, 1599}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{4, 4, 0}, "KitKat"},
			vars.DeviceTablet,
		},
	},
	{"Mozilla/5.0 (Linux; Android 4.4; Nexus 4 Build/KRT16E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.105 Mobile Safari",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserChrome, version.Version{30, 0, 1599}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{4, 4, 0}, "KitKat"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; Android 6.0.1; SM-G930V Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.98 Mobile Safari/537.36",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserChrome, version.Version{52, 0, 2743}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{6, 0, 1}, "Marshmallow"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; Android 7.0; Nexus 5X Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.98 Mobile Safari/537.36",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserChrome, version.Version{52, 0, 2743}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{7, 0, 0}, "Nougat"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; Android 7.0; Nexus 6P Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/52.0.2743.98 Mobile Safari/537.36",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserChrome, version.Version{52, 0, 2743}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{7, 0, 0}, "Nougat"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (Linux; Android 8.0.0; SM-G930F Build/R16NW; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.157 Mobile Safari/537.36",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserChrome, version.Version{74, 0, 3729}},
			useragent.OS{vars.PlatformLinux, vars.OSAndroid, version.Version{8, 0, 0}, "Oreo"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0(iPad; U; CPU iPhone OS 3_2 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Version/4.0.4 Mobile/7B314 Safari/531.21.10",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{4, 0, 4}},
			useragent.OS{vars.PlatformiPad, vars.OSiOS, version.Version{3, 2, 0}, ""},
			vars.DeviceTablet,
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
	{"Mozilla/5.0 (iPad; CPU OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{5, 1, 0}},
			useragent.OS{vars.PlatformiPad, vars.OSiOS, version.Version{5, 0, 0}, ""},
			vars.DeviceTablet,
		},
	},
	{"Mozilla/5.0 (iPad; CPU OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5355d Safari/8536.25",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{6, 0, 0}},
			useragent.OS{vars.PlatformiPad, vars.OSiOS, version.Version{6, 0, 0}, ""},
			vars.DeviceTablet,
		},
	},
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/546.10 (KHTML, like Gecko) Version/6.0 Mobile/7E18WD Safari/8536.25",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{6, 0, 0}},
			useragent.OS{vars.PlatformiPhone, vars.OSiOS, version.Version{7, 0, 0}, ""},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (iPad; CPU OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{7, 0, 0}},
			useragent.OS{vars.PlatformiPad, vars.OSiOS, version.Version{7, 0, 0}, ""},
			vars.DeviceTablet,
		},
	},
	{"Mozilla/5.0 (iPad; CPU OS 7_0_2 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A501 Safari/9537.53",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserSafari, version.Version{7, 0, 0}},
			useragent.OS{vars.PlatformiPad, vars.OSiOS, version.Version{7, 0, 2}, ""},
			vars.DeviceTablet,
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
	{"Mozilla/5.0 (X11; U; Linux x86_64; en; rv:1.9.0.14) Gecko/20080528 Ubuntu/9.10 (karmic) Epiphany/2.22 Firefox/3.0",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserFirefox, version.Version{3, 0, 0}},
			useragent.OS{vars.PlatformLinux, vars.OSLinux, version.Version{0, 0, 0}, ""},
			vars.DeviceComputer,
		},
	},
}
