package test_datasets_version

import (
	"github.com/Kiuryy/uaparser/test_datasets"
	"github.com/Kiuryy/uaparser/version"
)

var Datasets_Version = []test_datasets.TestVersionDataset{
	// simple checks
	{
		version.Version{0, 0, 0},
		version.Version{0, 0, 0},
		false,
	},
	{
		version.Version{1, 0, 0},
		version.Version{0, 0, 0},
		false,
	},
	{
		version.Version{0, 1, 0},
		version.Version{0, 0, 0},
		false,
	},
	{
		version.Version{0, 0, 1},
		version.Version{0, 0, 0},
		false,
	},
	{
		version.Version{1, 0, 0},
		version.Version{1, 0, 0},
		false,
	},
	{
		version.Version{1, 0, 0},
		version.Version{0, 1, 0},
		false,
	},
	{
		version.Version{1, 0, 0},
		version.Version{0, 0, 1},
		false,
	},
	// more complex checks
	{
		version.Version{0, 0, 1},
		version.Version{0, 0, 2},
		true,
	},
	{
		version.Version{325, 0, 1},
		version.Version{325, 0, 2},
		true,
	},
	{
		version.Version{5, 1000, 0},
		version.Version{6, 0, 0},
		true,
	},
	{
		version.Version{1, 1000, 0},
		version.Version{1, 1001, 0},
		true,
	},
	{
		version.Version{7, 0, 1000},
		version.Version{7, 1, 0},
		true,
	},
	// minus checks
	{
		version.Version{1, 0, 0},
		version.Version{-1, 0, 0},
		false,
	},
	{
		version.Version{-1, 0, 0},
		version.Version{-1, 0, 0},
		false,
	},
	{
		version.Version{1, 0, 0},
		version.Version{1, 0, -1},
		false,
	},
}
