package ports

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"gorm.io/gorm"

	"github.com/dpatsora/note-taker/domain"
	"github.com/dpatsora/note-taker/pkg/server/httperr"
)

type HttpServer struct {
	noteRepository domain.NoteRepository
}

func NewHttpServer(noteRepository domain.NoteRepository) HttpServer {
	return HttpServer{noteRepository}
}

func (h HttpServer) GetNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := h.noteRepository.GetAllNotes()
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	notesResp := notesToResponse(notes)

	render.Respond(w, r, notesResp)
}

func (h HttpServer) CreateNote(w http.ResponseWriter, r *http.Request) {
	postNote := PostNote{}
	if err := render.Decode(r, &postNote); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	note := domain.NewNote(postNote.Title, postNote.Description)
	err := h.noteRepository.Create(*note)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.Header().Set("content-location", "/notes/"+note.UUID.String())
	w.WriteHeader(http.StatusNoContent)
}

func (h HttpServer) DeleteNote(w http.ResponseWriter, r *http.Request, noteUUID openapi_types.UUID) {
	err := h.noteRepository.Delete(noteUUID)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h HttpServer) GetNote(w http.ResponseWriter, r *http.Request, noteUUID openapi_types.UUID) {
	note, err := h.noteRepository.GetNoteByUUID(noteUUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httperr.BadRequest("note-not-found", err, w, r)
			return
		}
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	noteResp := noteToResponse(note)
	render.Respond(w, r, noteResp)
}

func (h HttpServer) UpdateNote(w http.ResponseWriter, r *http.Request, noteUUID openapi_types.UUID) {
	postNote := PostNote{}
	if err := render.Decode(r, &postNote); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	err := h.noteRepository.Update(noteUUID, func(ctx context.Context, n *domain.Note) (*domain.Note, error) {
		n.Update(postNote.Title, postNote.Description)
		return n, nil
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httperr.BadRequest("note-not-found", err, w, r)
			return
		}
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func notesToResponse(notes []domain.Note) Notes {
	var responseNotes []Note
	for _, note := range notes {
		responseNotes = append(responseNotes, noteToResponse(note))
	}
	return Notes{responseNotes}
}

func noteToResponse(note domain.Note) Note {
	return Note{
		Description: note.Description,
		Time:        note.Timestamp,
		Title:       note.Title,
		Uuid:        note.UUID,
	}
}
