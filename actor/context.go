package actor

type Context struct {
  Name string
  behavior *behavior
}

func(context *Context) Become(newBehavior func(context Context)) {
  context.behavior.run = newBehavior
}

