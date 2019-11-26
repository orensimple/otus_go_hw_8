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
		tSt, _ := time.Parse("2006-01-02 15:04", "2019-11-01 20:00")
		tEn, _ := time.Parse("2006-01-02 15:04", "2019-11-21 20:59")
		eventService.CreateEvent(ctx, 1, `a`, `b`, `c`, tSt, tEn)
		_, err := eventService.CreateEvent(ctx, 1, `a`, `b`, `c`, tSt, tEn)
		fmt.Printf("%s\n", err)
		test, _ := eventService.GetEvents(ctx)
		fmt.Println(test[0].Owner)
		eventService.UpdateEvent(ctx, 1, `b`, `a`, `a`, tSt, tEn)
		test2, _ := eventService.GetEvents(ctx)
		fmt.Println(test2[0].Owner)
		//test3, _ := eventService.GetEventsByTime(ctx, "day")
		//fmt.Println(test3[0].Owner)
		//test4, _ := eventService.GetEventsByTime(ctx, "week")
		//fmt.Println(test4[0].Owner)
		test5, _ := eventService.GetEventsByTime(ctx, "month")
		fmt.Println(test5[0].Owner)
		_, err = eventService.UpdateEvent(ctx, 5, `b5`, `a`, `a`, tSt, tEn)
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
