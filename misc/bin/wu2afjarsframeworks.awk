#!/usr/bin/awk -f


# Author: Steve Woods
# Purpose: Convert windup file to rule yaml file

#--- Initialize filename as empty string
BEGIN {
   filename = ""
   yamlFile = ""
   print "VMWare, copyright 2018. All rights reserved"
   FS="`"
   print "name: jar-libs-frameworks"  
   print "filetype: jar"
   print "type: regex" 
}

NR>1 {
  if ( FILENAME == "" ) {
    print "Usage:\nwu2af.awk <CSV file>"
    exit
  }
  if ( index($2,".jar") ){ 
    gsub(/{/,"",$2)
     gsub(/}/,"",$2)
  }	     
  next
}
