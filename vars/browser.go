package vars

import (
	"github.com/Kiuryy/uaparser/version"
)

// Browser contains the name and version of the browser
type Browser struct {
	Name    BrowserName
	Version version.Version
}