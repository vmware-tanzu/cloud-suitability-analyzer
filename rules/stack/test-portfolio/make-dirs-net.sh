#!/bin/bash

# Check if the file path is provided as an argument
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <path-to-file>"
    exit 1
fi

# Read the file path
file="$1"

# Initialize variables
stack_name=""
imports=""

# Iterate through each line in the file
while IFS= read -r line; do
    # If the line starts with '#CSharp', it's the name of the stack
    if [[ $line == "#CSharp"* ]]; then
        # If we already have a stack name, we should create the directory and file for the previous stack
        if [ ! -z "$stack_name" ]; then
            mkdir -p "$stack_name"
            echo -e "$imports" > "$stack_name/$stack_name.java"
        fi
        # Update the stack name and reset imports
        stack_name=$(echo "$line" | sed 's/#//;s/ //g')
        imports=""
    elif [ ! -z "$line" ]; then
        # Add the import statement to imports
        imports+="$(echo "$line" | sed 's/^[ \t]*//')\n"
    fi
done < "$file"

# Handle the last stack
if [ ! -z "$stack_name" ]; then
    mkdir -p "$stack_name"
    echo -e "$imports" > "$stack_name/$stack_name.java"
fi
