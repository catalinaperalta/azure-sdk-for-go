// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package durations

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	p, err := ParseDuration("P123DT22H14M12.011S")
	if err != nil {
		t.Fatal(err)
	}
	if p != time.Duration(10707252011000000) {
		t.Fatalf("Did not receive the correct value for the duration.\nExpected: 10707252011000000, Received: %s", p.String())
	}
	p, err = ParseDuration("PT12H")
	if err != nil {
		t.Fatal(err)
	}
	if p != time.Duration(12*time.Hour) {
		t.Fatalf("Did not receive the correct value for the duration.\nExpected: 10707252011000000, Received: %s", p.String())
	}
	p, err = ParseDuration("PT0S")
	if err != nil {
		t.Fatal(err)
	}
	if p != time.Duration(0) {
		t.Fatalf("Did not receive the correct value for the duration.\nExpected: 10707252011000000, Received: %s", p.String())
	}
	p, err = ParseDuration("PT12.5H10S")
	if err != nil {
		t.Fatal(err)
	}
	if p != time.Duration(12.5*float64(time.Hour)+float64(10*time.Second)) {
		t.Fatalf("Did not receive the correct value for the duration.\nExpected: 10707252011000000, Received: %s", p.String())
	}
	p, err = ParseDuration("PT0.5S")
	if err != nil {
		t.Fatal(err)
	}
	if p != time.Duration(0.5*float64(time.Second)) {
		t.Fatalf("Did not receive the correct value for the duration.\nExpected: 10707252011000000, Received: %s", p.String())
	}
}
