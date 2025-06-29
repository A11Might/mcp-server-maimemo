package main

import (
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func NewMaimemoServer() (*mcp.Server, error) {

	h, err := NewMaimemoHanlder(os.Getenv("MAIMEMO_TOKEN"))
	if err != nil {
		return nil, err
	}

	s := mcp.NewServer("maimemo", "v1.0.0", nil)

	s.AddTools(
		mcp.NewServerTool("list_notepads", "查询墨墨云词本", h.ListNotepad, mcp.Input(
			mcp.Property("ids", mcp.Required(false), mcp.Description("云词本 id 列表")),
			mcp.Property("limit", mcp.Required(true), mcp.Description("每页数量（最大值：10）")),
			mcp.Property("offset", mcp.Required(true), mcp.Description("偏移量")),
		)),

		mcp.NewServerTool("create_notepad", "创建墨墨云词本", h.CreateNotepad, mcp.Input(
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
		)),

		mcp.NewServerTool("get_notepad", "获取墨墨云词本", h.GetNotepad, mcp.Input(
			mcp.Property("notepad_id", mcp.Required(true), mcp.Description("云词本 id")),
		)),

		mcp.NewServerTool("update_notepad", "更新墨墨云词本", h.UpdateNotepad, mcp.Input(
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
		)),

		mcp.NewServerTool("delete_notepad", "删除墨墨云词本", h.DeleteNotepad, mcp.Input(
			mcp.Property("notepad_id", mcp.Required(true), mcp.Description("云词本 id")),
		)),
	)

	return s, nil
}
