package main

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func SayHiTool() *mcp.ServerTool {
	return mcp.NewServerTool("greet", "say hi", SayHi, mcp.Input(
		mcp.Property("name", mcp.Description("the name of the person to greet")),
	))
}

func GetNotepadTool() *mcp.ServerTool {
	return mcp.NewServerTool("get_notepad", "根据云词本 id 获取指定的墨墨云词本", GetNotepad, mcp.Input(
		mcp.Property("notepad_id", mcp.Required(true), mcp.Description("云词本 id")),
	))
}

type HiParams struct {
	Name string `json:"name"`
}

func SayHi(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[HiParams]) (*mcp.CallToolResultFor[any], error) {
	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: "Hi " + params.Arguments.Name}},
	}, nil
}

type GetNotepadParams struct {
	NotepadId string `json:"notepad_id"`
}

// 获取云词本
func GetNotepad(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[GetNotepadParams]) (*mcp.CallToolResultFor[any], error) {
	notepad, err := DefaultMaimemoClient.GetNotepad(params.Arguments.NotepadId)
	if err != nil {
		return nil, err
	}

	text, err := json.Marshal(notepad)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: string(text)}},
	}, nil
}
