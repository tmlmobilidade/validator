#!/bin/bash

# Path to the source file containing the map
GTFS_FILE="src/types/gtfs.go"
RULES_DIR="src/rules"
MAIN_GO="$RULES_DIR/main.go"

# Create rules directory if it doesn't exist
mkdir -p "$RULES_DIR"

# Extract function names and create folders
grep -o '".*\.txt"' "$GTFS_FILE" | while read -r line; do
    # Remove quotes and .txt extension
    filename=$(echo "$line" | tr -d '"' | sed 's/\.txt$//')
    
    # Create folder
    mkdir -p "$RULES_DIR/$filename"
    
    # Convert to CamelCase function name
    function_name=$(echo "$filename" | sed -r 's/(^|_)([a-z])/\U\2/g')Rules
    
    # Check if function already exists
    if ! grep -q "func $function_name()" "$MAIN_GO"; then
        # Append new function to main.go
        echo -e "\nfunc $function_name() {\n\tfmt.Println(\"$function_name\")\n}" >> "$MAIN_GO"
    fi
done

echo "Script completed successfully!"