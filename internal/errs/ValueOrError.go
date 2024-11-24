package errs

func ValueOrError[T any](valueOrErrorFuncs ...valueOrErrorFunc[T]) (T, error) {
	p := newErrPipeline[T]()
	for _, f := range valueOrErrorFuncs {
		p.add(f)
	}
	return p.run()
}
