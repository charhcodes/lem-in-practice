// prints one path

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	name    string
	links   []*Vertex // adjacent
	visited bool
	ants    int
	path    []*Vertex
}

type Graph struct {
	vertices []*Vertex
}

// open file (os.Args[1]) and split into separate lines
func readFile() []string {
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// return number of ants
func antCount() int {
	data := readFile()
	ants := data[0]
	if ants <= "0" {
		err := fmt.Errorf("ERROR: invalid number of ants")
		fmt.Println(err.Error())
	}
	count, _ := strconv.Atoi(string(ants))
	return count
}

// returns starting room
func getStart() string {
	var start string
	file := readFile()

	for i, _ := range file {
		if file[i] == "##start" {
			start = strings.Split(string(file[i+1]), " ")[0]
		}
	}
	return start
}

// returns ending room
func getEnd() string {
	var end string
	file := readFile()

	for i, _ := range file {
		if file[i] == "##end" {
			end = strings.Split(string(file[i+1]), " ")[0]
		}
	}
	return end
}

// return a pointer to the Vertex with its name
func (g *Graph) getVertex(name string) *Vertex {
	for i, v := range g.vertices {
		if v.name == name {
			return g.vertices[i]
		}
	}
	return nil
}

var (
	startroom = getStart()
	endroom   = getEnd()
)

// checks if there are repeated rooms
func contains(s []*Vertex, name string) bool {
	for _, v := range s {
		if name == v.name {
			return true
		}
	}
	return false
}

// add vertex to graph
func (g *Graph) AddVertex(name string) { //*Node
	if contains(g.vertices, name) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", name)
		fmt.Println(err.Error())
	} else {
		vertices := &Vertex{name: name} // creates a pointer to a new Vertex struct, initialises with a field called name
		// and sets its value to the value of the name variable
		g.vertices = append(g.vertices, vertices)
	}
}

// add edges to graph
func (g *Graph) AddEdge(from, to string) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil { // if edges are valid
		err := fmt.Errorf("ERROR: invalid edges")
		fmt.Println(err.Error())
	} else if contains(fromVertex.links, to) { // if vertex already exists
		err := fmt.Errorf("ERROR: edge already exists")
		fmt.Println(err.Error())
	} else if fromVertex == toVertex { // if edges are the same
		err := fmt.Errorf("ERROR: cannot connect room to itself")
		fmt.Println(err.Error())
	} else if fromVertex.name == endroom { // if 'from' room  = end
		toVertex.links = append(toVertex.links, fromVertex)
	} else if toVertex.name == startroom { // if 'to' room = start
		toVertex.links = append(toVertex.links, fromVertex)
	} else {
		fromVertex.links = append(fromVertex.links, toVertex)
	}
}

// print graph out
func (g *Graph) Print() {
	fmt.Printf("Number of Ants: %v", antCount())
	fmt.Printf("\nStarting room: %v", getStart())
	fmt.Printf("\nEnding room: %v", getEnd())
	fmt.Println()

	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v: ", v.name)
		for _, v := range v.links {
			fmt.Printf("%v ", v.name)
		}
	}
}

var verticesMap = make(map[string]*Vertex)

func (g *Graph) addtoMap() map[string]*Vertex {
	// iterate over the vertices in the graph and add them to the map
	for _, v := range g.vertices {
		verticesMap[v.name] = v
	}
	return verticesMap
}

func (g *Graph) DFS(start, end string) map[string][]*Vertex {
	// create a map to store the paths
	paths := make(map[string][]*Vertex)

	// find the start and end vertices
	startVertex := g.getVertex(start)
	endVertex := g.getVertex(end)

	// run the DFS algorithm
	g.dfsHelper(startVertex, endVertex, []*Vertex{}, paths)

	return paths
}

func (g *Graph) dfsHelper(current, end *Vertex, path []*Vertex, paths map[string][]*Vertex) {
	// add the current vertex to the path
	path = append(path, current)
	// mark the current vertex as visited
	current.visited = true

	// if we reach the end vertex, add the path to the map
	if current == end {
		paths[end.name] = make([]*Vertex, len(path))
		copy(paths[end.name], path)
	} else {
		// otherwise, continue the DFS on the adjacent vertices
		for _, v := range current.links {
			if !v.visited {
				g.dfsHelper(v, end, path, paths)
			}
		}
	}

	// backtrack
	current.visited = false
	path = path[:len(path)-1]
}

func main() {
	test := Graph{}
	for i, line := range readFile() {
		if strings.Contains(string(line), " ") {
			test.AddVertex(strings.Split(readFile()[i], " ")[0])
		}
		if strings.Contains(string(line), "-") {
			test.AddEdge(strings.Split(readFile()[i], "-")[0], strings.Split(readFile()[i], "-")[1])
		}
	}

	test.Print()
	fmt.Println()

	test.addtoMap()
	fmt.Println()
	fmt.Println(verticesMap)
	fmt.Println()

	paths := test.DFS(startroom, endroom)

	for end, path := range paths {
		fmt.Printf("Path to %s: ", end)
		for _, v := range path {
			fmt.Printf("%s ", v.name)
		}
		fmt.Println()
	}
}

// https://www.youtube.com/watch?v=bSZ57h7GN2w
