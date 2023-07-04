package main

import (
	"fmt"
	"main/actor"
	"math/rand"
	"time"
)

type Sender struct{}

func (t *Sender) Recieve(context actor.Context) {
	msg := context.Message

	if msg != nil {
		switch msg.GetReciver() {
		case nil:
			fmt.Println("This is ONLY my Message")
			fmt.Printf("MSG: {%s}\n", msg.GetValue())
		default:
			fmt.Println("FORWARD MESSAGE")
			context.Send(msg.GetReciver(),
				&Message{
					Value: msg.GetValue(),
				})
		}
	} else {
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		fmt.Println("NO MESSAGE AVAILABLE,Sleeping")

	}
}

// func (t *Sender) SecondState(context actor.Context) {
// 	fmt.Printf("Hi im %s and this is my second state\n", context.Name)
// 	context.Become(t.Recieve)
// }

type Adder struct {
	sum int
}

func (a *Adder) Recieve(context actor.Context) {
	msg := context.Message
	if msg != nil {
		switch msg.GetReciver() {
		case nil:
			//DONT SEND A MESSAGE BACK!
			switch msg.GetValue().(type) {
			case string:
				context.Become(a.Add)
			default:
				fmt.Println("NOT CORRECT TYPE!")
			}
		default:
			//RETURNS A MESSAGE
		}
	} else {
		//NO MESSAGE AVAILABLE
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		fmt.Printf("Current sum is %v\n", a.sum)
	}
}

func (a *Adder) Add(context actor.Context) {

	msg := context.Message
	if msg != nil {
		switch msg.GetReciver() {
		case nil:
			//DONT SEND A MESSAGE BACK!
			switch msg.GetValue().(type) {
			case int:
				a.sum = a.sum + msg.GetValue().(int)
			case bool:
				context.Become(a.Recieve)
			default:
				fmt.Printf("NOT CORRECT TYPE! Message is... {%v}\n", msg.GetValue())
			}
		default:
			//RETURNS A MESSAGE
		}
	} else {
		//NO MESSAGE AVAILABLE

	}
}

type Message struct {
	Reciver *actor.Pid
	Value   actor.IValue
}

func (m *Message) GetValue() actor.IValue {
	return m.Value
}

func (m *Message) GetReciver() *actor.Pid {
	return m.Reciver
}

type ComplexValue struct {
	Name    string
	Surname string
	Age     int
}

func main() {
	system := actor.ActorSystem{}
	system.InitSystem()

	pid1 := system.InitActor(&Sender{})
	pid2 := system.InitActor(&Adder{})
	system.PrintValues()

	time.Sleep(3 * time.Second)

	pid1.MessageMyChannel(&Message{
		Value:   "Hello Sender",
		Reciver: nil,
	})

	time.Sleep(3 * time.Second)

	pid1.MessageMyChannel(&Message{
		Value:   "Hello Adder, Please Add!",
		Reciver: pid2,
	})

	time.Sleep(3 * time.Second)

	pid1.MessageMyChannel(&Message{
		Value:   "Ups!",
		Reciver: pid2,
	})

	time.Sleep(3 * time.Second)

	pid1.MessageMyChannel(&Message{
		Value:   3,
		Reciver: pid2,
	})

	time.Sleep(3 * time.Second)

	pid1.MessageMyChannel(&Message{
		Value:   2,
		Reciver: pid2,
	})

	pid1.MessageMyChannel(&Message{
		Value:   true,
		Reciver: pid2,
	})

	// pid1.MessageMyChannel(&Message{
	// 	Value: ComplexValue{
	// 		Name:    "Luka",
	// 		Surname: "Licina",
	// 		Age:     22,
	// 	},
	// 	Reciver: pid2,
	// })

	time.Sleep(60 * time.Second)
}
