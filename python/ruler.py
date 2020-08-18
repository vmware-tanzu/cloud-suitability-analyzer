#!/usr/bin/env python3

#
#          FILE:  ruler.py
#
#
#   DESCRIPTION:  Manage CSA rules
#
#       OPTIONS:  ---
#  REQUIREMENTS:  ---
#          BUGS:  ---
#        AUTHOR:  App Tx Team (Steve Woods)
#       COMPANY:  VMWare
#       VERSION:  1.0
#       CREATED:  08/16/2020
#      REVISION:  ---
# ===============================================================================

import argparse
import os

# --- process command lines
import sys

parser = argparse.ArgumentParser(description='Manage CSA rules')

parser.add_argument('-d', '--dir', metavar='directory', action='store', required=True, dest='directory')

parser.add_argument('-m', '--mode', metavar='mode', action='store', choices={'add', 'replace', 'verify'}, dest='mode')

args = parser.parse_args()

# --- add rules to existing rules
if args.mode == 'add':
    print(f'adding rules from {args.directory}')
    exitCode = os.system(f'csa rules import --rules-dir {args.directory}')
    if exitCode == 0:
        print("Rule import succeeded")
        sys.exit(0)
    else:
        print("Rule import failed")
        sys.exit(1)

# --- replace existing rules
if args.mode == 'replace':
    print(f'deleting all rules from csa')
    exitCode = os.system('csa rules delete-all')
    if exitCode == 0:
        print("Rules deletion successful")
    else:
        print("Rule deletion failed")
        sys.exit(0)

    print(f'adding rules from {args.directory}')
    exitCode = os.system(f'csa rules import --rules-dir {args.directory}')
    if exitCode == 0:
        print('rule import succeeded')
        sys.exit(0)
    else:
        print('rule import failed')
        sys.exit(1)
