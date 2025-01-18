package experiences

import "resume-sys/database/models"

type ResponseType map[string]any

type Experience struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Title     string `json:"title"`
	Company   string `json:"company"`
	Status    string `json:"status"`
	StartedAt string `json:"started_at"`
	EndedAt   string `json:"ended_at"`
}

// one experience response

func ExperienceTransform(experience models.Experience) Experience {
	return Experience{
		ID:        experience.ID,
		UserID:    experience.UserID,
		Title:     experience.Title,
		Company:   experience.Company,
		Status:    experience.Status,
		StartedAt: experience.StartedAt,
		EndedAt:   experience.EndedAt,
	}
}

type ExperienceResponseType struct {
	Experience Experience `json:"experience"`
}

func ExperienceResponse(exp models.Experience) ResponseType {
	return ResponseType{
		"experience": ExperienceTransform(exp),
	}
}

// list of experience response

func ExperiencesTransform(experiences []models.Experience) []Experience {
	var data = []Experience{}

	for _, exp := range experiences {
		data = append(data, ExperienceTransform(exp))
	}

	return data
}

type ExperiencesResponseType struct {
	Experiences []Experience `json:"experiences"`
}

func ExperiencesResponse(experiences []models.Experience) ResponseType {
	return ResponseType{
		"experiences": ExperiencesTransform(experiences),
	}
}
