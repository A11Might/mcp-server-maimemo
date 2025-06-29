package main

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func SayHiTool() *mcp.ServerTool {
	return mcp.NewServerTool("greet", "say hi", SayHi, mcp.Input(
		mcp.Property("name", mcp.Description("the name of the person to greet")),
	))
}

func ListNotepadsTool() *mcp.ServerTool {
	return mcp.NewServerTool("list_notepads", "查询墨墨云词本", ListNotepad, mcp.Input(
		mcp.Property("ids", mcp.Required(false), mcp.Description("云词本 id 列表")),
		mcp.Property("limit", mcp.Required(true), mcp.Description("每页数量")),
		mcp.Property("offset", mcp.Required(true), mcp.Description("偏移量")),
	))
}

func CreateNotepadTool() *mcp.ServerTool {
	return mcp.NewServerTool("create_notepad", "创建墨墨云词本", CreateNotepad, mcp.Input(
		mcp.Property("status",
			mcp.Required(true),
			mcp.Enum("PUBLISHED", "UNPUBLISHED", "DELETED"),
			mcp.Description("状态"),
		),
		mcp.Property("chapter_name", mcp.Required(false), mcp.Description("章节名称")),
		mcp.Property("words", mcp.Required(true), mcp.Description("章节中的单词列表")),
		mcp.Property("title", mcp.Required(true), mcp.Description("标题")),
		mcp.Property("brief", mcp.Required(true), mcp.Description("摘要")),
		mcp.Property("tags", mcp.Required(true), mcp.Description("标签")),
	))
}

func GetNotepadTool() *mcp.ServerTool {
	return mcp.NewServerTool("get_notepad", "获取墨墨云词本", GetNotepad, mcp.Input(
		mcp.Property("notepad_id", mcp.Required(true), mcp.Description("云词本 id")),
	))
}

func UpdateNotepadTool() *mcp.ServerTool {
	return mcp.NewServerTool("update_notepad", "更新墨墨云词本", UpdateNotepad, mcp.Input(
		mcp.Property("notepad_id", mcp.Required(true), mcp.Description("云词本 id")),
		mcp.Property("status",
			mcp.Required(true),
			mcp.Enum("PUBLISHED", "UNPUBLISHED", "DELETED"),
			mcp.Description("状态"),
		),
		mcp.Property("chapter_name", mcp.Required(false), mcp.Description("章节名称")),
		mcp.Property("words", mcp.Required(true), mcp.Description("章节中的单词列表")),
		mcp.Property("title", mcp.Required(true), mcp.Description("标题")),
		mcp.Property("brief", mcp.Required(true), mcp.Description("摘要")),
		mcp.Property("tags", mcp.Required(true), mcp.Description("标签")),
	))
}

func DeleteNotepadTool() *mcp.ServerTool {
	return mcp.NewServerTool("delete_notepad", "删除墨墨云词本", DeleteNotepad, mcp.Input(
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

type ListNotepadParams struct {
	Ids    []string `json:"ids"`
	Limit  int      `json:"limit"`
	Offset int      `json:"offset"`
}

// 查询云词本
func ListNotepad(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[ListNotepadParams]) (*mcp.CallToolResultFor[any], error) {
	notepadList, err := DefaultMaimemoClient.ListNotepads(params.Arguments.Ids, params.Arguments.Limit, params.Arguments.Offset)
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
func CreateNotepad(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[CreateNotepadParams]) (*mcp.CallToolResultFor[any], error) {
	content := FormateNotepadContent(params.Arguments.ChapterName, params.Arguments.Words)
	notepad, err := DefaultMaimemoClient.CreateNotepad(params.Arguments.Status, content, params.Arguments.Title, params.Arguments.Brief, params.Arguments.Tags)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(notepad)
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
func UpdateNotepad(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[UpdateNotepadParams]) (*mcp.CallToolResultFor[any], error) {
	content := FormateNotepadContent(params.Arguments.ChapterName, params.Arguments.Words)
	notepad, err := DefaultMaimemoClient.UpdateNotepad(params.Arguments.NotepadId, params.Arguments.Status, content, params.Arguments.Title, params.Arguments.Brief, params.Arguments.Tags)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(notepad)
}

type DeleteNotepadParams struct {
	NotepadId string `json:"notepad_id"`
}

func DeleteNotepad(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[DeleteNotepadParams]) (*mcp.CallToolResultFor[any], error) {
	notepad, err := DefaultMaimemoClient.DeleteNotepad(params.Arguments.NotepadId)
	if err != nil {
		return nil, err
	}

	return OriginToTextContent(notepad)
}
