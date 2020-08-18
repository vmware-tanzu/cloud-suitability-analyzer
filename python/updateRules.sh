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
message="VMWare App Tx\nUpdate CSA rules\nUsage: %s  -g <github url>  -m (replace,add,validate)"


#--- example of a two param script, one value, one flag
if (( $# < 2 ))
then
   printf "${message}\n" $(basename $0) >&2
   exit 1
fi

packageFlag=0

#--- input parameters into variables
while getopts 'd:m:' OPTION
do
   case $OPTION in
   d)    ruleDirFlag=g
         ruleDir="$OPTARG"
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

validModes="rav"

if [[ $mode != "a" && $mode != "r" && $mode != "v" ]]; then
   echo "Invalid mode: $mode"
   echo "r - replace existing rules"
   echo "a - add to existing rules"
   echo "v - validate rule yaml files"
   exit 1
fi

if [[ "$mode" == "r" ]]; then
   csa rules delete-all
   if [ $? -eq 0 ]; then
      echo "Successfully deleted all CSA rules"
   else
      echo "Error deleting CSA rules"
      exit 1
   fi 

   csa rules import --rules-dir $ruleDir

   if [ $? -eq 0 ]; then
      echo "Successfully imported rules from directory: $ruleDir"
      exit 0
   else
      echo "Error importing rules from directory: $ruleDir"
      exit 1
   fi 
   echo "Replacing existing CSA rules with yaml files from $rulesDir"
   exit 0
fi

if [[ "$mode" == "a" ]]; then
   echo "Adding to existing CSA rules with yaml files from $rulesDir"
   csa rules import --rules-dir $ruleDir
   if [ $? -eq 0 ]; then
      echo "Successfully imported rules from directory: $ruleDir"
      exit 0
   else
      echo "Error importing rules from directory: $ruleDir"
      exit 1
   fi 
fi


if [[ "$mode" == "v" ]]; then

      ruleList=$(find $ruleDir -name "*.yaml")
      rulesPassed=0
      rulesFailed=0
      for ruleFile in $ruleList
      do
      echo -n "Validating $ruleFile"
      result="$(csa rules validate $ruleFile 2>&1)"
      if echo "$result" | grep -q "error"; then
            echo ": ***** NOT VALID!! *****"
            rulesFailed=$(($rulesFailed + 1))
      else
            echo ": valid"
            rulesPassed=$(($rulesPassed + 1))
      fi
      done

      echo "Rules passed: $rulesPassed"
      echo "Rules failed: $rulesFailed"
      exit 0
fi