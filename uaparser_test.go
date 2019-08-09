package uaparser

import (
	"fmt"
	"github.com/Kiuryy/uaparser/browser"
	"github.com/Kiuryy/uaparser/device"
	"github.com/Kiuryy/uaparser/system"
	"github.com/Kiuryy/uaparser/test_datasets"
	testDatasetsBot "github.com/Kiuryy/uaparser/test_datasets/bot"
	testDatasetsBrowser "github.com/Kiuryy/uaparser/test_datasets/browser"
	testDatasetsSystem "github.com/Kiuryy/uaparser/test_datasets/system"
	"github.com/Kiuryy/uaparser/vars"
	"testing"
)

//
//
//

// validateParserResult will test the given datasets and checks whether the system, browser and platform, aswell as the versions are correct
func validateParserResult(t *testing.T, datasets []test_datasets.TestDataset) {
	for i, dataset := range datasets {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			ua := Parse(dataset.UA)

			if ua.Browser.Name != dataset.Browser.Name {
				t.Errorf("Browser name: got %v, wanted %v", ua.Browser.Name, dataset.Browser.Name)
				t.Logf("agent: %s", dataset.UA)
			}

			if ua.Browser.Version != dataset.Browser.Version {
				t.Errorf("Browser version: got %d, wanted %d", ua.Browser.Version, dataset.Browser.Version)
				t.Logf("agent: %s", dataset.UA)
			}

			if ua.OS.Platform != dataset.OS.Platform {
				t.Errorf("Platform: got %v, wanted %v", ua.OS.Platform, dataset.OS.Platform)
				t.Logf("agent: %s", dataset.UA)
			}

			if ua.OS.Name != dataset.OS.Name {
				t.Errorf("OS: got %v, wanted %v", ua.OS.Name, dataset.OS.Name)
				t.Logf("agent: %s", dataset.UA)
			}

			if ua.OS.Version != dataset.OS.Version {
				t.Errorf("OS version: got %d, wanted %d", ua.OS.Version, dataset.OS.Version)
				t.Logf("agent: %s", dataset.UA)
			}

			if ua.OS.VersionAlias != dataset.OS.VersionAlias {
				t.Errorf("OS version alias: got %s, wanted %s", ua.OS.VersionAlias, dataset.OS.VersionAlias)
				t.Logf("agent: %s", dataset.UA)
			}

			if ua.OS.String() != dataset.OSAlias {
				t.Errorf("OS toString: got %s, wanted %s", ua.OS.String(), dataset.OSAlias)
				t.Logf("agent: %s", dataset.UA)
			}

			if ua.DeviceType != dataset.DeviceType {
				t.Errorf("Device type: got %v, wanted %v", ua.DeviceType, dataset.DeviceType)
				t.Logf("agent: %s", dataset.UA)
			}
		})
	}
}

// getCompleteDataset returns a concatenated list of all test datasets
func getCompleteDataset() []test_datasets.TestDataset {
	var all []test_datasets.TestDataset

	for _, dataset := range [][]test_datasets.TestDataset{
		test_datasets.Datasets_Incomplete,
		test_datasets.Datasets_Malformed,
		//
		testDatasetsBot.Datasets_Bot,
		//
		testDatasetsBrowser.Datasets_AndroidWebview,
		testDatasetsBrowser.Datasets_Brave,
		testDatasetsBrowser.Datasets_Chrome,
		testDatasetsBrowser.Datasets_Chromium,
		testDatasetsBrowser.Datasets_CocCoc,
		testDatasetsBrowser.Datasets_Edge,
		testDatasetsBrowser.Datasets_Firefox,
		testDatasetsBrowser.Datasets_IE,
		testDatasetsBrowser.Datasets_Opera,
		testDatasetsBrowser.Datasets_QQ,
		testDatasetsBrowser.Datasets_Safari,
		testDatasetsBrowser.Datasets_Samsung,
		testDatasetsBrowser.Datasets_UCBrowser,
		testDatasetsBrowser.Datasets_Vivaldi,
		testDatasetsBrowser.Datasets_Yandex,
		//
		testDatasetsSystem.Datasets_Android,
		testDatasetsSystem.Datasets_ChromeOS,
		testDatasetsSystem.Datasets_iPad,
		testDatasetsSystem.Datasets_iPhone,
		testDatasetsSystem.Datasets_iPod,
		testDatasetsSystem.Datasets_Linux,
		testDatasetsSystem.Datasets_Windows,
		testDatasetsSystem.Datasets_Mac,
		testDatasetsSystem.Datasets_TV,
	} {
		all = append(all, dataset...)
	}

	return all
}

//
// GENERIC TESTS
//

func TestSystemMalformed(t *testing.T) {
	validateParserResult(t, test_datasets.Datasets_Malformed)
}

