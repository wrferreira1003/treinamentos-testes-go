package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`   // Driver do banco de dados
	DBHost     string `mapstructure:"DB_HOST"`     // Host do banco de dados
	DBPort     string `mapstructure:"DB_PORT"`     // Porta do banco de dados
	DBUser     string `mapstructure:"DB_USER"`     // Usuário do banco de dados
	DBPassword string `mapstructure:"DB_PASSWORD"` // Senha do banco de dados
	DBName     string `mapstructure:"DB_NAME"`     // Nome do banco de dados

	RabbitMQUser     string `mapstructure:"RABBITMQ_USER"`     // Usuário do RabbitMQ
	RabbitMQPassword string `mapstructure:"RABBITMQ_PASSWORD"` // Senha do RabbitMQ
	RabbitMQHost     string `mapstructure:"RABBITMQ_HOST"`     // Host do RabbitMQ
	RabbitMQPort     string `mapstructure:"RABBITMQ_PORT"`     // Porta do RabbitMQ

	WebServerPort string `mapstructure:"WEB_SERVER_PORT"` // Porta do servidor web
	GRPCPort      string `mapstructure:"GRPC_PORT"`       // Porta do servidor gRPC
}

// Gerar string de conexão com o banco de dados PostgreSQL
func (cfg *Config) GetPostgresConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,     // Usuário
		cfg.DBPassword, // Senha
		cfg.DBHost,     // Host
		cfg.DBPort,     // Porta
		cfg.DBName,     // Nome do banco
	)
}

func LoadConfig(path string) (*Config, error) {
	var cfg *Config
	viper.AutomaticEnv() // Carrega as variáveis de ambiente

	// Definir o caminho e nome do arquivo de configuração
	if path != "" {
		viper.SetConfigName(".env") // Nome do arquivo de configuração
		viper.SetConfigType("env")  // Tipo do arquivo
		viper.AddConfigPath(path)   // Caminho do diretório onde o arquivo está localizado

		// Tentar carregar o arquivo .env
		if err := viper.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("LoadConfig - erro ao carregar as configuracoes do arquivo .env: %w", err)
		}
	}

	// Deserializar as configurações no struct Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("LoadConfig - erro ao desserializar as configuracoes: %w", err)
	}

	return cfg, nil
}
