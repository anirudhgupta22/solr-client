package connectionFactory

import (
	"context"
	"github.com/Zomato/go-solr-client/solr"
	"github.com/anirudhgupta22/solr-client/pkg/solr/entities"
	"log"

	"net/url"
)

type Connection struct {
	protocol       string
	host           string
	port           string
	solrInterface  *solr.SolrInterface
	collectionName string
	//itype          solr2.InstanceType
}

func NewConnection(protocol string, host string, port string, collectionName string) *Connection {
	c := &Connection{protocol: protocol, host: host, port: port, collectionName: collectionName}
	solr, err := getSolrInterface(c)
	if err != nil {

	}
	c.solrInterface = solr
	return c
}

func getSolrInterface(config *Connection) (solrInterface *solr.SolrInterface, err error) {
	solrUrl := config.protocol + "://" + config.host + ":" + config.port + "/solr"
	solrInterface, err = solr.NewSolrInterface(solrUrl, config.collectionName)
	if err != nil {
		log.Printf("Failed to get solr connection for solrUrl:%s and collection:%s, error:%s", solrUrl, config.collectionName, err)
		return
	}
	return
}

func (s *Connection) Add(docs []entities.Document, chunkSize int, params *url.Values) (*solr.SolrUpdateResponse, error) {
	docList := make([]solr.Document, 0)
	for _, doc := range docs {
		docList = append(docList, *doc.Document)
	}

	return s.solrInterface.Add(docList, 0, nil)
}

func (s *Connection) Search(ctx context.Context, query entities.Query) (*solr.SolrResult, error) {
	search := s.solrInterface.Search(query.Query)
	//seg := StartSolrDbSegment(ctx, search.QueryString())
	//defer seg.End()
	return search.Result(nil)
}

func (s *Connection) Delete(context context.Context, data map[string]interface{}) (*solr.SolrUpdateResponse, error) {
	return s.solrInterface.Delete(data, nil)
}

/*
// Starts a new solr segment in the current txn
func StartSolrDbSegment(ctx context.Context, query string) *newrelic2.DatastoreSegment {
	txn := newrelic2.FromContext(ctx)
	return &newrelic2.DatastoreSegment{
		StartTime:          newrelic2.StartSegmentNow(txn),
		Product:            newrelic2.DatastoreSolr,
		ParameterizedQuery: query,
	}
}*/
