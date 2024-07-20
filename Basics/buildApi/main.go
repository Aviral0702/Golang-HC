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
	CourseId    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrice string `json:"price"`
	Author      Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

var mycourses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	//seeding data
	mycourses = append(mycourses, Course{CourseId: "1", CourseName: "Java", CoursePrice: "1000", Author: Author{FullName: "John Doe", Website: "www.johndoe.com"}})
	mycourses = append(mycourses, Course{CourseId: "2", CourseName: "Python", CoursePrice: "2000", Author: Author{FullName: "Jane Doe", Website: "www.janedoe.com"}})
	mycourses = append(mycourses, Course{CourseId: "3", CourseName: "Golang", CoursePrice: "3000", Author: Author{FullName: "John Doe", Website: "www.johndoe.com"}})

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/course", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", addOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the home page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all course")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mycourses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "applicaton/json")

	//grab id from request
	params := mux.Vars(r)

	for _, course := range mycourses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")

}
func addOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding one course here")
	w.Header().Set("Content-Type", "application/json")

	//if the response is nil then:
	if r.Body == nil {
		json.NewEncoder(w).Encode("Empty body has been passed")
	}

	var newCourse Course
	//if the given array is empty
	_ = json.NewDecoder(r.Body).Decode(&newCourse)

	if newCourse.IsEmpty() {
		json.NewEncoder(w).Encode("Course list is empty")
	}

	//check only if title is duplicate
	for _, course := range mycourses {
		if course.CourseName == newCourse.CourseName {
			json.NewEncoder(w).Encode("Course already exists")
			return
		}

	}

	//generate unique id for every course

	rand.Seed(time.Now().UnixNano())
	newCourse.CourseId = strconv.Itoa(rand.Intn(100))
	mycourses = append(mycourses, newCourse)
	json.NewEncoder(w).Encode(newCourse)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	//grab the id of the course to be updated
	params := mux.Vars(r)

	for index, course := range mycourses {
		if course.CourseId == params["id"] {
			mycourses = append(mycourses[:index], mycourses[index+1:]...)
			var updatedCourse Course
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse)
			mycourses = append(mycourses, updatedCourse)
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}

	//if no id is found
	json.NewEncoder(w).Encode("Found no course with the given Id...")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting a course")
	w.Header().Set("Content-Type", "application/json")

	// Grab the id of the course
	params := mux.Vars(r)
	courseID := params["id"]

	// Check if the confirmation parameter is present
	confirmDelete := r.URL.Query().Get("confirm")

	if confirmDelete != "true" {
		// If not confirmed, send a confirmation request
		json.NewEncoder(w).Encode(map[string]string{
			"message":  "Are you sure you want to delete this course? Add '?confirm=true' to the URL to confirm.",
			"courseID": courseID,
		})
		return
	}

	// If confirmed, proceed with deletion
	for index, course := range mycourses {
		if course.CourseId == courseID {
			mycourses = append(mycourses[:index], mycourses[index+1:]...)
			json.NewEncoder(w).Encode(map[string]string{
				"message":  "Course deleted successfully",
				"courseID": courseID,
			})
			return
		}
	}

	// If no course found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "No course found with the given id",
		"courseID": courseID,
	})
}
