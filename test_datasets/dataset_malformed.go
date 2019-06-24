package test_datasets

import (
	"github.com/Kiuryy/uaparser/useragent"
)

var Datasets_Malformed = []TestDataset{
	// Empty
	{"",
		useragent.UserAgent{},
	},

	// Single char
	{"a",
		useragent.UserAgent{},
	},

	// Some random string
	{"some random string",
		useragent.UserAgent{},
	},

	// Potentially malformed ua
	{")(",
		useragent.UserAgent{},
	},
}
