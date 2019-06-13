package uasurfer

import "testing"

var testUAVars = []struct {
	UA string
	UserAgent
}{
	// Empty
	{"",
		UserAgent{}},

	// Single char
	{"a",
		UserAgent{}},

	// Some random string
	{"some random string",
		UserAgent{}},

	// Potentially malformed ua
	{")(",
		UserAgent{}},

	// iPhone
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/546.10 (KHTML, like Gecko) Version/6.0 Mobile/7E18WD Safari/8536.25",
		UserAgent{
			Browser{BrowserSafari, Version{6, 0, 0}}, OS{PlatformiPhone, OSiOS, Version{7, 0, 0}}, DevicePhone}},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 8_0_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12A405 Safari/600.1.4",
		UserAgent{
			Browser{BrowserSafari, Version{8, 0, 0}}, OS{PlatformiPhone, OSiOS, Version{8, 0, 2}}, DevicePhone}},

	// iPad
	{"Mozilla/5.0(iPad; U; CPU iPhone OS 3_2 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Version/4.0.4 Mobile/7B314 Safari/531.21.10",
		UserAgent{
			Browser{BrowserSafari, Version{4, 0, 4}}, OS{PlatformiPad, OSiOS, Version{3, 2, 0}}, DeviceTablet}},

	{"Mozilla/5.0 (iPad; CPU OS 9_0 like Mac OS X) AppleWebKit/601.1.17 (KHTML, like Gecko) Version/8.0 Mobile/13A175 Safari/600.1.4",
		UserAgent{
			Browser{BrowserSafari, Version{8, 0, 0}}, OS{PlatformiPad, OSiOS, Version{9, 0, 0}}, DeviceTablet}},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 10_0 like Mac OS X) AppleWebKit/602.1.32 (KHTML, like Gecko) Version/10.0 Mobile/14A5261v Safari/602.1",
		UserAgent{
			Browser{BrowserSafari, Version{10, 0, 0}}, OS{PlatformiPhone, OSiOS, Version{10, 0, 0}}, DevicePhone}},

	// Chrome
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.130 Safari/537.36",
		UserAgent{
			Browser{BrowserChrome, Version{43, 0, 2357}}, OS{PlatformMac, OSMacOS, Version{10, 10, 4}}, DeviceComputer}},

	{"Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; en) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/534.48.3",
		UserAgent{
			Browser{BrowserChrome, Version{19, 0, 1084}}, OS{PlatformiPhone, OSiOS, Version{5, 1, 1}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; Android 6.0; Nexus 5X Build/MDB08L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.76 Mobile Safari/537.36",
		UserAgent{
			Browser{BrowserChrome, Version{46, 0, 2490}}, OS{PlatformLinux, OSAndroid, Version{6, 0, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36",
		UserAgent{
			Browser{BrowserChrome, Version{45, 0, 2454}}, OS{PlatformMac, OSMacOS, Version{10, 11, 0}}, DeviceComputer}},

	// Chromium
	{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/69.0.3497.81 Chrome/69.0.3497.81 Safari/537.36",
		UserAgent{
			Browser{BrowserChromium, Version{69, 0, 3497}}, OS{PlatformLinux, OSLinux, Version{0, 0, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/75.0.3731.0 Chrome/75.0.3731.0 Safari/537.36",
		UserAgent{
			Browser{BrowserChromium, Version{75, 0, 3731}}, OS{PlatformWindows, OSWindows, Version{10, 0, 0}}, DeviceComputer}},

	// Vivaldi
	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36 Vivaldi/2.1.1337.36",
		UserAgent{
			Browser{BrowserVivaldi, Version{2, 1, 1337}}, OS{PlatformWindows, OSWindows, Version{10, 0, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36 Vivaldi/1.2.490.43",
		UserAgent{
			Browser{BrowserVivaldi, Version{1, 2, 490}}, OS{PlatformLinux, OSLinux, Version{0, 0, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.88 Safari/537.36 Vivaldi/2.4.1488.35",
		UserAgent{
			Browser{BrowserVivaldi, Version{2, 4, 1488}}, OS{PlatformMac, OSMacOS, Version{10, 14, 1}}, DeviceComputer}},

	// Safari
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/600.7.12 (KHTML, like Gecko) Version/8.0.7 Safari/600.7.12",
		UserAgent{
			Browser{BrowserSafari, Version{8, 0, 7}}, OS{PlatformMac, OSMacOS, Version{10, 10, 4}}, DeviceComputer}},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_5; en-us) AppleWebKit/525.26.2 (KHTML, like Gecko) Version/3.2 Safari/525.26.12",
		UserAgent{
			Browser{BrowserSafari, Version{3, 2, 0}}, OS{PlatformMac, OSMacOS, Version{10, 5, 5}}, DeviceComputer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12) AppleWebKit/602.1.32 (KHTML, like Gecko) Version/10.0 Safari/602.1.32", // macOS Sierra dev beta
		UserAgent{
			Browser{BrowserSafari, Version{10, 0, 0}}, OS{PlatformMac, OSMacOS, Version{10, 12, 0}}, DeviceComputer}},

	// Firefox
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) FxiOS/1.0 Mobile/12F69 Safari/600.1.4",
		UserAgent{
			Browser{BrowserFirefox, Version{1, 0, 0}}, OS{PlatformiPhone, OSiOS, Version{8, 3, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Android 4.4; Tablet; rv:41.0) Gecko/41.0 Firefox/41.0",
		UserAgent{
			Browser{BrowserFirefox, Version{41, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{4, 4, 0}}, DeviceTablet}},

	{"Mozilla/5.0 (Android; Mobile; rv:40.0) Gecko/40.0 Firefox/40.0",
		UserAgent{
			Browser{BrowserFirefox, Version{40, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{0, 0, 0}}, DevicePhone}},

	{"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:38.0) Gecko/20100101 Firefox/38.0",
		UserAgent{
			Browser{BrowserFirefox, Version{38, 0, 0}}, OS{PlatformLinux, OSLinux, Version{0, 0, 0}}, DeviceComputer}},

	// Opera
	{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36 OPR/18.0.1284.68",
		UserAgent{
			Browser{BrowserOpera, Version{18, 0, 1284}}, OS{PlatformWindows, OSWindows, Version{6, 1, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 8_4 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) OPiOS/10.2.0.93022 Mobile/12H143 Safari/9537.53",
		UserAgent{
			Browser{BrowserOpera, Version{10, 2, 0}}, OS{PlatformiPhone, OSiOS, Version{8, 4, 0}}, DevicePhone}},

	// Brave
	{"Mozilla/5.0 (Linux; Android 9; ONEPLUS A5010) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.157 Mobile Safari/537.36 Brave/74",
		UserAgent{
			Browser{BrowserBrave, Version{74, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{9, 0, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.38 Safari/537.36 Brave/75",
		UserAgent{
			Browser{BrowserBrave, Version{75, 0, 0}}, OS{PlatformWindows, OSWindows, Version{10, 0, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) brave/0.7.10 Chrome/47.0.2526.110 Brave/0.36.5 Safari/537.36",
		UserAgent{
			Browser{BrowserBrave, Version{0, 7, 10}}, OS{PlatformWindows, OSWindows, Version{6, 1, 0}}, DeviceComputer}},

	// Internet Explorer -- https://msdn.microsoft.com/en-us/library/hh869301(v=vs.85).aspx
	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)",
		UserAgent{
			Browser{BrowserIE, Version{10, 0, 0}}, OS{PlatformWindows, OSWindows, Version{6, 2, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Windows NT 6.3; Trident/7.0; .NET4.0E; .NET4.0C; rv:11.0) like Gecko",
		UserAgent{
			Browser{BrowserIE, Version{11, 0, 0}}, OS{PlatformWindows, OSWindows, Version{6, 3, 0}}, DeviceComputer}},

	{"Mozilla/4.0 (compatible; MSIE 5.01; Windows NT 5.0; SV1; .NET CLR 1.1.4322; .NET CLR 1.0.3705; .NET CLR 2.0.50727)",
		UserAgent{
			Browser{BrowserIE, Version{5, 0, 1}}, OS{PlatformWindows, OSWindows, Version{5, 0, 0}}, DeviceComputer}},

	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/4.0; GTB6.4; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; OfficeLiveConnector.1.3; OfficeLivePatch.0.0; .NET CLR 1.1.4322)",
		UserAgent{
			Browser{BrowserIE, Version{7, 0, 0}}, OS{PlatformWindows, OSWindows, Version{6, 1, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0; Touch)", //Windows Surface RT tablet
		UserAgent{
			Browser{BrowserIE, Version{10, 0, 0}}, OS{PlatformWindows, OSWindows, Version{6, 2, 0}}, DeviceTablet}},

	// Edge
	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.123",
		UserAgent{
			Browser{BrowserEdge, Version{12, 123, 0}}, OS{PlatformWindows, OSWindows, Version{10, 0, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 12_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 EdgiOS/44.3.5 Mobile/15E148 Safari/605.1.15",
		UserAgent{
			Browser{BrowserEdge, Version{12, 0, 0}}, OS{PlatformiPhone, OSiOS, Version{12, 3, 1}}, DevicePhone}},

	{"Mozilla/5.0 (iPad; CPU OS 12_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 EdgiOS/44.3.2 Mobile/15E148 Safari/605.1.15",
		UserAgent{
			Browser{BrowserEdge, Version{12, 0, 0}}, OS{PlatformiPad, OSiOS, Version{12, 3, 1}}, DeviceTablet}},

	{"Mozilla/5.0 (Linux; Android 9; motorola one) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.90 Mobile Safari/537.36 EdgA/42.0.2.3728",
		UserAgent{
			Browser{BrowserEdge, Version{42, 0, 2}}, OS{PlatformLinux, OSAndroid, Version{9, 0, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3800.0 Safari/537.36 Edg/76.0.172.0",
		UserAgent{
			Browser{BrowserEdge, Version{76, 0, 172}}, OS{PlatformWindows, OSWindows, Version{10, 0, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3803.0 Safari/537.36 Edg/76.0.176.0",
		UserAgent{
			Browser{BrowserEdge, Version{76, 0, 176}}, OS{PlatformMac, OSMacOS, Version{10, 14, 5}}, DeviceComputer}},

	// UC Browser
	{"Mozilla/5.0 (Linux; U; Android 2.3.4; en-US; MT11i Build/4.0.2.A.0.62) AppleWebKit/534.31 (KHTML, like Gecko) UCBrowser/9.0.1.275 U3/0.8.0 Mobile Safari/534.31",
		UserAgent{
			Browser{BrowserUCBrowser, Version{9, 0, 1}}, OS{PlatformLinux, OSAndroid, Version{2, 3, 4}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 4.0.4; en-US; Micromax P255 Build/IMM76D) AppleWebKit/534.31 (KHTML, like Gecko) UCBrowser/9.2.0.308 U3/0.8.0 Mobile Safari/534.31",
		UserAgent{
			Browser{BrowserUCBrowser, Version{9, 2, 0}}, OS{PlatformLinux, OSAndroid, Version{4, 0, 4}}, DevicePhone}},

	{"UCWEB/2.0 (Java; U; MIDP-2.0; en-US; MicromaxQ5) U2/1.0.0 UCBrowser/9.4.0.342 U2/1.0.0 Mobile",
		UserAgent{
			Browser{BrowserUCBrowser, Version{9, 4, 0}}, OS{PlatformUnknown, OSUnknown, Version{0, 0, 0}}, DevicePhone}},

	// ChromeOS
	{"Mozilla/5.0 (X11; U; CrOS i686 9.10.0; en-US) AppleWebKit/532.5 (KHTML, like Gecko) Chrome/4.0.253.0 Safari/532.5",
		UserAgent{
			Browser{BrowserChrome, Version{4, 0, 253}}, OS{PlatformLinux, OSChromeOS, Version{0, 0, 0}}, DeviceComputer}},

	// iPod, iPod Touch
	{"mozilla/5.0 (ipod touch; cpu iphone os 9_3_3 like mac os x) applewebkit/601.1.46 (khtml, like gecko) version/9.0 mobile/13g34 safari/601.1",
		UserAgent{
			Browser{BrowserSafari, Version{9, 0, 0}}, OS{PlatformiPod, OSiOS, Version{9, 3, 3}}, DeviceTablet}},

	{"mozilla/5.0 (ipod; cpu iphone os 6_1_6 like mac os x) applewebkit/536.26 (khtml, like gecko) version/6.0 mobile/10b500 safari/8536.25",
		UserAgent{
			Browser{BrowserSafari, Version{6, 0, 0}}, OS{PlatformiPod, OSiOS, Version{6, 1, 6}}, DeviceTablet}},

	// Android WebView (Android <= 4.3)
	{"Mozilla/5.0 (Linux; U; Android 2.2; en-us; DROID2 GLOBAL Build/S273) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		UserAgent{
			Browser{BrowserAndroid, Version{4, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{2, 2, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 4.0.3; de-ch; HTC Sensation Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari53/4.30",
		UserAgent{
			Browser{BrowserAndroid, Version{4, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{4, 0, 3}}, DevicePhone}},

	// Smart TVs and TV dongles
	{"Mozilla/5.0 (CrKey armv7l 1.4.15250) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.0 Safari/537.36", // Chromecast
		UserAgent{
			Browser{BrowserChrome, Version{31, 0, 1650}}, OS{PlatformUnknown, OSUnknown, Version{0, 0, 0}}, DeviceTV}},

	{"Mozilla/5.0 (Linux; GoogleTV 3.2; VAP430 Build/MASTER) AppleWebKit/534.24 (KHTML, like Gecko) Chrome/11.0.696.77 Safari/534.24", // Google TV
		UserAgent{
			Browser{BrowserChrome, Version{11, 0, 696}}, OS{PlatformLinux, OSAndroid, Version{0, 0, 0}}, DeviceTV}},

	{"Mozilla/5.0 (Linux; Android 5.0; ADT-1 Build/LPX13D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.89 Mobile Safari/537.36", // Android TV
		UserAgent{
			Browser{BrowserChrome, Version{40, 0, 2214}}, OS{PlatformLinux, OSAndroid, Version{5, 0, 0}}, DeviceTV}},

	{"Mozilla/5.0 (Linux; Android 4.2.2; AFTB Build/JDQ39) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.173 Mobile Safari/537.22", // Amazon Fire
		UserAgent{
			Browser{BrowserChrome, Version{25, 0, 1364}}, OS{PlatformLinux, OSAndroid, Version{4, 2, 2}}, DeviceTV}},

	{"Mozilla/5.0 (Unknown; Linux armv7l) AppleWebKit/537.1+ (KHTML, like Gecko) Safari/537.1+ LG Browser/6.00.00(+mouse+3D+SCREEN+TUNER; LGE; GLOBAL-PLAT5; 03.07.01; 0x00000001;); LG NetCast.TV-2013/03.17.01 (LG, GLOBAL-PLAT4, wired)", // LG TV
		UserAgent{
			Browser{BrowserUnknown, Version{0, 0, 0}}, OS{PlatformLinux, OSLinux, Version{0, 0, 0}}, DeviceTV}},

	{"Mozilla/5.0 (X11; FreeBSD; U; Viera; de-DE) AppleWebKit/537.11 (KHTML, like Gecko) Viera/3.10.0 Chrome/23.0.1271.97 Safari/537.11", // Panasonic Viera
		UserAgent{
			Browser{BrowserChrome, Version{23, 0, 1271}}, OS{PlatformLinux, OSLinux, Version{0, 0, 0}}, DeviceTV}},

	{"Roku/DVP-5.2 (025.02E03197A)", // Roku
		UserAgent{
			Browser{BrowserUnknown, Version{0, 0, 0}}, OS{PlatformUnknown, OSUnknown, Version{0, 0, 0}}, DeviceTV}},

	{"mozilla/5.0 (smart-tv; linux; tizen 2.3) applewebkit/538.1 (khtml, like gecko) samsungbrowser/1.0 tv safari/538.1", // Samsung SmartTV
		UserAgent{
			Browser{BrowserSamsung, Version{0, 0, 0}}, OS{PlatformLinux, OSLinux, Version{0, 0, 0}}, DeviceTV}},

	{"mozilla/5.0 (linux; u) applewebkit/537.36 (khtml, like gecko) version/4.0 mobile safari/537.36 smarttv/6.0 (netcast)",
		UserAgent{
			Browser{BrowserUnknown, Version{4, 0, 0}}, OS{PlatformLinux, OSLinux, Version{0, 0, 0}}, DeviceTV}},

	// Google search app (GSA) for iOS -- it's Safari in disguise as of v6
	{"Mozilla/5.0 (iPad; CPU OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) GSA/6.0.51363 Mobile/12F69 Safari/600.1.4",
		UserAgent{
			Browser{BrowserSafari, Version{8, 3, 0}}, OS{PlatformiPad, OSiOS, Version{8, 3, 0}}, DeviceTablet}},

	// OCSP fetchers
	{"Microsoft-CryptoAPI/10.0",
		UserAgent{
			Browser{BrowserUnknown, Version{0, 0, 0}}, OS{PlatformWindows, OSUnknown, Version{0, 0, 0}}, DeviceComputer}},
	{"trustd (unknown version) CFNetwork/811.7.2 Darwin/16.7.0 (x86_64)",
		UserAgent{
			Browser{BrowserUnknown, Version{0, 0, 0}}, OS{PlatformMac, OSUnknown, Version{0, 0, 0}}, DeviceComputer}},
	{"ocspd (unknown version) CFNetwork/520.5.3 Darwin/11.4.2 (x86_64)(MacBookAir5%2C2)",
		UserAgent{
			Browser{BrowserUnknown, Version{0, 0, 0}}, OS{PlatformMac, OSUnknown, Version{0, 0, 0}}, DeviceComputer}},
	// Bots
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1; +http://www.apple.com/go/applebot)",
		UserAgent{
			Browser{BrowserAppleBot, Version{0, 0, 0}}, OS{PlatformBot, OSBot, Version{10, 10, 1}}, DeviceComputer}},

	{"Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)",
		UserAgent{
			Browser{BrowserBaiduBot, Version{0, 0, 0}}, OS{PlatformBot, OSBot, Version{0, 0, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)",
		UserAgent{
			Browser{BrowserBingBot, Version{0, 0, 0}}, OS{PlatformBot, OSBot, Version{0, 0, 0}}, DeviceComputer}},

	{"DuckDuckBot/1.0; (+http://duckduckgo.com/duckduckbot.html)",
		UserAgent{
			Browser{BrowserDuckDuckGoBot, Version{0, 0, 0}}, OS{PlatformBot, OSBot, Version{0, 0, 0}}, DeviceComputer}},

	{"facebookexternalhit/1.1 (+http://www.facebook.com/externalhit_uatext.php)",
		UserAgent{
			Browser{BrowserFacebookBot, Version{0, 0, 0}}, OS{PlatformBot, OSBot, Version{0, 0, 0}}, DeviceComputer}},

	{"Facebot/1.0",
		UserAgent{
			Browser{BrowserFacebookBot, Version{0, 0, 0}}, OS{PlatformBot, OSBot, Version{0, 0, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		UserAgent{
			Browser{BrowserGoogleBot, Version{0, 0, 0}}, OS{PlatformBot, OSBot, Version{0, 0, 0}}, DeviceComputer}},

	{"LinkedInBot/1.0 (compatible; Mozilla/5.0; Jakarta Commons-HttpClient/3.1 +http://www.linkedin.com)",
		UserAgent{
			Browser{BrowserLinkedInBot, Version{0, 0, 0}}, OS{PlatformBot, OSBot, Version{0, 0, 0}}, DeviceComputer}},

	{"msnbot/2.0b (+http://search.msn.com/msnbot.htm)",
		UserAgent{
			Browser{BrowserMsnBot, Version{0, 0, 0}}, OS{PlatformBot, OSBot, Version{0, 0, 0}}, DeviceComputer}},

	{"Pingdom.com_bot_version_1.4_(http://www.pingdom.com/)",
		UserAgent{
			Browser{BrowserPingdomBot, Version{0, 0, 0}}, OS{PlatformBot, OSBot, Version{0, 0, 0}}, DeviceComputer}},

	{"Twitterbot/1.0",
		UserAgent{
			Browser{BrowserTwitterBot, Version{0, 0, 0}}, OS{PlatformBot, OSBot, Version{0, 0, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)",
		UserAgent{
			Browser{BrowserYandexBot, Version{0, 0, 0}}, OS{PlatformBot, OSBot, Version{0, 0, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)",
		UserAgent{
			Browser{BrowserYahooBot, Version{0, 0, 0}}, OS{PlatformBot, OSBot, Version{0, 0, 0}}, DeviceComputer}},

	{"{UA:Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)}, ua: &{Browser:{Name:BrowserGoogleBot Version:{Major:0 Minor:0 Patch:0}} OS:{Platform:PlatformBot Name:OSBot Version:{Major:6 Minor:0 Patch:1}} DeviceType:DeviceComputer}",
		UserAgent{
			Browser{BrowserGoogleBot, Version{0, 0, 0}}, OS{PlatformBot, OSBot, Version{6, 0, 1}}, DeviceComputer}},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		UserAgent{
			Browser{BrowserGoogleBot, Version{0, 0, 0}}, OS{PlatformBot, OSBot, Version{6, 0, 0}}, DeviceComputer}},

	{"mozilla/5.0 (unknown; linux x86_64) applewebkit/538.1 (khtml, like gecko) phantomjs/2.1.1 safari/538.1",
		UserAgent{
			Browser{BrowserBot, Version{0, 0, 0}}, OS{PlatformBot, OSBot, Version{0, 0, 0}}, DeviceComputer}},

	// Unknown or partially handled
	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.4; en-US; rv:1.9.1b3pre) Gecko/20090223 SeaMonkey/2.0a3", //Seamonkey (~FF)
		UserAgent{
			Browser{BrowserFirefox, Version{0, 0, 0}}, OS{PlatformMac, OSMacOS, Version{10, 4, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en; rv:1.9.0.8pre) Gecko/2009022800 Camino/2.0b3pre", //Camino (~FF)
		UserAgent{
			Browser{BrowserUnknown, Version{0, 0, 0}}, OS{PlatformMac, OSMacOS, Version{10, 5, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Mobile; rv:26.0) Gecko/26.0 Firefox/26.0", //firefox OS
		UserAgent{
			Browser{BrowserFirefox, Version{26, 0, 0}}, OS{PlatformUnknown, OSUnknown, Version{0, 0, 0}}, DevicePhone}},

	{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.45 Safari/535.19", //chrome for android having requested desktop site
		UserAgent{
			Browser{BrowserChrome, Version{18, 0, 1025}}, OS{PlatformLinux, OSLinux, Version{0, 0, 0}}, DeviceComputer}},

	{"Opera/9.80 (S60; SymbOS; Opera Mobi/352; U; de) Presto/2.4.15 Version/10.00",
		UserAgent{
			Browser{BrowserOpera, Version{10, 0, 0}}, OS{PlatformUnknown, OSUnknown, Version{0, 0, 0}}, DevicePhone}},

	// BrowserQQ
	{"Mozilla/5.0 (Windows NT 6.2; WOW64; Trident/7.0; Touch; .NET4.0E; .NET4.0C; .NET CLR 3.5.30729; .NET CLR 2.0.50727; .NET CLR 3.0.30729; InfoPath.3; Tablet PC 2.0; QQBrowser/7.6.21433.400; rv:11.0) like Gecko",
		UserAgent{
			Browser{BrowserQQ, Version{7, 6, 21433}}, OS{PlatformWindows, OSWindows, Version{6, 2, 0}}, DeviceTablet}},

	{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.124 Safari/537.36 QQBrowser/9.0.2191.400",
		UserAgent{
			Browser{BrowserQQ, Version{9, 0, 2191}}, OS{PlatformWindows, OSWindows, Version{6, 1, 0}}, DeviceComputer}},

	{"mozilla/5.0 (iphone; cpu iphone os 8_1_2 like mac os x) applewebkit/600.1.4 (khtml, like gecko) mobile/12b440 qq/5.3.0.319 nettype/wifi mem/205",
		UserAgent{
			Browser{BrowserQQ, Version{5, 3, 0}}, OS{PlatformiPhone, OSiOS, Version{8, 1, 2}}, DevicePhone}},

	// ANDROID TESTS

	{"Mozilla/5.0 (Linux; U; Android 1.0; en-us; dream) AppleWebKit/525.10+ (KHTML,like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		UserAgent{
			Browser{BrowserAndroid, Version{3, 0, 4}}, OS{PlatformLinux, OSAndroid, Version{1, 0, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 1.0; en-us; generic) AppleWebKit/525.10 (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		UserAgent{
			Browser{BrowserAndroid, Version{3, 0, 4}}, OS{PlatformLinux, OSAndroid, Version{1, 0, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 1.0.3; de-de; A80KSC Build/ECLAIR) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17",
		UserAgent{
			Browser{BrowserAndroid, Version{4, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{1, 0, 3}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 1.5; en-gb; T-Mobile G1 Build/CRC1) AppleWebKit/528.5+ (KHTML, like Gecko) Version/3.1.2 Mobile Safari/525.20.1",
		UserAgent{
			Browser{BrowserAndroid, Version{3, 1, 2}}, OS{PlatformLinux, OSAndroid, Version{1, 5, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 1.5; es-; FBW1_4 Build/MASTER) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		UserAgent{
			Browser{BrowserAndroid, Version{3, 0, 4}}, OS{PlatformLinux, OSAndroid, Version{1, 5, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Linux U; Android 1.5 en-us hero) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		UserAgent{
			Browser{BrowserAndroid, Version{3, 0, 4}}, OS{PlatformLinux, OSAndroid, Version{1, 5, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 1.5; en-us; Opus One Build/RBE.00.00) AppleWebKit/528.18.1 (KHTML, like Gecko) Version/3.1.1 Mobile Safari/525.20.1",
		UserAgent{
			Browser{BrowserAndroid, Version{3, 1, 1}}, OS{PlatformLinux, OSAndroid, Version{1, 5, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 1.6; ar-us; SonyEricssonX10i Build/R2BA026) AppleWebKit/528.5+ (KHTML, like Gecko) Version/3.1.2 Mobile Safari/525.20.1",
		UserAgent{
			Browser{BrowserAndroid, Version{3, 1, 2}}, OS{PlatformLinux, OSAndroid, Version{1, 6, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 1.6; en-gb; HTC Tattoo Build/DRC79) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		UserAgent{
			Browser{BrowserAndroid, Version{3, 0, 4}}, OS{PlatformLinux, OSAndroid, Version{1, 6, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 1.6; ja-jp; Docomo HT-03A Build/DRD08) AppleWebKit/525.10 (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		UserAgent{
			Browser{BrowserAndroid, Version{3, 0, 4}}, OS{PlatformLinux, OSAndroid, Version{1, 6, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 2.1; en-us; Nexus One Build/ERD62) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17",
		UserAgent{
			Browser{BrowserAndroid, Version{4, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{2, 1, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 2.1-update1; en-au; HTC_Desire_A8183 V1.16.841.1 Build/ERE27) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17",
		UserAgent{
			Browser{BrowserAndroid, Version{4, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{2, 1, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 2.1; en-us; generic) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2",
		UserAgent{
			Browser{BrowserAndroid, Version{3, 0, 4}}, OS{PlatformLinux, OSAndroid, Version{2, 1, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 2.2; en-sa; HTC_DesireHD_A9191 Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		UserAgent{
			Browser{BrowserAndroid, Version{4, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{2, 2, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 2.2.1; en-gb; HTC_DesireZ_A7272 Build/FRG83D) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		UserAgent{
			Browser{BrowserAndroid, Version{4, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{2, 2, 1}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 2.3.3; en-us; Sensation_4G Build/GRI40) AppleWebKit/533.1 (KHTML, like Gecko) Version/5.0 Safari/533.16",
		UserAgent{
			Browser{BrowserAndroid, Version{5, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{2, 3, 3}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 2.3.5; ko-kr; SHW-M250S Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		UserAgent{
			Browser{BrowserAndroid, Version{4, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{2, 3, 5}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 2.3.7; ja-jp; L-02D Build/GWK74) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		UserAgent{
			Browser{BrowserAndroid, Version{4, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{2, 3, 7}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 3.0; en-us; Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
		UserAgent{
			Browser{BrowserAndroid, Version{4, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{3, 0, 0}}, DeviceTablet}},

	{"Mozilla/5.0 (Linux; U; Android 4.0.1; en-us; sdk Build/ICS_MR0) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		UserAgent{
			Browser{BrowserAndroid, Version{4, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{4, 0, 1}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 4.1.1; en-us; Nexus S Build/JRO03E) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		UserAgent{
			Browser{BrowserAndroid, Version{4, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{4, 1, 1}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 4.1; en-gb; Build/JRN84D) AppleWebKit/534.30 (KHTML like Gecko) Version/4.0 Mobile Safari/534.30",
		UserAgent{
			Browser{BrowserAndroid, Version{4, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{4, 1, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 4.1.1; el-gr; MB525 Build/JRO03H; CyanogenMod-10) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		UserAgent{
			Browser{BrowserAndroid, Version{4, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{4, 1, 1}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 4.1.1; fr-fr; MB525 Build/JRO03H; CyanogenMod-10) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		UserAgent{
			Browser{BrowserAndroid, Version{4, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{4, 1, 1}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; U; Android 4.2; en-us; Nexus 10 Build/JVP15I) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30",
		UserAgent{
			Browser{BrowserAndroid, Version{4, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{4, 2, 0}}, DeviceTablet}},

	{"Mozilla/5.0 (Linux; U; Android 4.2; ro-ro; LT18i Build/4.1.B.0.431) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		UserAgent{
			Browser{BrowserAndroid, Version{4, 0, 0}}, OS{PlatformLinux, OSAndroid, Version{4, 2, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; Android 4.3; Nexus 7 Build/JWR66D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.111 Safari/537.36",
		UserAgent{
			Browser{BrowserChrome, Version{27, 0, 1453}}, OS{PlatformLinux, OSAndroid, Version{4, 3, 0}}, DeviceTablet}},

	{"Mozilla/5.0 (Linux; Android 4.4; Nexus 7 Build/KOT24) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.105 Safari/537.36",
		UserAgent{
			Browser{BrowserChrome, Version{30, 0, 1599}}, OS{PlatformLinux, OSAndroid, Version{4, 4, 0}}, DeviceTablet}},

	{"Mozilla/5.0 (Linux; Android 4.4; Nexus 4 Build/KRT16E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.105 Mobile Safari",
		UserAgent{
			Browser{BrowserChrome, Version{30, 0, 1599}}, OS{PlatformLinux, OSAndroid, Version{4, 4, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; Android 6.0.1; SM-G930V Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.98 Mobile Safari/537.36",
		UserAgent{
			Browser{BrowserChrome, Version{52, 0, 2743}}, OS{PlatformLinux, OSAndroid, Version{6, 0, 1}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; Android 7.0; Nexus 5X Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.98 Mobile Safari/537.36",
		UserAgent{
			Browser{BrowserChrome, Version{52, 0, 2743}}, OS{PlatformLinux, OSAndroid, Version{7, 0, 0}}, DevicePhone}},

	{"Mozilla/5.0 (Linux; Android 7.0; Nexus 6P Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/52.0.2743.98 Mobile Safari/537.36",
		UserAgent{
			Browser{BrowserChrome, Version{52, 0, 2743}}, OS{PlatformLinux, OSAndroid, Version{7, 0, 0}}, DevicePhone}},

	{"Mozilla/5.0 (X11; U; CrOS i686 9.10.0; en-US) AppleWebKit/532.5 (KHTML, like Gecko) Chrome/4.0.253.0 Safari/532.5",
		UserAgent{
			Browser{BrowserChrome, Version{4, 0, 253}}, OS{PlatformLinux, OSChromeOS, Version{0, 0, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (X11; CrOS armv7l 5500.100.6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.120 Safari/537.36",
		UserAgent{
			Browser{BrowserChrome, Version{34, 0, 1847}}, OS{PlatformLinux, OSChromeOS, Version{0, 0, 0}}, DeviceComputer}},

	{"Mozilla/5.0(iPad; U; CPU iPhone OS 3_2 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Version/4.0.4 Mobile/7B314 Safari/531.21.10",
		UserAgent{
			Browser{BrowserSafari, Version{4, 0, 4}}, OS{PlatformiPad, OSiOS, Version{3, 2, 0}}, DeviceTablet}},

	{"Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_0 like Mac OS X; en-us) AppleWebKit/532.9 (KHTML, like Gecko) Version/4.0.5 Mobile/8A293 Safari/6531.22.7",
		UserAgent{
			Browser{BrowserSafari, Version{4, 0, 5}}, OS{PlatformiPhone, OSiOS, Version{4, 0, 0}}, DevicePhone}},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3",
		UserAgent{
			Browser{BrowserSafari, Version{5, 1, 0}}, OS{PlatformiPhone, OSiOS, Version{5, 0, 0}}, DevicePhone}},

	{"Mozilla/5.0 (iPad; CPU OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3",
		UserAgent{
			Browser{BrowserSafari, Version{5, 1, 0}}, OS{PlatformiPad, OSiOS, Version{5, 0, 0}}, DeviceTablet}},

	{"Mozilla/5.0 (iPad; CPU OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5355d Safari/8536.25",
		UserAgent{
			Browser{BrowserSafari, Version{6, 0, 0}}, OS{PlatformiPad, OSiOS, Version{6, 0, 0}}, DeviceTablet}},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/546.10 (KHTML, like Gecko) Version/6.0 Mobile/7E18WD Safari/8536.25",
		UserAgent{
			Browser{BrowserSafari, Version{6, 0, 0}}, OS{PlatformiPhone, OSiOS, Version{7, 0, 0}}, DevicePhone}},

	{"Mozilla/5.0 (iPad; CPU OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53",
		UserAgent{
			Browser{BrowserSafari, Version{7, 0, 0}}, OS{PlatformiPad, OSiOS, Version{7, 0, 0}}, DeviceTablet}},

	{"Mozilla/5.0 (iPad; CPU OS 7_0_2 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A501 Safari/9537.53",
		UserAgent{
			Browser{BrowserSafari, Version{7, 0, 0}}, OS{PlatformiPad, OSiOS, Version{7, 0, 2}}, DeviceTablet}},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 10_2_1 like Mac OS X) AppleWebKit/602.4.6 (KHTML, like Gecko) Mobile/14D27 [FBAN/FBIOS;FBAV/86.0.0.48.52;FBBV/53842252;FBDV/iPhone9,1;FBMD/iPhone;FBSN/iOS;FBSV/10.2.1;FBSS/2;FBCR/Verizon;FBID/phone;FBLC/en_US;FBOP/5;FBRV/0]",
		UserAgent{
			Browser{BrowserSafari, Version{10, 2, 1}}, OS{PlatformiPhone, OSiOS, Version{10, 2, 1}}, DevicePhone}},

	{"Mozilla/5.0 (iPhone; CPU iPhone OS 8_0_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12A405 Safari/600.1.4",
		UserAgent{
			Browser{BrowserSafari, Version{8, 0, 0}}, OS{PlatformiPhone, OSiOS, Version{8, 0, 2}}, DevicePhone}},

	{"Mozilla/5.0 (X11; U; Linux x86_64; en; rv:1.9.0.14) Gecko/20080528 Ubuntu/9.10 (karmic) Epiphany/2.22 Firefox/3.0",
		UserAgent{
			Browser{BrowserFirefox, Version{3, 0, 0}}, OS{PlatformLinux, OSLinux, Version{0, 0, 0}}, DeviceComputer}},

	// Can't parse browser due to limitation of user agent library
	{"Mozilla/5.0 (X11; U; Linux x86_64; zh-TW; rv:1.9.0.8) Gecko/2009032712 Ubuntu/8.04 (hardy) Firefox/3.0.8 GTB5",
		UserAgent{
			Browser{BrowserFirefox, Version{3, 0, 8}}, OS{PlatformLinux, OSLinux, Version{0, 0, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (compatible; Konqueror/3.5; Linux; x86_64) KHTML/3.5.5 (like Gecko) (Debian)",
		UserAgent{
			Browser{BrowserUnknown, Version{0, 0, 0}}, OS{PlatformLinux, OSLinux, Version{0, 0, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (X11; U; Linux i686; de; rv:1.9.1.5) Gecko/20091112 Iceweasel/3.5.5 (like Firefox/3.5.5; Debian-3.5.5-1)",
		UserAgent{
			Browser{BrowserFirefox, Version{3, 5, 5}}, OS{PlatformLinux, OSLinux, Version{0, 0, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.4; en-US; rv:1.9.1b3pre) Gecko/20090223 SeaMonkey/2.0a3",
		UserAgent{
			Browser{BrowserFirefox, Version{0, 0, 0}}, OS{PlatformMac, OSMacOS, Version{10, 4, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_5; en-us) AppleWebKit/525.26.2 (KHTML, like Gecko) Version/3.2 Safari/525.26.12",
		UserAgent{
			Browser{BrowserSafari, Version{3, 2, 0}}, OS{PlatformMac, OSMacOS, Version{10, 5, 5}}, DeviceComputer}},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en; rv:1.9.0.8pre) Gecko/2009022800 Camino/2.0b3pre",
		UserAgent{
			Browser{BrowserUnknown, Version{0, 0, 0}}, OS{PlatformMac, OSMacOS, Version{10, 5, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_2; en-US) AppleWebKit/533.1 (KHTML, like Gecko) Chrome/5.0.329.0 Safari/533.1",
		UserAgent{
			Browser{BrowserChrome, Version{5, 0, 329}}, OS{PlatformMac, OSMacOS, Version{10, 6, 2}}, DeviceComputer}},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.1.6) Gecko/20091201 Firefox/3.5.6 (.NET CLR 3.5.30729)",
		UserAgent{
			Browser{BrowserFirefox, Version{3, 5, 6}}, OS{PlatformMac, OSMacOS, Version{10, 6, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/534.52.7 (KHTML, like Gecko) Version/5.1.2 Safari/534.52.7",
		UserAgent{
			Browser{BrowserSafari, Version{5, 1, 2}}, OS{PlatformMac, OSMacOS, Version{10, 7, 2}}, DeviceComputer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.7; rv:9.0) Gecko/20111222 Thunderbird/9.0.1",
		UserAgent{
			Browser{BrowserUnknown, Version{0, 0, 0}}, OS{PlatformMac, OSMacOS, Version{10, 7, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/535.7 (KHTML, like Gecko) Chrome/16.0.912.75 Safari/535.7",
		UserAgent{
			Browser{BrowserChrome, Version{16, 0, 912}}, OS{PlatformMac, OSMacOS, Version{10, 7, 2}}, DeviceComputer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8) AppleWebKit/535.18.5 (KHTML, like Gecko) Version/5.2 Safari/535.18.5",
		UserAgent{
			Browser{BrowserSafari, Version{5, 2, 0}}, OS{PlatformMac, OSMacOS, Version{10, 8, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_8; en-US) AppleWebKit/532.5 (KHTML, like Gecko) Chrome/4.0.249.0 Safari/532.5",
		UserAgent{
			Browser{BrowserChrome, Version{4, 0, 249}}, OS{PlatformMac, OSMacOS, Version{10, 8, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9) AppleWebKit/537.35.1 (KHTML, like Gecko) Version/6.1 Safari/537.35.1",
		UserAgent{
			Browser{BrowserSafari, Version{6, 1, 0}}, OS{PlatformMac, OSMacOS, Version{10, 9, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10) AppleWebKit/538.34.48 (KHTML, like Gecko) Version/8.0 Safari/538.35.8",
		UserAgent{
			Browser{BrowserSafari, Version{8, 0, 0}}, OS{PlatformMac, OSMacOS, Version{10, 10, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10) AppleWebKit/538.32 (KHTML, like Gecko) Version/7.1 Safari/538.4",
		UserAgent{
			Browser{BrowserSafari, Version{7, 1, 0}}, OS{PlatformMac, OSMacOS, Version{10, 10, 0}}, DeviceComputer}},

	{"Opera/9.80 (S60; SymbOS; Opera Mobi/352; U; de) Presto/2.4.15 Version/10.00",
		UserAgent{
			Browser{BrowserOpera, Version{10, 0, 0}}, OS{PlatformUnknown, OSUnknown, Version{0, 0, 0}}, DevicePhone}},

	{"Opera/9.80 (S60; SymbOS; Opera Mobi/352; U; de) Presto/2.4.15 Version/10.00",
		UserAgent{
			Browser{BrowserOpera, Version{10, 0, 0}}, OS{PlatformUnknown, OSUnknown, Version{0, 0, 0}}, DevicePhone}},

	{"Mozilla/4.0 (compatible; MSIE 5.01; Windows NT 5.0; SV1; .NET CLR 1.1.4322; .NET CLR 1.0.3705; .NET CLR 2.0.50727)",
		UserAgent{
			Browser{BrowserIE, Version{5, 0, 1}}, OS{PlatformWindows, OSWindows, Version{5, 0, 0}}, DeviceComputer}},

	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/4.0; GTB6.4; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; OfficeLiveConnector.1.3; OfficeLivePatch.0.0; .NET CLR 1.1.4322)",
		UserAgent{
			Browser{BrowserIE, Version{7, 0, 0}}, OS{PlatformWindows, OSWindows, Version{6, 1, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Windows; U; Windows NT 6.1; sk; rv:1.9.1.7) Gecko/20091221 Firefox/3.5.7",
		UserAgent{
			Browser{BrowserFirefox, Version{3, 5, 7}}, OS{PlatformWindows, OSWindows, Version{6, 1, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)",
		UserAgent{
			Browser{BrowserIE, Version{10, 0, 0}}, OS{PlatformWindows, OSWindows, Version{6, 2, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/536.5 (KHTML, like Gecko) YaBrowser/1.0.1084.5402 Chrome/19.0.1084.5402 Safari/536.5",
		UserAgent{
			Browser{BrowserYandex, Version{1, 0, 1084}}, OS{PlatformWindows, OSWindows, Version{6, 2, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.15 (KHTML, like Gecko) Chrome/24.0.1295.0 Safari/537.15",
		UserAgent{
			Browser{BrowserChrome, Version{24, 0, 1295}}, OS{PlatformWindows, OSWindows, Version{6, 2, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko",
		UserAgent{
			Browser{BrowserIE, Version{11, 0, 0}}, OS{PlatformWindows, OSWindows, Version{6, 3, 0}}, DeviceTablet}},

	{"Mozilla/5.0 (IE 11.0; Windows NT 6.3; Trident/7.0; .NET4.0E; .NET4.0C; rv:11.0) like Gecko",
		UserAgent{
			Browser{BrowserIE, Version{11, 0, 0}}, OS{PlatformWindows, OSWindows, Version{6, 3, 0}}, DeviceComputer}},

	{"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; SLCC1; .NET CLR 2.0.50727; .NET CLR 1.1.4322; InfoPath.2; .NET CLR 3.5.21022; .NET CLR 3.5.30729; MS-RTC LM 8; OfficeLiveConnector.1.4; OfficeLivePatch.1.3; .NET CLR 3.0.30729)",
		UserAgent{
			Browser{BrowserIE, Version{8, 0, 0}}, OS{PlatformWindows, OSWindows, Version{6, 0, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Windows; U; Windows NT 5.1; cs; rv:1.9.1.8) Gecko/20100202 Firefox/3.5.8",
		UserAgent{
			Browser{BrowserFirefox, Version{3, 5, 8}}, OS{PlatformWindows, OSWindows, Version{5, 1, 0}}, DeviceComputer}},

	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; )",
		UserAgent{
			Browser{BrowserIE, Version{7, 0, 0}}, OS{PlatformWindows, OSWindows, Version{5, 1, 0}}, DeviceComputer}},

	// desktop mode for Windows Phone 7
	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; XBLWP7; ZuneWP7)",
		UserAgent{
			Browser{BrowserIE, Version{7, 0, 0}}, OS{PlatformWindows, OSWindows, Version{6, 1, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) coc_coc_browser/42.0 CoRom/36.0.1985.144 Chrome/36.0.1985.144 Safari/537.36",
		UserAgent{
			Browser{BrowserCocCoc, Version{42, 0, 0}}, OS{PlatformWindows, OSWindows, Version{6, 3, 0}}, DeviceComputer}},

	{"Mozilla/5.0 (compatible; coccocbot/1.0; +http://help.coccoc.com/searchengine)",
		UserAgent{
			Browser{BrowserCocCocBot, Version{0, 0, 0}}, OS{PlatformBot, OSBot, Version{0, 0, 0}}, DeviceComputer}},
}

func TestAgentSurfer(t *testing.T) {
	for _, determined := range testUAVars {
		t.Run("", func(t *testing.T) {
			testFuncs := []func(string) *UserAgent{
				Parse,
				func(ua string) *UserAgent {
					u := new(UserAgent)
					ParseUserAgent(ua, u)
					return u
				},
			}

			for _, f := range testFuncs {
				ua := f(determined.UA)

				if ua.Browser.Name != determined.Browser.Name {
					t.Errorf("browserName: got %v, wanted %v", ua.Browser.Name, determined.Browser.Name)
					t.Logf("agent: %s", determined.UA)
				}

				if ua.Browser.Version != determined.Browser.Version {
					t.Errorf("browser version: got %d, wanted %d", ua.Browser.Version, determined.Browser.Version)
					t.Logf("agent: %s", determined.UA)
				}

				if ua.OS.Platform != determined.OS.Platform {
					t.Errorf("platform: got %v, wanted %v", ua.OS.Platform, determined.OS.Platform)
					t.Logf("agent: %s", determined.UA)
				}

				if ua.OS.Name != determined.OS.Name {
					t.Errorf("os: got %s, wanted %s", ua.OS.Name, determined.OS.Name)
					t.Logf("agent: %s", determined.UA)
				}

				if ua.OS.Version != determined.OS.Version {
					t.Errorf("os version: got %d, wanted %d", ua.OS.Version, determined.OS.Version)
					t.Logf("agent: %s", determined.UA)
				}

				if ua.DeviceType != determined.DeviceType {
					t.Errorf("device type: got %v, wanted %v", ua.DeviceType, determined.DeviceType)
					t.Logf("agent: %s", determined.UA)
				}
			}
		})
	}
}

func BenchmarkAgentSurfer(b *testing.B) {
	num := len(testUAVars)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(testUAVars[i%num].UA)
	}
}

func BenchmarkAgentSurferReuse(b *testing.B) {
	dest := new(UserAgent)
	num := len(testUAVars)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dest.Reset()
		ParseUserAgent(testUAVars[i%num].UA, dest)
	}
}

func BenchmarkEvalSystem(b *testing.B) {
	num := len(testUAVars)
	v := UserAgent{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.evalOS(testUAVars[i%num].UA)
	}
}

func BenchmarkEvalBrowserName(b *testing.B) {
	num := len(testUAVars)
	v := UserAgent{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.evalBrowserName(testUAVars[i%num].UA)
	}
}

func BenchmarkEvalBrowserVersion(b *testing.B) {
	num := len(testUAVars)
	v := UserAgent{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.Browser.Name = testUAVars[i%num].Browser.Name
		v.evalBrowserVersion(testUAVars[i%num].UA)
	}
}

func BenchmarkEvalDevice(b *testing.B) {
	num := len(testUAVars)
	v := UserAgent{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.OS.Name = testUAVars[i%num].OS.Name
		v.OS.Platform = testUAVars[i%num].OS.Platform
		v.Browser.Name = testUAVars[i%num].Browser.Name
		v.evalDevice(testUAVars[i%num].UA)
	}
}

// Chrome for Mac
func BenchmarkParseChromeMac(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.130 Safari/537.36")
	}
}

// Chrome for Windows
func BenchmarkParseChromeWin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse("Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.134 Safari/537.36")
	}
}

// Chrome for Android
func BenchmarkParseChromeAndroid(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse("Mozilla/5.0 (Linux; Android 4.4.2; GT-P5210 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.93 Safari/537.36")
	}
}

// Safari for Mac
func BenchmarkParseSafariMac(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/600.7.12 (KHTML, like Gecko) Version/8.0.7 Safari/600.7.12")
	}
}

// Safari for iPad
func BenchmarkParseSafariiPad(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse("Mozilla/5.0 (iPad; CPU OS 8_1_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B440 Safari/600.1.4")
	}
}

func TestStringTrimPrefix(t *testing.T) {
	testCases := []struct {
		f        func() string
		expected string
	}{
		{
			f:        DeviceUnknown.StringTrimPrefix,
			expected: "Unknown",
		},
		{
			f:        BrowserUnknown.StringTrimPrefix,
			expected: "Unknown",
		},
		{
			f:        OSUnknown.StringTrimPrefix,
			expected: "Unknown",
		},
		{
			f:        PlatformUnknown.StringTrimPrefix,
			expected: "Unknown",
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			s := tc.f()
			if tc.expected != s {
				t.Fatalf("Expected %q, got %q", tc.expected, s)
			}
		})
	}
}
