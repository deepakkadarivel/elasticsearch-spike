package elasticsearch

import (
	"elasticsearch-spike/config"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAddLogToIndex(t *testing.T) {
	esClient := NewElasticSearchClient(config.ElasticSearchBaseURL)
	err := AddLogToIndex(esClient, logsArray()[0])
	assert.Nil(t, err)
}

func TestAddBulkLogsToIndex(t *testing.T) {
	esClient := NewElasticSearchClient(config.ElasticSearchBaseURL)
	err := AddBulkLogsToIndex(esClient, logsArray())
	assert.Nil(t, err)
}

func TestDeleteIndex(t *testing.T) {
	esClient := NewElasticSearchClient(config.ElasticSearchBaseURL)
	err := DeleteIndex(esClient)
	assert.Nil(t, err)
}

func TestSearchAllLogsFromIndex(t *testing.T) {
	esClient := NewElasticSearchClient(config.ElasticSearchBaseURL)
	logs, err := SearchAllLogsFromIndex(esClient)
	assert.Nil(t, err)
	assert.NotNil(t, logs)
}

func logsArray() []Log {
	logs := make([]Log, 10)
	for i := 1; i <= 10; i++ {
		l := Log{
			"app_name": config.AppName,
			"message":  fmt.Sprintf("Mesage with Log id : %v", i),
			"time":     time.Now(),
		}
		logs[i-1] = l
	}
	return logs
}
