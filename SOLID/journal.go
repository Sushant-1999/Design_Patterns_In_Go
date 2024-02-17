package main

import (
	"fmt"
)

//For Single Responsibility Principle, our package must have single primary responsibility whatever struct we are using

// Creating a Journal which has single responsibility of Maintaining the entries within
var entryCount = 0

type Journal struct {
	entries []string
}

// creating methods on Journal for adding / removing the entries from the Journal
func (j *Journal) addEntries(name string) {
	entryCount++
	j.entries = append(j.entries, name)
}

func (j *Journal) removeEntries(index int) {
	entryCount--
	j.entries = append(j.entries[:index], j.entries[index+1:]...)
}

func main() {
	school := Journal{}
	school.addEntries("Exam")
	school.addEntries("Class")
	school.addEntries("Games")
	school.removeEntries(2)
	fmt.Println(school, entryCount)
}
