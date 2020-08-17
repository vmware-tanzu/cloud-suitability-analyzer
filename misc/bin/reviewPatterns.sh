#!/bin/bash
#
#          FILE:  reviewPatterns.sh
#
#
#   DESCRIPTION:  Take a list of potential patterns and use googler to
#                 review uniquenness of patterns
#       OPTIONS:  ---
#  REQUIREMENTS:  ---
#          BUGS:  ---
#        AUTHOR:  App Tx Team (Steve Woods)
#       COMPANY:  VMWare
#       VERSION:  1.0
#       CREATED:  1/7/19
#      REVISION:  ---
#===============================================================================

#--- exit with usage message if number of command line parameters of less than required


#--- set the usage message
message="Pivotal App Tx\nReview Patterns\nUsage: %s  -p <pattern list>  -o <output file>"


#--- example of a two param script
if (( $# < 2 ))
then
   printf "${message}\n" $(basename $0) >&2
   exit 1
fi


#--- input parameters into variables
while getopts 'p:o:' OPTION
do
   case $OPTION in
   p)    patternFlag=p
         patternFile="$OPTARG"
         ;;
   o)    outputFlag=o
         outputFile="$OPTARG"
         ;;
   #---- exit with error if an unexpected parameter appears
   ?)    printf "$message\n" $(basename $0) >&2
         exit 2
         ;;
   esac
done


while read LINE
do
    googler -n 20 --np "$LINE"
    echo -n "Add pattern $LINE [y/n]? "
    read -u 3 yn
    if [ $yn = "y" ]; then
       echo "Added"
       echo $LINE >> $outputFile
    fi
done 3<&0 <$patternFile
