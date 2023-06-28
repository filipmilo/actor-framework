package actor

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Actor interface {
	//Mozda ovo da vrati Actora? Pa da ActorSystem sacuva PID njegov u mapu neku recimo
	//I onda da pozove metodu Live() konkurento -> go , koji potom ceka na neke poruke
	Birth() uuid.UUID
	Live()
	Kill()
}

type ActorStatus int8

const (
	ActorLiving ActorStatus = 1
	ActorEnd    ActorStatus = 2
)

type BasicActor struct {
	Pid    uuid.UUID
	Name   string
	Status ActorStatus
}

func (a *BasicActor) Birth() uuid.UUID {
	a.Pid = uuid.New()
	a.Name = fmt.Sprintf("%s%d", "BasicActor", a.Pid)

	fmt.Printf("I, %s am BORN!", a.Name)
	fmt.Println()

	go a.Live()
	return a.Pid
}

func (a *BasicActor) Live() {
	//Runs function after current function finishes
	defer a.Kill()

	for a.Status = ActorLiving; a.Status == ActorLiving; {
		//Some business logic - message recieves
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)

		//Ako mu neko posalje recimo poruku da se samoubije ili ako on tako odluci recimo
		//Ovo ce zavisiti od toga kakve poruke primi
		if rand.Intn(100) > 90 {
			a.Status = ActorEnd
		}
		fmt.Printf("%s waiting for message", a.Name)
		fmt.Println()
	}
}

func (a *BasicActor) Kill() {
	fmt.Printf("I,%s have died... ARGHHHH!", a.Name)
	fmt.Println()
}

type ActorSystem struct {
	environment map[uuid.UUID]Actor
}

func (as *ActorSystem) InitSystem() {
	as.environment = make(map[uuid.UUID]Actor)
}

func (as *ActorSystem) InitActor() {
	a := BasicActor{}
	pid := a.Birth()

	_, ok := as.environment[pid]
	if ok {
		a.Status = ActorEnd
		return
	}

	as.environment[pid] = &a
}

func (as *ActorSystem) PrintValues() {
	fmt.Println(as.environment)
}
