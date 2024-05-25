package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseName  string  `json:"courseName"`
	CourseId    string  `json:"courseId"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:fullName`
	Email    string `json:email`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

	fmt.Println("Welcome to Clint server")
	r := mux.NewRouter()
	courses = append(courses, Course{CourseName: "Introduction to Programming",
		CourseId:    "101",
		CoursePrice: 499,
		Author: &Author{FullName: "John Doe",
			Email: "john.doe@example.com"}})
	courses = append(courses, Course{CourseName: "Web Development Fundamentals",
		CourseId:    "102",
		CoursePrice: 999,
		Author: &Author{FullName: "Jane Smith",
			Email: "jane.smith@example.com"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	//Listen to port 4000
	log.Fatal(http.ListenAndServe(":4000", r))

}

// Controller - file
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Clint's Server! :)</h2>"))
	json.NewEncoder(w).Encode("Welcome to Clint's Server! :)")
}

// GET Request
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// GET Request
func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Course by id")

	// grab id from request url
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found!")
}

// POST Request
func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Course")
	w.Header().Set("content-type", "application/json")

	//if response data is empty or not Valid
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send valid data")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside Json")
		return
	}

	for _, courseItem := range courses {
		if courseItem.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course already exist!")
			return
		}
	}

	rand.NewSource(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

// PUT Request
func updateCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Update Course")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

// DELETE Request
func deleteCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Delete Course")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("The course has been removed")
			return
		}
	}
	json.NewEncoder(w).Encode("No course Found")
}
