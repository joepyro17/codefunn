package models

type Tutorials struct {
	Tutorials []Tutorial `json:"tutorials"`
}

type Tutorial struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Subtitle string `json:"subtitle"`
	ImgURL string `json:"img_url"`
}