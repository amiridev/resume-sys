package skills

import "resume-sys/database/models"

type ResponseType map[string]any

type Skill struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Level       string `json:"level"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
}

// one skill response

func SkillTransform(skill models.Skill) Skill {
	return Skill{
		ID:          skill.ID,
		UserID:      skill.UserID,
		Title:       skill.Title,
		Level:       skill.Level,
		ImageUrl:    skill.ImageUrl,
		Description: skill.Description,
	}
}

type SkillResponseType struct {
	Skill Skill `json:"skill"`
}

func SkillResponse(skill models.Skill) ResponseType {
	return ResponseType{
		"skill": SkillTransform(skill),
	}
}

// list of skill response

func SkillsTransform(skills []models.Skill) []Skill {
	var data = []Skill{}

	for _, skill := range skills {
		data = append(data, SkillTransform(skill))
	}

	return data
}

type SkillsResponseType struct {
	Skills []Skill `json:"skills"`
}

func SkillsResponse(skills []models.Skill) ResponseType {
	return ResponseType{
		"skills": SkillsTransform(skills),
	}
}
