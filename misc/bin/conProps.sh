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
message="VMWare App Tx\nConsolidate properties\nUsage: %s  -d <directory>  -p <prefix -n <service name> -s <suffix>"


#--- example of a two param script, one value, one flag
if (( $# < 4 ))
then
   printf "${message}\n" $(basename $0) >&2
   exit 1
fi

packageFlag=0

#--- input parameters into variables
while getopts 'd:p:n:s:' OPTION
do
   case $OPTION in
   d)    directoryFlag=d
         directory="$OPTARG"
         ;;
   p)    prefixFlag=p
         prefix="$OPTARG"
         ;;
   n)    serviceNameFlag=n
         serviceName="$OPTARG"
         ;;
   s)    suffixFlag=s
         suffix="$OPTARG"
         ;;

   #---- exit with error if an unexpected parameter appears
   ?)    printf "$message\n" $(basename $0) >&2
         exit 2
         ;;
   esac
done

#--- find all properties patterns
egrep -Rh --include="*.properties" "^[A-Za-z.]+=" ${directory} | \

#--- strip the value from expression
sed 's/=.*/=/' | \

#--- sort, keep all
sort |

#--- count the occurances
uniq -c | \

#--- filter out those with only one occurance
awk '{ if ($1 >  1) print $2 }' | \

#--- filter out duplicates
sort -u | \

#--- set values
sed -r "s/([A-Za-z.]*)=.*/\1=${prefix}${serviceName}\1${suffix}/"
