package uaparser

import (
	"fmt"
	"github.com/Kiuryy/uaparser/test_datasets"
	testDatasetsBot "github.com/Kiuryy/uaparser/test_datasets/bot"
	testDatasetsBrowser "github.com/Kiuryy/uaparser/test_datasets/browser"
	testDatasetsSystem "github.com/Kiuryy/uaparser/test_datasets/system"
	"testing"
)

// validateParserResult will test the given datasets and checks whether the system, browser and platform, aswell as the versions are correct
func validateParserResult(t *testing.T, datasets []test_datasets.TestParserDataset) {
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

//
// PARSER TESTS -> GENERIC
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
// PARSER TESTS -> BROWSER
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

func TestBrowserOpera(t *testing.T) {
	validateParserResult(t, testDatasetsBrowser.Datasets_Opera)
}

func TestBrowserUCBrowser(t *testing.T) {
	validateParserResult(t, testDatasetsBrowser.Datasets_UCBrowser)
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
// PARSER TESTS -> SYSTEM
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
