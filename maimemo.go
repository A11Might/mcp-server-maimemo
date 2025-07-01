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

// 释义
//
// 获取释义
func (c *MaimemoClient) ListInterpretations(vocId string) ([]*Interpretation, error) {
	var resp Response[struct{ Interpretations []*Interpretation }]
	_, err := c.cli.R().
		SetQueryParam("voc_id", vocId).
		SetResult(&resp).
		SetError(&resp).
		Get("/interpretations")

	if err := ProcessMaimemoResponeError(err, resp); err != nil {
		return nil, fmt.Errorf("获取释义列表失败, error: %w", err)
	}

	return resp.Data.Interpretations, nil
}

// 创建释义
func (c *MaimemoClient) CreateInterpretation(vocId, interpretation string, tags []InterpretationTag, status InterpretationStatus) (*Interpretation, error) {
	var resp Response[struct{ Interpretation *Interpretation }]
	_, err := c.cli.R().
		SetBody(map[string]any{
			"interpretation": map[string]any{
				"voc_id":         vocId,
				"interpretation": interpretation,
				"tags":           tags,
				"status":         status,
			},
		}).
		SetResult(&resp).
		SetError(&resp).
		Post("/interpretations")

	if err := ProcessMaimemoResponeError(err, resp); err != nil {
		return nil, fmt.Errorf("创建释义失败, error: %w", err)
	}

	return resp.Data.Interpretation, nil
}

// 更新释义
func (c *MaimemoClient) UpdateInterpretation(interpretationId, interpretation string, tags []InterpretationTag, status InterpretationStatus) (*Interpretation, error) {
	var resp Response[struct{ Interpretation *Interpretation }]
	_, err := c.cli.R().
		SetPathParam("id", interpretationId).
		SetBody(map[string]any{
			"interpretation": map[string]any{
				"interpretation": interpretation,
				"tags":           tags,
				"status":         status,
			},
		}).
		SetResult(&resp).
		SetError(&resp).
		Put("/interpretations/{id}")

	if err := ProcessMaimemoResponeError(err, resp); err != nil {
		return nil, fmt.Errorf("更新释义失败, error: %w", err)
	}

	return resp.Data.Interpretation, nil
}

// 删除释义
func (c *MaimemoClient) DeleteInterpretation(interpretationId string) (*Interpretation, error) {
	var resp Response[struct{ Interpretation *Interpretation }]
	_, err := c.cli.R().
		SetPathParam("id", interpretationId).
		SetResult(&resp).
		SetError(&resp).
		Delete("/interpretations/{id}")

	if err := ProcessMaimemoResponeError(err, resp); err != nil {
		return nil, fmt.Errorf("删除释义失败, error: %w", err)
	}

	return resp.Data.Interpretation, nil
}

// 助记
//
// 获取助记
func (c *MaimemoClient) ListNotes(vocId string) ([]*Note, error) {
	var resp Response[struct{ Notes []*Note }]
	_, err := c.cli.R().
		SetQueryParam("voc_id", vocId).
		SetResult(&resp).
		SetError(&resp).
		Get("/notes")

	if err := ProcessMaimemoResponeError(err, resp); err != nil {
		return nil, fmt.Errorf("获取助记列表失败, error: %w", err)
	}

	return resp.Data.Notes, nil
}

// 创建助记
func (c *MaimemoClient) CreateNote(vocId string, noteType NoteType, note string) (*Note, error) {
	var resp Response[struct{ Note *Note }]
	_, err := c.cli.R().
		SetBody(map[string]any{
			"note": map[string]any{
				"voc_id":    vocId,
				"note_type": noteType,
				"note":      note,
			},
		}).
		SetResult(&resp).
		SetError(&resp).
		Post("/notes")

	if err := ProcessMaimemoResponeError(err, resp); err != nil {
		return nil, fmt.Errorf("创建助记失败, error: %w", err)
	}

	return resp.Data.Note, nil
}

// 更新助记
func (c *MaimemoClient) UpdateNote(noteId string, noteType NoteType, note string) (*Note, error) {
	var resp Response[struct{ Note *Note }]
	_, err := c.cli.R().
		SetPathParam("id", noteId).
		SetBody(map[string]any{
			"note": map[string]any{
				"note_type": noteType,
				"note":      note,
			},
		}).
		SetResult(&resp).
		SetError(&resp).
		Put("/notes/{id}")

	if err := ProcessMaimemoResponeError(err, resp); err != nil {
		return nil, fmt.Errorf("更新助记失败, error: %w", err)
	}

	return resp.Data.Note, nil
}

