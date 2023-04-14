// Null Object Design Pattern

// WHAT
// WHY
// HOW

package main

import "fmt"

// DependenceBase
type studentName interface {
	getName() string
}

//ConcreteDependency
type student struct {
	name string
}

func getNewStudent(name string) *student {
	return &student{
		name: name,
	}
}

func (s *student) getName() string {
	return s.name
}

// Null Implementation
type nullStudent struct {
}

func newNullStudent() *student {
	return &student{
		name: "Name of student Not Present",
	}
}

// Factory
type studentFactory struct {
	studentNames []string
}

func getNewStudentFactory() *studentFactory {
	return &studentFactory{
		studentNames: []string{"A", "B", "C"},
	}
}

func (sf *studentFactory) getStudentName(name string) studentName {
	studentList := getNewStudentFactory()

	for _, value := range studentList.studentNames {
		if name == value {
			return getNewStudent(name)
		}
	}

	return newNullStudent()
}

// Client
func main() {
	students := getNewStudentFactory()

	s1 := students.getStudentName("A")
	s2 := students.getStudentName("B")
	s3 := students.getStudentName("C")
	s4 := students.getStudentName("D")
	s5 := students.getStudentName("E")

	fmt.Println("Student Name : ", s1.getName())
	fmt.Println("Student Name : ", s2.getName())
	fmt.Println("Student Name : ", s3.getName())
	fmt.Println("Student Name : ", s4.getName())
	fmt.Println("Student Name : ", s5.getName())
}
