package errs

type errFunc func() error

type pipeline struct {
	errFuncs []errFunc
}

func newPipeline() *pipeline {
	return &pipeline{
		errFuncs: []errFunc{},
	}
}

func (p *pipeline) add(errFunc errFunc) {
	p.errFuncs = append(p.errFuncs, errFunc)
}

func (p *pipeline) run() error {
	var err error

	for _, errFunc := range p.errFuncs {
		err = errFunc()
		if err != nil {
			return err
		}
	}

	return nil
}
