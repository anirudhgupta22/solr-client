package solr

import (
	"context"
	"github.com/Zomato/go-solr-client/solr"
	"github.com/anirudhgupta22/solr-client/pkg/solr/connectionFactory"
	"github.com/anirudhgupta22/solr-client/pkg/solr/entities"
	"net/url"
)

type SolrClient struct {
	conn *connectionFactory.CollectionConfig
}

func NewSolrClient(conn *connectionFactory.CollectionConfig) *SolrClient {
	return &SolrClient{conn: conn}
}

func (s *SolrClient) Add(docs []entities.Document, chunkSize int, params *url.Values) (*solr.SolrUpdateResponse, error) {
	client := s.conn.Master()
	return client.Add(docs, chunkSize, params)
}

func (s *SolrClient) Search(ctx context.Context, query entities.Query) (*solr.SolrResult, error) {
	client := s.conn.Slave()
	return client.Search(ctx, query)
}

func (s *SolrClient) Delete(ctx context.Context, data map[string]interface{}) (*solr.SolrUpdateResponse, error) {
	client := s.conn.Master()
	return client.Delete(ctx, data)
}
