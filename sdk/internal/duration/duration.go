// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package duration

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func parseFloat64(value string) float64 {
	if len(value) == 0 {
		return 0
	}
	parsed, err := strconv.ParseFloat(value[:len(value)-1], 64)
	if err != nil {
		return 0
	}
	return parsed
}

// parseISO8601 takes a slice of token strings that it will parse into a time.Duration.
func parseISO8601(matches []string) (time.Duration, error) {
	years := parseFloat64(matches[1]) * float64(time.Hour*24*365)
	months := parseFloat64(matches[2]) * float64(time.Hour*24*30)
	days := parseFloat64(matches[3]) * float64(time.Hour*24)
	hours := parseFloat64(matches[4]) * float64(time.Hour)
	minutes := parseFloat64(matches[5]) * float64(time.Minute)
	seconds := parseFloat64(matches[6]) * float64(time.Second)
	// problem here are years and months
	return time.Duration(years + months + days + hours + minutes + seconds), nil
}

// Parse will check for a valid duration format and parse according to the identified format in order
// to return a valid time.Duration. If a valid format is not received it will fail.
func Parse(s string) (time.Duration, error) {
	iso8601Format := regexp.MustCompile(`P(\d+[\.\d]+?Y)?(\d+[\.\d]+?M)?(\d+[\.\d]+?D)?T?(\d+[\.\d]+?H)?(\d+[\.\d]+?M)?(\d+[\.\d]+?S)?`)
	if matches := iso8601Format.FindStringSubmatch(s); len(matches) > 0 {
		return parseISO8601(matches)
	} else if d, err := time.ParseDuration(s); err == nil {
		return d, nil
	}
	return 0, errors.New("Could not find a valid duration format")
}

// ToISO8601 parses a time.Duration into a string in the format of ISO 8601 duration.
// This function provides up to millisecond precision.
func ToISO8601(duration time.Duration) string {
	s := "P"
	days, remainder := duration.Milliseconds()/8.64e+7, duration.Milliseconds()%8.64e+7
	hours, remainder := remainder/3.6e+6, remainder%3.6e+6
	minutes, remainder := remainder/60000, remainder%60000
	seconds, remainder := remainder/1000, remainder%1000
	if days > 0 {
		s += fmt.Sprintf("%dD", days)
	}
	if hours > 0 || minutes > 0 || seconds > 0 {
		s += "T"
		if hours > 0 {
			s += fmt.Sprintf("%dH", hours)
		}
		if minutes > 0 {
			s += fmt.Sprintf("%dM", minutes)
		}
		if seconds > 0 {
			if remainder > 0 {
				// get remainder up to millisecond precision
				sec := float64(remainder) / 1000
				sec += float64(seconds)
				floatStr := strconv.FormatFloat(sec, 'f', -1, 64)
				s += fmt.Sprintf("%sS", floatStr)
			} else {
				s += fmt.Sprintf("%dS", seconds)
			}
		}
	}
	return s
}
