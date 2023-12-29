package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/adde/advent-of-code/utils"
)

const (
	TRIALS = 1000
)

type Component struct {
	Name        string
	Connections []string
}

type QueueNode struct {
	Node string
	Seen []string
}

func (c *Component) AddConnection(connection string) {
	if c.Connections == nil {
		c.Connections = make([]string, 0)
	}

	c.Connections = append(c.Connections, connection)
}

func main() {
	startTime := time.Now()
	lines := utils.ReadInput("input.txt")

	fmt.Println("\nComponent group size:", getComponentGroupSize(parseInput(lines)))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func parseInput(lines []string) map[string]Component {
	components := make(map[string]Component)

	for _, line := range lines {
		componentParts := strings.Split(line, ": ")
		componentName := componentParts[0]
		componentConnections := strings.Split(componentParts[1], " ")

		// Add component if it doesn't exist
		var component Component
		if _, ok := components[componentName]; ok {
			component = components[componentName]
		} else {
			component = Component{Name: componentName}
		}
		for _, connection := range componentConnections {
			component.AddConnection(connection)

			// Add the connections as their own components
			var otherComponent Component
			if _, ok := components[connection]; ok {
				otherComponent = components[connection]
			} else {
				otherComponent = Component{Name: connection}
			}
			otherComponent.AddConnection(componentName)
			components[connection] = otherComponent
		}

		components[componentName] = component
	}

	return components
}

func getComponentGroupSize(components map[string]Component) int {
	totalComps := len(components)
	compInOneGrp := getComponentGroup(components)

	// Since the trial is selecting random components,
	// we are sometimes unlucky and cut the wrong wires.
	// Retry until we get the correct result.
	// We make an assumption that the difference between the
	// total number of components and the number of components
	// in one group should be greater than 100,
	// because the groups should be somewhat evenly distributed.
	// (Very ugly but it works on my input)
	for (totalComps - compInOneGrp) < 100 {
		compInOneGrp = getComponentGroup(components)
	}

	return (totalComps - compInOneGrp) * compInOneGrp
}

func getComponentGroup(components map[string]Component) int {
	seen := utils.OrderedMap{Values: make(map[string]int)}

	// Get some number of random component pairs
	for i := 0; i < TRIALS; i++ {
		c1, c2 := getRandComponent(components), getRandComponent(components)

		// Avoid checking against itself
		for c1 == c2 {
			c2 = getRandComponent(components)
		}

		// Run BFS on them to find the three bridges that we need to cut
		seen = findBridges(c1, c2, components, seen)
	}

	// Sort seen components, highest first
	seen.Sort(true)

	// Make a copy of the components
	componentsCopy := make(map[string]Component)
	for k, v := range components {
		componentsCopy[k] = v
	}

	// Remove the six top seen components (the three bridges)
	for i := 0; i < 6; i++ {
		delete(componentsCopy, seen.Keys[i])
	}

	// Run BFS to get the number of components in one group
	return findAllInGroup(getRandComponent(componentsCopy), componentsCopy)
}

func findBridges(start, end string, components map[string]Component, seen utils.OrderedMap) utils.OrderedMap {
	visited := make(map[string]bool)
	queue := []QueueNode{{Node: start, Seen: []string{}}}

	for len(queue) > 0 {
		comp := queue[0]
		node := comp.Node
		nSeen := comp.Seen
		queue = queue[1:]

		if visited[node] {
			continue
		}

		visited[node] = true

		if node == end {
			for _, n := range nSeen {
				v, ok := seen.Get(n)
				if !ok {
					v = 0
				}
				seen.Set(n, v+1)
			}
			break
		}

		for _, neighbor := range components[node].Connections {
			if !visited[neighbor] {
				nSeen = append(nSeen, node)
				queue = append(queue, QueueNode{Node: neighbor, Seen: nSeen})
			}
		}
	}

	return seen
}

func findAllInGroup(start string, components map[string]Component) int {
	visited := make(map[string]bool)
	queue := []string{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if visited[node] {
			continue
		}

		visited[node] = true

		for _, neighbor := range components[node].Connections {
			if _, ok := components[neighbor]; !visited[neighbor] && ok {
				queue = append(queue, neighbor)
			}
		}
	}

	// Add three since we removed three bridges
	return len(visited) + 3
}

func getRandComponent(components map[string]Component) string {
	keys := make([]string, 0, len(components))

	for k := range components {
		keys = append(keys, k)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return keys[r.Intn(len(keys))]
}
