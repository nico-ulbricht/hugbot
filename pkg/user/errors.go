package user

import (
	"fmt"

	"github.com/google/uuid"
)

type ErrNotFound struct {
	ExternalID string
	UserID     uuid.UUID
}

func (err ErrNotFound) Error() string {
	nullID := uuid.UUID{}
	if err.ExternalID != "" {
		return fmt.Sprintf("unable to find user with external id %s", err.ExternalID)
	} else if err.UserID != nullID {
		return fmt.Sprintf("unable to find user with id %s", err.UserID)
	}

	return "unable to find user"
}
