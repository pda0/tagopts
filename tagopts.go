// Package tagopts is a small package containing functions that make it
// easier to use struct tags.
package tagopts

import "strings"

// TagOptions is a comma-separated string containing the values ​​after
// the first comma in the struct tag. Or an empty string if there are
// no commas in the tag.
type TagOptions []string

// ParseTag reads the tag value at the first comma, returning the first
// value as the tag value and the rest as options. If the tag value
// is not specified, the default value is returned.
//   - tag: struct field tag value.
//   - defaultValue: the default value.
func ParseTag(tag, defaultValue string) (string, TagOptions) {
	parsed := strings.Split(tag, ",")

	if parsed[0] != "" {
		return parsed[0], parsed[1:]
	} else {
		return defaultValue, parsed[1:]
	}
}

// Contains allows you to check whether the specified option is in
// the tag options.
func (o TagOptions) Contains(option string) bool {
	if len(o) > 0 {
		for _, opt := range o {
			if opt == option {
				return true
			}
		}
	}

	return false
}

// Count returns the number of tag options.
func (o TagOptions) Count() int {
	return len(o)
}

// Option returns the value of an option by its index or an empty string
// if there is no option with that index value.
func (o TagOptions) Option(optIndex int) string {
	if optIndex >= 0 && optIndex < len(o) {
		return o[optIndex]
	} else {
		return ""
	}
}
