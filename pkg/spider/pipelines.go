package spider

import (
	"errors"
	"github.com/lizongying/go-crawler/internal"
)

func (s *BaseSpider) GetPipelines() (pipelines map[int]string) {
	pipelines = make(map[int]string)
	for k, v := range s.pipelines {
		pipelines[k] = v.GetName()
	}

	return
}

func (s *BaseSpider) ReplacePipelines(pipelines map[int]internal.Pipeline) (err error) {
	pipelinesNameMap := make(map[string]struct{})
	pipelinesOrderMap := make(map[int]struct{})
	for k, v := range pipelines {
		if _, ok := pipelinesNameMap[v.GetName()]; ok {
			err = errors.New("pipeline name duplicate")
			s.Logger.Error(err)
			return
		}
		pipelinesNameMap[v.GetName()] = struct{}{}
		if _, ok := pipelinesOrderMap[k]; ok {
			err = errors.New("pipeline order duplicate")
			s.Logger.Error(err)
			return
		}
		pipelinesOrderMap[k] = struct{}{}
	}

	s.pipelines = pipelines

	return
}

func (s *BaseSpider) SetPipeline(pipeline internal.Pipeline, order int) {
	for k, v := range s.pipelines {
		if v.GetName() == pipeline.GetName() && k != order {
			delete(s.pipelines, k)
			break
		}
	}

	s.pipelines[order] = pipeline

	return
}

func (s *BaseSpider) DelPipeline(name string) {
	for k, v := range s.pipelines {
		if v.GetName() == name {
			delete(s.pipelines, k)
			break
		}
	}

	return
}

func (s *BaseSpider) CleanPipelines() {
	s.pipelines = make(map[int]internal.Pipeline)

	return
}
