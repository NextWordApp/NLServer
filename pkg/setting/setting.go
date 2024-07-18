package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
)

var Config = new(AppConfig)

type AppConfig struct {
	*ServerConfig `mapsturcture:"server"`
	//*LogConfig    `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*Eth         `mapstructure:"eth"`
	*AI          `mapstructure:"ai"`
	*RES         `mapstructure:"res"`
}

type ServerConfig struct {
	Name       string `mapstructure:"name"`
	Mode       string `mapstructure:"mode"`
	Port       string `mapstructure:"port"`
	PrivateKey string `mapstructure:"privateKey"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"poolSize"`
}

//type LogConfig struct {
//	Level      string `mapstructure:"level"`
//	Filename   string `mapstructure:"filename"`
//	MaxSize    int    `mapstructure:"maxSize"`
//	MaxAge     int    `mapstructure:"maxAge"`
//	MaxBackups int    `mapstructure:"maxBackups"`
//}

type Eth struct {
	apiKeyUrl   string `mapstructure:"apiKeyUrl"`
	contractUrl string `mapstructure:"ContractUrl"`
}

type AI struct {
	ApiKey string `mapstructure:"apiKey"`
	Url    string `mapstructure:"url"`
}

type RES struct {
	WordPath string `mapstructure:"wordPath"`
}

// Init 初始化配置
func Init(configPath string) error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("viper.ReadInConfig() failed, err: %v", err)
	}

	overrideWithEnv()

	if err := viper.Unmarshal(Config); err != nil {
		return fmt.Errorf("viper.Unmarshal failed, err: %v", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
		overrideWithEnv()
		if err := viper.Unmarshal(Config); err != nil {
			fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
		}
	})
	return nil
}

// overrideWithEnv 使用环境变量覆盖配置文件中的值
func overrideWithEnv() {
	viper.AutomaticEnv()
	for _, key := range viper.AllKeys() {
		value := viper.GetString(key)
		viper.Set(key, expandEnv(value))
	}
}

// expandEnv 展开环境变量
func expandEnv(value string) string {
	return os.ExpandEnv(value)
}
