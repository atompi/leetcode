#!/bin/bash

cat words.txt | egrep -o "\b[[:alpha:]]+\b" | sort | uniq -c | sort -r | awk '{print $2,$1}'
