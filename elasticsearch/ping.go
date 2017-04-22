package elasticsearch

import (
	es "gopkg.in/olivere/elastic.v5"
	"context"
	"net/http"
	"errors"
)

func NewElasticSearchClient(baseURL string) *es.Client {
	client, err := getClient(baseURL)
	if err != nil {
		panic("Unable to reach Elastic Search running at - " + baseURL)
	}

	if err = ping(client, baseURL); err != nil {
		panic("Elastic Search Ping failed. - " + err.Error())
	}

	return client
}

func getClient(hostURL string) (*es.Client, error) {
	// Sniffing is disables to avoid Client to reach all elastic search nodes.
	return es.NewClient(es.SetURL(hostURL), es.SetSniff(false))
}

func ping(client *es.Client, hostURL string) error {
	res, code, err := client.Ping(hostURL).Do(context.TODO())

	if err != nil {
		return err
	}

	if code != http.StatusOK {
		return errors.New("Http status not OK.")
	}

	if res == nil || res.Name == "" {
		return errors.New("No valid Ping response.")
	}
	return nil
}
