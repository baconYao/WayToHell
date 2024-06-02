package streamer

import "fmt"

// VideoDispatcher holds info for a dispatcher
type VideoDispatcher struct {
	WorkerPool chan chan VideoProcessingJob
	maxWorkers int
	jobQueue   chan VideoProcessingJob
	Processor  Processor
}

// videoWorker holds info for a pool worker. It has the numeric id of the worker,
// the job queue, and the worker pool chan. A chan chan is used when the thing you want to
// send down a channel is another channel to send things back.
// See http://tleyden.github.io/blog/2013/11/23/understanding-chan-chans-in-go/
type videoWorker struct {
	id         int
	jobQueue   chan VideoProcessingJob
	workerPool chan chan VideoProcessingJob
}

// newVideoWorker takes a numeric id and a channel of chan VideoProcessingJob, and returns a videoWorker object.
func newVideoWorker(id int, workerPool chan chan VideoProcessingJob) videoWorker {
	fmt.Println("newVideoWorker: creating video worker id", id)
	return videoWorker{
		id:         id,
		jobQueue:   make(chan VideoProcessingJob),
		workerPool: workerPool,
	}
}

// start starts an individual worker.
func (w videoWorker) start() {
	fmt.Println("w.start(): starting worker id", w.id)
	go func() {
		for {
			// Add jobQueue to the worker pool.
			w.workerPool <- w.jobQueue

			// Wait for a job to come back.
			job := <-w.jobQueue

			// Process the job.
			w.processVideoJob(job.Video)
		}
	}()
}

// Run runs the workers, and makes the worker pool active. Once we run
// everything, we can push jobs onto the channel to process them.
func (vd *VideoDispatcher) Run() {
	fmt.Println("vd.Run: starting worker pool by running workers")
	for i := 0; i < vd.maxWorkers; i++ {
		worker := newVideoWorker(i+1, vd.WorkerPool)
		worker.start()
	}

	go vd.dispatch()
}

// dispatch dispatches a worker to handle a job.
func (vd *VideoDispatcher) dispatch() {
	for {
		// Wait for a job to come in
		job := <-vd.jobQueue
		fmt.Println("vd.dispatch: sending job", job.Video.ID, "to worker job queue")
		go func() {
			workJobQueue := <-vd.WorkerPool
			workJobQueue <- job
		}()
	}
}

// processVideoJob processes the main queue job with a particular worker.
func (w videoWorker) processVideoJob(video Video) {
	fmt.Println("w.processVideoJob: starting encode on video", video.ID)
	video.encode()
}
