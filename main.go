package main

import (
	"fmt"
	"log"
	"main/actor"
	"time"

	"github.com/google/uuid"
)

type Sender struct{
  adderPid uuid.UUID
}

type AdderMessage struct {
  add bool
  amount int32
}

type SenderMessage struct {
  amount int32
}

func (t *Sender) Recieve(context *actor.ActorContext) {
  switch msg := context.Message.Message().(type) {
  case SenderMessage:
    context.Send(t.adderPid, AdderMessage{
      add: true,
      amount: msg.amount,
    })

    context.Become(t.Subtract)
  default:
    log.Printf("Ivalid message type")
  }
}


func (t *Sender) Subtract(context *actor.ActorContext) {
  switch msg := context.Message.Message().(type) {
  case int32:
    context.Send(t.adderPid, AdderMessage{
      add: false,
      amount: msg,
    })

    context.Become(t.Recieve)
  default:
    log.Printf("Ivalid message type")
  }
}

type Adder struct {
	sum int32
}

func (a *Adder) Recieve(context *actor.ActorContext) {
  switch msg := context.Message.Message().(type) {
  case AdderMessage:
    if msg.add {
      a.sum += msg.amount
    } else {
      a.sum -= msg.amount
    }
    fmt.Printf("Current sum is: %d", a.sum)
  default:
    fmt.Printf("Ivalid message type")
  }
}

type ComplexValue struct {
	Name    string
	Surname string
	Age     int
}

func main() {
	system := actor.NewSystem()
  context := system.Root

  adder := context.InitActor(&Adder{}, "Adder")
  sender := context.InitActor(&Sender{
    adderPid: *adder,
  }, "Sender")

	time.Sleep(3 * time.Second)

  context.Send(*sender, SenderMessage{amount: 6})
  context.Send(*sender, SenderMessage{amount: 1})
  context.Send(*sender, SenderMessage{amount: 8})
  context.Send(*sender, SenderMessage{amount: 1})
  context.Send(*sender, SenderMessage{amount: 4})
  context.Send(*sender, SenderMessage{amount: 10})
  context.Send(*sender, SenderMessage{amount: 89})
  context.Send(*sender, SenderMessage{amount: 6})

	system.PrintValues()
	time.Sleep(60 * time.Second)
}
