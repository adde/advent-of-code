package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/adde/advent-of-code/utils"
)

const (
	RATING_MIN = 1
	RATING_MAX = 4000
)

type Rule struct {
	Category string
	Compare  string
	Rating   int
	Target   string
}

func main() {
	workflows := make(map[string][]Rule)
	parts := []map[string]int{}

	startTime := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	isParts := false
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			isParts = true
			continue
		}

		if !isParts {
			// Parse and store workflows
			lineParts := strings.Split(line, "{")
			workflowName := lineParts[0]
			workflowParts := lineParts[1][:len(lineParts[1])-1]
			workflowRules := strings.Split(workflowParts, ",")
			rules := []Rule{}

			for _, rule := range workflowRules {
				if strings.Contains(rule, ":") {
					ruleParts := strings.Split(rule, ":")
					ruleCond := ruleParts[0]
					ruleTarget := ruleParts[1]

					rules = append(rules, Rule{
						Category: string(ruleCond[0]),
						Compare:  string(ruleCond[1]),
						Rating:   utils.ToInt(ruleCond[2:]),
						Target:   ruleTarget,
					})
				} else {
					rules = append(rules, Rule{Target: rule, Rating: -1})
				}
			}

			workflows[workflowName] = rules
		} else {
			// Parse and store machine parts
			lineParts := strings.Split(line[1:len(line)-1], ",")
			p := map[string]int{}

			for _, part := range lineParts {
				partName := strings.Split(part, "=")[0]
				partRating := utils.ToInt(strings.Split(part, "=")[1])
				p[partName] = partRating
			}

			parts = append(parts, p)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nSum of accepted part ratings:", getTotalPartsRating(workflows, parts))
	fmt.Println("Distinct combinations of ratings:", getDistinctCombinations(getInitialRanges(), workflows, "in"))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func getTotalPartsRating(workflows map[string][]Rule, parts []map[string]int) int {
	sum := 0

	for _, part := range parts {
		if isPartAccepted(workflows, "in", part) {
			for _, key := range []string{"x", "m", "a", "s"} {
				sum += part[key]
			}
		}
	}

	return sum
}

func isPartAccepted(workflows map[string][]Rule, workflowName string, part map[string]int) bool {
	if workflowName == "A" {
		return true
	} else if workflowName == "R" {
		return false
	}

	for _, rule := range workflows[workflowName] {
		if rule.Compare == ">" {
			if part[rule.Category] > rule.Rating {
				return isPartAccepted(workflows, rule.Target, part)
			}
		} else if rule.Compare == "<" {
			if part[rule.Category] < rule.Rating {
				return isPartAccepted(workflows, rule.Target, part)
			}
		} else {
			return isPartAccepted(workflows, rule.Target, part)
		}
	}

	return false
}

func getDistinctCombinations(ranges map[string][2]int, workflows map[string][]Rule, name string) int {
	total := 0
	lastRule := ""

	if name == "R" {
		return 0
	}
	if name == "A" {
		combinations := 1
		for _, r := range ranges {
			combinations *= (r[1] - r[0] + 1)
		}
		return combinations
	}

	for _, rule := range workflows[name] {
		matching, notMatching := [2]int{}, [2]int{}

		if rule.Compare == "<" {
			matching = [2]int{ranges[rule.Category][0], rule.Rating - 1}
			notMatching = [2]int{rule.Rating, ranges[rule.Category][1]}
		} else if rule.Compare == ">" {
			matching = [2]int{rule.Rating + 1, ranges[rule.Category][1]}
			notMatching = [2]int{ranges[rule.Category][0], rule.Rating}
		} else {
			lastRule = rule.Target
			break
		}

		if matching[0] <= matching[1] {
			copyRanges := copyMap(ranges)
			copyRanges[rule.Category] = matching
			total += getDistinctCombinations(copyRanges, workflows, rule.Target)
		}
		if notMatching[0] <= notMatching[1] {
			copyRanges := copyMap(ranges)
			copyRanges[rule.Category] = notMatching
			ranges = copyRanges
		} else {
			break
		}
	}

	if lastRule != "" {
		total += getDistinctCombinations(ranges, workflows, lastRule)
	}

	return total
}

func getInitialRanges() map[string][2]int {
	ranges := map[string][2]int{}

	for _, r := range []string{"x", "m", "a", "s"} {
		ranges[r] = [2]int{RATING_MIN, RATING_MAX}
	}

	return ranges
}

func copyMap(m map[string][2]int) map[string][2]int {
	copyMap := make(map[string][2]int)

	for key, value := range m {
		copyMap[key] = value
	}

	return copyMap
}
