# Author: Steve Woods
# Syntax: gawk -f removeCommit.awk 
BEGIN {
commitCount = 0
printLine = 0 
}

#--- If the line begins with VMEMBER
/^commit/ {
  commit++
  if (commit > limit ) {
     printLine = 1
  }
  next
}

{
  if ( printLine ) {
     print $0
  }
  next
}
