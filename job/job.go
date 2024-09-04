package job


import (
	c "github.com/ottoMuller1/base/concurrency"
	"fmt"
	"sync"
	"time"
	l "github.com/ottoMuller1/base/logger"
)

// job
type Job struct {
	Name 	   string
	PeriodSecs float64
	Handle     func()
	schedule   *sync.WaitGroup
	Mutex      *sync.Mutex			// jobs can block each other if required
}








// job handler
// TODO: create a special error here
func execJob(j Job) {

	if j.schedule == nil {
		l.DefaultLogger{
			Name: "job " + j.Name,
			Message: "has no schedule",
		}.Error()
	}

	j.schedule.Add(1)

	waitingTime := time.Duration(j.PeriodSecs) * time.Second

	c.GoSync(j.Mutex, func() {
		for {

			func() {
				
				defer func() {

					if r := recover(); r != nil {
						l.DefaultLogger{
							Name: "job " + j.Name,
							Message: fmt.Sprint(r),
						}.Error()
					}
			
				}()
	
				time.Sleep(waitingTime)

				l.DefaultLogger{
					Name: "start job",
					Message: j.Name,
				}.Info()
			
				j.Handle()
			
				l.DefaultLogger{
					Name: "end job",
					Message: j.Name,
				}.Info()

			}()
	
		}
	})

}





// job list handler
func ExecSchedule(jobs []Job) {

	schedule := &sync.WaitGroup{}

	for _, job := range(jobs) {

		job.schedule = schedule

		execJob(job)

	}

	schedule.Wait()

}

