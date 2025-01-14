package projects

import "resume-sys/database/models"

type ResponseType map[string]any

type Project struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	ImageUrl    string `json:"image_url"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// one project response

func ProjectTransform(project models.Project) Project {
	return Project{
		ID:          project.ID,
		UserID:      project.UserID,
		Title:       project.Title,
		ImageUrl:    project.ImageUrl,
		Link:        project.Link,
		Description: project.Description,
		Status:      project.Status,
	}
}

type ProjectResponseType struct {
	Project Project `json:"project"`
}

func ProjectResponse(project models.Project) ResponseType {
	return ResponseType{
		"project": ProjectTransform(project),
	}
}

// list of project response

func ProjectsTransform(projects []models.Project) []Project {
	var data = []Project{}

	for _, project := range projects {
		data = append(data, ProjectTransform(project))
	}

	return data
}

type ProjectsResponseType struct {
	Projects []Project `json:"projects"`
}

func ProjectsResponse(projects []models.Project) ResponseType {
	return ResponseType{
		"projects": ProjectsTransform(projects),
	}
}
