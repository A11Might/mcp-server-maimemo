package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func FormateNotepadContent(chapterName string, words []string) string {
	content := strings.Join(words, ",")
	if chapterName != "" {
		content = fmt.Sprintf("# %s\n%s", chapterName, content)
	}

	return content
}

func ProcessMaimemoResponeError[T any](err error, respone Response[T]) error {
	if err != nil {
		return err
	}
	if !respone.Success {
		return fmt.Errorf("请求墨墨开放 API 失败, error: %v", respone.Errors)
	}

	return nil
}

func OriginToTextContent(origin any) (*mcp.CallToolResultFor[any], error) {
	text, ok := origin.(string)
	if ok {
		return &mcp.CallToolResultFor[any]{
			Content: []mcp.Content{&mcp.TextContent{Text: string(text)}},
		}, nil
	}

	bytes, err := json.Marshal(origin)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: string(bytes)}},
	}, nil
}
