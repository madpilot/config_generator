package descriptor

type Descriptor struct {
	FilePath string                 `json:"filePath"`
	Fields   map[string]interface{} `json:"fields"`
}
