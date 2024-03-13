#!/usr/bin/awk -f

# Purpose: Convert csv file to rule yaml file

#--- Initialize filename as empty string
BEGIN {
   filename = ""
   yamlFile = ""
   print "VMware, copyright 2018. All rights reserved"
   FS=","
}

NR>1 {
  if ( FILENAME == "" ) {
    print "Usage:\ncsv2rule.awk <CSV file>"
    exit
  }

  #--- if rulename if 'F:' prefix, this is full rule
  #    if not, then it is not a full rule, just a pattern value
  if ( substr($1,1,2) == "F:" ) {
    #--- after the first time through, close the current yaml file
    #    before creating a new yamlFile
    if (yamlFile != "") close(yamlFile)
    yamlFile = substr($1,3)  ".yaml"
    print "Creating " yamlFile
    printFullRule = 1
  } else {
    printFullRule = 0
  }

  #--- put column values into holding variables
  if (printFullRule) {
    rulename = substr($1,3)
    filetype = $2
    target = $3
    type = $4
    defaultpattern = $5
    category = $6
    advice = $7
    if ( $8 != "" )
      uri = "\"" $8 "\""
    else
      uri = ""
  
    effort = $9
    criticality = $10
    tags = $11
    pattern = $12
    gsub(/"/,"",pattern)
    gsub(/"/,"",tags)
    print "rulename: " rulename > yamlFile
    print "filetype: " filetype > yamlFile
    print "target: " target > yamlFile
    print "type: " type > yamlFile
    print "defaultpattern: " defaultpattern > yamlFile
    print "category: " category > yamlFile
    print "advice: " advice > yamlFile
    if ( uri != "") {
      print "recipes: " > yamlFile
      print "- uri: " uri > yamlFile
    }
    print "effort: " effort > yamlFile
    print "readiness: 0" > yamlFile
    print "criticality: " criticality > yamlFile
    print "tags: " tags > yamlFile
    print "patterns: "  > yamlFile
    print "- value: " pattern > yamlFile
  } else {
    #--- capture and print the pattern value
    pattern = $1
    gsub(/"/,"",pattern)
    print "- value: " pattern > yamlFile
  }
  next
}
