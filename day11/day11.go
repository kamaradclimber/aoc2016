package main

import (
	"fmt"
	"log"
	"sort"
	// "os"
	"github.com/thoas/go-funk"
	"strings"
)

type GenOrMicro struct {
	madeof    string
	generator bool
}

func (m GenOrMicro) String() string {
	char := "M"
	if m.generator {
		char = "G"
	}
	return fmt.Sprintf("%s%s", strings.ToUpper(string(m.madeof[0])), char)
}

type Floor struct {
	objects []GenOrMicro
}

func (f Floor) String() string {
	objects := make([]string, len(f.objects))
	for i, obj := range f.objects {
		objects[i] = obj.String()
	}
	sort.Slice(objects, func(i, j int) bool {
		return objects[i] < objects[j]
	})
	return strings.Join(objects, " ")

}

type World struct {
	floors         [4]Floor
	elevator_level int
}

func (w World) String() string {
	var lines [4]string
	for i, floor := range w.floors {
		s := fmt.Sprintf("F%d ", i+1)
		if w.elevator_level == i {
			s += "E "
		}
		s += fmt.Sprintf("%v", floor)
		lines[i] = s
	}
	return strings.Join(funk.ReverseStrings(lines[:]), "\n") + "\n"
}

func (f Floor) valid() bool {
	couples := make(map[string][]GenOrMicro, 0)
	rtg_count := 0
	for _, obj := range f.objects {
		couples[obj.madeof] = append(couples[obj.madeof], obj)
		if obj.generator {
			rtg_count += 1
		}
	}
	valid := true
	for _, els := range couples {
		if len(els) > 2 {
			log.Fatal("Big big error!")
		}
		if len(els) == 1 && !els[0].generator {
			// alone chip
			if rtg_count > 0 {
				// any RTG is dangerous since this microchip is not powered
				valid = false
			}
		}
	}
	return valid
}

func (w World) valid() bool {
	valid := true
	for _, floor := range w.floors {
		if !floor.valid() {
			valid = false
		}
	}
	return valid
}

func (w World) PossibleMoves() []World {
	candidates := make([]World, 0)
	for _, elevator_offset := range []int{-1, 1} {
		candidate_level := elevator_offset + w.elevator_level
		if candidate_level < 0 || candidate_level >= 4 {
			continue
		}
		// first we add a world where we just move the elevator without anything
		new_world := World{floors: w.floors, elevator_level: candidate_level}
		if new_world.valid() {
			candidates = append(candidates, new_world)
		}
		for _, object := range w.floors[w.elevator_level].objects {
			for _, object2 := range w.floors[w.elevator_level].objects {
				// second we add a world where we move two items in the elevator
				// case to move just one object is handled when object and object2 are identical
				// we should notice that we add each case two times, but we'll filter that later on
				new_world := w.move(object, w.elevator_level, candidate_level)
				new_world = w.move(object2, w.elevator_level, candidate_level)
				if new_world.valid() {
					candidates = append(candidates, new_world)
				}
			}
		}
	}
	return candidates
}

func (w World) move(object GenOrMicro, old_level int, new_level int) World {
	old_floor := w.floors[old_level]
	old_floor = old_floor.remove(object, object)
	new_floor := Floor{append(w.floors[new_level].objects, object)}
	var floors [4]Floor
	for i := 0; i < len(floors); i++ {
		if i == old_level {
			floors[i] = old_floor
		} else if i == new_level {
			floors[i] = new_floor
		} else {
			floors[i] = w.floors[i]
		}
	}
	return World{floors: floors, elevator_level: new_level}
}

func (f Floor) remove(o1, o2 GenOrMicro) Floor {
	objects := make([]GenOrMicro, 0)
	for _, o := range f.objects {
		if o != o1 && o != o2 {
			objects = append(objects, o)
		}
	}
	return Floor{objects}
}

type WorldExploration struct {
	world World
	step  int
}

func main() {
	// for once let's hardcode input

	//first_floor := Floor{generators: make([]Generator, 0), microchips: []Microchip{Microchip{"hydrogen"}, Microchip{"lithium"}}}
	//second_floor := Floor{generators: []Generator{Generator{"hydrogen"}}, microchips: make([]Microchip, 0)}
	//third_floor := Floor{generators: []Generator{Generator{"lithium"}}, microchips: make([]Microchip, 0)}
	//fourth_floor := Floor{generators: make([]Generator, 0), microchips: make([]Microchip, 0)}
	first_floor := Floor{objects: []GenOrMicro{GenOrMicro{"hydrogen", false}, GenOrMicro{"lithium", false}}}
	second_floor := Floor{objects: []GenOrMicro{GenOrMicro{"hydrogen", true}}}
	third_floor := Floor{objects: []GenOrMicro{GenOrMicro{"lithium", true}}}
	fourth_floor := Floor{objects: make([]GenOrMicro, 0)}

	initial_world := World{floors: [4]Floor{first_floor, second_floor, third_floor, fourth_floor}, elevator_level: 0}

	fmt.Printf("%v", initial_world)

	to_explore := make([]WorldExploration, 0)
	to_explore = append(to_explore, WorldExploration{initial_world, 0})

	explored := make(map[string]bool, 0)

	all_objects := make([]GenOrMicro, 0)

	all_objects = append(all_objects, first_floor.objects...)
	all_objects = append(all_objects, second_floor.objects...)
	all_objects = append(all_objects, third_floor.objects...)
	all_objects = append(all_objects, fourth_floor.objects...)

	floor_with_all_objects := Floor{objects: all_objects}
	empty_floor := Floor{objects: make([]GenOrMicro, 0)}

	target_world := World{elevator_level: 3, floors: [4]Floor{empty_floor, empty_floor, empty_floor, floor_with_all_objects}}

	var best WorldExploration

	for len(to_explore) > 0 {
		world := to_explore[0].world
		count := to_explore[0].step
		to_explore = to_explore[1:]
		if best.step > 0 && best.step < count {
			// it's important taht this check is before the "explored" inclusion test
			// to avoid dismiss a potential shorter path
			continue // we already know a faster way to go there
		}
		if explored[world.String()] {
			continue
		}
		explored[world.String()] = true

		if world.String() == target_world.String() {
			log.Printf("Found a path to target in %d steps", count)
			best = WorldExploration{world, count}
			continue
		}

		for _, move := range world.PossibleMoves() {
			if !explored[move.String()] {
				to_explore = append(to_explore, WorldExploration{move, count + 1})
			}
		}
	}
}
