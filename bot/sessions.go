package main

import "errors"

// TODO add redis usage (after/before commit create also issue and delete comment)

type Session interface {
	getData() any
	setNext(any) (bool, error) // returns true whether it is end
	getState() int
}

type SessionManager struct {
	sessions map[int64]Session
}

func (sm *SessionManager) getSessionState(id int64) int {
	if session, ok := sm.sessions[id]; ok {
		return session.getState()
	}
	return -1
}

func (sm *SessionManager) register(id int64, session Session) {
	sm.sessions[id] = session
}

func (sm *SessionManager) putNextData(id int64, data any) (storedData any, end bool, err error) { // TODO May exsist method put concrete data (f.e. when by using inline keyboards you change fields)
	end, err = sm.sessions[id].setNext(data)
	storedData = sm.sessions[id].getData()
	return
}

func NewSessionManager() (sm *SessionManager) {
	sm = new(SessionManager)
	sm.sessions = make(map[int64]Session)
	return
}

type AdvertiseSession struct {
	state     int
	advertise Advertise
}

func (as *AdvertiseSession) getState() int {
	return as.state
}

var ErrorDataType = errors.New("Invalid type passed")

func (as *AdvertiseSession) setNext(data any) (end bool, err error) {
	as.state++
	var ok bool
	end = false
	switch as.state {
	case 0:
		as.advertise.Name, ok = data.(string)
	case 1:
		as.advertise.Link, ok = data.(string)
		end = true
	}

	if !ok {
		err = ErrorDataType
	}

	return
}

func (as *AdvertiseSession) getData() any {
	return as.advertise
}

func NewAdvertiseSession() (as *AdvertiseSession) {
	as = new(AdvertiseSession)
	as.state = 0
	return
}
