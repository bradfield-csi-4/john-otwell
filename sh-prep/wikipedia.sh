#!/bin/bash


wiki () {
    WIKI_BASE_URL="https://en.wikipedia.org/wiki/"
    capitalized=$(echo $1 | sed 's/\b[A-Za-z]/\U&/')
    res=$(curl -s $WIKI_BASE_URL$capitalized)
    thing=$(echo $res | sed -e "/<p>/,/<\/p>/P")
    echo $thing
}

wiki $1
