package config

type Config interface {
	GetDBHost() string
	GetDBPort() int64
	GetDBName() string
	GetDBUser() string
	GetDBPass() string
}

type config struct {
	dbHost string
	dbPort int64
	dbName string
	dbUser string
	dbPass string
}

func (c *config) GetDBHost() string {
	return c.dbHost
}

func (c *config) GetDBPort() int64 {
	return c.dbPort
}

func (c *config) GetDBName() string {
	return c.dbName
}
func (c *config) GetDBUser() string {
	return c.dbUser
}
func (c *config) GetDBPass() string {
	return c.dbPass
}

// NewConfig config init
func NewConfig() Config {
	return &config{
		dbHost: "localhost",
		dbPort: 5432,
		dbName: "blog",
		dbUser: "postgres",
		dbPass: "",
	}
}
