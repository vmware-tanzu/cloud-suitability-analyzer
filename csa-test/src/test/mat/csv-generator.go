package matt

import (
	"encoding/csv"
	"log"
	"os"
	"test/csa-app/model"
)

var outputDir = "/Users/scarbonell/Workspace/bankofamerica/csa-test/mat_export.csv"

func Export(rules []model.Rule) {

	file, err := os.Create(outputDir)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()
	// Using Write
	w.Write([]string{"SNO", "Rule ID", "Category", "Title", "Rule Group", "Rule Description"})
	for _, rule := range rules {
		row := []string{rule.Mat_sno, rule.Name, rule.Mat_category, rule.Mat_title, rule.Mat_group, rule.Mat_description}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
