package test_datasets

import (
	"github.com/Kiuryy/uaparser/vars"
	"github.com/Kiuryy/uaparser/version"
)

type TestParserDataset struct {
	UA string
	vars.UserAgent
	OSAlias string
}

type TestVersionDataset struct {
	Version1 version.Version
	Version2 version.Version
	Less bool
}
