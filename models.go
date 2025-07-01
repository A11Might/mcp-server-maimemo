package main

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

// 释义
// https://open.maimemo.com/#/schemas/Interpretation
type Interpretation struct {
	Id             string               `json:"id"`
	Interpretation string               `json:"interpretation"`
	Tags           []InterpretationTag  `json:"tags"`
	Status         InterpretationStatus `json:"status"`
	CreatedTime    string               `json:"created_time"`
	UpdatedTime    string               `json:"updated_time"`
}

// 释义标签
type InterpretationTag string

// 简明、详细、英英、小学、初中、高中、四级、六级、专升本、专四、专八、考研、考博、雅思、托福、托业、新概念、GRE、GMAT、BEC、MBA、SAT、ACT、法学、医学
const (
	InterpretationTagSimple               InterpretationTag = "简明"
	InterpretationTagDetailed             InterpretationTag = "详细"
	InterpretationTagEnglishEnglish       InterpretationTag = "英英"
	InterpretationTagPrimarySchool        InterpretationTag = "小学"
	InterpretationTagJuniorHighSchool     InterpretationTag = "初中"
	InterpretationTagSeniorHighSchool     InterpretationTag = "高中"
	InterpretationTagFourLevelExamination InterpretationTag = "四级"
	InterpretationTagSixLevelExamination  InterpretationTag = "六级"
	InterpretationTagUndergraduateDegree  InterpretationTag = "专升本"
	InterpretationTagFourYearDegree       InterpretationTag = "专四"
	InterpretationTagEightYearDegree      InterpretationTag = "专八"
	InterpretationTagGraduateDegree       InterpretationTag = "考研"
	InterpretationTagPostgraduateDegree   InterpretationTag = "考博"
	InterpretationTagIELTS                InterpretationTag = "雅思"
	InterpretationTagTOEFL                InterpretationTag = "托福"
	InterpretationTagTOEIC                InterpretationTag = "托业"
	InterpretationTagNewConcept           InterpretationTag = "新概念"
	InterpretationTagGRE                  InterpretationTag = "GRE"
	InterpretationTagGMAT                 InterpretationTag = "GMAT"
	InterpretationTagBEC                  InterpretationTag = "BEC"
	InterpretationTagMBA                  InterpretationTag = "MBA"
	InterpretationTagSAT                  InterpretationTag = "SAT"
	InterpretationTagACT                  InterpretationTag = "ACT"
	InterpretationTagLaw                  InterpretationTag = "法学"
	InterpretationTagMedical              InterpretationTag = "医学"
)

// 释义状态
// https://open.maimemo.com/#/schemas/InterpretationStatus
type InterpretationStatus string

// PUBLISHED, UNPUBLISHED, DELETED
const (
	InterpretationStatusPublished   InterpretationStatus = "PUBLISHED"
	InterpretationStatusUnpublished InterpretationStatus = "UNPUBLISHED"
	InterpretationStatusDeleted     InterpretationStatus = "DELETED"
)

// 助记
// https://open.maimemo.com/#/schemas/Note
type Note struct {
	Id          string     `json:"id"`
	NoteType    NoteType   `json:"note_type"`
	Note        string     `json:"note"`
	Status      NoteStatus `json:"status"`
	CreatedTime string     `json:"created_time"`
	UpdatedTime string     `json:"updated_time"`
}

// 助记类型
type NoteType string

// 词根词缀、固定搭配、近反义词、派生、词源、辨析、语法、联想、谐音、串记、口诀、扩展、合成、其他
const (
	NoteTypeRootWord       NoteType = "词根词缀"
	NoteTypeFixedPhrase    NoteType = "固定搭配"
	NoteTypeNearAntonym    NoteType = "近反义词"
	NoteTypeDerivation     NoteType = "派生"
	NoteTypeWordOrigin     NoteType = "词源"
	NoteTypeDistinction    NoteType = "辨析"
	NoteTypeGrammar        NoteType = "语法"
	NoteTypeAssociation    NoteType = "联想"
	NoteTypePhonetic       NoteType = "谐音"
	NoteTypeString         NoteType = "串记"
	NoteTypeMnemonicPhrase NoteType = "口诀"
	NoteTypeExpansion      NoteType = "扩展"
	NoteTypeComposition    NoteType = "合成"
	NoteTypeOther          NoteType = "其他"
)

// 助记状态
// https://open.maimemo.com/#/schemas/NoteStatus
type NoteStatus string

const (
	NoteStatusPublished NoteStatus = "PUBLISHED"
	NoteStatusDeleted   NoteStatus = "DELETED"
)

// 云词本
// https://open.maimemo.com/#/schemas/Notepad
type Notepad struct {
	ID          string              `json:"id"`
	Type        NotepadType         `json:"type"`
	Creator     int                 `json:"creator"`
	Status      NotepadStatus       `json:"status"`
	Content     string              `json:"content,omitempty"`
	Title       string              `json:"title"`
	Brief       string              `json:"brief"`
	Tags        []NotepadTag        `json:"tags"`
	List        []NotepadParsedItem `json:"list"`
	CreatedTime string              `json:"created_time"`
	UpdatedTime string              `json:"updated_time"`
}

// 云词本类型
// https://open.maimemo.com/#/schemas/NotepadType
type NotepadType string

const (
	NotepadTypeFavorite NotepadType = "FAVORITE"
	NotepadTypeNotepad  NotepadType = "NOTEPAD"
)

