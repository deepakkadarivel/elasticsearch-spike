package config

const (
	ElasticSearchBaseURL        = "http://localhost:9200/"
	InvalidElasticSearchBaseURL = "http://remotehost:9200/"
	IndexName                   = "esspike"
	IndexType                   = "log"
	AppName                     = "ESSpike"
	IndexMapping                = Mappings
	MappingsJson                = "./config/mappings.json"
)

const Mappings = `{
  "mapping": {
    "log": {
      "properties": {
        "app" : {"type": "string"},
        "message": {"type": "string"},
        "time": {"time": "date"}
      }
    }
  }
}`