package errs

func Fails(errFuncs ...errFunc) bool {
	err := Validate(errFuncs...)
	return err != nil
}

func Pass(errFuncs ...errFunc) bool {
	err := Validate(errFuncs...)
	return err == nil
}

func Validate(errFuncs ...errFunc) error {
	p := newPipeline()

	for _, errFunc := range errFuncs {
		p.add(errFunc)
	}

	return p.run()
}

func ValidateAndGet[T any](valueOrErrorFuncs ...valueOrErrorFunc[T]) (T, error) {
	p := newErrPipeline[T]()

	for _, valueOrErrorFunc := range valueOrErrorFuncs {
		p.add(valueOrErrorFunc)
	}

	return p.run()
}
