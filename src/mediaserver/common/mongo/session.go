package mongo

import (
	"sync"

	"gopkg.in/mgo.v2"
)

type SessionManager struct {
	sessions   map[string]*mgo.Session
	accessLock *sync.RWMutex
}

func NewSessionManager(defaultAlias string) *SessionManager {

	sessionManager := &SessionManager{
		sessions:   make(map[string]*mgo.Session),
		accessLock: &sync.RWMutex,
	}
	return sessionManager
}

func (s *SessionManager) Get() (*mgo.Session, bool, error) {

}
