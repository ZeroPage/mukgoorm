package session

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func PerformRequest(r http.Handler, method, path, cookie string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	req.Header.Add("Cookie", cookie)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestClearSessions(t *testing.T) {
	sessionId := ""
	cookie := ""

	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		sess, _ := GlobalSessions.SessionStart(c.Writer, c.Request)
		defer sess.SessionRelease(c.Writer)

		if sessionId == "" {
			sessionId = sess.SessionID()
		} else {
			assert.NotEqual(t, sessionId, sess.SessionID())
		}
		sess.Set("key", "value")
	})

	res := PerformRequest(r, "GET", "/test", cookie)
	cookie = res.Header().Get("Set-Cookie")

	r.DELETE("/flush", func(c *gin.Context) {
		assert.Equal(t, 1, GlobalSessions.GetActiveSession())

		ClearSessions()
		assert.Equal(t, 0, GlobalSessions.GetActiveSession())

		sess, _ := GlobalSessions.SessionStart(c.Writer, c.Request)
		defer sess.SessionRelease(c.Writer)
		assert.Equal(t, 1, GlobalSessions.GetActiveSession())
		assert.NotEqual(t, sess.Get("key"), "value")
	})

	res = PerformRequest(r, "DELETE", "/flush", cookie)
	cookie = res.Header().Get("Set-Cookie")

	res = PerformRequest(r, "GET", "/test", cookie)
}
