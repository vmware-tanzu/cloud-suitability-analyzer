#! /bin/bash

while read -r url; 
    do  
    if curl --output /dev/null --silent --head --fail "$url"; then
        echo "URL exists: $url"
    else
    echo "URL does not exist: $url"
    fi
done < uris.txt
