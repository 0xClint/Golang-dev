package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseName  string  `json: courseName`
	CourseId    string  `json: courseId`
	CoursePrice int     `json:price`
	Author      *Author `json:author`
}

type Author struct {
	FullName string `json:fullName`
	Website  string `json:website`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

// Controller - file

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Clint's Server! :)</h2>"))
}

// POST Request
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// GET Request
func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Course by id")

	// grab id from request
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

	//if response data isempty or not Valid
	if r.Body != nil {
		json.NewEncoder(w).Encode("Please send valid data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside Json")
	}

	rand.NewSource(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}
