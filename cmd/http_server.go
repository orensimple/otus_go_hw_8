package cmd

import (
	"net/http"
	"strings"
	"time"

	"github.com/orensimple/otus_hw1_8/config"
	"github.com/orensimple/otus_hw1_8/internal/domain/handlers"
	"github.com/orensimple/otus_hw1_8/internal/domain/middlewares"
	"github.com/orensimple/otus_hw1_8/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addr string

// HTTPServerCmd init
var HTTPServerCmd = &cobra.Command{
	Use:   "http_server",
	Short: "Run http server",
	Run: func(cmd *cobra.Command, args []string) {
		config.Init(addr)
		logger.InitLogger()

		var serverAddr strings.Builder
		serverAddr.WriteString(viper.GetString("http_listen.ip"))
		serverAddr.WriteString(":")
		serverAddr.WriteString(viper.GetString("http_listen.port"))
		logger.ContextLogger.Infof("Starting http server", viper.GetString("http_listen.ip"), viper.GetString("http_listen.port"))
		router := &handlers.Handler{}
		router.AddForTest()
		mux := http.NewServeMux()
		mux.HandleFunc("/events_for_day", middlewares.HTTPLogger(middlewares.WithTimeout(router.GetEventsByDay, 50*time.Millisecond)))
		mux.HandleFunc("/events_for_week", middlewares.HTTPLogger(middlewares.WithTimeout(router.GetEventsByWeek, 50*time.Millisecond)))
		mux.HandleFunc("/events_for_month", middlewares.HTTPLogger(middlewares.WithTimeout(router.GetEventsByMonth, 50*time.Millisecond)))
		mux.HandleFunc("/create_event", middlewares.HTTPLogger(middlewares.WithTimeout(router.CreateEvent, 50*time.Millisecond)))
		mux.HandleFunc("/update_event", middlewares.HTTPLogger(middlewares.WithTimeout(router.UpdateEvent, 50*time.Millisecond)))
		mux.HandleFunc("/delete_event", middlewares.HTTPLogger(middlewares.WithTimeout(router.DeleteEvent, 50*time.Millisecond)))

		server := &http.Server{
			Addr:           serverAddr.String(),
			Handler:        mux,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		server.ListenAndServe()
	},
}

func init() {
	HTTPServerCmd.Flags().StringVar(&addr, "config", "./config", "")
	RootCmd.AddCommand(HTTPServerCmd)
}
