package webserver

const (
	DefaultPort          int = 8080
	DefaultStopTimeoutMS int = 5000
)

var DefaultWebServerConfig = WebServerConfig{
	Port:          DefaultPort,
	StopTimeoutMS: DefaultStopTimeoutMS,
	HttpHandler: HttpHandlerConfig{
		UseLogger:   true,
		UseRecovery: true,
	},
}

type Configuration struct {
	WebServer WebServerConfig
}

type WebServerConfig struct {
	Port          int
	StopTimeoutMS int
	HttpHandler   HttpHandlerConfig
}

type HttpHandlerConfig struct {
	UseLogger   bool
	UseRecovery bool
}

//func LoadConfiguration(path string) (config *Configuration) {
//	data, err := os.ReadFile(path)
//	if err != nil {
//		fmt.Printf("error reading config: %v\n", err)
//		config = loadDefaultConfiguration()
//		return
//	}
//
//	err = json.Unmarshal(data, &config)
//	if err != nil {
//		fmt.Printf("error unmarshalling configuration: %v\n", err)
//		config = loadDefaultConfiguration()
//		return
//	}
//	return
//}
//
//func loadDefaultConfiguration() *Configuration {
//	fmt.Println("loading default configuration...")
//	return &Configuration{
//		WebServer: WebServerConfig{
//			StopTimeoutMS: DefaultPort,
//			Port:          DefaultStopTimeoutMS,
//		}}
//}
