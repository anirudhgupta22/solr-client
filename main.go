package main

import (
	"fmt"
	"github.com/anirudhgupta22/solr-client/pkg/solr"
	connectionfactory "github.com/anirudhgupta22/solr-client/pkg/solr/connectionFactory"

	"github.com/spf13/viper"
	"strings"
)

func initializeSolr() {
	collection := connectionfactory.NewEmptyCollectionConfig()
	conn := getSolrConfig("solr_collection_user_activity")
	collection.SetMaster(conn)
	collection.AddSlaves(conn)
	coll := solr.NewEmptySolrCollections()
	solr.SetCollections(coll)
	fmt.Printf("%+v", coll)

}

func getSolrConfig(identifier string) *connectionfactory.Connection {
	protocol := viper.GetString(identifier + ".protocol")
	host := viper.GetString(identifier + ".host")
	port := viper.GetString(identifier + ".port")
	collection := viper.GetString(identifier + ".collection")
	solrConfig := connectionfactory.NewConnection(protocol, host, port, collection)
	return solrConfig
}

func initializeConfig() {
	// name of config file (without extension)
	viper.SetConfigName("config")

	// path to look for the config file in
	viper.AddConfigPath("./configs/")

	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// Define Prefix
	viper.SetEnvPrefix("activity.service")

	// Define Replacer
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// Check and load env variables
	viper.AutomaticEnv()
}
