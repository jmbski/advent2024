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
	Title: profile
	Description: Code for running task profiling
	Author: Joseph Bochinski
	Date: 2024-12-10

********************************************************************************
*/
package cmn

import (
	"fmt"
	"time"
)

// Profile logs the time it took to execute a task with the given name and start time
func Profile(start time.Time, name string) {
	elapsed := time.Since(start)
	elapsedSeconds := elapsed.Seconds()
	fmt.Printf("Process [%s] took %.5f seconds\n", name, elapsedSeconds)
}

// StartProfile returns a function that logs the time it took to execute a task with the given name
// this can be used in combination defer to log the time it took to execute a task
func StartProfile(taskName string) func() {
	start := time.Now()

	return func() {
		Profile(start, taskName)
	}
}
