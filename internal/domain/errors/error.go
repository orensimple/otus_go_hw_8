package errors

// EventError type
type EventError string

func (ee EventError) Error() string {
	return string(ee)
}

var (
	//ErrOverlaping comment
	ErrOverlaping = EventError("another event exists for this date")
	//ErrIncorrectEndDate comment
	ErrIncorrectEndDate = EventError("end_date is incorrect")
)
