package worker

type IamDetails struct {
	ProjectID string `json:"project_id"`
	User      string `json:"user"`
	Role      string `json:"role"`
}
