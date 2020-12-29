package consumer

import (
	"github.com/galehuang/message-service-project/config"
	"github.com/galehuang/message-service-project/consumer/worker"
	"sync"
	"time"
)

type EmailSendConsumer struct {
	workerPool *worker.EmailSendWorkerPool

	emailReceiveCh chan string

	routineWait *sync.WaitGroup
}

func NewEmailSendConsumer() *EmailSendConsumer  {
	return &EmailSendConsumer{
		workerPool:     &worker.EmailSendWorkerPool{},
		emailReceiveCh: make(chan string),
		routineWait:    &sync.WaitGroup{},
	}
}

func (c *EmailSendConsumer) Start()  {
	consumerConfig := config.GetConfig().Consumer.Email
	// 开启工作线程池
	for i := 0; i< consumerConfig.WorkerNum; i++ {
		intervalMs := 1000 / consumerConfig.WorkerNum

		c.routineWait.Add(1)
		go c.workerPool.DispatchWorker(i, c.emailReceiveCh, c.routineWait)
		time.Sleep(time.Millisecond * time.Duration(intervalMs))
	}
	// 开启kafka监听线程

}

func (c *EmailSendConsumer) Consume()  {
	defer c.routineWait.Done()
	defer close(c.emailReceiveCh)

}
