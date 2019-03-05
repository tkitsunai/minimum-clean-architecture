package config

type DbConfig struct {
	Source string `yaml:"source"`
}

func AppPort() string {
	return ":8080"
}

func DBConnection() *DbConfig {
	return &DbConfig{
		Source: "root:test@/testdb?charset=utf8&parseTime=True&loc=Local",
	}
}
