package main

type (
	Workspace struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	Account struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	User struct {
		ID             int    `json:"id,omitempty"`
		Email          string `json:"email"`
		Name           string `json:"name"`
		ForeignID      string `json:"foreign_id,omitempty"`
		SendInvitation bool   `json:"send_invitation,omitempty"`
	}

	Client struct {
		ID        int    `json:"id,omitempty"`
		Name      string `json:"name"`
		ForeignID string `json:"foreign_id,omitempty"`
	}

	Project struct {
		ID       int    `json:"id,omitempty"`
		Name     string `json:"name,omitempty"`
		Active   bool   `json:"active,omitempty"`
		Billable bool   `json:"billable,omitempty"`
		ClientID int    `json:"cid,omitempty"`

		ForeignID       string `json:"foreign_id,omitempty"`
		foreignClientID int
	}

	Task struct {
		ID               int    `json:"id,omitempty"`
		Name             string `json:"name"`
		Active           bool   `json:"active"`
		ForeignID        string `json:"foreign_id,omitempty"`
		ProjectID        int    `json:"pid"`
		foreignProjectID int
	}

	TimeEntry struct {
		ID                int    `json:"id"`
		ProjectID         int    `json:"pid,omitempty"`
		TaskID            int    `json:"tid,omitempty"`
		UserID            int    `json:"uid,omitempty"`
		Billable          bool   `json:"billable"`
		Start             string `json:"start"`
		Stop              string `json:"stop,omitempty"`
		DurationInSeconds int    `json:"duration"`
		Description       string `json:"description,omitempty"`
		foreignID         int
		foreignTaskID     int
		foreignUserID     int
		foreignProjectID  int
	}

	AccountsResponse struct {
		Error    string     `json:"error"`
		Accounts []*Account `json:"accounts"`
	}

	UsersResponse struct {
		Error string  `json:"error"`
		Users []*User `json:"users"`
	}

	ClientsResponse struct {
		Error   string    `json:"error"`
		Clients []*Client `json:"clients"`
	}

	ProjectsResponse struct {
		Error    string     `json:"error"`
		Projects []*Project `json:"projects"`
	}

	TasksResponse struct {
		Error string  `json:"error"`
		Tasks []*Task `json:"tasks"`
	}
)
