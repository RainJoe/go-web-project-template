package config

//TOMLConfig holds config using toml
type TOMLConfig struct {
	Title       string
	Port        string
	RedisConfig RedisConfig `toml:"redis"`
	DBConfig    DBConfig    `toml:"db"`
}

//RedisConfig represents its name says
type RedisConfig struct {
	Addr     string
	PassWord string
	DB       int
}

//DBConfig represents its name says
type DBConfig struct {
	Host     string
	Port     int
	DBName   string
	UserName string
	PassWord string
}
