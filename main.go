package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Station struct {
	Name  string
	Max   float64
	Min   float64
	Sum   float64
	Count int
}

func main() {
	stations := make(map[string]*Station)
	f, err := os.Open("./data/mm.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines := scanner.Text()
		parts := strings.Split(lines, ";")

		name := parts[0]
		value, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			panic(err)
		}

		if _, ok := stations[name]; !ok {
			stations[name] = &Station{
				Name:  name,
				Max:   value,
				Min:   value,
				Sum:   value,
				Count: 1,
			}
		} else {
			if value > stations[name].Max {
				stations[name].Max = value
			} else if value < stations[name].Min {
				stations[name].Min = value
			}
			stations[name].Sum += value
			stations[name].Count++
		}

	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}

	keys := make([]string, 0, len(stations))

	for station := range stations {
		keys = append(keys, station)
	}

	sort.Strings(keys)

	fmt.Fprint(os.Stdout, "{")
	for i, station := range keys {
		if i > 0 {
			fmt.Fprint(os.Stdout, ", ")
		}
		s := stations[station]
		mean := s.Sum / float64(s.Count)
		fmt.Fprintf(os.Stdout, "%s=%.1f/%.1f/%.1f", station, s.Min, mean, s.Max)
	}
	fmt.Fprint(os.Stdout, "}\n")
}
