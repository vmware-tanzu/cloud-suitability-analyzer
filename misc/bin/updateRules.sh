#!/bin/bash
#
#          FILE:  updateRules.sh
#
#
#   DESCRIPTION:  Update CSA rules using a github repo url
#
#       OPTIONS:  ---
#  REQUIREMENTS:  ---
#          BUGS:  ---
#        AUTHOR:  App Tx Team (Steve Woods)
#       COMPANY:  VMWare
#       VERSION:  1.0
#       CREATED:  08/16/2020
#      REVISION:  ---
#===============================================================================

#--- exit with usage message if number of command line parameters of less than required


#--- set the usage message
message="VMWare App Tx\Update CSA rules\nUsage: %s  -g <github url>  -m (replace,add)"


#--- example of a two param script, one value, one flag
if (( $# < 2 ))
then
   printf "${message}\n" $(basename $0) >&2
   exit 1
fi

packageFlag=0

#--- input parameters into variables
while getopts 'g:m:' OPTION
do
   case $OPTION in
   g)    gitHubUrlFlag=g
         gitHubUrl="$OPTARG"
         ;;
   m)    modeFlag=m
         mode="$OPTARG"
         ;;

   #---- exit with error if an unexpected parameter appears
   ?)    printf "$message\n" $(basename $0) >&2
         exit 2
         ;;
   esac
done

git clone $gitHubUrl tmp

