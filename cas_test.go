package cascaron_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	gc "launchpad.net/gocheck"

	"github.com/cmars/cascaron"
)

func TestServerSuite(t *testing.T) { gc.TestingT(t) }

type CasSuite struct {
	*httptest.Server
}

var _ = gc.Suite(&CasSuite{})

func (cs *CasSuite) SetUpTest(c *gc.C) {
	s := cascaron.NewService()
	cs.Server = httptest.NewServer(s)
}

func (cs *CasSuite) TestNotFound(c *gc.C) {
	resp, err := http.Get(cs.URL + "/bad-taste")
	c.Assert(err, gc.IsNil)
	c.Check(resp.StatusCode, gc.Equals, http.StatusNotFound)
}

func (cs *CasSuite) TestLogin(c *gc.C) {
	resp, err := http.Get(cs.URL + "/login")
	c.Assert(err, gc.IsNil)
	c.Check(resp.StatusCode, gc.Equals, http.StatusNotImplemented)
}
