package pool

import (
	"log"
)

type worker struct {
	ID         int
	StopChan   chan bool
	JobChannel *(chan Job)
	Action     WorkerAction
}

// WorkerAction is a function that a worker can be registered to perform
type WorkerAction func(Job) error

func (w worker) spawn() {
	go func() {
		for {
			select {
			case job := <-*(w.JobChannel):
				log.Println("Job recieved by worker", w.ID)
				err := w.Action(job)
				if err != nil {
					job.StatusChan <- err
					w.respawn(err)
					return
				}
			case <-w.StopChan:
				return
			}
		}
	}()
}

func (w worker) respawn(err error) {
	log.Println("Worker", w.ID, "threw an error. Restarting.")
	log.Println(err)
	w.spawn()
	log.Println("Restarted worker", w.ID)
}
