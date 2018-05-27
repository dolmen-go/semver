package semver

import (
	"errors"
	"strconv"
)

var ErrInvalidVersion = errors.New("invalid version")

// Version represent a semver style version ("v1.2.4").
type Version struct {
	major Major
	minor uint
	patch uint
}

// V is a constructor for Version
func V(major Major, minor uint, patch uint) Version {
	return Version{major, minor, patch}
}

func (v Version) Major() Major {
	return v.major
}

func (v Version) Minor() uint {
	return v.minor
}

func (v Version) Patch() uint {
	return v.patch
}

func (v Version) String() string {
	// Allocate for the minimal (and most common) length
	b := make([]byte, 1, 6) // v0.0.0
	b[0] = 'v'
	b = strconv.AppendUint(b, uint64(v.major), 10)
	b = append(b, '.')
	b = strconv.AppendUint(b, uint64(v.minor), 10)
	b = append(b, '.')
	return string(strconv.AppendUint(b, uint64(v.patch), 10))
}

func parseDotUint(str string) (uint, error) {
	return 0, nil
}

func (v *Version) Set(str string) error {
	if len(str) < 6 || str[0] != 'v' {
		return ErrInvalidVersion
	}
	var err error
	var major, minor, patch uint64
	if str[1] == '0' {
		str = str[2:]
	} else {
		if str[1] < '1' || str[1] > '9' {
			return ErrInvalidVersion
		}
		i := 2
		for str[i] != '.' {
			if str[i] < '0' || str[i] > '9' {
				return ErrInvalidVersion
			}
			i++
			if i >= len(str)-4 {
				return ErrInvalidVersion
			}
		}
		if major, err = strconv.ParseUint(str[1:i], 10, 0); err != nil {
			return ErrInvalidVersion
		}
		str = str[i:]
	}
	if str[0] != '.' {
		return ErrInvalidVersion
	}
	if str[1] == '0' {
		str = str[2:]
	} else {
		if str[1] < '1' || str[1] > '9' {
			return ErrInvalidVersion
		}
		i := 2
		for str[i] != '.' {
			i++
			if i >= len(str)-2 {
				return ErrInvalidVersion
			}
		}
		if minor, err = strconv.ParseUint(str[1:i], 10, 0); err != nil {
			return ErrInvalidVersion
		}
		str = str[i:]
	}
	if str[0] != '.' {
		return ErrInvalidVersion
	}
	if str[1] == '0' {
		if len(str) != 2 {
			return ErrInvalidVersion
		}
	} else {
		if len(str) < 2 || str[1] < '1' || str[1] > '9' {
			return ErrInvalidVersion
		}
		if patch, err = strconv.ParseUint(str[1:], 10, 0); err != nil {
			return ErrInvalidVersion
		}
	}
	*v = Version{Major(major), uint(minor), uint(patch)}
	return nil
}

type Major uint

func (m Major) String() string {
	b := make([]byte, 1, 3)
	b[0] = 'v'
	return string(strconv.AppendUint(b, uint64(m), 10))
}

func (m *Major) Set(str string) error {
	if len(str) < 2 || str[0] != 'v' {
		return ErrInvalidVersion
	}
	n, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		return ErrInvalidVersion
	}
	*m = Major(n)
	return nil
}
