package main

import (
	"github.com/trylotus/connector/config"
	"github.com/trylotus/connector/kafkautils"
	"github.com/trylotus/connectors/candymachine"

	_ "go.uber.org/automaxprocs"
)

var conf = config.GetConfig()

func init() {
	kafkautils.TopicTypeRegistry.Load(candymachine.TopicTypes)
	conf.SetDefault("candymachine.kafka.txID", "candymachine")
	conf.SetDefault("candymachine.id", "candymachine")
	conf.SetDefault("candymachine.programid", "cndy3Z4yapfJBmL3ShUp5exZKqR3z33thTzeNMm2gRZ")
	conf.SetDefault("candymachine.backfillLimit", 250) // Maximum value is 1000
}

func main() {
	connector := candymachine.NewConnector(conf)

	connector.Start()
}
