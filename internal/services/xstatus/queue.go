package xstatus

type QueueStatus struct {
	Delayed    int
	Processing int
	Waiting    int
}

type QueueStatusService struct{}
