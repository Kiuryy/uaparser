package test_datasets

import (
	"github.com/Kiuryy/uaparser/vars"
)

type TestDataset struct {
	UA string
	vars.UserAgent
	OSAlias string
}
