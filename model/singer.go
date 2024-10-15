package model

type SingerID int

type Singer struct {
	ID   SingerID `json:"id"`
	Name string   `json:"name"`
}

func (s *Singer) Validate() error {
	if s.Name == "" {
		return ErrInvalidParam
	}
	if len(s.Name) > 255 {
		return ErrInvalidParam
	}
	return nil
}
