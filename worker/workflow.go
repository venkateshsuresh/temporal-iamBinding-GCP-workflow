package worker

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"net/http"
	"time"
)

func IamWorkFlow(c *gin.Context) {
	var details IamDetails
	err := LoadData(c, &details)

	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, err)

		return
	}

	err = IamSvc.IamBindingService(details)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, err)

		return
	}

	c.JSON(http.StatusOK, err)
}

func IamBindingGoogle(ctx workflow.Context, details IamDetails) (string, error) {

	iamCtx := workflow.WithActivityOptions(
		ctx,
		workflow.ActivityOptions{
			StartToCloseTimeout:    1 * time.Hour,
			ScheduleToCloseTimeout: 1 * time.Hour,
			RetryPolicy: &temporal.RetryPolicy{
				MaximumAttempts: 3,
			},
			TaskQueue: IAMTASKQUEUE,
		})

	err := workflow.ExecuteActivity(iamCtx, AddIAMBinding, details).Get(ctx, nil)

	return "", err
}
