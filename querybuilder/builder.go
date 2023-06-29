package querybuilder

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

const (
	INVALID_FILE_TYPE = "invalid file type"
)

type QueryError struct {
	QueryFile string
	Message   string
}

func (e *QueryError) Error() string {
	return " Reason: " + e.Message + " , Target: " + e.QueryFile
}

type QueryBuilder struct {
	ScriptPath string
	Query      BindEntry
}

var queryBuilder *QueryBuilder

func QueryBuilderInitialize(scriptPath string) error {
	queryBuilder = &QueryBuilder{
		ScriptPath: scriptPath,
	}
	return queryBuilder.Load()
}

func GetQueryBuilder() *QueryBuilder {
	return queryBuilder
}

func (qb *QueryBuilder) Load() error {
	qb.Query = BindEntry{}
	if fileInfo, err := os.Stat(qb.ScriptPath); err != nil {
		return err
	} else {
		if fileInfo.IsDir() {
			return qb.loadFiles()
		} else {
			return qb.loadFile()
		}
	}
}

func (qb *QueryBuilder) loadFileData(scriptFilePath string) *QueryError {
	if fBody, err := ioutil.ReadFile(scriptFilePath); err != nil {
		return &QueryError{
			QueryFile: scriptFilePath,
			Message:   err.Error(),
		}
	} else {
		holder := BindEntry{}
		if marshalErr := yaml.Unmarshal(fBody, holder); marshalErr != nil {
			return &QueryError{
				QueryFile: scriptFilePath,
				Message:   marshalErr.Error(),
			}
		} else {
			for key, entry := range holder {
				qb.Query[key] = entry
			}
		}
	}
	return nil
}

func validFileType(filename string) bool {
	isValid := strings.HasSuffix(filename, ".yml") || strings.HasSuffix(filename, ".yaml")
	return isValid
}

func (qb *QueryBuilder) loadFile() error {
	if !validFileType(qb.ScriptPath) {
		return &QueryError{
			QueryFile: qb.ScriptPath,
			Message:   INVALID_FILE_TYPE,
		}
	}
	if err := qb.loadFileData(qb.ScriptPath); err != nil {
		return err
	}
	return nil
}

func (qb *QueryBuilder) loadFiles() error {
	if files, err := ioutil.ReadDir(qb.ScriptPath); err != nil {
		return err
	} else {
		for _, fileInfo := range files {
			if fileInfo.IsDir() {
				continue
			}
			fileName := fileInfo.Name()
			if validFileType(qb.ScriptPath) {
				continue
			}
			scriptFilePath := fmt.Sprintf("%s/%s", qb.ScriptPath, fileName)
			if err := qb.loadFileData(scriptFilePath); err != nil {
				return err
			}
		}
		return nil
	}
}

func (qb *QueryBuilder) GetStatement(name string) *Statement {
	if item, exists := qb.Query[name]; exists {
		return &Statement{
			Name:        name,
			Type:        item.Type,
			Description: item.Description,
			Script:      item.Script,
			Columns:     item.Columns,
		}
	}
	return nil
}
