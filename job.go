package pool

// Job contains the data that a WorkerAction will consume
type Job struct {
	DataPoint  interface{}
	StatusChan chan interface{}
}

// CreateJob creates a new job that can be dispatched
// to the thread pool
func CreateJob(data interface{}) Job {
	statusChan := make(chan interface{})
	return Job{
		StatusChan: statusChan,
		DataPoint:  data,
	}
}
