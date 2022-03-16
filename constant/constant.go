package constant

// Auth
const (
	Salt          = "lpms Secret"
	Authorization = "Bearer"
)

// pagination key
const (
	Page       = "page"
	PageSize   = "page_size"
	TextSearch = "keywords"
)

// http request
const (
	ID               = "id"
	IDS              = "ids"
	Name             = "name"
	Level            = "level"
	ProjectType      = "project_type"
	ConstructSubject = "construct_subject"
	BeginAt          = "begin_at"
	EndAt            = "end_at"
	Status           = "status"
)

// project status
const (
	// 草稿
	Draft = 0
	// 已入库
	EnteredDB = 1
	// 前期计划
	EarlyPlan = 2
	// 已发文
	Posted = 3
	// 已入政府项目实施库
	Implementation = 4
)
