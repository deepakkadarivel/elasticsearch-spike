package elasticsearch

import (
	"context"
	"elasticsearch-spike/config"
	"encoding/json"
	"errors"
	es "gopkg.in/olivere/elastic.v5"
)

type Log map[string]interface{}

func checkIfIndexAlreadyExists(client *es.Client) (bool, error) {
	exist, err := client.IndexExists(config.IndexName).Do(context.TODO())
	if err != nil {
		return exist, err
	}
	return exist, nil
}

func createIndex(client *es.Client) error {
	res, err := client.CreateIndex(config.IndexName).Do(context.TODO())
	if err != nil {
		return err
	}

	if !res.Acknowledged {
		return errors.New("Create Index " + config.IndexName + " was not acknowledged, check the timeout value.")
	}
	return nil
}

func setClient(client *es.Client) error {
	exists, err := checkIfIndexAlreadyExists(client)
	if err != nil {
		return err
	}

	if !exists {
		err = createIndex(client)
		if err != nil {
			return err
		}
	}
	return nil
}

func AddLogToIndex(client *es.Client, log Log) error {
	err := setClient(client)

	if err != nil {
		return err
	}

	_, err = client.Index().Index(config.IndexName).Type(config.IndexType).BodyJson(log).Do(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func AddBulkLogsToIndex(client *es.Client, logs []Log) error {
	err := setClient(client)

	if err != nil {
		return err
	}

	processor, err := client.BulkProcessor().Do(context.TODO())
	if err != nil {
		return err
	}

	defer processor.Close()

	for _, log := range logs {
		request := es.NewBulkIndexRequest().Index(config.IndexName).Type(config.IndexType).Doc(log)
		processor.Add(request)
	}

	err = processor.Flush()
	if err != nil {
		return err
	}
	return nil
}

func SearchAllLogsFromIndex(client *es.Client) ([]Log, error) {
	searchResult, err := client.Search().Index(config.IndexName).Do(context.TODO())
	if err != nil {
		return nil, err
	}
	var logs = make([]Log, searchResult.Hits.TotalHits)
	if searchResult.Hits.TotalHits > 0 {
		for i, hit := range searchResult.Hits.Hits {
			var log Log
			err := json.Unmarshal(*hit.Source, &log)
			if err != nil {
				return nil, err
			}
			logs[i] = log
		}
	} else {
		return nil, errors.New("Serch results returned 0 hits.")
	}
	return logs, nil
}

func DeleteIndex(client *es.Client) error {
	deletedIndex, err := client.DeleteIndex(config.IndexName).Do(context.TODO())
	if err != nil {
		return err
	}
	if !deletedIndex.Acknowledged {
		return errors.New("Deleted Index " + config.IndexName + " was not acknowledged. Check timeout value.")
	}
	return nil
}
