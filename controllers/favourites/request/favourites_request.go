package request

import "injar/usecase/favourites"

type Favourites struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	WebinarID int `json:"webinar_id"`
}

func (req *Favourites) ToDomain() *favourites.Domain {
	return &favourites.Domain{
		UserID:    req.UserID,
		WebinarID: req.WebinarID,
	}
}
