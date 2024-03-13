#!/usr/local/bin/gawk -f

BEGIN {
  usage = "Usage:\ngenRuleTests.awk *.yaml <output directory>"	
  patternsFound=0
  if ( ARGC == 1 ) {
    print usage
    exit 1    
  }
  directory =  ARGV[ARGC-1]
  delete ARGV[ARGC-1]
  if (system("test -d " directory)) {
     print "Directory: " directory " does not exist"
     print usage
     exit 1		
  }
  #--- trim trailing "/" if necessary
  if ( substr(directory,length(directory)) == "/" ){
    directory = substr(directory,1, length(directory)-1)
  }
  print "Generating rule test targets to " directory
}

END {
  close( outputFile )
}
#--- gather rule data 
/^name:/ {
  outputFile = $2
  outputFile = outputFile 
  print "Proceessing rule" outputFile 
}

/^patterns:/ {
  patternsFound=1	
}

/^defaultpattern:/ {
  if (index($2, "import")) {
    prefix = "import "
  } else {
    prefix = ""
  }  
}
/^filetype:/ {
  fileExtension=$2
  searchIndex = index( fileExtension, "(cs|vb" )
  if ( searchIndex  > 0   )
    fileExtension = "cs"
  searchIndex = index( fileExtension, "xm[li]")
  if (searchIndex > 0)
	  fileExtension = "xml"
  outputFile = directory "/" outputFile "." fileExtension  
}
/^##/ {
if ( patternsFound ) {
    ruleTargetString = prefix $2 " " $3 " " $4	
    print ruleTargetString > outputFile
  }
}
/^##F/ {
  outputFile = directory "/" $2 
  print "Creating file: " outputFile	
  print "Dummy test file" > outputFile 
  close (outputFIle)
}
{
if ( currentFile == "") {	
   currentFile = FILENAME
 } else {
   if (FILENAME != currentFile) {
     print "Closing " currentFile
     close(currentFile)
     currentFile = FILENAME
   }
 }
 next	
}	
