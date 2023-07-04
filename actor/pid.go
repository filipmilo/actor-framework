package actor

import "github.com/google/uuid"

// Da li PID treba da bude svestan svog kanala
type Pid struct {
	Value uuid.UUID
	// actor   *actor
	channel chan IMessage
}

// PRETPOSTAVLJAM DA DIREKTAN PRISTUP PID-A AKTORU NIJE ONO STO SE OD NAS TRAZI!
// Ovo se kosi sa LifeCycle-om aktora i moguce da remeti njegovo pravilno funkcionisanje i rad
func (p *Pid) MessageMethod(reciver *Pid, message IMessage) {
	// p.actor.sendMessage()
}

// Neki od ova dva nacina je ispravan
// OVO SUMNJAM DA JESTE JER ZAOBILAZI PONASANJE AKTORA NEGO DIREKT PID SALJE NA
// KANAL DRUGOG AKTORA PORUKU ZANEMARUJUCI SVOJE PONASANJE!
func (p *Pid) MessageDirect(reciver *Pid, message IMessage) {
	reciver.channel <- message
}

// Najverovatnije je ispravno!
func (p *Pid) MessageMyChannel(message IMessage) {
	p.channel <- message
}
