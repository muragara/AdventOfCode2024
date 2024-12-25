package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
)

func main() {
    fmt.Println("Part1: ", solutionPart1())
}


func solutionPart1() int {
    slice := fileToArrays()
   
    safeCount := 0

    for _, line := range slice {
        desc := true
        safe := true

        if line[0] < line[1] {
            desc = false
        } else if line[0] == line[1] {
            continue
        }

        for i := 0; i < len(line) - 1; i++ {
            dist := Abs(line[i] - line[i+1]) 
            if line[i] < line[i+1] && desc {
                safe = false
                break
            } else if line[i] > line[i+1] && !desc {
                safe = false
                break
            } else if dist < 1 || dist > 3 {
                safe = false
                break
            }
        }

        if safe {
            safeCount++
        }
    }

    return safeCount
    
}


func fileToArrays() [][]int {
    content, err := os.Open("input.txt")

    if err != nil {
        panic(err)
    }

    defer content.Close()

    var slice [][]int
    
    scanner := bufio.NewScanner(content)
    scanner.Split(bufio.ScanLines)

    counter := 0
    for scanner.Scan() {
        stringSlice := strings.Split(scanner.Text(), " ")
        intSlice := make([]int, len(stringSlice))

        for index, str := range stringSlice {
            intSlice[index], _ = strconv.Atoi(str) 
        }

        slice = append(slice, intSlice)
        counter++
    }
    return slice 

}

func Abs(x int) int{
	if x < 0 {
		return -x
	}
	return x
}
