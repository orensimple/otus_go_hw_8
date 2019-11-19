package errors

// EventError type
type EventError string

func (ee EventError) Error() string {
	return string(ee)
}

var (
	//ErrOverlaping comment
	ErrOverlaping = EventError("another event exists for this date")
	//ErrEventExist comment
	ErrEventExist = EventError("event alredy exist")
	//ErrEventNotFound comment
	ErrEventNotFound = EventError("not found event")
)
