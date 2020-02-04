package uaparser

import (
	"fmt"
	test_datasets_version "github.com/Kiuryy/uaparser/test_datasets/version"
	"testing"
)

func TestVersionCompare(t *testing.T) {
	datasets := test_datasets_version.Datasets_Version

	for i, dataset := range datasets {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			result1 := dataset.Version1.Less(dataset.Version2)

			if result1 != dataset.Less {
				t.Errorf("Version comparison failed: %v and %v should be %v, got %v", dataset.Version1, dataset.Version2, dataset.Less, result1)
			}

			if dataset.Version1 != dataset.Version2 { // check the reverse comparison, too, if the versions are not the same
				result2 := dataset.Version2.Less(dataset.Version1)

				if result2 == dataset.Less {
					t.Errorf("Version comparison failed: %v and %v should be %v, got %v", dataset.Version2, dataset.Version1, !dataset.Less, result2)
				}
			}
		})
	}
}
