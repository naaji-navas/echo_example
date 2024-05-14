package handlers

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
) 


type EndToEndTestSuite struct {
	suite.Suite
}


func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndToEndTestSuite))
}

func (s *EndToEndTestSuite) TestEndToEnd() {
	// TestEndToEnd function
	c := http.Client{}
	r,_ := c.Get("http://localhost:1323/posts")
	s.Equal(http.StatusOK, r.StatusCode)
} 

func (s *EndToEndTestSuite) TestPostHandler() {
	// TestEndToEndSingle function
	c := http.Client{}
	r,_ := c.Get("http://localhost:1323/posts/5334")
	s.Equal(http.StatusOK, r.StatusCode)
	b, _ := ioutil.ReadAll(r.Body)
	s.JSONEq(`{"status":"ok","data":[]}`, string(b) )
	
}