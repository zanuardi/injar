package request

import "injar/usecase/transactions"

type Transactions struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	WebinarID int    `json:"webinar_id"`
	Status    string `json:"status"`
}

func (req *Transactions) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		UserID:    req.UserID,
		WebinarID: req.WebinarID,
		Status:    "pending",
	}
}
