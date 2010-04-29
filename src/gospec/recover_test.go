// Copyright © 2009-2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package gospec

import (
	"nanospec"
)


func RecoverSpec(c nanospec.Context) {

	c.Specify("When the called function panics", func() {
		err := recoverOnPanic(boom2)

		c.Specify("the cause is returned", func() {
			c.Expect(err.Cause).Equals("boom!")
		})
		c.Specify("the stack trace begins with the panicking line", func() {
			c.Expect(err.StackTrace[0].Name()).Equals("gospec.boom0")
		})
		c.Specify("the stack trace includes all parent functions", func() {
			c.Expect(err.StackTrace[1].Name()).Equals("gospec.boom1")
		})
		c.Specify("the stack trace ends with the called function", func() {
			lastEntry := err.StackTrace[len(err.StackTrace)-1]
			c.Expect(lastEntry.Name()).Equals("gospec.boom2")
		})
	})

	c.Specify("When the called function does not panic", func() {
		err := recoverOnPanic(noBoom)

		c.Specify("there is no error", func() {
			c.Expect(err == nil).IsTrue()
		})
	})
}

func boom2() {
	boom1()
}
func boom1() {
	boom0()
}
func boom0() {
	panic("boom!")
}
func noBoom() {
}