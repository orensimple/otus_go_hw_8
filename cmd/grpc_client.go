package cmd

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/orensimple/otus_hw1_8/internal/grpc/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"time"
)

var server, title, text, startTime, endTime string

const tsLayout = "2006-01-02T15:04:05"

func parseTs(s string) (*timestamp.Timestamp, error) {
	t, err := time.Parse(tsLayout, s)
	if err != nil {
		return nil, err
	}
	ts, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

var GRPCClientCmd = &cobra.Command{
	Use:   "grpc_client",
	Short: "Run grpc client",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(server, grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		client := api.NewCalendarServiceClient(conn)
		st, err := parseTs(startTime)
		if err != nil {
			log.Fatal(err)
		}
		et, err := parseTs(endTime)
		if err != nil {
			log.Fatal(err)
		}
		req := &api.CreateEventRequest{
			Title:     title,
			Text:      text,
			StartTime: st,
			EndTime:   et,
		}
		resp, err := client.CreateEvent(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		if resp.GetError() != "" {
			log.Fatal(resp.GetError())
		} else {
			log.Println(resp.GetEvent().ID)
		}
	},
}

func init() {
	GRPCClientCmd.Flags().StringVar(&server, "server", "localhost:8088", "host:port to connect to")
	GRPCClientCmd.Flags().StringVar(&title, "title", "", "event title")
	GRPCClientCmd.Flags().StringVar(&text, "text", "", "event text")
	GRPCClientCmd.Flags().StringVar(&startTime, "start-time", "", "event start time, format: "+tsLayout)
	GRPCClientCmd.Flags().StringVar(&endTime, "end-time", "", "event end time, format: "+tsLayout)
	RootCmd.AddCommand(GRPCClientCmd)
}
