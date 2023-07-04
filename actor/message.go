package actor

// Messages should be IMMUTABLE!
// Setters should not be defined, create messages only through constructor
type IMessage interface {
	GetReciver() *Pid
	GetValue() IValue
}

// CAN BE OF ANY TYPE! Stavio sam ovo da bude jasno da je neki Value
type IValue interface{}

// type MessageType int8

// const (
// 	Notification MessageType = 1
// 	Request      MessageType = 2
// )

// type Message struct {
// 	Reciver *Pid
// 	Value   IValue
// }

// func (m *Message) GetValue() IValue {
// 	return m.Value
// }

// func (m *Message) GetReciver() *Pid {
// 	return m.Reciver
// }
