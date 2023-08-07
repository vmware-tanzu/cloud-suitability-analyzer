#!/bin/bash
for yamlfile in *.yaml; do
  cypherFile="$(basename $yamlfile .yaml).cypher"
  printf "MATCH (n:Application)-[c:CONSUMES]->(a:ApplicationService)-[r:CONSISTS_OF]->(s:SourceModel)-[:IMPORTS]->(i:Import)\nWHERE " > $cypherFile
  printf "\nProcessing $cypherFile"
  
  #--- yaml file contents
  name=$(yq eval '.name' $yamlfile)
  filetype=$(yq eval '.filetype' $yamlfile)
  target=$(yq eval '.target' $yamlfile)
  type=$(yq eval '.type' $yamlfile)
  defaultpattern=$(yq eval '.defaultpattern' $yamlfile)
  advice=$(yq eval '.advice' $yamlfile)
  effort=$(yq eval '.effort' $yamlfile) 
  impact=$(yq eval '.impact' $yamlfile)
  readiness=$(yq eval '.readiness' $yamlfile)
  category=$(yq eval '.category' $yamlfile)
  criticality=$(yq eval '.criticality' $yamlfile)
  tags=$(yq eval '.tags[].value' $yamlfile)
  recipes==$(yq eval '.recipes[].value' $yamlfile)
  patterns=$(yq eval '.patterns[].value' $yamlfile)
  regex=$(yq eval '.regex' $yamlfile)

  patternCount=$(echo "$patterns" | wc -w)
  currentLine=1

  for import in $patterns; do
      
      if [ $currentLine -lt $patternCount ]; then
          printf "\ni.name CONTAINS \\'$import\\' OR  " >> $cypherFile
      else
          printf "\ni.name CONTAINS \\'$import\\' " >> $cypherFile
      fi
      let currentLine++   
  done

  printf "\nRETURN\n n.appName as app, \'import \' + i.name as import, " >> $cypherFile
  printf "\'${name}\' as rule, \'${defaultpattern}\' as pattern, \'${advice}\' as advice, $effort as effort, $impact as impact, " >> $cypherFile
  printf "$readiness as readiness, \'${category}\' as category, $criticality as criticality, " >> $cypherFile
  printf "s.location as fdn, s.name as filename, apoc.text.regexGroups(s.name, \'\.[a-z]*$\')[0][0] as ext, " >> $cypherFile 
  printf "null as line, 1 as run_id, null as note, null as result" >> $cypherFile 

  if [ ! -z "$tags" ]
  then 
     tagclause=""
     for tag in $tags; do
        tag=${tag//[-\/]/_}
        tag=" 3000 as $tag,"
        tagclause="$tagclause$tag"
     done
     temp=$tagclause

     tagclause=${temp%?}

     printf ", $tagclause" >> $cypherFile
  fi

  
done 
