package entities

import "github.com/Zomato/go-solr-client/solr"

type Query struct {
	*solr.Query
}

func NewQuery() *Query {
	return &Query{solr.NewQuery()}
}
