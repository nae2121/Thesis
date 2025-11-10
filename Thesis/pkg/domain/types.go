package domain


type Chunk struct {
	Text string `json:"text"`
	Page int `json:"page"`
	Section string `json:"section"`
	Heading string `json:"heading"`
}