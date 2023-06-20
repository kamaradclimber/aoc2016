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

func (c Coordinates) Edges() []dijkstra.Edge {
	favorite_number := 1364
	my_edges := make([]dijkstra.Edge, 0)
	candidates := [][2]int{[2]int{c.x, c.y - 1}, [2]int{c.x - 1, c.y}, [2]int{c.x + 1, c.y}, [2]int{c.x, c.y + 1}}
	for _, candidate := range candidates {
		if candidate[0] >= 0 && candidate[1] >= 0 {
			x, y := candidate[0], candidate[1]
			sum := x*x + 3*x + 2*x*y + y + y*y + favorite_number
			bin_value := strconv.FormatInt(int64(sum), 2)
			count_of_ones := 0
			for i := 0; i < len(bin_value); i++ {
				if bin_value[i] == '1' {
					count_of_ones++
				}
			}
			if count_of_ones%2 == 0 {
				my_edges = append(my_edges, MyEdge{c, Coordinates{x, y}})
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
}
