package test_datasets_bot

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/useragent"
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Bot = []test_datasets.TestDataset{
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1; +http://www.apple.com/go/applebot)",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserAppleBot, version.Version{8, 0, 2}},
			useragent.OS{vars.PlatformBot, vars.OSBot, version.Version{10, 10, 1}, "Yosemite"},
			vars.DeviceComputer,
		},
	},
	{"Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserBaiduBot, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformBot, vars.OSBot, version.Version{0, 0, 0}, ""},
			vars.DeviceUnknown,
		},
	},
	{"Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserBingBot, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformBot, vars.OSBot, version.Version{0, 0, 0}, ""},
			vars.DeviceUnknown,
		},
	},
	{"DuckDuckBot/1.0; (+http://duckduckgo.com/duckduckbot.html)",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserDuckDuckGoBot, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformBot, vars.OSBot, version.Version{0, 0, 0}, ""},
			vars.DeviceUnknown,
		},
	},
	{"facebookexternalhit/1.1 (+http://www.facebook.com/externalhit_uatext.php)",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserFacebookBot, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformBot, vars.OSBot, version.Version{0, 0, 0}, ""},
			vars.DeviceUnknown,
		},
	},
	{"Facebot/1.0",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserFacebookBot, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformBot, vars.OSBot, version.Version{0, 0, 0}, ""},
			vars.DeviceUnknown,
		},
	},
	{"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserGoogleBot, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformBot, vars.OSBot, version.Version{0, 0, 0}, ""},
			vars.DeviceUnknown,
		},
	},
	{"LinkedInBot/1.0 (compatible; Mozilla/5.0; Jakarta Commons-HttpClient/3.1 +http://www.linkedin.com)",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserLinkedInBot, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformBot, vars.OSBot, version.Version{0, 0, 0}, ""},
			vars.DeviceUnknown,
		},
	},
	{"msnbot/2.0b (+http://search.msn.com/msnbot.htm)",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserMsnBot, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformBot, vars.OSBot, version.Version{0, 0, 0}, ""},
			vars.DeviceUnknown,
		},
	},
	{"Pingdom.com_bot_version_1.4_(http://www.pingdom.com/)",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserPingdomBot, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformBot, vars.OSBot, version.Version{0, 0, 0}, ""},
			vars.DeviceUnknown,
		},
	},
	{"Twitterbot/1.0",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserTwitterBot, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformBot, vars.OSBot, version.Version{0, 0, 0}, ""},
			vars.DeviceUnknown,
		},
	},
	{"Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserYandexBot, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformBot, vars.OSBot, version.Version{0, 0, 0}, ""},
			vars.DeviceUnknown,
		},
	},
	{"Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserYahooBot, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformBot, vars.OSBot, version.Version{0, 0, 0}, ""},
			vars.DeviceUnknown,
		},
	},
	{"Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserGoogleBot, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformBot, vars.OSBot, version.Version{6, 0, 1}, "Marshmallow"},
			vars.DevicePhone,
		},
	},
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserGoogleBot, version.Version{6, 0, 0}},
			useragent.OS{vars.PlatformBot, vars.OSBot, version.Version{6, 0, 0}, ""},
			vars.DevicePhone,
		},
	},
	{"mozilla/5.0 (unknown; linux x86_64) applewebkit/538.1 (khtml, like gecko) phantomjs/2.1.1 safari/538.1",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserBot, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformBot, vars.OSBot, version.Version{0, 0, 0}, ""},
			vars.DeviceComputer,
		},
	},
	{"Mozilla/5.0 (compatible; coccocbot/1.0; +http://help.coccoc.com/searchengine)",
		useragent.UserAgent{
			useragent.Browser{vars.BrowserCocCocBot, version.Version{0, 0, 0}},
			useragent.OS{vars.PlatformBot, vars.OSBot, version.Version{0, 0, 0}, ""},
			vars.DeviceUnknown,
		},
	},
}
