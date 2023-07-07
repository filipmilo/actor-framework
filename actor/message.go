package actor

import (
	"github.com/google/uuid"
)

type IMessage interface{}

type Envelope struct {
	reciver uuid.UUID
	sender  *uuid.UUID
	message IMessage
}

func (e *Envelope) Reciever() uuid.UUID {
  return e.reciver;
}

func (e *Envelope) Sender() *uuid.UUID {
  return e.sender;
}

func (e *Envelope) Message() IMessage{
  return e.message;
}

type CreatedMessage struct {
  pid uuid.UUID
  channel *chan Envelope
}
