package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Dir struct {
	Files    []File
	Dirs     []*Dir
	Path     string
	Name     string
	Size     int
	Previous *Dir
}

type File struct {
	Name string
	Path string
	Dir  *Dir
	Size int
}

func readInput(path string) []string {
	ret := []string{}

	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return ret
}

func cd(cwd *Dir, root *Dir, path string) *Dir {
	if path == "/" {
		return root
	} else if path == ".." {
		return cwd.Previous
	} else {
		for _, dir := range cwd.Dirs {
			if dir.Name == path {
				return dir
			}
		}
	}

	return nil
}

func dirSize(dir *Dir) int {
	totalSize := 0
	for _, file := range dir.Files {
		totalSize += file.Size
	}

	for _, subDir := range dir.Dirs {
		totalSize += dirSize(subDir)
	}

	dir.Size = totalSize
	return totalSize
}

func parseInput(input []string) Dir {
	rootDir := Dir{}
	rootDir.Path = "/"
	rootDir.Name = "/"

	cwd := &rootDir

	for _, line := range input {
		splitLine := strings.Split(line, " ")
		if splitLine[0] == "$" {
			if splitLine[1] == "cd" {
				cwd = cd(cwd, &rootDir, splitLine[2])
			} else if splitLine[1] == "ls" {
				// Skip for now
			}
		} else if splitLine[0] == "dir" {
			// New directory
			dirName := splitLine[1]
			newDir := Dir{}
			newDir.Name = dirName
			newDir.Path = filepath.Join(cwd.Path, dirName)
			newDir.Previous = cwd
			cwd.Dirs = append(cwd.Dirs, &newDir)
		} else {
			// New file
			newFile := File{}
			fileSize, _ := strconv.Atoi(splitLine[0])
			fileName := splitLine[1]
			newFile.Name = fileName
			newFile.Size = fileSize
			newFile.Path = filepath.Join(cwd.Path, fileName)
			newFile.Dir = cwd
			cwd.Files = append(cwd.Files, newFile)
		}
	}

	dirSize(&rootDir)
	return rootDir
}

func printTree(dir *Dir, depth int) {
	fmt.Printf("%s- %s (dir, size=%d)\n", strings.Repeat("  ", depth), dir.Name, dir.Size)
	for _, file := range dir.Files {
		fmt.Printf("%s- %s (file, size=%d)\n", strings.Repeat("  ", depth+1), file.Name, file.Size)

	}

	for _, dir := range dir.Dirs {
		printTree(dir, depth+1)
	}
}

func part1(dir *Dir, maxSize int) int {

	totalSize := 0

	if dir.Size < maxSize {
		totalSize += dir.Size
	}

	for _, subdir := range dir.Dirs {
		totalSize += part1(subdir, maxSize)
	}

	return totalSize
}

func flattenAndFilter(dir *Dir, minSize int) []*Dir {
	ret := []*Dir{}
	for _, subDir := range dir.Dirs {
		if subDir.Size >= minSize {
			ret = append(ret, subDir)
		}
		ret = append(ret, flattenAndFilter(subDir, minSize)...)
	}

	return ret
}

func part2(dir *Dir, maxSize int, sizeRequired int) int {
	freeSpace := maxSize - dir.Size
	toFree := sizeRequired - freeSpace
	fmt.Printf("Used: %d, Free: %d, Needed: %d\n", dir.Size, freeSpace, toFree)
	flatDirs := flattenAndFilter(dir, toFree)

	min := maxSize
	for _, fDir := range flatDirs {
		if fDir.Size < min {
			min = fDir.Size
		}
	}

	return min
}

func main() {
	input := readInput("./input.txt")
	dirTree := parseInput(input)
	printTree(&dirTree, 0)
	fmt.Printf("Part 1: %d\n", part1(&dirTree, 100000))
	fmt.Printf("Part 2: %d\n", part2(&dirTree, 70000000, 30000000))
}
