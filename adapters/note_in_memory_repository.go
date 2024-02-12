package adapters

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"

	"github.com/dpatsora/note-taker/domain"
)

type NoteInMemoryRepository struct {
	notes map[uuid.UUID]domain.Note

	mu sync.RWMutex
}

func NewNoteInMemoryRepository() *NoteInMemoryRepository {
	return &NoteInMemoryRepository{
		notes: make(map[uuid.UUID]domain.Note),
	}
}

func (r *NoteInMemoryRepository) Create(note domain.Note) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.notes[note.UUID] = note

	return nil
}

func (r *NoteInMemoryRepository) GetAllNotes() ([]domain.Note, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	notes := make([]domain.Note, 0, len(r.notes))
	for _, note := range r.notes {
		notes = append(notes, note)
	}

	return notes, nil
}

func (r *NoteInMemoryRepository) GetNoteByUUID(noteUUID uuid.UUID) (domain.Note, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	note, ok := r.notes[noteUUID]
	if !ok {
		return domain.Note{}, fmt.Errorf("note with UUID %s not found", noteUUID)
	}

	return note, nil
}

func (r *NoteInMemoryRepository) Delete(noteUUID uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.notes, noteUUID)

	return nil
}

func (r *NoteInMemoryRepository) Update(
	noteUUID uuid.UUID,
	updateFn func(ctx context.Context, n *domain.Note) (*domain.Note, error),
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	note, ok := r.notes[noteUUID]
	if !ok {
		return fmt.Errorf("note with UUID %s not found", noteUUID)
	}

	updatedNote, err := updateFn(context.Background(), &note)
	if err != nil {
		return err
	}

	r.notes[noteUUID] = *updatedNote

	return nil
}
