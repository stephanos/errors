package errors

import (
	. "launchpad.net/gocheck"
)

// TESTS ==========================================================================================

func (s *S) TestNew(c *C) {

	err := New("this is %s", "nice!")
	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "this is nice!")
}
