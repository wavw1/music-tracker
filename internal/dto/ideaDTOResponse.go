package dto

type IdeaDTOResponse struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Bpm    int      `json:"bpm"`
	Key    string   `json:"key"`
	Status string   `json:"status"`
	Tags   []string `json:"tags"`
}
