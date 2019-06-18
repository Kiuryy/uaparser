package version

import (
	"strings"
)

// Version contains three segments: 'Major', 'Minor' and 'Patch'
type Version struct {
	Major int
	Minor int
	Patch int
}

// Less compares the given version with the current version object: `if ver1.Less(ver2) {}`
func (v Version) Less(c Version) bool {
	if v.Major < c.Major {
		return true
	}

	if v.Major > c.Major {
		return false
	}

	if v.Minor < c.Minor {
		return true
	}

	if v.Minor > c.Minor {
		return false
	}

	return v.Patch < c.Patch
}

func (v *Version) FindVersionNumber(s string, m string) bool {
	if ind := strings.Index(s, m); ind != -1 {
		return v.Parse(s[ind+len(m):])
	}
	return false
}

// parse accepts a string and extracts the version information
// with {0, 0, 0} being default.
func (v *Version) Parse(str string) bool {
	if len(str) == 0 || str[0] < '0' || str[0] > '9' {
		return false
	}
	for i := 0; i < 3; i++ {
		empty := true
		val := 0
		l := len(str) - 1

		for k, c := range str {
			if c >= '0' && c <= '9' {
				if empty {
					val = int(c) - 48
					empty = false
					if k == l {
						str = str[:0]
					}
					continue
				}

				if val == 0 {
					if c == '0' {
						if k == l {
							str = str[:0]
						}
						continue
					}
					str = str[k:]
					break
				}

				val = 10*val + int(c) - 48
				if k == l {
					str = str[:0]
				}
				continue
			}
			str = str[k+1:]
			break
		}

		switch i {
		case 0:
			v.Major = val

		case 1:
			v.Minor = val

		case 2:
			v.Patch = val
		}
	}
	return true
}
