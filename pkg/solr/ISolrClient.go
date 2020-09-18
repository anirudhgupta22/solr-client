package solr

import (
	"context"
	"github.com/Zomato/go-solr-client/solr"
	"github.com/Zomato/user-activity-service/pkg/solr/entities"
	"net/url"
)

type InstanceType string

const (
	MASTER InstanceType = "MASTER"
	SLAVE  InstanceType = "SLAVE"
)

type ISolrClient interface {
	Add(docs []entities.Document, chunkSize int, params *url.Values) (*solr.SolrUpdateResponse, error)
	Search(ctx context.Context, query entities.Query) (*solr.SolrResult, error)
	Delete(ctx context.Context, data map[string]interface{}) (*solr.SolrUpdateResponse, error)
}
