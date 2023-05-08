package mat

import (
	"encoding/csv"
	"log"
	"os"
	"test/csa-app/model"
)

var outputDir = "/Users/scarbonell/Workspace/boa-csa/cloud-suitability-analyzer/csa-test/mat_export.csv"

func Export(rules []model.Rule) {

	file, err := os.Create(outputDir)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()
	// Using Write
	w.Write([]string{"SNO", "Rule ID", "Category", "Platform", "Title", "Rule Group", "Rule Description(From CSA)"})
	for _, rule := range rules {
		row := []string{rule.Mat_sno, rule.Name, rule.Mat_category, "Windows", rule.Mat_title, rule.Mat_group, rule.Mat_description}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
