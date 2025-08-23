package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Station struct {
	Name  string
	Max   int64
	Min   int64
	Sum   int64
	Count int
}

func parseLines(line string) (string, int64) {
	
	i := strings.Index(line, ";")
	name := line[:i]

	chunk := line[i+1:]
	neg := false

	if chunk[0] == '-' {
		neg = true
	}

	var temp int64

	for i := 0; i < len(chunk); i++ {
		if chunk[i] == '\n' {
			break
		}
		if chunk[i]=='-' {
			continue
		}
		
		if chunk[i]=='.' {
			continue
		}

		temp = temp*10 + int64(chunk[i]-'0')
	}
	if neg {
		temp = -temp
	}
	return name, temp

}

func main() {
	stations := make(map[string]*Station)
	f, err := os.Open("./data/measurements.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		
		name, temp := parseLines(line)

		if _, ok := stations[name]; !ok {
			stations[name] = &Station{
				Name:  name,
				Max:   temp,
				Min:   temp,
				Sum:   temp,
				Count: 1,
			}
		} else {
			if temp > stations[name].Max {
				stations[name].Max = temp
			} else if temp < stations[name].Min {
				stations[name].Min = temp
			}
			stations[name].Sum += temp
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
		mean := s.Sum / int64(s.Count)
		fmt.Fprintf(os.Stdout, "%s=%.1f/%.1f/%.1f", station, (float64)(s.Min/10), (float64)(mean/10), (float64)(s.Max/10))
	}
	fmt.Fprint(os.Stdout, "}\n")
}
