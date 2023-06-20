package main

import (
	"github.com/flemeur/go-shortestpath/dijkstra"
	"log"
	"strconv"
)

type Coordinates struct {
	x int
	y int
}

func (c Coordinates) OpenSpace() bool {
	favorite_number := 1364
	x, y := c.x, c.y
	sum := x*x + 3*x + 2*x*y + y + y*y + favorite_number
	bin_value := strconv.FormatInt(int64(sum), 2)
	count_of_ones := 0
	for i := 0; i < len(bin_value); i++ {
		if bin_value[i] == '1' {
			count_of_ones++
		}
	}
	return count_of_ones%2 == 0
}

func (c Coordinates) Edges() []dijkstra.Edge {
	my_edges := make([]dijkstra.Edge, 0)
	candidates := [][2]int{[2]int{c.x, c.y - 1}, [2]int{c.x - 1, c.y}, [2]int{c.x + 1, c.y}, [2]int{c.x, c.y + 1}}
	for _, candidate := range candidates {
		if candidate[0] >= 0 && candidate[1] >= 0 {
			dest := Coordinates{candidate[0], candidate[1]}
			if dest.OpenSpace() {
				my_edges = append(my_edges, MyEdge{c, dest})
			}
		}
	}

	return my_edges
}

type MyEdge struct {
	source      Coordinates
	destination Coordinates
}

func (e MyEdge) Destination() dijkstra.Node {
	return e.destination
}

func (e MyEdge) Weight() float64 {
	return 1
}

func main() {
	path, err := dijkstra.ShortestPath(Coordinates{1, 1}, Coordinates{31, 39})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Part1: %d", len(path)-1)

	// for part 2, we are clearly at a disavantage since I choose to use a package that implement dijkstra without exposing the intermediate computation of shortest path.
	// Let's bruteforce!

	reachable_count := 0
	for x := 0; x <= 52; x++ {
		for y := 0; y <= 52; y++ {
			if (Coordinates{x, y}).OpenSpace() {
				path, err := dijkstra.ShortestPath(Coordinates{1, 1}, Coordinates{x, y})
				if err != nil { // ok this is quite ugly!
					// log.Print(Coordinates{x, y}, err)
					continue
				}
				if len(path)-1 <= 50 {
					reachable_count += 1
				}
			}

		}
	}
	log.Printf("Part2: %d", reachable_count)

}
