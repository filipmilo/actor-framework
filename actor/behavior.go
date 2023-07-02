package actor

type behavior struct {
  run func(context Context)
}

func initBehavior(handler func(context Context)) *behavior {
  return &behavior{
    run: handler,
  }
}
