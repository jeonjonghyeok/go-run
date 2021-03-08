package chat

import (
	"sync"
)

type Room struct {
	participant map[int]Conn
	n           int
	mux         sync.Mutex
}

type Conn interface {
	Send([]byte)
}

var globalRoom *Room

func GetRoom() *Room {
	if globalRoom != nil {
		return globalRoom
	}
	globalRoom = &Room{
		participant: make(map[int]Conn),
	}
	return globalRoom
}

func (r *Room) AddParticipant(c Conn) (id int) {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.participant[r.n] = c
	id = r.n
	r.n++
	return
}

func (r *Room) RemoveParticipant(id int) {
	r.mux.Lock()
	defer r.mux.Unlock()

	delete(r.participant, id)
}

func (r *Room) SendMessage(msg []byte) {
	r.mux.Lock()
	defer r.mux.Unlock()

	for _, v := range r.participant {
		v.Send(msg)
	}
}
