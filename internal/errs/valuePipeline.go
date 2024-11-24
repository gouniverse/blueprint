package errs

type valueOrErrorFunc[T any] func() (T, error)

type valuePipeline[T any] struct {
	errFuncs []valueOrErrorFunc[T]
}

func newErrPipeline[T any]() *valuePipeline[T] {
	return &valuePipeline[T]{
		errFuncs: []valueOrErrorFunc[T]{},
	}
}

func (p *valuePipeline[T]) add(valueOrErrorFunc valueOrErrorFunc[T]) {
	p.errFuncs = append(p.errFuncs, valueOrErrorFunc)
}

func (p *valuePipeline[T]) run() (T, error) {
	var result T
	var err error

	for _, errFunc := range p.errFuncs {
		result, err = errFunc()
		if err != nil {
			return result, err
		}
	}

	return result, nil
}
