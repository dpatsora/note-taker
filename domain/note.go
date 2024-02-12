package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type NoteRepository interface {
	Create(note Note) error

	GetAllNotes() ([]Note, error)
	GetNoteByUUID(noteUUID uuid.UUID) (Note, error)

	Update(
		noteUUID uuid.UUID,
		updateFn func(ctx context.Context, n *Note) (*Note, error),
	) error

	Delete(noteUUID uuid.UUID) error
}

type Note struct {
	UUID        uuid.UUID
	Title       string
	Description string
	Timestamp   time.Time
}

func NewNote(title, description string) *Note {
	return &Note{
		UUID:        uuid.New(),
		Title:       title,
		Description: description,
		Timestamp:   time.Now(),
	}
}

func (n *Note) Update(title, description string) {
	n.Title = title
	n.Description = description
}
