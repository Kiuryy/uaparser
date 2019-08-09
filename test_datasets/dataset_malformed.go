package test_datasets

import (
	"github.com/Kiuryy/uaparser/vars"
)

var Datasets_Malformed = []TestDataset{
	// Empty
	{"",
		vars.UserAgent{},
		"",
	},

	// Single char
	{"a",
		vars.UserAgent{},
		"",
	},

	// Some random string
	{"some random string",
		vars.UserAgent{},
		"",
	},

	// Potentially malformed ua
	{")(",
		vars.UserAgent{},
		"",
	},
}
