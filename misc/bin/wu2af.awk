#!/usr/bin/awk -f


# Author: Steve Woods
# Purpose: Convert windup file to rule yaml file

#--- Initialize filename as empty string
BEGIN {
   filename = ""
   yamlFile = ""
   FS="`"
   print "name: jar-libs-frameworks"
   print "target: file"  
   print "filetype: jar"
   print "type: regex" 
   print "category: lib-framework"
   print "defaultpattern: \".*%s.*\\.jar\""
   print "tags:"
   print "- value: libraries"
   print "- value: frameworks"
   print "patterns:"
}

NR>1 {
  if ( FILENAME == "" ) {
    print "Usage:\nwu2af.awk <CSV file>"
    exit
  }
  if ( index($2,".jar") ){ 
    gsub(/{/,"",$2)
    gsub(/}/,"",$2)
    gsub(/\.jar/,"",$2)
    gsub(/\*/,"",$2)
    print "- value: " $2
    print "  advice: " $3
    print "  effort: " $4
    if ( $5 == "optional" ) 
       numeric = "0"
    else
       numeric = "10"	    
    print "  readiness: " numeric 
  }	     
  next
}
