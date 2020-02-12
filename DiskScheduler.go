// “I [Andres Felipe Santamaria] ([an190282]) affirm that
// this program is entirely my own work and that I have neither developed my code with any
// another person, nor copied any code from any other person, nor permitted my code to be copied
// or otherwise used by any other person, nor have I copied, modified, or otherwise used programs
// created by others. I acknowledge that any violation of the above terms will be treated as
// academic dishonesty.”

package main

import (
    "bufio"
    "fmt"
    "strings"
    "strconv"
    "os"
    "sort"
)

// User Defined Struct to hold the data read in from the file.
// This struct will serve as the hub for the data used to drive
// the algorithms
type disk_scheduler struct {
    method string
    upper int
    lower int
    init int
    cylreqs []int
}

type node struct {
    distance int
    checked bool
}

// Abs returns the absolute value of x.
func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

// Simple Check function to capture errors when performing certain operations
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func newNode(distance int) node {
    ret := node{}
    ret.distance = 0
    ret.checked = false
    return ret
}

func calculateDifference(reqs []int, head int, diffs []node) []node {
    for i := 0; i < len(diffs); i++ {
        diffs[i].distance = Abs(reqs[i] - head)
    }

    return diffs
}

func findMin(diffs []node) int{
    var index int = -1
    var min int = 1000000

    for i := 0; i < len(diffs); i++ {
        if !diffs[i].checked && min > diffs[i].distance {
            min = diffs[i].distance
            index = i
        }
    }

    return index
}

// Function to check whether or not a string is in a slice. Under normal circumstances,
// this return value is not relevant. However, if the program is given in invalid format,
// this function in conjunction with the panic function will halt the execution of the program
// instead of crashing and burning
func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}


func FCFS(data disk_scheduler) {

    var tc int = 0
    var diff int = 0
    fmt.Printf("Seek algorithm: FCFS\n")
    fmt.Printf("\tLower cylinder: %5d\n", data.lower)
    fmt.Printf("\tUpper cylinder: %5d\n", data.upper)
    fmt.Printf("\tInit cylinder:  %5d\n", data.init)
    fmt.Printf("\tCylinder requests:\n")

    for _, e := range data.cylreqs {
        fmt.Printf("\t\tCylinder %5d\n", e)
    }
    diff = Abs(data.init - data.cylreqs[0])
    tc += diff


    for i, e := range data.cylreqs {
        fmt.Printf("Servicing %5d\n", e)
        if i + 1 == len(data.cylreqs) {
            break
        }
        
        diff = Abs(data.cylreqs[i] - data.cylreqs[i+1])
        tc += diff
    }

    fmt.Printf("FCFS traversal count = %5d\n", tc)
}

func SSTF(data disk_scheduler) {

    copy := data
    var distance []node


    fmt.Printf("Seek algorithm: SSTF\n")
    fmt.Printf("\tLower cylinder: %5d\n", data.lower)
    fmt.Printf("\tUpper cylinder: %5d\n", data.upper)
    fmt.Printf("\tInit cylinder:  %5d\n", data.init)
    fmt.Printf("\tCylinder requests:\n")

    for i := 0; i < len(copy.cylreqs); i++ {
        fmt.Printf("\t\tCylinder %5d\n", copy.cylreqs[i])
    }

    for i := 0; i < len(copy.cylreqs); i++ {
        distance = append(distance, newNode(0))
    }

    current := copy.init
    var seek_count int = 0
    var sequence [100]int

    for i := 0; i < len(copy.cylreqs); i++ {
        sequence[i] = current
        distance = calculateDifference(copy.cylreqs, current, distance)
        var index int = findMin(distance)
        distance[index].checked = true
        seek_count = seek_count + distance[index].distance

        current = copy.cylreqs[index]
    }

    sequence[len(copy.cylreqs)] = current

    for i := 1; i < len(copy.cylreqs) + 1; i++ {
        fmt.Printf("Servicing %5d\n", sequence[i])
    }

    fmt.Printf("SSTF traversal count = %5d\n", seek_count)

}

func SCAN(data disk_scheduler) {

    copy := data
    var tc int = 0
    var diff int = 0
    var disk_head int = 0

    fmt.Printf("Seek algorithm: SCAN\n")
    fmt.Printf("\tLower cylinder: %5d\n", data.lower)
    fmt.Printf("\tUpper cylinder: %5d\n", data.upper)
    fmt.Printf("\tInit cylinder:  %5d\n", data.init)
    fmt.Printf("\tCylinder requests:\n")

    for _, e := range data.cylreqs {
        fmt.Printf("\t\tCylinder %5d\n", e)
    }

    copy.cylreqs = append(copy.cylreqs, copy.init)


    sort.Slice(copy.cylreqs, func(i, j int) bool {
        return copy.cylreqs[i] < copy.cylreqs[j]
    })

    for i := 0; i < len(copy.cylreqs); i++ {
        if copy.init == copy.cylreqs[i] {
            disk_head = i
            break
        }
    }

    for i := disk_head; i < len(copy.cylreqs) - 1; i++ {
        diff = Abs(copy.cylreqs[i] - copy.cylreqs[i + 1])
        tc = tc + diff
        fmt.Printf("Servicing %5d\n", copy.cylreqs[i+1])
    }
    
    if disk_head != 0 {
        diff = Abs(copy.cylreqs[len(copy.cylreqs) - 1] - copy.upper)
        tc = tc + diff
        diff = Abs(copy.cylreqs[0] - copy.upper)
        tc = tc + diff

        for i := disk_head - 1; i >= 0; i-- {
            fmt.Printf("Servicing %5d\n", copy.cylreqs[i])
        }   
    }

    fmt.Printf("SCAN traversal count = %5d\n", tc)


}

