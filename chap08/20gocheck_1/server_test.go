/**
$ go test -check.vv # 詳細表示オプション
*/

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/packer/common/json"

	. "gopkg.in/check.v1"
)

type PostTestSuite struct{}

func init() {
	Suite(&PostTestSuite{})
}

func Test(t *testing.T) { TestingT(t) }

func (s *PostTestSuite) TestHandleGet(c *C) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)

	c.Check(writer.Code, Equals, 200)
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	c.Check(post.ID, Equals, 1)
}
