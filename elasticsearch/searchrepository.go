package elasticsearch

import (
	"context"
	"elasticsearch-spike/config"
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

func AddLogToIndex(client *es.Client, log Log) error {
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

	_, err = client.Index().Index(config.IndexName).Type(config.IndexType).BodyJson(log).Do(context.TODO())
	if err != nil {
		return err
	}

	return nil
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
