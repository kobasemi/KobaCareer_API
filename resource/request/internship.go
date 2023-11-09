package request

type CreateInternship struct {
	ID          uint   `json:"company_id"`
	Company     string `json:"company"`
	Title       string `json:"title"`
	Salary      uint   `json:"salary"`
	Period      string `json:"period"`
	Select      string `json:"select"`
	Deadline    string `json:"deadline"`
	Contributor string `json:"contributor"`
	Detail      string `json:"detail"`
	FutureJob   string `json:"future_job"`
	Flow        string `json:"flow"`
	Method      string `json:"method"`
}

type UpdateInternship struct {
	ID          uint   `json:"id"`
	Company     string `json:"company"`
	Title       string `json:"title"`
	Salary      uint   `json:"salary"`
	Period      string `json:"period"`
	Select      string `json:"select"`
	Deadline    string `json:"deadline"`
	Contributor string `json:"contributor"`
	Detail      string `json:"detail"`
	FutureJob   string `json:"future_job"`
	Flow        string `json:"flow"`
	Method      string `json:"method"`
}
