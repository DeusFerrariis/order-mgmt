package internal

type (
	CustomerData struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	CustomerRecord struct {
		CustomerData
		Id int64 `json:"id"`
	}
)
