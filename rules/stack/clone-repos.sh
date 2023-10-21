#!/bin/bash

# List of Java application server names
servers=("glassfish" "jboss" "jetty" "payara" "resin" "tomcat" "weblogic" "websphere")

# Loop through each server
for server in "${servers[@]}"; do
    echo "Top 5 starred repositories for $server:"
    
    # Use GitHub CLI to search for repositories by server name and sort by stars in descending order
    gh repo list --language=Java "topic:$server" --limit 5 --sort=stars --json nameWithOwner,url --jq '.[] | "\(.nameWithOwner) - \(.url)"'
    
    echo "" # Add an empty line for separation
done
