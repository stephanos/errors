package errors

import (
	. "github.com/101loops/bdd"
)

var _ = Describe("Utility", func() {

	It("return image url", func() {
		err := New("this is %s", "nice!")
		Check(err, IsNil)
		Check(err.Error(), Equals, "this is nice!")
	})
})
