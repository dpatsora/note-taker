package adapters

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/dpatsora/note-taker/domain"
)

type NotePostgresqlRepository struct {
	db *gorm.DB
}

func NewNotePostgresqlRepository(db *gorm.DB) *NotePostgresqlRepository {
	return &NotePostgresqlRepository{db: db}
}

type NoteModel struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	Timestamp   time.Time `gorm:"column:timestamp"`
}

func (NoteModel) TableName() string {
	return "notes"
}

func (n NotePostgresqlRepository) Create(note domain.Note) error {
	model := NoteModel{
		ID:          note.UUID,
		Title:       note.Title,
		Description: note.Description,
		Timestamp:   note.Timestamp,
	}

	return n.db.Model(&NoteModel{}).Create(model).Error
}

func (n NotePostgresqlRepository) GetAllNotes() ([]domain.Note, error) {
	var notes []NoteModel

	err := n.db.Find(&notes).Error
	if err != nil {
		return nil, err
	}

	var result []domain.Note
	for _, note := range notes {
		result = append(result, modelToEntity(note))
	}

	return result, nil
}

func (n NotePostgresqlRepository) GetNoteByUUID(noteUUID uuid.UUID) (domain.Note, error) {
	var note NoteModel
	err := n.db.
		Where(NoteModel{ID: noteUUID}).
		First(&note).
		Error
	if err != nil {
		return domain.Note{}, err
	}
	return modelToEntity(note), err
}

func (n NotePostgresqlRepository) Update(noteUUID uuid.UUID, updateFn func(ctx context.Context, n *domain.Note) (*domain.Note, error)) error {
	tx := n.db.Begin()
	defer tx.Rollback()

	var note NoteModel
	err := tx.
		Where(NoteModel{ID: noteUUID}).
		First(&note).
		Error
	if err != nil {
		return err
	}

	entity := modelToEntity(note)
	updatedNote, err := updateFn(context.Background(), &entity)
	if err != nil {
		return err
	}

	model := toModel(*updatedNote)
	err = tx.Model(&NoteModel{}).Where(NoteModel{ID: noteUUID}).Updates(map[string]interface{}{"title": model.Title, "description": model.Description, "timestamp": model.Timestamp}).Error
	if err != nil {
		return err
	}

	return tx.Commit().Error
}

func (n NotePostgresqlRepository) Delete(noteUUID uuid.UUID) error {
	return n.db.Delete(&NoteModel{}, noteUUID).Error
}

func toModel(note domain.Note) NoteModel {
	return NoteModel{
		ID:          note.UUID,
		Title:       note.Title,
		Description: note.Description,
		Timestamp:   note.Timestamp,
	}
}

func modelToEntity(m NoteModel) domain.Note {
	return domain.Note{
		UUID:        m.ID,
		Title:       m.Title,
		Description: m.Description,
		Timestamp:   m.Timestamp,
	}
}
