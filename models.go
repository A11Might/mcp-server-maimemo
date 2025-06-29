package main

import "time"

type Response[T any] struct {
	Success bool    `json:"success"`
	Data    T       `json:"data,omitempty"`
	Errors  []Error `json:"errors,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
	Info    string `json:"info"`
}

/*
示例数据结构：
{
  "id": "5a7BFf4F63612e5AD9fdebB7a50D3881",
  "type": "FAVORITE",
  "creator": 192,
  "status": "PUBLISHED",
  "title": "常用词汇",
  "brief": "常用",
  "tags": ["考研"],
  "created_time": "2023-03-13T16:00:00.000Z",
  "updated_time": "2023-03-13T16:00:00.000Z"
}
*/

// Notepad 云词本数据结构
// 对应墨墨开放 API 文档：https://open.maimemo.com/#/schemas/Notepad
type Notepad struct {
	ID        string       `json:"id"`
	Type      Type         `json:"type"`                               // 词本类型，参见 NotepadType
	Creator   int          `json:"creator"`                            // 创建者 ID
	Status    Status       `json:"status"`                             // 词本状态，参见 NotepadStatus
	Content   string       `json:"content,omitempty"`                  // 词本内容（Markdown 格式）
	Title     string       `json:"title"`                              // 词本标题
	Brief     string       `json:"brief"`                              // 词本简介
	Tags      []string     `json:"tags"`                               // 标签列表
	List      []ParsedItem `json:"list"`                               // 解析后的内容项列表
	CreatedAt time.Time    `json:"created_time" time_format:"iso8601"` // ISO8601 创建时间
	UpdatedAt time.Time    `json:"updated_time" time_format:"iso8601"` // ISO8601 更新时间
}

// Type 词本类型枚举
// 文档：https://open.maimemo.com/#/schemas/NotepadType
type Type string

const (
	TypeNotepad  Type = "NOTEPAD"  // 普通词本
	TypeFavorite Type = "FAVORITE" // 收藏词本
)

// Status 词本状态枚举
// 文档：https://open.maimemo.com/#/schemas/NotepadStatus
type Status string

const (
	StatusPublished   Status = "PUBLISHED"   // 已发布
	StatusUnpublished Status = "UNPUBLISHED" // 未发布
	StatusDeleted     Status = "DELETED"     // 已删除
)

// ParsedItem 解析后的内容项
// 文档：https://open.maimemo.com/#/schemas/NotepadParsedItem
type ParsedItem struct {
	Type    ItemType `json:"type"`    // 内容项类型
	Chapter string   `json:"chapter"` // 章节
	Word    string   `json:"word"`    // 单词 当 type=WORD 时，该字段才有值
}

// ItemType 内容项类型枚举
type ItemType string

const (
	ItemChapter ItemType = "CHAPTER" // 章节标题
	ItemWord    ItemType = "WORD"    // 单词项
)

/*
示例数据结构：
{
  "id": "5a7BFf4F63612e5AD9fdebB7a50D3881",
  "type": "FAVORITE",
  "creator": 192,
  "status": "PUBLISHED",
  "content": "apple",
  "title": "常用词汇",
  "brief": "常用",
  "tags": ["考研"],
  "list": [{
    "type": "CHAPTER",
    "data": {}
  }],
  "created_time": "2023-03-13T16:00:00.000Z",
  "updated_time": "2023-03-13T16:00:00.000Z"
}
*/
