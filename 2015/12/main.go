package main

import (
	"encoding/json"
	"fmt"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	json := u.ReadAll("input.txt")

	// Part one
	numbers := u.GetIntsFromString(json, true)
	ansP1 := 0
	for _, n := range numbers {
		ansP1 += n
	}

	// Part two
	newJson, err := cleanJSON(json)
	if err != nil {
		fmt.Println("Error cleaning JSON:", err)
		return
	}
	numbers = u.GetIntsFromString(newJson, true)
	ansP2 := 0
	for _, n := range numbers {
		ansP2 += n
	}

	fmt.Println("\nPart one:", ansP1)
	fmt.Println("Part two:", ansP2)
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func removeRed(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		// Check if any value in this object is "red"
		hasRed := false
		for _, val := range v {
			if str, ok := val.(string); ok && str == "red" {
				hasRed = true
				break
			}
		}
		if hasRed {
			return nil
		}

		// Recursively process all values in the object
		result := make(map[string]interface{})
		for key, val := range v {
			if processed := removeRed(val); processed != nil {
				result[key] = processed
			}
		}
		if len(result) > 0 {
			return result
		}
		return nil

	case []interface{}:
		// Recursively process array elements
		result := make([]interface{}, 0)
		for _, val := range v {
			if processed := removeRed(val); processed != nil {
				result = append(result, processed)
			}
		}
		if len(result) > 0 {
			return result
		}
		return nil

	default:
		return data
	}
}

func cleanJSON(jsonStr string) (string, error) {
	var data interface{}

	// Parse JSON
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return "", err
	}

	// Remove objects containing "red"
	cleaned := removeRed(data)

	// Convert back to JSON
	result, err := json.Marshal(cleaned)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
