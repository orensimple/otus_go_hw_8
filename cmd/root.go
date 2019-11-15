package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/orensimple/otus_hw1_8/internal/domain/services"
	"github.com/spf13/cobra"
)

// RootCmd asdf
var RootCmd = &cobra.Command{
	Use:   "clncnd",
	Short: "CleanCalendar is a calendar micorservice demo",
}

// StartCmd asdf
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long:  `A longer description for the delete command.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()
		var service services.EventService
		t := time.Now()
		service.CreateEvent(ctx, 1, `a`, `b`, `c`, &t, &t)
		fmt.Println("This is the start command!")
	},
}

func init() {
	RootCmd.AddCommand()
	RootCmd.AddCommand(StartCmd)
}
