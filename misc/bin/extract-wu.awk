#!/usr/bin/awk -f
# Author: Steve Woods
BEGIN {
  printLine = 0 
}

#--- If the line begins with VMEMBER
/^ *<rule id/{
  printLine = 1
  print $0  
  next
}

/^ *<\/rule>/{
  printLine = 0
  print $0
  next
}

{
if ( printLine ) {
   print $0
}
  next
}
