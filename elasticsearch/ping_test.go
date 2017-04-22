package elasticsearch

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"elasticsearch-spike/config"
)

func TestNewElasticSearchClient(t *testing.T) {
	esClient := NewElasticSearchClient(config.ElasticSearchBaseURL)
	assert.NotNil(t, esClient)
}

func TestNewElasticSearchClientPanicOnUnreachableURl(t *testing.T) {
	assert.Panics(t, func() {
		NewElasticSearchClient(config.InvalidElasticSearchBaseURL)
	}, "NewElasticSeachClient should panic on Invalid URL.")
}
