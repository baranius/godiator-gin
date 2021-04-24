package pipelines

import (
	"github.com/baranius/godiator"
	"gopkg.in/go-playground/validator.v9"
)

type ValidationPipeline struct {
	godiator.Pipeline
}

var v = validator.New()

func (p *ValidationPipeline) Handle(request interface{}, params ...interface{}) (interface{}, error){
	if err := v.Struct(request); err != nil {
		return nil, err
	}

	return p.Next().Handle(request, params...)
}
