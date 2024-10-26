#!/bin/bash

# Check if a CSV file is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <csv_file>"
  exit 1
fi

CSV_FILE=$1
KEY_COUNTER=1

# Set IFS to newline and comma
IFS=$'\n,'

# Read the CSV file line by line
while read -r line; do
if [ -z "$line" ]; then
    continue
  fi

  # Split the line into category and word
  category=$(echo "$line" | cut -d',' -f1)
  word=$(echo "$line" | cut -d',' -f2 | tr -d '\n\r')
  
  # Add the values to Redis as a hash
  redis-cli HSET "$KEY_COUNTER" category "$category" word "$word"
  
  # Increment the key counter
  KEY_COUNTER=$((KEY_COUNTER + 1))
done < "$CSV_FILE"

echo "Data from $CSV_FILE has been added to Redis."