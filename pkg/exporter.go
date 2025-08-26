package pkg

type Exporter interface {
	Export(Item) error
	Names() map[uint8]string
	Pipelines() []Pipeline
	SetPipeline(pipeline Pipeline, order uint8)

	// Remove removes the pipeline at the specified index.
	//
	// Parameters:
	//   - index: the position of the pipeline in the pipeline slice to be removed.
	//
	// Behavior:
	//   - If the index is out of range, no operation is performed.
	//   - After deletion, the remaining pipelines shift to fill the gap.
	Remove(index int)

	// Clean removes all pipelines from the Spider.
	//
	// Behavior:
	//   - Clears the slice of pipelines, effectively removing all existing pipelines.
	//   - After calling this method, the spider will have no pipelines.
	Clean()
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
