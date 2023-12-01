package invoke

type Interceptor struct {
	Intercept func(inv *Invoker)
}
