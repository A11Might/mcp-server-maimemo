package main

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type MaimemoHandler struct {
	maimemoClient *MaimemoClient
}

func NewMaimemoHanlder(token string) (*MaimemoHandler, error) {
	client := NewMaiMemoClient(token)

	return &MaimemoHandler{
		maimemoClient: client,
	}, nil
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
	Status      string   `json:"status"`
	ChapterName string   `json:"chapter_name"`
	Words       []string `json:"words"`
	Title       string   `json:"title"`
	Brief       string   `json:"brief"`
	Tags        []string `json:"tags"`
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
	NotepadId   string   `json:"notepad_id"`
	Status      string   `json:"status"`
	ChapterName string   `json:"chapter_name"`
	Words       []string `json:"words"`
	Title       string   `json:"title"`
	Brief       string   `json:"brief"`
	Tags        []string `json:"tags"`
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
