package main

import (
	"fmt"
	"strings"

	"resty.dev/v3"
)

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
func (c *MaimemoClient) ListNotepads(ids []string, limit, offset int) ([]*Notepad, error) {
	queryParams := map[string]string{
		"limit":  fmt.Sprintf("%d", limit),
		"offset": fmt.Sprintf("%d", offset),
	}
	if len(ids) > 0 {
		queryParams["ids"] = strings.Join(ids, ",")
	}

	var resp Response[struct{ Notepads []*Notepad }]
	_, err := c.cli.R().
		SetQueryParams(queryParams).
		SetResult(&resp).
		SetError(&resp).
		Get("/notepads")

	if err := ProcessMaimemoResponeError(err, resp); err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("查询云词本列表失败: %w", err)
	}

	return resp.Data.Notepads, nil
}

// 创建云词本
func (c *MaimemoClient) CreateNotepad(status, content, title, brief string, tags []string) (*Notepad, error) {
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
		SetError(&resp).
		Post("/notepads")

	if err := ProcessMaimemoResponeError(err, resp); err != nil {
		return nil, fmt.Errorf("创建云词本失败, error: %w", err)
	}

	return resp.Data.Notepad, nil
}

// 获取云词本
func (c *MaimemoClient) GetNotepad(notepadId string) (*Notepad, error) {
	var resp Response[struct{ Notepad *Notepad }]
	_, err := c.cli.R().
		SetPathParam("notepadId", notepadId).
		SetResult(&resp).
		SetError(&resp).
		Get("/notepads/{notepadId}")

	if err := ProcessMaimemoResponeError(err, resp); err != nil {
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
		SetError(&resp).
		Post("/notepads/{id}")

	if err := ProcessMaimemoResponeError(err, resp); err != nil {
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
		SetError(&resp).
		Delete("/notepads/{id}")

	if err := ProcessMaimemoResponeError(err, resp); err != nil {
		return nil, fmt.Errorf("删除云词本失败, error: %w", err)
	}

	return resp.Data.Notepad, nil
}
