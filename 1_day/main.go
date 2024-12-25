package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {
    fmt.Println("Part1: ", solutionPart1())
    fmt.Println("Part2: ", solutionPart2())
}

func solutionPart1() int {
    slice1, slice2 := fileToArrays()
    mergeSort(slice1, 0, len(slice1) - 1)
    mergeSort(slice2, 0, len(slice2) - 1)

    distance := 0

    for i := 0; i < len(slice1); i++ {
        distance += Abs(slice1[i] - slice2[i])
    }

    return distance

}

func solutionPart2() int {
    slice1, slice2 := fileToArrays()

    slice2Count := make(map[int]int)
    
    for i := 0; i < len(slice2); i++ {
        slice2Count[slice2[i]]++
    }

    distance := 0

    for i := 0; i < len(slice1); i++ {
        distance += slice1[i] * slice2Count[slice1[i]]
    }

    return distance
}


func fileToArrays() ([]int, []int) {
    content, err := os.Open("input.txt")

    if err != nil {
        panic(err)
    }

    defer content.Close()

    var slice1 []int
    var slice2 []int
    
    scanner := bufio.NewScanner(content)
    scanner.Split(bufio.ScanWords)

    counter := 0
    for scanner.Scan() {
        i, err := strconv.Atoi(scanner.Text())
        if err != nil {
            panic(err)
        }

        if counter % 2 == 0 {
            slice1 = append(slice1, i)
        } else {
            slice2 = append(slice2, i)
        }

        counter++
    }
    return slice1, slice2

}

func mergeSort(s []int, beg int, end int) {
    if beg < end {
        mid := (beg + end) / 2
        mergeSort(s, beg, mid)
        mergeSort(s, mid + 1, end)
        merge(s, beg, mid, end) 
    }
}

func merge(s []int, beg int, mid int, end int) {
    i := 0
    j := 0
    k := beg 
    n1 := mid - beg + 1
    n2 := end - mid

    leftSlice := make([]int, n1)
    rightSlice := make([]int, n2)

    copy(leftSlice, s[beg:mid+1])
    copy(rightSlice, s[mid+1:end+1])

    for i < n1 && j < n2 {
        if leftSlice[i] < rightSlice[j] {
            s[k] = leftSlice[i]
            i++
        } else {
            s[k] = rightSlice[j]
            j++
        }
        k++
    }

    for i < n1 {
        s[k] = leftSlice[i]
        i++
        k++ 
    }

    for j < n2 {
        s[k] = rightSlice[j]
        j++
        k++ 
    }
}


func Abs(x int) int{
	if x < 0 {
		return -x
	}
	return x
}
