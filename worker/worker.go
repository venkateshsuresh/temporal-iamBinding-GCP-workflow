package worker

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"os"
)

const IAMTASKQUEUE = "IAM_TASK_QUEUE"

var IamWorker worker.Worker = newWorker()

func newWorker() worker.Worker {
	opts := client.Options{
		HostPort: os.Getenv("TEMPORAL_HOSTPORT"),
	}
	c, err := client.NewClient(opts)
	if err != nil {
		panic(err)
	}

	w := worker.New(c, IAMTASKQUEUE, worker.Options{})
	w.RegisterWorkflow(IamBindingGoogle)
	w.RegisterActivity(AddIAMBinding)

	return w
}
