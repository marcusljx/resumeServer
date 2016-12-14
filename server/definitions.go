package server

type ResumeObject struct {
	Name         string `json:"name"`
	CurrentJob   Job    `json:"currentJob"`
	PreviousJobs []Job  `json:"previousJobs"`
}

type Job struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
}
