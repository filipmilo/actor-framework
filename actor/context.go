package actor

type Context struct {
	Pid      Pid
	Name     string
	behavior *behavior
	Message  IMessage
}

func (context *Context) Become(newBehavior func(context Context)) {
	context.behavior.run = newBehavior
}

func (context *Context) Send(reciverPid *Pid, message IMessage) {
	reciverPid.channel <- message
}

func (context *Context) GetPid() Pid {
	return context.Pid
}