// 删除助记
func (c *MaimemoClient) DeleteNote(noteId string) (*Note, error) {
	var resp Response[struct{ Note *Note }]
	_, err := c.cli.R().
		SetPathParam("id", noteId).
		SetResult(&resp).
		SetError(&resp).
		Delete("/notes/{id}")

	if err := ProcessMaimemoResponeError(err, resp); err != nil {
		return nil, fmt.Errorf("删除助记失败, error: %w", err)
	}

	return resp.Data.Note, nil
}

// 云词本
//
// 查询云词本
func (c *MaimemoClient) ListNotepads(notepadIds []string, limit, offset int) ([]*Notepad, error) {
	queryParams := map[string]string{
		"limit":  fmt.Sprintf("%d", limit),
		"offset": fmt.Sprintf("%d", offset),
	}
	if len(notepadIds) > 0 {
		queryParams["ids"] = strings.Join(notepadIds, ",")
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
func (c *MaimemoClient) CreateNotepad(status NotepadStatus, content, title, brief string, tags []NotepadTag) (*Notepad, error) {
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
		SetPathParam("id", notepadId).
		SetResult(&resp).
		SetError(&resp).
		Get("/notepads/{id}")

	if err := ProcessMaimemoResponeError(err, resp); err != nil {
		return nil, fmt.Errorf("获取云词本失败, error: %w", err)
	}

	return resp.Data.Notepad, nil
}

// 更新云词本
func (c *MaimemoClient) UpdateNotepad(notepadId string, status NotepadStatus, content, title, brief string, tags []NotepadTag) (*Notepad, error) {
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

// 例句
//
// 获取例句
func (c *MaimemoClient) ListPhrases(vocId string) ([]*Phrase, error) {
	var resp Response[struct{ Phrases []*Phrase }]
	_, err := c.cli.R().
		SetQueryParam("voc_id", vocId).
		SetResult(&resp).
		SetError(&resp).
		Get("/phrases")

	if err := ProcessMaimemoResponeError(err, resp); err != nil {
		return nil, fmt.Errorf("获取例句列表失败, error: %w", err)
	}

	return resp.Data.Phrases, nil
}

// 创建例句
func (c *MaimemoClient) CreatePhrase(vocId, phrase, interpretation string, tags []PhraseTag, origin string) (*Phrase, error) {
	var resp Response[struct{ Phrase *Phrase }]
	_, err := c.cli.R().
		SetBody(map[string]any{
			"phrase": map[string]any{
				"voc_id":         vocId,
				"phrase":         phrase,
				"interpretation": interpretation,
				"tags":           tags,
				"origin":         origin,
			},
		}).
		SetResult(&resp).
		SetError(&resp).
		Post("/phrases")

	if err := ProcessMaimemoResponeError(err, resp); err != nil {
		return nil, fmt.Errorf("创建例句失败, error: %w", err)
	}

	return resp.Data.Phrase, nil
}

// 更新例句
func (c *MaimemoClient) UpdatePhrase(phraseId, phrase, interpretation string, tags []PhraseTag, origin string) (*Phrase, error) {
	var resp Response[struct{ Phrase *Phrase }]
	_, err := c.cli.R().
		SetPathParam("id", phraseId).
		SetBody(map[string]any{
			"phrase": map[string]any{
				"phrase":         phrase,
				"interpretation": interpretation,
				"tags":           tags,
				"origin":         origin,
			},
		}).
		SetResult(&resp).
		SetError(&resp).
		Post("/phrases/{id}")

	if err := ProcessMaimemoResponeError(err, resp); err != nil {
		return nil, fmt.Errorf("更新例句失败, error: %w", err)
	}

	return resp.Data.Phrase, nil
}

// 删除例句
func (c *MaimemoClient) DeletePhrase(phraseId string) (*Phrase, error) {
	var resp Response[struct{ Phrase *Phrase }]
	_, err := c.cli.R().
		SetPathParam("id", phraseId).
		SetResult(&resp).
		SetError(&resp).
		Delete("/phrases/{id}")

	if err := ProcessMaimemoResponeError(err, resp); err != nil {
		return nil, fmt.Errorf("删除例句失败, error: %w", err)
	}

	return resp.Data.Phrase, nil
}

// 单词
//
// 查询单词
func (c *MaimemoClient) GetVocabulary(spelling string) (*Vocabulary, error) {
	var resp Response[struct{ Voc *Vocabulary }]
	_, err := c.cli.R().
		SetQueryParam("spelling", spelling).
		SetResult(&resp).
		SetError(&resp).
		Get("/vocabulary")

	if err := ProcessMaimemoResponeError(err, resp); err != nil {
		return nil, fmt.Errorf("查询单词失败, error: %w", err)
	}

	return resp.Data.Voc, nil
}
