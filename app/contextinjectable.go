package app

type ContextInjectable interface {
	Inject(*Context)
}
