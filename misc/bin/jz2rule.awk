#!/usr/bin/awk -f


# Author: Steve Woods
# Purpose: Convert csv file to rule yaml file

#--- Initialize filename as empty string
BEGIN {
   filename = ""
   yamlFile = ""
   print "VMWare, copyright 2018. All rights reserved"
   FS=","
}

NR>0 {
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
    filetype = "xml" 
    target = "line" 
    type = "contains" 
    category = "version" 
    advice = "" 
    uri = ""
    effort = $4
    readiness = $5
    tags = "" 
    pattern = $3
    print "rulename: " rulename > yamlFile
    print "filetype: " filetype > yamlFile
    print "target: " target > yamlFile
    print "type: " type > yamlFile
    print "category: " category > yamlFile
    print "advice: " advice > yamlFile
    print "effort: " effort > yamlFile
    print "readiness: " readiness > yamlFile
    print "tags: " tags > yamlFile
    print "patterns: "  > yamlFile
    print "- value: " pattern > yamlFile
  } else {
    #--- capture and print the pattern value
    pattern = $1
    print "- value: " pattern > yamlFile
  }
  next
}
