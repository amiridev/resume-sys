package documents

type DocumentCreateDto struct {
	Title           string `json:"title" binding:"required,min=2,max=250"`
	Category        string `binding:"required"`
	Ratings         int    `binding:"gte=0,lte=5"`
	Description     string `json:"description" binding:"required"`
	Status          string `json:"status" binding:"required"`
	PublicationDate string `binding:"required,datetime=2006-01-02"`
}

type DocumentUpdateDto struct {
	Title           string `json:"title" binding:"required,min=2,max=250"`
	Category        string `binding:"required"`
	Ratings         int    `binding:"gte=0,lte=5"`
	Description     string `json:"description" binding:"required"`
	Status          string `json:"status" binding:"required"`
	PublicationDate string `binding:"required,datetime=2006-01-02"`
}
