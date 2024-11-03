package streamer

type VideoDispatcher struct {
	WorkerPool chan chan VideoProcessingJob
	maxWorkers int
	jobQueue   chan VideoProcessingJob
	Processor  Processor
}

// type video worker
type VideoWorker struct {
	ID         int
	jobQueue   chan VideoProcessingJob
	workerPool chan chan VideoProcessingJob
}

// new video worker
func NewVideoWorker(id int, workerPool chan chan VideoProcessingJob) VideoWorker {
	return VideoWorker{
		ID:         id,
		workerPool: workerPool,
		jobQueue:   make(chan VideoProcessingJob),
	}
}

// start()
func (w VideoWorker) start() {
	go func() {
		for {
			w.workerPool <- w.jobQueue
			job := <-w.jobQueue
			w.processVideoJob(job.Video)
		}
	}()
}

// Run()
func (vd *VideoDispatcher) Run() {
	for i := 0; i < vd.maxWorkers; i++ {
		worker := NewVideoWorker(i+1, vd.WorkerPool)
		worker.start()
	}

	go vd.dispatch()
}

// dispatch()
func (vd *VideoDispatcher) dispatch() {
	for {
		job := <-vd.jobQueue

		go func() {
			workerJobQueue := <-vd.WorkerPool
			workerJobQueue <- job
		}()
	}
}

// process video job
func (w VideoWorker) processVideoJob(video Video) {
	video.encode()
}
