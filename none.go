package utils

// Null is the zero length struct which usually used to check item existence along with None.
//
// Example:
//
//	nullMap := make(map[string]utils.Null)
//	nullMap["foo"] = utils.None
type Null struct{}

var None Null