func TestIncomplete(t *testing.T) {
	validateParserResult(t, test_datasets.Datasets_Incomplete)
}

func TestBot(t *testing.T) {
	validateParserResult(t, testDatasetsBot.Datasets_Bot)
}

//
// BROWSER TESTS
//

func TestBrowserChrome(t *testing.T) {
	validateParserResult(t, testDatasetsBrowser.Datasets_Chrome)
}

func TestBrowserChromium(t *testing.T) {
	validateParserResult(t, testDatasetsBrowser.Datasets_Chromium)
}

func TestBrowserSafari(t *testing.T) {
	validateParserResult(t, testDatasetsBrowser.Datasets_Safari)
}

func TestBrowserEdge(t *testing.T) {
	validateParserResult(t, testDatasetsBrowser.Datasets_Edge)
}

func TestBrowserFirefox(t *testing.T) {
	validateParserResult(t, testDatasetsBrowser.Datasets_Firefox)
}

func TestBrowserIE(t *testing.T) {
	validateParserResult(t, testDatasetsBrowser.Datasets_IE)
}

func TestBrowserAndroidWebview(t *testing.T) {
	validateParserResult(t, testDatasetsBrowser.Datasets_AndroidWebview)
}

func TestBrowserVivaldi(t *testing.T) {
	validateParserResult(t, testDatasetsBrowser.Datasets_Vivaldi)
}

func TestBrowserOpera(t *testing.T) {
	validateParserResult(t, testDatasetsBrowser.Datasets_Opera)
}

func TestBrowserUCBrowser(t *testing.T) {
	validateParserResult(t, testDatasetsBrowser.Datasets_UCBrowser)
}

func TestBrowserBrave(t *testing.T) {
	validateParserResult(t, testDatasetsBrowser.Datasets_Brave)
}

func TestBrowserYandex(t *testing.T) {
	validateParserResult(t, testDatasetsBrowser.Datasets_Yandex)
}

func TestBrowserSamsung(t *testing.T) {
	validateParserResult(t, testDatasetsBrowser.Datasets_Samsung)
}

func TestBrowserQQ(t *testing.T) {
	validateParserResult(t, testDatasetsBrowser.Datasets_QQ)
}

func TestBrowserCocCoc(t *testing.T) {
	validateParserResult(t, testDatasetsBrowser.Datasets_CocCoc)
}

//
// SYSTEM TESTS
//

func TestSystemAndroid(t *testing.T) {
	validateParserResult(t, testDatasetsSystem.Datasets_Android)
}

func TestSystemChromeOS(t *testing.T) {
	validateParserResult(t, testDatasetsSystem.Datasets_ChromeOS)
}

func TestSystemiPad(t *testing.T) {
	validateParserResult(t, testDatasetsSystem.Datasets_iPad)
}

func TestSystemiPhone(t *testing.T) {
	validateParserResult(t, testDatasetsSystem.Datasets_iPhone)
}

func TestSystemiPod(t *testing.T) {
	validateParserResult(t, testDatasetsSystem.Datasets_iPod)
}

func TestSystemLinux(t *testing.T) {
	validateParserResult(t, testDatasetsSystem.Datasets_Linux)
}

func TestSystemWindows(t *testing.T) {
	validateParserResult(t, testDatasetsSystem.Datasets_Windows)
}

func TestSystemMac(t *testing.T) {
	validateParserResult(t, testDatasetsSystem.Datasets_Mac)
}

func TestSystemTV(t *testing.T) {
	validateParserResult(t, testDatasetsSystem.Datasets_TV)
}

//
// BENCHMARKS
//

func BenchmarkParser(b *testing.B) {
	datasets := getCompleteDataset()
	num := len(datasets)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(datasets[i%num].UA)
	}
}

func BenchmarkEvalSystem(b *testing.B) {
	datasets := getCompleteDataset()
	num := len(datasets)
	v := vars.UserAgent{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		system.Eval(&v, datasets[i%num].UA)
	}
}

func BenchmarkEvalBrowserName(b *testing.B) {
	datasets := getCompleteDataset()
	num := len(datasets)
	v := vars.UserAgent{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		browser.EvalName(&v, datasets[i%num].UA)
	}
}

func BenchmarkEvalBrowserVersion(b *testing.B) {
	datasets := getCompleteDataset()
	num := len(datasets)
	v := vars.UserAgent{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.Browser.Name = datasets[i%num].Browser.Name
		browser.EvalVersion(&v, datasets[i%num].UA)
	}
}

func BenchmarkEvalDevice(b *testing.B) {
	datasets := getCompleteDataset()
	num := len(datasets)
	v := vars.UserAgent{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.OS.Name = datasets[i%num].OS.Name
		v.OS.Platform = datasets[i%num].OS.Platform
		v.Browser.Name = datasets[i%num].Browser.Name
		device.Eval(&v, datasets[i%num].UA)
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
