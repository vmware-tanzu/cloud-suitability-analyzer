#!/bin/bash
#
#          FILE:  text2slice.sh
#
#
#   DESCRIPTION:  Convert list of API pattern from a text file to a
#                 literal slice declaration
#
#       OPTIONS:  ---
#  REQUIREMENTS:  ---
#          BUGS:  ---
#        AUTHOR:  App Tx Team (Steve Woods)
#       COMPANY:  VMWare
#       VERSION:  1.0
#       CREATED:  2/23/18
#      REVISION:  ---
#===============================================================================

#--- exit with usage message if number of command line parameters of less than required


#--- set the usage message
message="Pivotal App Tx\nExtract unique properties\nUsage: %s  -d <directory> "


#--- example of a two param script, one value, one flag
if (( $# < 1 ))
then
   printf "${message}\n" $(basename $0) >&2
   exit 1
fi

packageFlag=0

#--- input parameters into variables
while getopts 'd:' OPTION
do
   case $OPTION in
   d)    directoryFlag=d
         directory="$OPTARG"
         ;;
   #---- exit with error if an unexpected parameter appears
   ?)    printf "$message\n" $(basename $0) >&2
         exit 2
         ;;
   esac
done

#--- find all properties, resursively
egrep -Rh --include="*.properties" "^[A-Za-z.]+=" ${directory} | \

#--- strip property values
sed 's/=.*/=/' | \

#--- count occurances
sort | uniq -c | \

#--- filter out only the unique properties (count == 1), strip the count column
awk '{ if ($1 ==  1) print $2 }' | \

#--- search for each property to restore property values
while read LINE; do egrep -Rh "^$LINE" ; done
