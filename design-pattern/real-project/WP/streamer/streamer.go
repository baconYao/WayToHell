package streamer

type ProcessingMessage struct {
	ID         int
	Successful bool
	Message    string
	OutputFile string
}

type VideoProcessingJob struct {
}

type Processor struct {
}

func New(jobQueue chan VideoProcessingJob, maxWorkers int) *VideoDispatcher {
	workerPool := make(chan chan VideoProcessingJob, maxWorkers)

	// TODO: implement processor logic
	p := Processor{}
	return &VideoDispatcher{
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		WorkerPool: workerPool,
		Processor:  p,
	}
}
