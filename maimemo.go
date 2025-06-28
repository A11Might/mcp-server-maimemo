package main

import (
	"fmt"
	"os"
	"strings"

	"resty.dev/v3"
)

var DefaultMaimemoClient = NewMaiMemoClient(os.Getenv("MAIMEMO_TOKEN"))

const (
	API_ENDPOINT = "https://open.maimemo.com/open/api/v1"
)

type MaimemoClient struct {
	cli *resty.Client
}

func NewMaiMemoClient(token string) *MaimemoClient {
	if token == "" {
		panic("MAIMEMO_TOKEN 不能为空")
	}

	c := resty.New()
	c.SetBaseURL(API_ENDPOINT)
	c.SetHeader("Content-Type", "application/json")
	token = strings.Replace(token, "Bearer ", "", 1)
	c.SetAuthToken(token)
	c.SetDebug(true)

	return &MaimemoClient{
		cli: c,
	}
}

func (client *MaimemoClient) Close() error {
	return client.cli.Close()
}

// 云词本
//
// 查询云词本
func (c *MaimemoClient) ListNotepads(ids []string, limit, offset int) ([]BriefNotepad, error) {
	var resp Response[struct{ Notepads []BriefNotepad }]
	_, err := c.cli.R().
		SetQueryParams(map[string]string{
			"ids":    strings.Join(ids, ","),
			"limit":  fmt.Sprintf("%d", limit),
			"offset": fmt.Sprintf("%d", offset),
		}).
		SetResult(&resp).
		Get("/notepads")

	if err != nil || !resp.Success {
		return nil, fmt.Errorf("查询云词本列表失败: %w", err)
	}

	return resp.Data.Notepads, nil
}

func FormatNotepadContent(chapterName string, words []string) string {
	content := strings.Join(words, ",")
	if chapterName != "" {
		content = fmt.Sprintf("# %s\n%s", chapterName, content)
	}
	return content
}

// 创建云词本
func (c *MaimemoClient) CreateNotepad(status, content, title, brief string, tags []string) error {
	var resp Response[struct{ Notepad *Notepad }]
	_, err := c.cli.R().
		SetBody(map[string]any{
			"notepad": map[string]any{
				"status":  status,
				"content": content,
				"title":   title,
				"brief":   brief,
				"tags":    tags,
			},
		}).
		SetResult(&resp).
		Post("/notepads")

	if err != nil || !resp.Success {
		return fmt.Errorf("创建云词本失败, error: %w", err)
	}

	return nil
}

// 获取云词本
func (c *MaimemoClient) GetNotepad(notepadId string) (*Notepad, error) {
	var resp Response[struct{ Notepad *Notepad }]
	_, err := c.cli.R().
		SetPathParam("notepadId", notepadId).
		SetResult(&resp).
		Get("/notepads/{notepadId}")

	if err != nil || !resp.Success {
		return nil, fmt.Errorf("获取云词本失败, error: %w", err)
	}

	return resp.Data.Notepad, nil
}

// 更新云词本
func (c *MaimemoClient) UpdateNotepad(notepadId, status, content, title, brief string, tags []string) (*Notepad, error) {
	var resp Response[struct{ Notepad *Notepad }]
	_, err := c.cli.R().
		SetPathParam("id", notepadId).
		SetBody(map[string]any{
			"notepad": map[string]any{
				"status":  status,
				"content": content,
				"title":   title,
				"brief":   brief,
				"tags":    tags,
			},
		}).
		SetResult(&resp).
		Post("/notepads/{id}")

	if err != nil || !resp.Success {
		return nil, fmt.Errorf("更新云词本失败, error: %w", err)
	}

	return resp.Data.Notepad, nil
}

// 删除云词本
func (c *MaimemoClient) DeleteNotepad(notepadId string) (*Notepad, error) {
	var resp Response[struct{ Notepad *Notepad }]
	_, err := c.cli.R().
		SetPathParam("id", notepadId).
		SetResult(&resp).
		Delete("/notepads/{id}")

	if err != nil || !resp.Success {
		return nil, fmt.Errorf("删除云词本失败, error: %w", err)
	}

	return resp.Data.Notepad, nil
}
