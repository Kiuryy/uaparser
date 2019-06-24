package test_datasets

import (
	"github.com/Kiuryy/uaparser/useragent"
)

type TestDataset struct {
	UA string
	useragent.UserAgent
}
