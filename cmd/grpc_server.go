package cmd

import (
	"github.com/orensimple/otus_hw1_8/internal/domain/services"
	"github.com/orensimple/otus_hw1_8/internal/grpc/api"
	"github.com/orensimple/otus_hw1_8/internal/maindb"
	"github.com/spf13/cobra"
	"log"
)

// TODO: dependency injection, orchestrator
func construct(dsn string) (*api.CalendarServer, error) {
	eventStorage, err := maindb.NewPgEventStorage(dsn)
	if err != nil {
		return nil, err
	}
	eventService := &services.EventService{
		EventStorage: eventStorage,
	}
	server := &api.CalendarServer{
		EventService: eventService,
	}
	return server, nil
}

var addrGRPC string
var dsn string

var GRPCServerCmd = &cobra.Command{
	Use:   "grpc_server",
	Short: "Run grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		server, err := construct(dsn)
		if err != nil {
			log.Fatal(err)
		}
		err = server.Serve(addrGRPC)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	GRPCServerCmd.Flags().StringVar(&addrGRPC, "addr", "localhost:8088", "host:port to listen")
	GRPCServerCmd.Flags().StringVar(&dsn, "dsn", "host=127.0.0.1 user=event_user password=event_pwd dbname=event_db", "database connection string")
	RootCmd.AddCommand(GRPCServerCmd)
}
