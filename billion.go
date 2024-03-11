package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	min := make(map[string]float64)
	max := make(map[string]float64)
	avg := make(map[string]float64)

	station_name_counts := make(map[string]int)

	readFile, _ := os.Open("measurements.txt")

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lineparts := strings.Split(line, ";")

		station_name, temp_as_string := lineparts[0], lineparts[1]

		temp_val, _ := strconv.ParseFloat(temp_as_string, 32)
		temp_reading := float64(temp_val)

		if min[station_name] > temp_reading {
			min[station_name] = temp_reading
		}

		if temp_reading > max[station_name] {
			max[station_name] = temp_reading
		}

		station_name_counts[station_name]++
		cur_avg := (avg[station_name]*float64(station_name_counts[station_name]) + temp_reading) / float64(station_name_counts[station_name]+1)

		avg[station_name] = cur_avg
	}

	readFile.Close()

	print("{")
	for station_name, avg_temp := range avg {
		weather_station := station_name + "=" + strconv.FormatFloat(min[station_name], 'f', 1, 64) + "/" + strconv.FormatFloat(avg_temp, 'f', 1, 64) + "/" + strconv.FormatFloat(max[station_name], 'f', 1, 64) + ","
		print(weather_station)
	}
	print("}")
}
