package entities

import "github.com/Zomato/go-solr-client/solr"

type Document struct {
	*solr.Document
}

func NewDocument() *Document {
	doc := solr.Document{}
	return &Document{&doc}
}

func (d *Document) SetMap(obj map[string]interface{}) {
	for k, v := range obj {
		d.Set(k, v)
	}
}
