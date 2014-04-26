package errors

import (
	. "github.com/101loops/bdd"
)

var _ = Describe("Factory", func() {

	It("create new error", func() {
		err := New("this is %s", "nice!")
		Check(err.Error(), Equals, "this is nice!")
	})
})
