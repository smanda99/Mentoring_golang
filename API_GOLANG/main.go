package main

// model for courses
type Author struct {
	Id      int    `json:"id"`
	Website string `json:"website"`
}

type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
}

func main() {
	

}
