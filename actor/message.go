package actor

import "github.com/google/uuid"

// Messages should be IMMUTABLE!
// Setters should not be defined, create messages only through constructor

// Ovaj interface je mozda i nepotreban? Pa sam za sve stavio strukturu Envelope
// type IEnvelope interface {
// 	GetReciver() *uuid.UUID
// 	GetSender() *uuid.UUID
// 	GetMessage() IMessage
// }

// CAN BE OF ANY TYPE! Stavio sam ovo da bude jasno da je neki Value
type IMessage interface{}

type Envelope struct {
	reciver *uuid.UUID
	sender  *uuid.UUID
	message IMessage
}

func (e *Envelope) GetMessage() IMessage {
	return e.message
}

func (e *Envelope) GetReciver() *uuid.UUID {
	return e.reciver
}

func (e *Envelope) GetSender() *uuid.UUID {
	return e.sender
}
