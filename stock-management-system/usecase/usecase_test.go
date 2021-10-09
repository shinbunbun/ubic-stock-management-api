package usecase

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/Yuto/ubic-stock-management-api/stock-management-system/domain"
)

func TestFindUserByID(t *testing.T) {
	t.Run("succesful find", func(t *testing.T) {
		users := &dummyUserRepository{
			users: []domain.User{
				{
					ID:       "1",
					Email:    "hoge@gmail.com",
					Name:     "User1",
					Password: "NoPassword",
				},
			},
		}
		it := Interactor{UserRepository: users}

		want := domain.User{
			Email:    "hoge@gmail.com",
			Name:     "User1",
			Password: "NoPassword",
		}
		got, err := it.FindUserByID("1")
		CheckError(t, err, nil)
		CheckUser(t, got, want)
	})
	t.Run("failed find", func(t *testing.T) {
		users := &dummyUserRepository{
			users: []domain.User{
				{
					ID:       "2",
					Email:    "hoge@gmail.com",
					Name:     "User1",
					Password: "NoPassword",
				},
			},
		}
		it := Interactor{UserRepository: users}

		_, err := it.FindUserByID("1")
		CheckError(t, err, UserNotFoundError)
	})
}

func TestFindUserByEmail(t *testing.T) {
	t.Run("succesful find", func(t *testing.T) {
		users := &dummyUserRepository{
			users: []domain.User{
				{
					ID:       "1",
					Email:    "hoge@gmail.com",
					Name:     "User1",
					Password: "NoPassword",
				},
			},
		}
		it := Interactor{UserRepository: users}

		want := domain.User{
			Email:    "hoge@gmail.com",
			Name:     "User1",
			Password: "NoPassword",
		}
		got, err := it.FindUserByEmail("hoge@gmail.com")
		CheckError(t, err, nil)
		CheckUser(t, got, want)
	})
	t.Run("failed find", func(t *testing.T) {
		users := dummyUserRepository{
			users: []domain.User{
				{
					ID:       "1",
					Email:    "hoge@yahoo.com",
					Name:     "User1",
					Password: "NoPassword",
				},
			},
		}
		it := Interactor{UserRepository: &users}
		_, err := it.FindUserByEmail("hoeg@gmail.com")
		CheckError(t, err, UserNotFoundError)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("succesful delete", func(t *testing.T) {
		users := dummyUserRepository{
			users: []domain.User{
				{
					ID:       "1",
					Email:    "hoge@gmail.com",
					Name:     "User1",
					Password: "NoPassword",
				},
			},
		}
		want := &dummyUserRepository{
			users: []domain.User{},
		}
		it := Interactor{UserRepository: &users}
		err := it.DeleteUserByID("1")
		got := it.UserRepository.(*dummyUserRepository)
		CheckError(t, err, nil)
		CheckUserRepository(t, *got, *want)
	})
	t.Run("failed delete", func(t *testing.T) {
		users := dummyUserRepository{
			users: []domain.User{
				{
					ID:       "2",
					Email:    "hoge@gmail.com",
					Name:     "User1",
					Password: "NoPassword",
				},
			},
		}
		want := &dummyUserRepository{
			users: []domain.User{
				{
					ID:       "2",
					Email:    "hoge@gmail.com",
					Name:     "User1",
					Password: "NoPassword",
				},
			},
		}
		it := Interactor{UserRepository: &users}
		err := it.DeleteUserByID("1")
		got := it.UserRepository.(*dummyUserRepository)
		CheckError(t, err, CantDeleteUserError)
		CheckUserRepository(t, *got, *want)
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("succesful create", func(t *testing.T) {
		users := &dummyUserRepository{
			users: []domain.User{},
		}
		email := "hoge@gmail.com"
		name := "john doe"
		password := "123456"
		want := &dummyUserRepository{
			users: []domain.User{
				{
					ID:       "1",
					Email:    email,
					Name:     name,
					Password: password,
				},
			},
		}
		it := Interactor{UserRepository: users}
		_, err := it.CreateUser(email, name, password)
		got := it.UserRepository.(*dummyUserRepository)

		CheckError(t, err, nil)
		CheckUserRepository(t, *got, *want)
	})
}

type dummyUserRepository struct {
	users []domain.User
}

func SameUserRepository(ur1 dummyUserRepository, ur2 dummyUserRepository) bool {
	if len(ur1.users) != len(ur2.users) {
		return false
	}

	sort.Slice(ur1.users, func(i, j int) bool { return ur1.users[i].Email < ur1.users[j].Email })
	sort.Slice(ur2.users, func(i, j int) bool { return ur2.users[i].Email < ur2.users[j].Email })

	for i, u1 := range ur1.users {
		u2 := ur2.users[i]
		if !SameUser(u1, u2) {
			return false
		}
	}

	return true
}

func CheckUserRepository(t testing.TB, got, want dummyUserRepository) {
	if !SameUserRepository(got, want) {
		t.Errorf("got %q dummyUserRepository,want %q", got, want)
	}

}

func SameUser(u1 domain.User, u2 domain.User) bool {
	if u1.Email != u2.Email {
		return false
	}
	if u1.Name != u2.Name {
		return false
	}
	if u1.Password != u2.Password {
		return false
	}
	return true
}

func CheckUser(t testing.TB, got, want domain.User) {
	t.Helper()
	if !SameUser(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func CheckError(t testing.TB, got, want error) {
	t.Helper()
	if want != nil && got != nil {
		if want.Error() != got.Error() {
			t.Errorf("got %q error, want %q", got, want)
		}
	} else {
		if want != nil {
			t.Fatal("want to got a error")
		}
		if got != nil {
			t.Error("got a error")
		}
	}
}

func (ur *dummyUserRepository) FindByID(id string) (domain.User, error) {
	for _, col := range ur.users {
		if col.ID == id {
			return col, nil
		}
	}
	return domain.User{}, UserNotFoundError
}

func (ur *dummyUserRepository) FindByEmail(email string) (domain.User, error) {
	for _, col := range ur.users {
		if col.Email == email {
			return col, nil
		}
	}
	return domain.User{}, UserNotFoundError
}

func (ur *dummyUserRepository) Delete(id string) error {
	for i, col := range ur.users {
		if col.ID == id {
			next := ur.users[:i]
			next = append(next, ur.users[i+1:]...)
			ur.users = next
			return nil
		}
	}
	return CantDeleteUserError
}

func (ur *dummyUserRepository) Create(email string, name string, password string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		uuid := GenerateUUID()
		_, err := ur.FindByEmail(uuid)
		if err == nil {
			continue
		}
		newUser := domain.User{ID: uuid, Email: email, Name: name, Password: password}
		ur.users = append(ur.users, newUser)
		return uuid, nil
	}
	return "", errors.New("No unused uuid error")
}

func GenerateUUID() string {
	var res string
	for i := 0; i < 30; i++ {
		res += fmt.Sprintf("%d", rand.Intn(10))
	}
	return res
}
