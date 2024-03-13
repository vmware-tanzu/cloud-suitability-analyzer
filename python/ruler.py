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
#       COMPANY:  VMware by Broadcom
#       VERSION:  1.0
#       CREATED:  08/16/2020
#      REVISION:  1.1
# ===============================================================================

import argparse
import os
import subprocess
import sys

parser = argparse.ArgumentParser(description='Manage CSA rules')

parser.add_argument('-d', '--dir', metavar='directory', action='store', required=True, dest='directory')

parser.add_argument('-m', '--mode', metavar='mode', action='store',
                    choices={'add', 'replace', 'verify', 'export'}, dest='mode',
                    help='modes: %(choices)s')

args = parser.parse_args()

# --- export rules to directory
if args.mode == 'export':
    print(f'exporting rules to {args.directory}')
    exitCode = os.system(f'csa rules export --rules-dir {args.directory}')
    if exitCode == 0:
        print("Rule export succeeded")
        sys.exit(0)
    else:
        print("Rule export failed")
        sys.exit(1)

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

# --- verify rules
if args.mode == 'verify':
    goodRuleCount = 0
    badRuleCount = 0
    expandedDir = os.path.expanduser(args.directory)
    for entry in os.scandir(expandedDir):
        if entry.path.endswith(".yaml") and entry.is_file():
            print(entry.path, end='')
            results = subprocess.check_output(['csa', 'rules', 'validate', entry.path], stderr=subprocess.STDOUT)
            if b"error" in results:
                print(" - failed")
                badRuleCount += 1
            else:
                goodRuleCount += 1
                print(" - verified")
print(f'Valid rules: {goodRuleCount}, Invalid rules: {badRuleCount}')