// 云词本状态
// https://open.maimemo.com/#/schemas/NotepadStatus
type NotepadStatus string

const (
	NotepadStatusPublished   NotepadStatus = "PUBLISHED"
	NotepadStatusUnpublished NotepadStatus = "UNPUBLISHED"
	NotepadStatusDeleted     NotepadStatus = "DELETED"
)

// 云词本标签
type NotepadTag string

// 小学、初中、高中、大学教科书、四级、六级、专四、专八、考研、新概念、SAT、托福、雅思、GRE、GMAT、托业、BEC、词典、词频、其他
const (
	NotepadTagPrimarySchool      NotepadTag = "小学"
	NotepadTagJuniorMiddleSchool NotepadTag = "初中"
	NotepadTagSeniorHighSchool   NotepadTag = "高中"
	NotepadTagUniversityTextbook NotepadTag = "大学教科书"
	NotepadTagFourLevel          NotepadTag = "四级"
	NotepadTagSixLevel           NotepadTag = "六级"
	NotepadTagFour               NotepadTag = "专四"
	NotepadTagEight              NotepadTag = "专八"
	NotepadTagGraduate           NotepadTag = "考研"
	NotepadTagNewConcept         NotepadTag = "新概念"
	NotepadTagSAT                NotepadTag = "SAT"
	NotepadTagTOEFL              NotepadTag = "托福"
	NotepadTagIELTS              NotepadTag = "雅思"
	NotepadTagGRE                NotepadTag = "GRE"
	NotepadTagGMAT               NotepadTag = "GMAT"
	NotepadTagTOEIC              NotepadTag = "托业"
	NotepadTagBEC                NotepadTag = "BEC"
	NotepadTagDictionary         NotepadTag = "词典"
	NotepadTagPhraseFrequency    NotepadTag = "词频"
	NotepadTagOther              NotepadTag = "其他"
)

// 云词本解析结果
// https://open.maimemo.com/#/schemas/NotepadParsedItem
type NotepadParsedItem struct {
	Type    NotepadParsedItemType `json:"type"`
	Chapter string                `json:"chapter"`
	Word    string                `json:"word"`
}

// 云词本解析结果类型
type NotepadParsedItemType string

const (
	NotepadParsedItemChapter NotepadParsedItemType = "CHAPTER"
	NotepadParsedItemWord    NotepadParsedItemType = "WORD"
)

// 例句
// https://open.maimemo.com/#/schemas/Phrase
type Phrase struct {
	ID             string                 `json:"id"`
	PhraseText     string                 `json:"phrase"`
	Interpretation string                 `json:"interpretation"`
	Tags           []Phrase               `json:"tags"`
	Highlight      []PhraseHighlightRange `json:"highlight"`
	Status         PhraseStatus           `json:"status"`
	CreatedTime    string                 `json:"created_time"`
	UpdatedTime    string                 `json:"updated_time"`
	Origin         string                 `json:"origin"`
}

// 例句标签
type PhraseTag string

// 小学、初中、高中、四级、六级、专升本、专四、专八、考研、考博、雅思、托福、托业、新概念、GRE、GMAT、BEC、MBA、SAT、ACT、法学、医学、词典、短语
const (
	PhraseTagPrimarySchool       PhraseTag = "小学"
	PhraseTagJuniorMiddleSchool  PhraseTag = "初中"
	PhraseTagSeniorHighSchool    PhraseTag = "高中"
	PhraseTagFourLevel           PhraseTag = "四级"
	PhraseTagSixLevel            PhraseTag = "六级"
	PhraseTagUndergraduate       PhraseTag = "专升本"
	PhraseTagFour                PhraseTag = "专四"
	PhraseTagEight               PhraseTag = "专八"
	PhraseTagGraduate            PhraseTag = "考研"
	PhraseTagGraduateExamination PhraseTag = "考博"
	PhraseTagIELTS               PhraseTag = "雅思"
	PhraseTagTOEFL               PhraseTag = "托福"
	PhraseTagTOEIC               PhraseTag = "托业"
	PhraseTagNewConcept          PhraseTag = "新概念"
	PhraseTagGRE                 PhraseTag = "GRE"
	PhraseTagGMAT                PhraseTag = "GMAT"
	PhraseTagBEC                 PhraseTag = "BEC"
	PhraseTagMBA                 PhraseTag = "MBA"
	PhraseTagSAT                 PhraseTag = "SAT"
	PhraseTagACT                 PhraseTag = "ACT"
	PhraseTagLaw                 PhraseTag = "法学"
	PhraseTagMedical             PhraseTag = "医学"
	PhraseTagDictionary          PhraseTag = "词典"
	PhraseTagPhrase              PhraseTag = "短语"
)

// 例句中的单词高亮区间 [start, end)
// https://open.maimemo.com/#/schemas/PhraseHighlightRange
type PhraseHighlightRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// 例句状态
// https://open.maimemo.com/#/schemas/PhraseStatus
type PhraseStatus string

const (
	PhraseStatusPublished PhraseStatus = "PUBLISHED"
	PhraseStatusDeleted   PhraseStatus = "DELETED"
)

// 单词
// https://open.maimemo.com/#/schemas/Vocabulary
type Vocabulary struct {
	Id       string `json:"id"`
	Spelling string `json:"spelling"`
}
