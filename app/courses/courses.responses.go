package courses

import "resume-sys/database/models"

type ResponseType map[string]any

type Course struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

//one course response

func CourseTransform(course models.Course) Course {
	return Course{
		ID:          course.ID,
		UserID:      course.UserID,
		Title:       course.Title,
		Category:    course.Category,
		Description: course.Description,
		Status:      course.Status,
	}
}

type CourseResponseType struct {
	Course Course `json:"course"`
}

func CourseResponse(course models.Course) ResponseType {
	return ResponseType{
		"course": CourseTransform(course),
	}
}

//list of course response

func CoursesTransform(courses []models.Course) []Course {
	var data = []Course{}

	for _, course := range courses {
		data = append(data, CourseTransform(course))
	}

	return data
}

type CoursesResponseType struct {
	Courses []Course `json:"courses"`
}

func CoursesResponse(courses []models.Course) ResponseType {
	return ResponseType{
		"courses": CoursesTransform(courses),
	}
}
