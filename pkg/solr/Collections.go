package solr

import (
	"errors"
	"fmt"
	"github.com/anirudhgupta22/solr-client/pkg/solr/connectionFactory"
)

var collections *Collections

type Collections struct {
	collections map[string]*connectionFactory.CollectionConfig
}

func NewEmptySolrCollections() *Collections {
	return &Collections{collections:map[string]*connectionFactory.CollectionConfig{}}
}

func (s *Collections) AddCollection(name string, config *connectionFactory.CollectionConfig) {
	s.collections[name] = config
}

func (s *Collections) GetClient(name string) (*connectionFactory.CollectionConfig, error) {
	if conn, exists := s.collections[name]; exists {
		return conn, nil
	}
	return nil, errors.New("collection not found")
}




func SetCollections(c *Collections)  {
	collections = c
}

func NewSolrClientForCollection(name string) *SolrClient {
	fmt.Println(collections)
	x, err := collections.GetClient(name)
	if err != nil {
		panic("error in solr")
	}
	return NewSolrClient(x)
}
