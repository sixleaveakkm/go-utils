// Package ctime
package ctime

import "time"

var cNow customNow

type customNow struct {
	customFn   func() time.Time
	customTime *time.Time
}

// Now provides an alternative to time.Now() in order to write testable code with no pain.
// You can set a custom time by calling SetNow() or SetNowFunc() (SetNowFunc() has higher priority if both was set).
// By default, it will return time.Now().
func Now() time.Time {
	if cNow.customFn != nil {
		return cNow.customFn()
	}
	if cNow.customTime != nil {
		return *cNow.customTime
	}
	return time.Now()
}

// SetNow allows you to set a custom time you want to return by Now().
// Note SetNowFunc() has higher priority so if you set both, SetNowFunc() will be used.
func SetNow(t time.Time) {
	cNow.customTime = &t
}

// SetNowFunc allows you to return a custom time you want to return by Now().
// This has higher priority, SetNow() will be ignored.
func SetNowFunc(fn func() time.Time) {
	cNow.customFn = fn
}

// ResetNow cleans the custom time and custom time function, just in case you need.
func ResetNow() {
	cNow.customTime = nil
	cNow.customFn = nil
}
