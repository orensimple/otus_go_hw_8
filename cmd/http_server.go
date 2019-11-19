package cmd

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/orensimple/otus_hw1_8/config"
	"github.com/orensimple/otus_hw1_8/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addr string
var contextLogger logger.Logger

// HTTPServerCmd init
var HTTPServerCmd = &cobra.Command{
	Use:   "http_server",
	Short: "Run http server",
	Run: func(cmd *cobra.Command, args []string) {
		config.Init(addr)
		InitLogger()
		var serverAddr strings.Builder
		serverAddr.WriteString(viper.GetString("http_listen.ip"))
		serverAddr.WriteString(":")
		serverAddr.WriteString(viper.GetString("http_listen.port"))
		contextLogger.Infof("Starting http server", viper.GetString("http_listen.ip"), viper.GetString("http_listen.port"))
		handler := &MyHandler{}
		// создаем HTTP сервер
		server := &http.Server{
			Addr:           serverAddr.String(),
			Handler:        handler,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		// запускаем сервер, это заблокирует текущую горутину
		server.ListenAndServe()
	},
}

func init() {
	HTTPServerCmd.Flags().StringVar(&addr, "config", "localhost:8080", "host:port to listen")
	RootCmd.AddCommand(HTTPServerCmd)
}

// MyHandler http handlers
type MyHandler struct {
	// все нужные вам объекты: конфиг, логер, соединение с базой и т.п.
}

func (h *MyHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/hello" {
		uri := req.URL.RequestURI()

		contextLogger.Debugf("Debug info do", "uri", uri)
		contextLogger.Infof("Starting with zap")

		resp.Header().Set("Content-Type", "application/json; charset=utf-8")
		resp.WriteHeader(200)
		json.NewEncoder(resp).Encode("ok")
	}
}

// InitLogger init
func InitLogger() {
	loggerConfig := logger.Configuration{
		EnableConsole:     true,
		ConsoleLevel:      viper.GetString("log_level.command"),
		ConsoleJSONFormat: true,
		EnableFile:        true,
		FileLevel:         viper.GetString("log_level.file"),
		FileJSONFormat:    true,
		FileLocation:      viper.GetString("log_file"),
	}
	err := logger.NewLogger(loggerConfig, logger.InstanceZapLogger)
	if err != nil {
		log.Fatalf("Could not instantiate log %s", err.Error())
	}

	contextLogger = logger.WithFields(logger.Fields{"version": "0.0.1"})
}
