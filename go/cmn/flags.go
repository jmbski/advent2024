/*
Copyright 2024 Joseph Bochinski

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the “Software”), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

********************************************************************************

	Package: cmn
	Title: flags
	Description: Collection of functions to retrieve cmd flag values
	Author: Joseph Bochinski
	Date: 2024-12-12

********************************************************************************
*/
package cmn

// GetFlagBool retrieves the bool value of the flag, defaults to false
func GetFlagBool(flagName string) bool {
	if value, err := ActiveCmd.Flags().GetBool(flagName); err != nil {
		return false
	} else {
		return value
	}
}

// GetFlagBoolD retrieves the bool value of the flag, defaults to the provided value
func GetFlagBoolD(flagName string, defaultVal bool) bool {
	if value, err := ActiveCmd.Flags().GetBool(flagName); err != nil {
		return defaultVal
	} else {
		return value
	}
}

// GetFlagInt retrieves the int value of the flag, defaults to 0
func GetFlagInt(flagName string) int {
	if value, err := ActiveCmd.Flags().GetInt(flagName); err != nil {
		return 0
	} else {
		return value
	}
}

// GetFlagIntD retrieves the int value of the flag, defaults to the provided value
func GetFlagIntD(flagName string, defaultVal int) int {
	if value, err := ActiveCmd.Flags().GetInt(flagName); err != nil {
		return defaultVal
	} else {
		return value
	}
}

// GetFlagString retrieves the string value of the flag, defaults to ""
func GetFlagString(flagName string) string {
	if value, err := ActiveCmd.Flags().GetString(flagName); err != nil {
		return ""
	} else {
		return value
	}
}

// GetFlagStringD retrieves the string value of the flag, defaults to the provided value
func GetFlagStringD(flagName, defaultVal string) string {
	if value, err := ActiveCmd.Flags().GetString(flagName); err != nil {
		return defaultVal
	} else {
		return value
	}
}
