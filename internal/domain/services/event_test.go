package services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEvents(t *testing.T) {
	id := "id"
	user := &models.User{ID: id}

	s := NewEventLocalStorage()

	for i := 0; i < 10; i++ {
		bm := &models.Event{
			ID:     fmt.Sprintf("id%d", i),
			UserID: user.ID,
		}

		err := s.CreateEvent(context.Background(), user, bm)
		assert.NoError(t, err)
	}

	returnedEvents, err := s.GetEvents(context.Background(), user)
	assert.NoError(t, err)

	assert.Equal(t, 10, len(returnedEvents))
}

/*func TestDeleteBookmark(t *testing.T) {
	id1 := "id1"
	id2 := "id2"

	user1 := &models.User{ID: id1}
	user2 := &models.User{ID: id2}

	bmID := "bmID"
	bm := &models.Bookmark{ID: bmID, UserID: user1.ID}

	s := NewBookmarkLocalStorage()

	err := s.CreateBookmark(context.Background(), user1, bm)
	assert.NoError(t, err)

	err = s.DeleteBookmark(context.Background(), user1, bmID)
	assert.NoError(t, err)

	err = s.CreateBookmark(context.Background(), user1, bm)
	assert.NoError(t, err)

	err = s.DeleteBookmark(context.Background(), user2, bmID)
	assert.Error(t, err)
	assert.Equal(t, err, bookmark.ErrBookmarkNotFound)
}*/
