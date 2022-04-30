package worker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

func LoadData(c *gin.Context, model interface{}) error {
	var body bytes.Buffer

	if _, err := io.Copy(&body, c.Request.Body); err != nil {
		customErr := fmt.Errorf("response parsing failed %w", err)

		return customErr
	}

	_ = json.Unmarshal(body.Bytes(), &model)

	return nil
}
