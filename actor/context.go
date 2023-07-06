package actor

import "github.com/google/uuid"

type Context struct {
	Pid      *uuid.UUID
	Name     string
	behavior *behavior
	Envelope Envelope
}

func (context *Context) Become(newBehavior func(context Context)) {
	context.behavior.run = newBehavior
}

// To be implemented
func (context *Context) Send(pid uuid.UUID, message Envelope) {
}

func (context *Context) GetPid() *uuid.UUID {
	return context.Pid
}
