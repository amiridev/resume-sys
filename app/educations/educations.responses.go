package educations

import (
	"resume-sys/database/models"
	"time"
)

type ResponseType map[string]any

type Education struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Field       string `json:"field"`
	Description string `json:"description"`
	StartedAt   string `json:"started_at"`
	EndedAt     string `json:"ended_at"`
}

// one education response

func EducationTransform(education models.Education) Education {
	return Education{
		ID:          education.ID,
		UserID:      education.UserID,
		Title:       education.Title,
		Field:       education.Field,
		Description: education.Description,
		StartedAt:   education.StartedAt.Format(time.RFC3339),
		EndedAt:     education.EndedAt.Format(time.RFC3339),
	}
}

type EducationResponseType struct {
	Education Education `json:"education"`
}

func EducationResponse(education models.Education) ResponseType {
	return ResponseType{
		"education": EducationTransform(education),
	}
}

// list of education response

func EducationsTransform(educations []models.Education) []Education {
	var data = []Education{}

	for _, education := range educations {
		data = append(data, EducationTransform(education))
	}

	return data
}

type EducationsResponseType struct {
	Educations []Education `json:"educations"`
}

func EducationsResponse(educations []models.Education) ResponseType {
	return ResponseType{
		"educations": EducationsTransform(educations),
	}
}
