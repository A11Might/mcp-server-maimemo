package main

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type MaimemoHandler struct {
	maimemoClient *MaimemoClient
}

func NewMaimemoHanlder(client *MaimemoClient) (*MaimemoHandler, error) {
	return &MaimemoHandler{
		maimemoClient: client,
	}, nil
}

type ListInterpretationsParams struct {
	VocId string `json:"voc_id"`
}

// 获取释义
func (handler *MaimemoHandler) ListInterpretations(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[ListInterpretationsParams]) (*mcp.CallToolResultFor[any], error) {
	interpretations, err := handler.maimemoClient.ListInterpretations(params.Arguments.VocId)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(interpretations)
}

type CreateInterpretationParams struct {
	VocId          string               `json:"voc_id"`
	Interpretation string               `json:"interpretation"`
	Tags           []InterpretationTag  `json:"tags"`
	Status         InterpretationStatus `json:"status"`
}

// 创建释义
func (handler *MaimemoHandler) CreateInterpretation(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[CreateInterpretationParams]) (*mcp.CallToolResultFor[any], error) {
	interpretation, err := handler.maimemoClient.CreateInterpretation(params.Arguments.VocId, params.Arguments.Interpretation, params.Arguments.Tags, params.Arguments.Status)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(interpretation)
}

type UpdateInterpretationParams struct {
	InterpretationId string               `json:"interpretation_id"`
	Interpretation   string               `json:"interpretation"`
	Tags             []InterpretationTag  `json:"tags"`
	Status           InterpretationStatus `json:"status"`
}

// 更新释义
func (handler *MaimemoHandler) UpdateInterpretation(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[UpdateInterpretationParams]) (*mcp.CallToolResultFor[any], error) {
	interpretation, err := handler.maimemoClient.UpdateInterpretation(params.Arguments.InterpretationId, params.Arguments.Interpretation, params.Arguments.Tags, params.Arguments.Status)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(interpretation)
}

type DeleteInterpretationParams struct {
	InterpretationId string `json:"interpretation_id"`
}

// 删除释义
func (handler *MaimemoHandler) DeleteInterpretation(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[DeleteInterpretationParams]) (*mcp.CallToolResultFor[any], error) {
	interpretation, err := handler.maimemoClient.DeleteInterpretation(params.Arguments.InterpretationId)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(interpretation)
}

type ListNotesParams struct {
	VocId string `json:"voc_id"`
}

// 获取助记
func (handler *MaimemoHandler) ListNotes(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[ListNotesParams]) (*mcp.CallToolResultFor[any], error) {
	notes, err := handler.maimemoClient.ListNotes(params.Arguments.VocId)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(notes)
}

type CreateNoteParams struct {
	VocId    string   `json:"voc_id"`
	NoteType NoteType `json:"note_type"`
	Note     string   `json:"note"`
}

// 创建助记
func (handler *MaimemoHandler) CreateNote(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[CreateNoteParams]) (*mcp.CallToolResultFor[any], error) {
	note, err := handler.maimemoClient.CreateNote(params.Arguments.VocId, params.Arguments.NoteType, params.Arguments.Note)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(note)
}

type UpdateNoteParams struct {
	NoteId   string   `json:"note_id"`
	NoteType NoteType `json:"note_type"`
	Note     string   `json:"note"`
}

// 更新助记
func (handler *MaimemoHandler) UpdateNote(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[UpdateNoteParams]) (*mcp.CallToolResultFor[any], error) {
	note, err := handler.maimemoClient.UpdateNote(params.Arguments.NoteId, params.Arguments.NoteType, params.Arguments.Note)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(note)
}

type DeleteNoteParams struct {
	NoteId string `json:"note_id"`
}

// 删除助记
func (handler *MaimemoHandler) DeleteNote(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[DeleteNoteParams]) (*mcp.CallToolResultFor[any], error) {
	note, err := handler.maimemoClient.DeleteNote(params.Arguments.NoteId)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(note)
}

type ListNotepadParams struct {
	Ids    []string `json:"ids"`
	Limit  int      `json:"limit"`
	Offset int      `json:"offset"`
}

// 查询云词本
func (handler *MaimemoHandler) ListNotepad(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[ListNotepadParams]) (*mcp.CallToolResultFor[any], error) {
	notepadList, err := handler.maimemoClient.ListNotepads(params.Arguments.Ids, params.Arguments.Limit, params.Arguments.Offset)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(notepadList)
}