func CSCAN(data disk_scheduler) {
    copy := data
    var tc int = 0
    var diff int = 0
    var disk_head int = 0

    fmt.Printf("Seek algorithm: C-SCAN\n")
    fmt.Printf("\tLower cylinder: %5d\n", data.lower)
    fmt.Printf("\tUpper cylinder: %5d\n", data.upper)
    fmt.Printf("\tInit cylinder:  %5d\n", data.init)
    fmt.Printf("\tCylinder requests:\n")

    for _, e := range data.cylreqs {
        fmt.Printf("\t\tCylinder %5d\n", e)
    }

    copy.cylreqs = append(copy.cylreqs, copy.init)


    sort.Slice(copy.cylreqs, func(i, j int) bool {
        return copy.cylreqs[i] < copy.cylreqs[j]
    })

    for i := 0; i < len(copy.cylreqs); i++ {
        if copy.init == copy.cylreqs[i] {
            disk_head = i
            break
        }
    }

    for i := disk_head; i < len(copy.cylreqs) - 1; i++ {
        diff = Abs(copy.cylreqs[i] - copy.cylreqs[i + 1])
        tc = tc + diff
        fmt.Printf("Servicing %5d\n", copy.cylreqs[i+1])
    }

    if disk_head != 0 {
        diff = Abs(copy.cylreqs[len(copy.cylreqs) - 1] - copy.upper)
        tc = tc + diff
        tc = tc + copy.upper
        tc = tc + Abs(copy.cylreqs[disk_head - 1])

        for i := 0; i < disk_head; i++ {
            fmt.Printf("Servicing %5d\n", copy.cylreqs[i])
        }   
    }

    fmt.Printf("C-SCAN traversal count = %5d\n", tc)
}

func LOOK(data disk_scheduler) {
    copy := data
    var tc int = 0
    var diff int = 0
    var disk_head int = 0

    fmt.Printf("Seek algorithm: LOOK\n")
    fmt.Printf("\tLower cylinder: %5d\n", data.lower)
    fmt.Printf("\tUpper cylinder: %5d\n", data.upper)
    fmt.Printf("\tInit cylinder:  %5d\n", data.init)
    fmt.Printf("\tCylinder requests:\n")

    for _, e := range data.cylreqs {
        fmt.Printf("\t\tCylinder %5d\n", e)
    }

    copy.cylreqs = append(copy.cylreqs, copy.init)


    sort.Slice(copy.cylreqs, func(i, j int) bool {
        return copy.cylreqs[i] < copy.cylreqs[j]
    })

    for i := 0; i < len(copy.cylreqs); i++ {
        if copy.init == copy.cylreqs[i] {
            disk_head = i
            break
        }
    }

    for i := disk_head; i < len(copy.cylreqs) - 1; i++ {
        diff = Abs(copy.cylreqs[i] - copy.cylreqs[i + 1])
        tc = tc + diff
        fmt.Printf("Servicing %5d\n", copy.cylreqs[i+1])
    }

    if disk_head != 0 {
        diff = Abs(copy.cylreqs[len(copy.cylreqs) - 1] - copy.cylreqs[0])
        tc = tc + diff

        for i := disk_head - 1; i >= 0; i-- {
            fmt.Printf("Servicing %5d\n", copy.cylreqs[i])
        }     
    }

    fmt.Printf("LOOK traversal count = %5d\n", tc)
} 

