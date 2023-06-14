package metadata

import (
	"encoding/csv"
	"log"
	"os"
	"csa-app/model"
)

var outputDir = "export_rule_list.csv"

func Export(rules []model.Rule) {

	work_dir := os.Getenv("WORK_DIR")

	if work_dir != "" {
		outputDir = work_dir + "/export_rule_list.csv"
	}

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
		row := []string{rule.Metadata_sno, rule.Name, rule.Metadata_category, "Windows", rule.Metadata_title, rule.Metadata_group, rule.Metadata_description}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
