package picnic

type TemplatedContent struct {
	Type        string   `json:"type"`
	TemplateId  string   `json:"template_id"`
	VersionId   string   `json:"version_id"`
	VersionName string   `json:"version_name"`
	Content     string   `json:"content"`
	Parameters  []string `json:"parameters"`
	Actions     []string `json:"actions"`
}
