package queue

import (
	"fmt"
	"time"
)

type Job interface {
	Process()
}

type Worker struct {
	WorkerId   int
	Done       chan bool
	JobRunning chan Job
}

func NewWorker(workerID int, jobChan chan Job) *Worker {
	return &Worker{
		WorkerId:   workerID,
		Done:       make(chan bool),
		JobRunning: jobChan,
	}
}

func (w *Worker) Run() {
	fmt.Println("Run worker id", w.WorkerId)
	go func() {
		for {
			select {
			case job := <-w.JobRunning:
				fmt.Println("Job running", w.WorkerId)
				job.Process()
			case <-w.Done:
				fmt.Println("Stop worker", w.WorkerId)
				return
			}
		}
	}()
}

func (w *Worker) StopWorker() {
	w.Done <- true
}

type JobQueue struct {
	Worker     []*Worker
	JobRunning chan Job
	Done       chan bool
}

func NewJobQueue(numOfWorker int) JobQueue {
	workers := make([]*Worker, numOfWorker, numOfWorker)
	jobRunning := make(chan Job)
	for i := 0; i < numOfWorker; i++ {
		workers[i] = NewWorker(i, jobRunning)
	}
	return JobQueue{
		Worker:     workers,
		JobRunning: jobRunning,
		Done:       make(chan bool),
	}
}

func (jq *JobQueue) Push(job Job) {
	jq.JobRunning <- job
}

func (jq *JobQueue) Stop() {
	jq.Done <- true
}

func (jq *JobQueue) Start() {
	go func() {
		for i := 0; i < len(jq.Worker); i++ {
			jq.Worker[i].Run() // lấy từng worker ra
		}
	}()

	go func() {
		for {
			select {
			case <-jq.Done:
				for i := 0; i < len(jq.Worker); i++ {
					jq.Worker[i].StopWorker()
				}
				return
			}
		}
	}()
}

type Sender struct {
	Email string
}

func (s Sender) Process() {
	fmt.Println(s.Email)
}

func main() {
	emails := []string{
		"a@gmail.com",
		"b@gmail.com",
		"c@gmail.com",
		"d@gmail.com",
		"e@gmail.com",
	}

	JobQueue := NewJobQueue(100)
	JobQueue.Start()

	for _, email := range emails {
		sender := Sender{Email: email}
		JobQueue.Push(sender)
	}

	time.AfterFunc(time.Second*2, func() {
		JobQueue.Stop()
	})

	time.Sleep(time.Second * 6)
}
