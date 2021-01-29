package uaparser

import (
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

// getCompleteDataset returns a concatenated list of all test datasets
func getCompleteDataset() []test_datasets.TestParserDataset {
	var all []test_datasets.TestParserDataset

	for _, dataset := range [][]test_datasets.TestParserDataset{
		test_datasets.Datasets_Incomplete,
		test_datasets.Datasets_Malformed,
		//
		testDatasetsBot.Datasets_Bot,
		//
		testDatasetsBrowser.Datasets_AndroidWebview,
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
		testDatasetsSystem.Datasets_Console,
	} {
		all = append(all, dataset...)
	}

	return all
}

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
