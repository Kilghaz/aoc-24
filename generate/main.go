package main

import (
	"aoc24/aoc"
	"fmt"
	"os"
	"path"
	"regexp"
	"slices"
	"strings"
	"time"
)

var fileTemplate = "package day_%d\n\nfunc Part%d(input string) {\n}\n"
var inputsPath = "inputs"

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func main() {
	client := aoc.CreateClient()
	date := time.Now().Local()
	day := date.Day()

	inputFileName := path.Join(inputsPath, fmt.Sprintf("input-%d.txt", day))
	packageName := fmt.Sprintf("day_%d", day)
	file, _ := os.ReadFile("main/main.go")
	mainSrc := string(file)

	var isMainSetup = func() bool {
		exp, _ := regexp.Compile(fmt.Sprintf("%d:\\s+{\\s*day_%d.Part1,\\s*day_%d.Part2\\s*},", day, day, day))
		return exp.MatchString(mainSrc)
	}

	var updateImports = func() {
		exp, _ := regexp.Compile("(?s)import \\(.*?\\)\n")
		importSrc := exp.FindString(mainSrc)
		lines := strings.Split(importSrc, "\n")
		imports := lines[1 : len(lines)-2]
		imports = append(imports, fmt.Sprintf("\t\"aoc24/day_%d\"", day))
		slices.Sort(imports)
		importSrcUpdated := fmt.Sprintf("import (\n%s\n)\n", strings.Join(imports, "\n"))
		mainSrc = strings.Replace(mainSrc, importSrc, importSrcUpdated, -1)
	}

	var updateMap = func() {
		exp, _ := regexp.Compile("(?s)var advents =.*?},\\s+}\\s")
		mapSrc := exp.FindString(mainSrc)
		lines := strings.Split(mapSrc, "\n")
		days := lines[1 : len(lines)-2]
		days = append(days, fmt.Sprintf("\t%d: {day_%d.Part1, day_%d.Part2},", day, day, day))
		mapSrcUpdated := fmt.Sprintf("var advents = map[int][]func(input string){\n%s\n}\n", strings.Join(days, "\n"))
		mainSrc = strings.Replace(mainSrc, mapSrc, mapSrcUpdated, -1)
	}

	if !exists(packageName) {
		_ = os.Mkdir(packageName, 0755)
		_ = os.WriteFile(path.Join(packageName, "part-1.go"), []byte(fmt.Sprintf(fileTemplate, day, 1)), 0755)
		_ = os.WriteFile(path.Join(packageName, "part-2.go"), []byte(fmt.Sprintf(fileTemplate, day, 2)), 0755)
	}

	if !exists(inputFileName) {
		input := client.LoadInput(date.Year(), day)
		_ = os.WriteFile(path.Join(inputsPath, fmt.Sprintf("input-%d.txt", day)), []byte(input), 0755)
	}

	if !isMainSetup() {
		updateImports()
		updateMap()
		_ = os.WriteFile("main/main.go", []byte(mainSrc), 0755)
	}
}
