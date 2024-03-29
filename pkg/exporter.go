package pkg

type Exporter interface {
	Export(Item) error
	PipelineNames() map[uint8]string
	Pipelines() []Pipeline
	SetPipeline(Pipeline, uint8)
	DelPipeline(int)
	CleanPipelines()
	WithDumpPipeline()
	WithFilePipeline()
	WithImagePipeline()
	WithFilterPipeline()
	WithNonePipeline()
	WithCsvPipeline()
	WithJsonLinesPipeline()
	WithMongoPipeline()
	WithSqlitePipeline()
	WithMysqlPipeline()
	WithKafkaPipeline()
	WithCustomPipeline(Pipeline)
}
