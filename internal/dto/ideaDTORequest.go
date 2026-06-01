package dto

type IdeaDTORequest struct {
	Title string   `json:"title" binding:"required"`
	Bpm   int      `json:"bpm" binding:"required"`
	Key   string   `json:"key"`
	Tags  []string `json:"tags"`
}
