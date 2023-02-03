package config

type MongoType struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type ServerType struct {
	Port   int
	Host   string
	Domain string
}

type ConfigType struct {
	Mongo  MongoType
	Server ServerType
}
