package cmd

import (
	"net/http"
	"strings"
	"time"

	"github.com/orensimple/otus_hw1_8/config"
	"github.com/orensimple/otus_hw1_8/internal/domain/handlers"
	"github.com/orensimple/otus_hw1_8/internal/domain/mw"
	"github.com/orensimple/otus_hw1_8/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addr string

const TimeOut = 50 * time.Millisecond

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
		router.InitDB()
		mux := http.NewServeMux()
		mux.HandleFunc("/events_for_day", mw.HTTPLogger(mw.WithTimeout(router.GetEventsByDay, TimeOut)))
		mux.HandleFunc("/events_for_week", mw.HTTPLogger(mw.WithTimeout(router.GetEventsByWeek, TimeOut)))
		mux.HandleFunc("/events_for_month", mw.HTTPLogger(mw.WithTimeout(router.GetEventsByMonth, TimeOut)))
		mux.HandleFunc("/create_event", mw.HTTPLogger(mw.WithTimeout(router.CreateEvent, TimeOut)))
		mux.HandleFunc("/update_event", mw.HTTPLogger(mw.WithTimeout(router.UpdateEvent, TimeOut)))
		mux.HandleFunc("/delete_event", mw.HTTPLogger(mw.WithTimeout(router.DeleteEvent, TimeOut)))

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
