package worker

import (
	"github.com/galehuang/message-service-project/config"
	"github.com/galehuang/message-service-project/model"
	"github.com/galehuang/message-service-project/services/log"
	"math/rand"
	"sync"
	"time"
)

type EmailSendWorkerPool struct {
}

func (w *EmailSendWorkerPool) DispatchWorker(id int, receiveCh <-chan string, routineWait *sync.WaitGroup)  {
	defer routineWait.Done()

	workerConfig := config.GetConfig().Consumer.Email
	EmailSendModel := model.NewEmailSendModel(workerConfig.BucketSize)
	tickerInterval := workerConfig.IntervalMs + int(rand.Int31n(int32(workerConfig.OffsetRangeMs))) + id
	ticker := time.NewTicker(time.Millisecond * time.Duration(tickerInterval))

	defer func() {
		ticker.Stop()
		_ = EmailSendModel.FlushData()
	}()

	for{
		select{
			case <- ticker.C:
				log.GLogger.Info("email send worker[%d] tick", id)

				err := EmailSendModel.FlushData()
				if err != nil{
					log.GLogger.Error("email send err=[%v]", err)
				}
			case data, ok := <-receiveCh:
				if !ok{
					log.GLogger.Error("email send worker[%d] receive chan is closed", id)
					return
				}
				log.GLogger.Info("email send worker[%d] receive data [%v]", id, data)

				err := EmailSendModel.ProcessData(data)
				if err != nil{
					log.GLogger.Error("email send worker err=[%v] data=[%v]", err, data)
				}
			}
	}

}