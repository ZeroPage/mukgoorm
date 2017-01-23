package session

import (
	"github.com/astaxie/beego/session"
)

var GlobalSessions *session.Manager

func ClearSessions() {
	config := timeConfig(-1, -1, -1)
	GlobalSessions, _ = session.NewManager("memory", config)

	GlobalSessions.GC()

	config = timeConfig(3600, 3600, 3600)
	GlobalSessions, _ = session.NewManager("memory", config)
}

func timeConfig(gctime, maxtime int64, cookietime int) *session.ManagerConfig {
	return &session.ManagerConfig{
		Gclifetime:     gctime,
		Maxlifetime:    maxtime,
		CookieLifeTime: cookietime,
	}
}

func init() {
	config := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "",
	}

	GlobalSessions, _ = session.NewManager("memory", config)
	go GlobalSessions.GC()
}
