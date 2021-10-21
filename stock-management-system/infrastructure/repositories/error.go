package repositories

type RepositoryError string

func (e RepositoryError) Error() string {
	return string(e)
}

const (
	AlreadyExistsErr = RepositoryError("Can't make user because the email was already used")
	UserNotFoundErr  = RepositoryError("There are no corresond user")
)