func CLOOK(data disk_scheduler) {
    copy := data
    var tc int = 0
    var diff int = 0
    var disk_head int = 0

    fmt.Printf("Seek algorithm: C-LOOK\n")
    fmt.Printf("\tLower cylinder: %5d\n", data.lower)
    fmt.Printf("\tUpper cylinder: %5d\n", data.upper)
    fmt.Printf("\tInit cylinder:  %5d\n", data.init)
    fmt.Printf("\tCylinder requests:\n")

    for _, e := range data.cylreqs {
        fmt.Printf("\t\tCylinder %5d\n", e)
    }

    copy.cylreqs = append(copy.cylreqs, copy.init)
    sort.Slice(copy.cylreqs, func(i, j int) bool {
        return copy.cylreqs[i] < copy.cylreqs[j]
    })

    for i := 0; i < len(copy.cylreqs); i++ {
        if copy.init == copy.cylreqs[i] {
            disk_head = i
            break
        }
    }

    for i := disk_head; i < len(copy.cylreqs) - 1; i++ {
        diff = Abs(copy.cylreqs[i] - copy.cylreqs[i + 1])
        tc = tc + diff
        fmt.Printf("Servicing %5d\n", copy.cylreqs[i+1])
    }

    if disk_head != 0 {

        diff = Abs(copy.cylreqs[len(copy.cylreqs) - 1] - copy.cylreqs[0])
        tc = tc + diff
        diff = Abs(copy.cylreqs[0] - copy.cylreqs[disk_head - 1])
        tc = tc + diff

        for i := 0; i < disk_head; i++ {
            fmt.Printf("Servicing %5d\n", copy.cylreqs[i])
        }    
    }

    fmt.Printf("C-LOOK traversal count = %5d\n", tc)
}


func main() {

	args := os.Args

    var methods []string
    methods = append(methods, "fcfs")
    methods = append(methods, "sstf")
    methods = append(methods, "scan")
    methods = append(methods, "c-scan")
    methods = append(methods, "look")
    methods = append(methods, "c-look")

	fin, err := os.Open(args[1])
    check(err)

    sc := bufio.NewScanner(fin)
    sc.Split(bufio.ScanLines)

    var lines []string
    var scheduler_data disk_scheduler

    for sc.Scan() {
    	lines = append(lines, sc.Text())
    }

    for index, element := range lines {
        currentLine := bufio.NewScanner(strings.NewReader(element))
        currentLine.Split(bufio.ScanWords)
        // Index tells us which line of text we are looking at.
        // The formatting includes import information about the state 
        // of the disk and which algorithm to use in the first several lines.
        // Since this is known, we can parse out this information based on the
        // first 4 indices.
        var count int = 0
        var temp int = 0
        // keep index in current line.
        // anything past 0,1 is a comment and should be disregarded
        for currentLine.Scan() {
            
            if count > 1 {
                break
            }
            if index == 0 && count == 0 {
                // continue to relevant info
                count = count + 1
                continue
            }
            // Find algorithm to use from first line
            if index == 0 && count == 1 {
                // make sure to capture this value
                if stringInSlice(currentLine.Text(), methods) == false {
                    panic("ERROR: Invalid algorithm input.")
                }
                scheduler_data.method = currentLine.Text();
                count = count + 1
            }
            // pass over lowerCYL string
            if index == 1 && count == 0 {
                count = count + 1
                continue
            }
            // capture lowerCYL value
            if index == 1 && count == 1 {
                scheduler_data.lower, err = strconv.Atoi(currentLine.Text())
                check(err)
                count = count + 1
            }
            // pass over upperCYL string
            if index == 2 && count == 0 {
                count = count + 1
                continue
            }
            // capture upperCYL value
            if index == 2 && count == 1 {
                // make sure to capture this value
                scheduler_data.upper, err = strconv.Atoi(currentLine.Text())
                check(err)
                count = count + 1
            }

            // pass over initCYL string
            if index == 3 && count == 0 {
                count = count + 1
                continue
            }

            // capture initCYL value
            if index == 3 && count == 1 {
                // make sure to capture this value
                scheduler_data.init, err = strconv.Atoi(currentLine.Text())
                check(err)
                count = count + 1
            }

            // all other lines until "end" contain information
            // about cylreqs
            // amount of cylreqs = len(lines) - 5
            // the first four lines contain information about
            // the algorithm and state of the disk
            // the last line contains the end keyword

            // skip over cylreq string
            if (index > 3 && index < len(lines) - 1) && count == 0 {
                count = count + 1
                continue
            }
            // capture cylreq information
            if (index > 3 && index < len(lines) - 1) && count == 1 {
                temp, err = strconv.Atoi(currentLine.Text())
                check(err)
                scheduler_data.cylreqs = append(scheduler_data.cylreqs, temp)

            }
        }
    }

    fin.Close()

    // once the appropriate values have been captured, run the methodd
    // for the appropriate algorithm.
    if strings.Compare(scheduler_data.method, "fcfs") == 0 {
        FCFS(scheduler_data)
    }

    if strings.Compare(scheduler_data.method, "sstf") == 0 {
        SSTF(scheduler_data)
    }

    if strings.Compare(scheduler_data.method, "scan") == 0 {
        SCAN(scheduler_data)
    }

    if strings.Compare(scheduler_data.method, "c-scan") == 0 {
        CSCAN(scheduler_data)
    }

    if strings.Compare(scheduler_data.method, "look") == 0 {
        LOOK(scheduler_data)
    }

    if strings.Compare(scheduler_data.method, "c-look") == 0 {
        CLOOK(scheduler_data)
    }
}