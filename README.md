# <img src="https://raw.githubusercontent.com/Kiuryy/uaparser/master/__asset/logo-small.png" height="88" align="left" /> UAParser

[![Build Status](https://travis-ci.org/Kiuryy/uaparser.svg?branch=master)](https://travis-ci.org/Kiuryy/uaparser)
[![Code Coverage](https://img.shields.io/codecov/c/github/Kiuryy/uaparser)](https://codecov.io/gh/Kiuryy/uaparser)
[![GoDoc](https://godoc.org/github.com/Kiuryy/uaparser?status.svg)](https://godoc.org/github.com/Kiuryy/uaparser)
[![Go Report Card](https://goreportcard.com/badge/github.com/Kiuryy/uaparser)](https://goreportcard.com/report/github.com/Kiuryy/uaparser)
[![License: Apache v2](https://img.shields.io/badge/License-Apache%20v2-lightgray.svg)](https://www.apache.org/licenses/LICENSE-2.0)

---

`UAParser` is a lightweight Golang package that parses and abstracts [HTTP User-Agent strings](https://en.wikipedia.org/wiki/User_agent) with the focus on the currently popular browsers and OS.

This project is a fork of `uasurfer`. For more information, you can take a look at their [repository](https://github.com/avct/uasurfer).

## Usage

### Parse(ua string) Function

The `Parse()` function accepts a user agent `string` and returns UserAgent struct with named constants and integers for versions (minor, major and patch separately). A string can be retrieved by adding `.String()` to a variable, such as `uaparser.BrowserName.String()`.

```
// Define a user agent string
uaStr := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3824.6 Safari/537.36"

// Parse() returns an UserAgent object
ua := uaparser.Parse(uaStr)
```

where example UserAgent is:
```
{
    Browser {
        BrowserName: BrowserChrome,
        Version: {
            Major: 77,
            Minor: 0,
            Patch: 3824,
        },
    },
    OS {
        Platform: PlatformMac,
        Name: OSMacOS,
        Version: {
            Major: 10,
            Minor: 15,
            Patch: 1,
        },
        VersionAlias: "Catalina",
    },
    DeviceType: DeviceComputer,
}
```

**Usage note:** There are some OS that do not return a version, see docs below. Linux is typically not reported with a specific Linux distro name or version.

#### Browser Name
* `BrowserChrome` - Google [Chrome](https://en.wikipedia.org/wiki/Google_Chrome)
* `BrowserSafari` - Apple [Safari](https://en.wikipedia.org/wiki/Safari_(web_browser)), Google Search ([GSA](https://itunes.apple.com/us/app/google/id284815942))
* `BrowserEdge` - Microsoft [Edge](https://en.wikipedia.org/wiki/Microsoft_Edge)
* `BrowserFirefox` - Mozilla [Firefox](https://en.wikipedia.org/wiki/Firefox)
* `BrowserIE` - Microsoft [Internet Explorer](https://en.wikipedia.org/wiki/Internet_Explorer)
* `BrowserAndroid` - Android [WebView](https://developer.chrome.com/multidevice/webview/overview) (Android OS <4.4 only)
* `BrowserOpera` - [Opera](https://en.wikipedia.org/wiki/Opera_(web_browser))
* `BrowserChromium` - Google [Chromium](https://en.wikipedia.org/wiki/Chromium_(web_browser))
* `BrowserUCBrowser` - [UC Browser](https://en.wikipedia.org/wiki/UC_Browser)
* `BrowserQQ` - Tencent [QQ](https://en.wikipedia.org/wiki/Tencent_QQ)
* `BrowserYandex` - [Yandex](https://en.wikipedia.org/wiki/Yandex_Browser)
* `BrowserSamsung` - [Samsung Internet](https://en.wikipedia.org/wiki/Samsung_Internet)
* `BrowserCocCoc`- [Cốc Cốc](https://en.wikipedia.org/wiki/C%E1%BB%91c_C%E1%BB%91c)
* `BrowserUnknown` - Unknown

#### Browser Version

Browser version returns an `unint8` of the major version attribute of the User-Agent String. For example Chrome 45.0.23423 would return `45`. The intention is to support math operators with versions, such as "do XYZ for Chrome version >23".

Unknown version is returned as `0`.

#### Platform
* `PlatformWindows` - Microsoft Windows
* `PlatformMac` - Apple Mac
* `PlatformLinux` - Linux, including Android and other OS
* `PlatformiPad` - Apple iPad
* `PlatformiPhone` - Apple iPhone
* `PlatformUnknown` - Unknown

#### OS Name
* `OSWindows`
* `OSMacOS` - includes "macOS and OS X"
* `OSiOS`
* `OSAndroid`
* `OSChromeOS`
* `OSLinux`
* `OSUnknown`

#### OS Version

MacOS major version is always 10 with consecutive minor versions indicating release releases (10 - Yosemite, 11 - El Capitain, 12 Sierra, etc). Windows version is NT version. `Version{0, 0, 0}` indicated version is unknown or not evaluated.
Versions can be compared using `Less` function: `if ver1.Less(ver2) {}`

Here are some examples across the platform, os.name, and os.version:

* For Windows XP (Windows NT 5.1), "`PlatformWindows`" is the platform, "`OSWindows`" is the name, and `{5, 1, 0}` the version.
* For OS X 10.5.1, "`PlatformMac`" is the platform, "`OSMacOS`" the name, and `{10, 5, 1}` the version.
* For Android 5.1, "`PlatformLinux`" is the platform, "`OSAndroid`" is the name, and `{5, 1, 0}` the version.
* For iOS 5.1, "`PlatformiPhone`" or "`PlatformiPad`" is the platform, "`OSiOS`" is the name, and `{5, 1, 0}` the version.

#### OS Version Alias

Some OS use version aliases, which are commonly used to identify specific versions of OS. Besides the internal representation of the version, `UAParser` also returns these version aliases. 

###### macOS Version Aliases

* macOS 10.15 - `Catalina` - `{10, 15, 0}`
* macOS 10.14 - `Mojave` - `{10, 14, 0}`
* macOS 10.13 - `High Sierra` - `{10, 13, 0}`
* macOS 10.12 - `Sierra` - `{10, 12, 0}`
* OS X 10.11 - `El Capitan` -  `{10, 11, 0}`
* OS X 10.10 - `Yosemite` -  `{10, 10, 0}`
* ...

###### Android Version Aliases

* Android 9 - `Pie` - `{9, 0, 0}`
* Android 8 - `Oreo` - `{8, 0, 0}`
* Android 7 - `Nougat` - `{7, 0, 0}`
* Android 6 - `Marshmallow` - `{6, 0, 0}`
* Android 5 - `Lollipop` - `{5, 0, 0}`
* Android 4.4 - `KitKat` - `{4, 4, 0}`
* ...

###### Windows Version Aliases

* Windows 10 - `{10, 0, 0}`
* Windows 8.1 - `{6, 3, 0}`
* Windows 8 - `{6, 2, 0}`
* Windows 7 - `{6, 1, 0}`
* Windows Vista - `{6, 0, 0}`
* Windows XP - `{5, 1, 0}` or `{5, 2, 0}`
* Windows 2000 - `{5, 0, 0}`

Windows 95, 98, and ME represent 0.01% of traffic worldwide and are not available through this package.

#### DeviceType
DeviceType is typically quite accurate, though determining between phones and tablets on Android is not always possible due to how some vendors design their UA strings. A mobile Android device without tablet indicator defaults to being classified as a phone. DeviceTV supports major brands such as Philips, Sharp, Vizio and steaming boxes such as Apple, Google, Roku, Amazon.

* `DeviceComputer`
* `DevicePhone`
* `DeviceTablet`
* `DeviceTV`
* `DeviceConsole`
* `DeviceWearable`
* `DeviceUnknown`
