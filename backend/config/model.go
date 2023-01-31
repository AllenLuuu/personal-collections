package config

type MongoUriType string

type ServerType struct {
	Port   int
	Host   string
	Domain string
}

type ConfigType struct {
	MongoUri MongoUriType
	Server   ServerType
}
