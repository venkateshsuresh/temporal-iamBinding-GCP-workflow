package worker

import (
	"context"
	"go.temporal.io/sdk/client"
	"os"
)

var (
	IamSvc IamServiceI = &iamServiceStruct{}
)

type IamServiceI interface {
	IamBindingService(details IamDetails) error
}

type iamServiceStruct struct {
}

type iamServiceModel struct {
	client     client.Client
	workflowID string
}

func (*iamServiceStruct) IamBindingService(details IamDetails) error {
	cr := new(iamServiceModel)
	opts := client.Options{
		HostPort: os.Getenv("TEMPORAL_HOSTPORT"),
	}
	c, err := client.NewClient(opts)
	if err != nil {
		panic(err)
	}

	cr.client = c

	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: IAMTASKQUEUE,
	}

	_, err = cr.client.ExecuteWorkflow(context.Background(), workflowOptions, IamBindingGoogle, details)
	if err != nil {
		return err
	}

	return nil
}
