package configs

import (
	"github.com/go-chi/jwtauth" // Esse pacote é para autenticação JWT
	"github.com/spf13/viper"
)

// A struct nao esta exportado, pois nao queremos que seja acessado diretamente
type conf struct {
	// Configurações do servidor
	DBDriver      string           `mapstructure:"DB_DRIVER"`       // Driver do banco de dados
	DBHost        string           `mapstructure:"DB_HOST"`         // Host do banco de dados
	DBPort        string           `mapstructure:"DB_PORT"`         // Porta do banco de dados
	DBUser        string           `mapstructure:"DB_USER"`         // Usuário do banco de dados
	DBPassword    string           `mapstructure:"DB_PASSWORD"`     // Senha do banco de dados
	DBName        string           `mapstructure:"DB_NAME"`         // Nome do banco de dados
	WebServerPort string           `mapstructure:"WEB_SERVER_PORT"` // Porta do servidor web
	JWTSecret     string           `mapstructure:"JWT_SECRET"`      // Chave secreta do JWT
	JWTExpiresIn  int              `mapstructure:"JWT_EXPIRES_IN"`  // Tempo de expiração do JWT
	TokenAuth     *jwtauth.JWTAuth // Token de autenticação
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")                                                           // Nome do arquivo de configuração
	viper.SetConfigType("env")                                                                  // Tipo do arquivo de configuração
	viper.AddConfigPath("./APIS/")                                                              // Caminho do arquivo de configuração
	viper.SetConfigFile("/Users/wrferreira/Documents/treinamentos-testes/APIS/cmd/server/.env") // Caminho do arquivo de configuração
	viper.AutomaticEnv()                                                                        // Carrega as variáveis de ambiente

	// Carrega as configurações do arquivo .env
	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file: " + err.Error())
	}

	// Converte as configurações do arquivo .env para a struct conf
	if err := viper.Unmarshal(&cfg); err != nil {
		panic("Unable to decode into struct, %v" + err.Error())
	}

	// Criando o token de autenticação
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return cfg, nil
}

//
