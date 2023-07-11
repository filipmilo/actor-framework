package actor

type behavior struct {
	run func(context *ActorContext)
}

func initBehavior(handler func(context *ActorContext)) *behavior {
	return &behavior{
		run: handler,
	}
}
