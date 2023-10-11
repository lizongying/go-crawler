package pkg

type Exporter interface {
	Export(ItemWithContext) error
	PipelineNames() map[uint8]string
	Pipelines() []Pipeline
	SetPipeline(Pipeline, uint8)
	DelPipeline(int)
	CleanPipelines()
	WithDumpPipeline()
	WithFilePipeline()
	WithImagePipeline()
	WithFilterPipeline()
	WithCsvPipeline()
	WithJsonLinesPipeline()
	WithMongoPipeline()
	WithSqlitePipeline()
	WithMysqlPipeline()
	WithKafkaPipeline()
	WithCustomPipeline(Pipeline)
}
