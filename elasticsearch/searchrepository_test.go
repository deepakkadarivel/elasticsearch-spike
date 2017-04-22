package elasticsearch

import (
	"testing"
	"elasticsearch-spike/config"
	"fmt"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestAddLogToIndex(t *testing.T) {
	esClient := NewElasticSearchClient(config.ElasticSearchBaseURL)
	err := AddLogToIndex(esClient, logsArray()[0])
	assert.Nil(t, err)
}

func logsArray() []Log {
	logs := make([]Log, 10)
	for i := 1; i <= 10; i++ {
		l := Log{
			"app_name": config.AppName,
			"message": fmt.Sprintf("Mesage with Log id : %v", i),
			"time": time.Now(),
		}
		logs[i-1] = l
	}
	fmt.Println("Logs", logs)
	return logs
}
