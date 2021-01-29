package vars

import (
	"fmt"
	"github.com/Kiuryy/uaparser/version"
	"strconv"
	"strings"
)

// OS contains the name, version and platform of the OS
type OS struct {
	Platform     Platform
	Name         OSName
	Version      version.Version
	VersionAlias string
}

// String returns the name of the OS in human readable form (e.g. 'macOS Catalina 10.15')
func (o *OS) String() string {
	name := strings.TrimPrefix(o.Name.String(), "OS")
	_version := strconv.Itoa(o.Version.Major)

	if len(name) == 0 || name == "Unknown" || name == "Bot" { // could not determine OS
		return ""
	}

	if o.Version.Minor > 0 && (o.Platform != PlatformMac || o.Version.Major < 11) { // add minor version, except for macOS 11+
		_version = fmt.Sprintf("%s.%d", _version, o.Version.Minor)
	}

	if o.Platform == PlatformMac {
		name = "macOS"

		if o.Version.Major == 10 && o.Version.Minor < 12 { // handle different macOS names, 'OS X' for _version <= 10.11, 'macOS' for _version >= 10.12
			name = "OS X"
		}
	}

	if len(o.VersionAlias) > 0 { // alias for this version of the OS exists
		if o.Platform == PlatformWindows { // Replace internal Windows version with alias
			_version = o.VersionAlias
		} else if o.Platform == PlatformMac && o.Version.Major < 11 { // prepend alias string to the version for macOS < 11
			_version = fmt.Sprintf("%s %s", o.VersionAlias, _version)
		} else { // append alias string to the version for all other platforms
			_version = fmt.Sprintf("%s %s", _version, o.VersionAlias)
		}
	} else if _version == "0" {
		_version = ""
	}

	ret := fmt.Sprintf("%s %s", name, _version)
	return strings.Trim(ret, " ")
}
