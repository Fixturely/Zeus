package config

func GetTestConfig(cfg *Config) {
	cfg.Environment = "test"
	cfg.Database.Host = "127.0.0.1"
	cfg.Database.Port = 5441
	cfg.Database.User = "admin_test"
	cfg.Database.Password = "admin_test"
	cfg.Database.Name = "zeus-test-db"
	cfg.Database.SSLMode = false
}