type CreateNotepadParams struct {
	Status      NotepadStatus `json:"status"`
	ChapterName string        `json:"chapter_name"`
	Words       []string      `json:"words"`
	Title       string        `json:"title"`
	Brief       string        `json:"brief"`
	Tags        []NotepadTag  `json:"tags"`
}

// 创建云词本
func (handler *MaimemoHandler) CreateNotepad(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[CreateNotepadParams]) (*mcp.CallToolResultFor[any], error) {
	content := FormateNotepadContent(params.Arguments.ChapterName, params.Arguments.Words)
	notepad, err := handler.maimemoClient.CreateNotepad(params.Arguments.Status, content, params.Arguments.Title, params.Arguments.Brief, params.Arguments.Tags)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(notepad)
}

type GetNotepadParams struct {
	NotepadId string `json:"notepad_id"`
}

// 获取云词本
func (handler *MaimemoHandler) GetNotepad(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[GetNotepadParams]) (*mcp.CallToolResultFor[any], error) {
	notepad, err := handler.maimemoClient.GetNotepad(params.Arguments.NotepadId)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(notepad)
}

type UpdateNotepadParams struct {
	NotepadId   string        `json:"notepad_id"`
	Status      NotepadStatus `json:"status"`
	ChapterName string        `json:"chapter_name"`
	Words       []string      `json:"words"`
	Title       string        `json:"title"`
	Brief       string        `json:"brief"`
	Tags        []NotepadTag  `json:"tags"`
}

// 更新云词本
func (handler *MaimemoHandler) UpdateNotepad(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[UpdateNotepadParams]) (*mcp.CallToolResultFor[any], error) {
	content := FormateNotepadContent(params.Arguments.ChapterName, params.Arguments.Words)
	notepad, err := handler.maimemoClient.UpdateNotepad(params.Arguments.NotepadId, params.Arguments.Status, content, params.Arguments.Title, params.Arguments.Brief, params.Arguments.Tags)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(notepad)
}

type DeleteNotepadParams struct {
	NotepadId string `json:"notepad_id"`
}

func (handler *MaimemoHandler) DeleteNotepad(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[DeleteNotepadParams]) (*mcp.CallToolResultFor[any], error) {
	notepad, err := handler.maimemoClient.DeleteNotepad(params.Arguments.NotepadId)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(notepad)
}

type ListPhrasesParams struct {
	VocId string `json:"voc_id"`
}

// 获取例句
func (handler *MaimemoHandler) ListPhrases(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[ListPhrasesParams]) (*mcp.CallToolResultFor[any], error) {
	phrases, err := handler.maimemoClient.ListPhrases(params.Arguments.VocId)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(phrases)
}

type CreatePhraseParams struct {
	VocId          string      `json:"voc_id"`
	Phrase         string      `json:"phrase"`
	Interpretation string      `json:"interpretation"`
	Tags           []PhraseTag `json:"tags"`
	Origin         string      `json:"origin"`
}

// 创建例句
func (handler *MaimemoHandler) CreatePhrase(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[CreatePhraseParams]) (*mcp.CallToolResultFor[any], error) {
	phrase, err := handler.maimemoClient.CreatePhrase(params.Arguments.VocId, params.Arguments.Phrase, params.Arguments.Interpretation, params.Arguments.Tags, params.Arguments.Origin)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(phrase)
}

type UpdatePhraseParams struct {
	PhraseId       string      `json:"phrase_id"`
	Phrase         string      `json:"phrase"`
	Interpretation string      `json:"interpretation"`
	Tags           []PhraseTag `json:"tags"`
	Origin         string      `json:"origin"`
}

// 更新例句
func (handler *MaimemoHandler) UpdatePhrase(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[UpdatePhraseParams]) (*mcp.CallToolResultFor[any], error) {
	phrase, err := handler.maimemoClient.UpdatePhrase(params.Arguments.PhraseId, params.Arguments.Phrase, params.Arguments.Interpretation, params.Arguments.Tags, params.Arguments.Origin)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(phrase)
}

type DeletePhraseParams struct {
	PhraseId string `json:"phrase_id"`
}

// 删除例句
func (handler *MaimemoHandler) DeletePhrase(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[DeletePhraseParams]) (*mcp.CallToolResultFor[any], error) {
	phrase, err := handler.maimemoClient.DeletePhrase(params.Arguments.PhraseId)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(phrase)
}

type GetVocabularyParams struct {
	Spelling string `json:"spelling"`
}

// 获取单词
func (handler *MaimemoHandler) GetVocabulary(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[GetVocabularyParams]) (*mcp.CallToolResultFor[any], error) {
	voc, err := handler.maimemoClient.GetVocabulary(params.Arguments.Spelling)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(voc)
}
