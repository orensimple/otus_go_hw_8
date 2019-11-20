package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/orensimple/otus_hw1_8/internal/domain/services"
	"github.com/orensimple/otus_hw1_8/internal/memory"
	"github.com/spf13/cobra"
)

// RootCmd asdf
var RootCmd = &cobra.Command{
	Use:   "clncnd",
	Short: "CleanCalendar is a calendar micorservice demo",
}

// TestCmd asdf
var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long:  `A longer description for the delete command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the Test command!")
		eventStorage := memory.NewMemEventStorage()

		eventService := &services.EventService{
			EventStorage: eventStorage,
		}
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()
		t := time.Now()
		eventService.CreateEvent(ctx, 1, `a`, `b`, `c`, t, t)
		_, err := eventService.CreateEvent(ctx, 1, `a`, `b`, `c`, t, t)
		fmt.Printf("%s\n", err)
		test, _ := eventService.GetEvents(ctx)
		fmt.Println(test[0].Owner)
		eventService.UpdateEvent(ctx, 1, `b`, `a`, `a`, t, t)
		test2, _ := eventService.GetEvents(ctx)
		fmt.Println(test2[0].Owner)
		_, err = eventService.UpdateEvent(ctx, 5, `b5`, `a`, `a`, t, t)
		fmt.Printf("%s\n", err)
		err = eventService.DeleteEvent(ctx, 1)
		fmt.Printf("%s\n", err)
		err = eventService.DeleteEvent(ctx, 2)
		fmt.Printf("%s\n", err)
	},
}

func init() {
	RootCmd.AddCommand()
	RootCmd.AddCommand(TestCmd)
}
