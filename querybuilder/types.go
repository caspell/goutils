package querybuilder

type BindEntry map[string]Entry

type Entry struct {
	Description string   `yaml:"description"`
	Type        string   `yaml:"type"`
	Columns     []string `yaml:"columns"`
	Script      string   `yaml:"script"`
}

type Statement struct {
	Name        string
	Type        string
	Description string
	Script      string
	Columns     []string
	Parameter   map[string]interface{}
	Parameters  []map[string]interface{}
}

type QueryExecute interface {
	Connect()
	Query()
}
