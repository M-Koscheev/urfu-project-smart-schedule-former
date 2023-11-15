package app

import (
	"strings"
)

func (app App) AddKnowledge(input string) error {
	knowledgeList := strings.Split(input, ", ")
	for _, elem := range knowledgeList {
		_, err := app.AddData("knowledge", elem, "knowledge_pk")
		if err != nil {
			return err
		}
	}

	return nil
}
