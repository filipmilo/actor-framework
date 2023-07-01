package actor

type Context struct {
  behavior *behavior
}

func(context *Context) Become(newBehavior func(context Context)) {
  context.behavior.run = newBehavior
}

