package main

import (
	"os"

	"github.com/modelcontextprotocol/go-sdk/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/samber/lo"
)

func NewMaimemoServer() (*mcp.Server, error) {

	client := NewMaiMemoClient(os.Getenv("MAIMEMO_TOKEN"))
	h, err := NewMaimemoHanlder(client)
	if err != nil {
		return nil, err
	}

	s := mcp.NewServer("maimemo", "v1.0.0", nil)

	s.AddTools(
		// Interpretations
		mcp.NewServerTool(
			"list_interpretations",
			"获取指定单词下所有自己创建的释义列表。",
			h.ListInterpretations,
			mcp.Input(
				mcp.Property("voc_id",
					mcp.Required(true),
					mcp.Description("单词的唯一ID，可通过`get_vocabulary`获取"),
				),
			),
		),
		mcp.NewServerTool(
			"create_interpretation",
			"为指定单词创建一个新的释义。成功后会返回新创建的释义对象。",
			h.CreateInterpretation,
			mcp.Input(
				mcp.Property("voc_id",
					mcp.Required(true),
					mcp.Description("要添加释义的单词的唯一ID"),
				),
				mcp.Property("interpretation",
					mcp.Required(true),
					mcp.Description("释义的具体内容"),
				),
				mcp.Property("tags",
					mcp.Required(true),
					mcp.Enum("简明", "详细", "英英", "小学", "初中", "高中", "四级", "六级", "专升本", "专四", "专八", "考研", "考博", "雅思", "托福", "托业", "新概念", "GRE", "GMAT", "BEC", "MBA", "SAT", "ACT", "法学", "医学"),
					mcp.Description("与此释义关联的标签数组"),
				),
				mcp.Property("status",
					mcp.Required(true),
					mcp.Enum("PUBLISHED", "UNPUBLISHED", "DELETED"),
					mcp.Description("释义的状态"),
				),
			),
		),
		mcp.NewServerTool(
			"update_interpretation",
			"全量更新一个已有的释义。重要：由于API限制，此操作流程固定：1. 调用`get_vocabulary`获取`voc_id`。2. 调用`list_interpretations`获取所有释义。3. 在本地修改后调用本工具。如果用户只提供interpretation_id，你必须向用户询问该释义所属的单词。",
			h.UpdateInterpretation,
			mcp.Input(
				mcp.Property("interpretation_id",
					mcp.Required(true),
					mcp.Description("要更新的释义的唯一ID"),
				),
				mcp.Property("interpretation",
					mcp.Required(true),
					mcp.Description("新的释义具体内容"),
				),
				mcp.Property("tags",
					mcp.Required(true),
					mcp.Description("更新后的完整标签数组"),
				),
				mcp.Property("status",
					mcp.Required(true),
					mcp.Enum("PUBLISHED", "UNPUBLISHED", "DELETED"),
					mcp.Description("释义的状态"),
				),
			),
		),
		mcp.NewServerTool(
			"delete_interpretation",
			"删除一个指定的释义。注意：这是一个不可恢复的危险操作，请谨慎使用。",
			h.DeleteInterpretation,
			mcp.Input(
				mcp.Property("interpretation_id",
					mcp.Required(true),
					mcp.Description("要删除的释义的唯一ID"),
				),
			),
		),

		// Notes
		mcp.NewServerTool(
			"list_notes",
			"获取指定单词的所有助记列表。",
			h.ListNotes,
			mcp.Input(
				mcp.Property("voc_id",
					mcp.Required(true),
					mcp.Description("单词的唯一ID"),
				),
			),
		),
		mcp.NewServerTool(
			"create_note",
			"为指定单词创建一个新的助记。成功后会返回新创建的助记对象。",
			h.CreateNote,
			mcp.Input(
				mcp.Property("voc_id",
					mcp.Required(true),
					mcp.Description("要添加助记的单词的唯一ID"),
				),
				mcp.Property("note_type",
					mcp.Required(true),
					mcp.Enum("词根词缀", "固定搭配", "近反义词", "派生", "词源", "辨析", "语法", "联想", "谐音", "串记", "口诀", "扩展", "合成", "其他"),
					mcp.Description("助记的类型"),
				),
				mcp.Property("note",
					mcp.Required(true),
					mcp.Description("助记的内容"),
				),
			),
		),
		mcp.NewServerTool(
			"update_note",
			"全量更新一个已有的助记。重要：操作流程类似于更新释义，需要先获取再更新。如果只提供note_id，必须向用户询问其所属单词。",
			h.UpdateNote,
			mcp.Input(
				mcp.Property("note_id",
					mcp.Required(true),
					mcp.Description("要更新的助记的唯一ID"),
				),
				mcp.Property("note_type",
					mcp.Required(true),
					mcp.Enum("词根词缀", "固定搭配", "近反义词", "派生", "词源", "辨析", "语法", "联想", "谐音", "串记", "口诀", "扩展", "合成", "其他"),
					mcp.Description("助记的类型"),
				),
				mcp.Property("note",
					mcp.Required(true),
					mcp.Description("助记的内容"),
				),
			),
		),
		mcp.NewServerTool(
			"delete_note",
			"删除一个指定的助记。注意：这是一个不可恢复的危险操作，请谨慎使用。",
			h.DeleteNote,
			mcp.Input(
				mcp.Property("note_id",
					mcp.Required(true),
					mcp.Description("要删除的助记的唯一ID"),
				),
			),
		),

		// Notepads
		mcp.NewServerTool(
			"list_notepads",
			"查询云词本列表，支持分页。",
			h.ListNotepad,
			mcp.Input(
				mcp.Property("ids",
					mcp.Required(false),
					mcp.Description("要查询的词本ID列表，如果为空则查询所有"),
				),
				mcp.Property("limit",
					mcp.Required(true),
					mcp.Schema(&jsonschema.Schema{
						Description: "单次查询返回的最大数量（例如：10）",
						Type:        "integer",
						Minimum:     lo.ToPtr(float64(1)),
						Maximum:     lo.ToPtr(float64(10)),
					}),
				),
				mcp.Property("offset",
					mcp.Required(true),
					mcp.Schema(&jsonschema.Schema{
						Description: "查询结果的起始偏移量（例如：0）",
						Type:        "integer",
						Minimum:     lo.ToPtr(float64(0)),
					}),
				),
			),
		),
		mcp.NewServerTool(
			"create_notepad",
			"创建一个新的云词本。成功后会返回新创建的云词本对象。",
			h.CreateNotepad,
			mcp.Input(
				mcp.Property("status",
					mcp.Required(true),
					mcp.Enum("PUBLISHED", "UNPUBLISHED", "DELETED"),
					mcp.Description("词本状态"),
				),
				mcp.Property("chapter_name",
					mcp.Required(false),
					mcp.Description("初始章节的名称"),
				),
				mcp.Property("words",
					mcp.Required(true),
					mcp.Description("初始章节中的单词列表"),
				),
				mcp.Property("title",
					mcp.Required(true),
					mcp.Description("词本的标题"),
				),
				mcp.Property("brief",
					mcp.Required(true),
					mcp.Description("词本的摘要或简介"),
				),
				mcp.Property("tags",
					mcp.Required(true),
					mcp.Enum("小学", "初中", "高中", "大学教科书", "四级", "六级", "专四", "专八", "考研", "新概念", "SAT", "托福", "雅思", "GRE", "GMAT", "托业", "BEC", "词典", "词频", "其他"),
					mcp.Description("与词本关联的标签数组"),
				),
			)),
		mcp.NewServerTool(
			"get_notepad",
			"获取单个云词本的完整信息。返回的结果可用于'update_notepad'工具。",
			h.GetNotepad,
			mcp.Input(
				mcp.Property("notepad_id",
					mcp.Required(true),
					mcp.Description("要获取的云词本的唯一ID"),
				),
			)),
		mcp.NewServerTool(
			"update_notepad",
			"全量更新一个已有的云词本。注意：这是一个全量替换操作，必须提供云词本的所有字段。推荐的操作流程是：1. 先使用 `get_notepad` 获取云词本的当前完整信息。 2. 在获取到的信息基础上修改（例如增删单词）。 3. 使用修改后的完整对象调用本工具。",
			h.UpdateNotepad,
			mcp.Input(
				mcp.Property("notepad_id",
					mcp.Required(true),
					mcp.Description("要更新的云词本的唯一ID"),
				),
				mcp.Property("status",
					mcp.Required(true),
					mcp.Enum("PUBLISHED", "UNPUBLISHED", "DELETED"),
					mcp.Description("更新后的词本状态"),
				),
				mcp.Property("chapter_name",
					mcp.Required(false),
					mcp.Description("更新后的章节名称"),
				),
				mcp.Property("words",
					mcp.Required(true),
					mcp.Description("更新后的完整单词列表"),
				),
				mcp.Property("title",
					mcp.Required(true),
					mcp.Description("更新后的标题"),
				),
				mcp.Property("brief",
					mcp.Required(true),
					mcp.Description("更新后的摘要"),
				),
				mcp.Property("tags",
					mcp.Required(true),
					mcp.Enum("小学", "初中", "高中", "大学教科书", "四级", "六级", "专四", "专八", "考研", "新概念", "SAT", "托福", "雅思", "GRE", "GMAT", "托业", "BEC", "词典", "词频", "其他"),
					mcp.Description("更新后的完整标签数组"),
				),
			)),
		mcp.NewServerTool(
			"delete_notepad",
			"删除一个指定的云词本。注意：这是一个不可恢复的危险操作，请谨慎使用。",
			h.DeleteNotepad,
			mcp.Input(
				mcp.Property("notepad_id",
					mcp.Required(true),
					mcp.Description("要删除的云词本的唯一ID"),
				),
			)),

		// Phrases
		mcp.NewServerTool(
			"list_phrases",
			"获取指定单词的所有例句列表。",
			h.ListPhrases,
			mcp.Input(
				mcp.Property("voc_id",
					mcp.Required(true),
					mcp.Description("单词的唯一ID"),
				),
			),
		),
		mcp.NewServerTool(
			"create_phrase",
			"为指定单词创建一个新的例句。成功后会返回新创建的例句对象。",
			h.CreatePhrase,
			mcp.Input(
				mcp.Property("voc_id",
					mcp.Required(true),
					mcp.Description("要添加例句的单词的唯一ID"),
				),
				mcp.Property("phrase",
					mcp.Required(true),
					mcp.Description("例句的原文"),
				),
				mcp.Property("interpretation",
					mcp.Required(true),
					mcp.Description("例句的翻译"),
				),
				mcp.Property("tags",
					mcp.Required(true),
					mcp.Enum("小学", "初中", "高中", "四级", "六级", "专升本", "专四", "专八", "考研", "考博", "雅思", "托福", "托业", "新概念", "GRE", "GMAT", "BEC", "MBA", "SAT", "ACT", "法学", "医学", "词典", "短语"),
					mcp.Description("与例句关联的标签数组"),
				),
				mcp.Property("origin",
					mcp.Required(false),
					mcp.Description("例句的来源"),
				),
			),
		),
		mcp.NewServerTool(
			"update_phrase",
			"全量更新一个已有的例句。重要：操作流程类似于更新释义，需要先获取再更新。如果只提供phrase_id，必须向用户询问其所属单词。",
			h.UpdatePhrase,
			mcp.Input(
				mcp.Property("phrase_id",
					mcp.Required(true),
					mcp.Description("要更新的例句的唯一ID"),
				),
				mcp.Property("phrase",
					mcp.Required(true),
					mcp.Description("更新后的例句原文"),
				),
				mcp.Property("interpretation",
					mcp.Required(true),
					mcp.Description("更新后的例句翻译"),
				),
				mcp.Property("tags",
					mcp.Required(true),
					mcp.Description("更新后的完整标签数组"),
				),
				mcp.Property("origin",
					mcp.Required(false),
					mcp.Description("更新后的来源"),
				),
			),
		),
		mcp.NewServerTool(
			"delete_phrase",
			"删除一个指定的例句。注意：这是一个不可恢复的危险操作，请谨慎使用。",
			h.DeletePhrase,
			mcp.Input(
				mcp.Property("phrase_id",
					mcp.Required(true),
					mcp.Description("要删除的例句的唯一ID"),
				),
			),
		),

		// Vocabularies
		mcp.NewServerTool(
			"get_vocabulary",
			"通过拼写获取一个单词的核心信息，主要是它的唯一ID(`voc_id`)。这是执行其他需要`voc_id`操作（如`list_interpretations`）的前置步骤。",
			h.GetVocabulary,
			mcp.Input(
				mcp.Property("spelling",
					mcp.Required(true),
					mcp.Description("要查询的单词的拼写"),
				),
			),
		),
	)

	return s, nil
}
