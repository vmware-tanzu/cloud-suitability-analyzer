#!/usr/bin/awk -f


# Author: Steve Woods
# Purpose: Convert file list into AF config file

#--- Initialize filename as empty string
BEGIN {
   print "{"
   FS="/"
   print "\"runName\": \".\","
   print "\"applications\": " "["
}

{
  fullPath = $0
  #--- determine last slash position in
  lastSlashPos = index(fullPath, "pom.xml")
  appDir = substr(fullPath,1,lastSlashPos-1)
  slashCount = gsub(/\//, fullPath)
  split(fullPath,dirs,"/")

  print "   {"
  print "      \"Name\": " " \"" dirs[slashCount] "\","
  print "      \"Path\": " " \"" appDir "\","
  print "      \"business-domain\": " " \"\","
  print "      \"business-value\": " "0"
  if(getline == 0)
     print "   }"
  else
      print "   },"
  next
}

END {
  print "],"
  print "   \"rule-include-tags\": " "\"\","
  print "   \"rule-exclude-tags\": " "\"\""
  print "}"
}
