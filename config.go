package orange

import(
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
	"strings"
	"time"
)

var defaultReplacer =  strings.NewReplacer(".", "_")

type Config struct{
	replacer *strings.Replacer
	prefix   string
	autoenv bool
	filename string
	path string
	filetype string
	app *App
	vconfig *viper.Viper
}


// load config file
func (config *Config) load() error{
	var (
		err       error
	)
	config.vconfig = viper.New()
	config.vconfig.SetEnvKeyReplacer(config.replacer)
	config.vconfig.SetEnvPrefix(config.prefix)
	if config.autoenv{
		config.vconfig.AutomaticEnv()
	}
	config.vconfig.SetConfigName(config.filename)
	config.vconfig.AddConfigPath(config.path)
	config.vconfig.SetConfigType(config.filetype)
	if err = config.vconfig.ReadInConfig(); err != nil {
		return err
	}
	config.vconfig.WatchConfig()
	config.vconfig.OnConfigChange(func(e fsnotify.Event) {
		colorLog("[INFO] " + config.filename + " file changed %s:", e.Name)
	})
	return err
}

func (config *Config) SetReplacer(replacer strings.Replacer) {
	config.replacer = &replacer
}

func (config *Config) GetReplacer() strings.Replacer{
	return *config.replacer
}

func (config *Config) SetENVPrefix(prefix string){
	config.prefix = prefix
}

func (config *Config) GetENVPrefix() string{
	return config.prefix
}

func (config *Config) Path() string{
	return config.path
}

func (config *Config) Filetype() string{
	return config.filetype
}

func (config *Config) Filename() string{
	return config.filename
}

// Get value by key
func (config *Config) GetInt(key string) int{
	return config.vconfig.GetInt(key)
}
// Get value by key
func (config *Config) GetInt64(key string) int64{
	return config.vconfig.GetInt64(key)
}

// Get value by key
func (config *Config) GetFloat(key string) float64{
	return config.vconfig.GetFloat64(key)
}

// Get value by key
func (config *Config) GetString(key string) string{
	return config.vconfig.GetString(key)
}

// Get value by key
func (config *Config) GetBool(key string) bool{
	return config.vconfig.GetBool(key)
}

// Get value by key
func (config *Config) GetTimeDuration(key string) time.Duration {
	return config.vconfig.GetDuration(key)
}

func (config *Config) GetStringMap(key string) map[string]interface{} {
	return config.vconfig.GetStringMap(key)
}

func (config *Config) GetStringMapString(key string) map[string]string{
	return config.vconfig.GetStringMapString(key)
}

func (config *Config) GetStringMapStringSlice(key string) map[string][]string{
	return config.vconfig.GetStringMapStringSlice(key)
}

func (config *Config) GetStringSlice(key string) []string{
	return config.vconfig.GetStringSlice(key)
}

func (config *Config) AllKeys() []string{
	return config.vconfig.AllKeys()
}

// Set value by key
func (config *Config) Set(key string, i interface{}) bool{
	config.vconfig.Set(key, i)
	return config.vconfig.IsSet(key)
}




