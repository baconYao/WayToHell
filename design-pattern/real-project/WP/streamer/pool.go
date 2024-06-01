package streamer

type VideoDispatcher struct {
	WorkerPool chan chan VideoProcessingJob
	maxWorkers int
	jobQueue   chan VideoProcessingJob
	Processor  Processor
}
