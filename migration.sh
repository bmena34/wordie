#!/bin/bash

# Check if a CSV file is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <csv_file>"
  exit 1
fi

CSV_FILE=$1
KEY_COUNTER=1

# Read the CSV file line by line
while IFS=, read -r category word; do

  # Add the values to Redis as a hash
  redis-cli HSET "key_$KEY_COUNTER" category "$category" word "$word"
  KEY_COUNTER=$((KEY_COUNTER + 1))
done < "$CSV_FILE"

echo "Data from $CSV_FILE has been added to Redis."