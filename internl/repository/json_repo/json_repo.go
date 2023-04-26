package jsonrepo

import (
	"encoding/json"
	"io/fs"
	"os"
	"refactoring/internl/entity"
	rerror "refactoring/internl/repository/repo_error"
	"strconv"
	"sync"
	"time"
)

type jsonStore struct {
	sync.RWMutex
	filename string
}

func NewJsonStore(filename string) (*jsonStore, error) {
	_, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return nil, err
	}

	return &jsonStore{
		filename: filename,
	}, nil
}

func (repo *jsonStore) userStore() (*entity.UserStore, error) {
	f, err := os.ReadFile(repo.filename)
	if err != nil {
		return nil, err
	}
	s := entity.UserStore{}
	if len(f) == 0 {
		s.List = make(entity.UserList)
		return &s, nil
	}
	err = json.Unmarshal(f, &s)
	if err != nil {
		return nil, err
	}
	return &s, err
}

func (repo *jsonStore) commit(s *entity.UserStore) error {
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return os.WriteFile(repo.filename, b, fs.ModePerm)
}

func (repo *jsonStore) GetUsers() (*entity.UserStore, error) {
	repo.RLock()
	defer repo.RUnlock()

	return repo.userStore()
}

func (repo *jsonStore) CreateUser(user *entity.User) (string, error) {
	repo.Lock()
	defer repo.Unlock()

	s, err := repo.userStore()
	if err != nil {
		return "", err
	}
	s.Increment++
	u := entity.User{
		CreatedAt:   time.Now(),
		DisplayName: user.DisplayName,
		Email:       user.Email,
	}

	id := strconv.Itoa(s.Increment)
	s.List[id] = u

	return id, repo.commit(s)
}

func (repo *jsonStore) GetUser(id string) (*entity.User, error) {
	repo.RLock()
	defer repo.RUnlock()

	s, err := repo.userStore()
	if err != nil {
		return nil, err
	}
	user := s.List[id]

	return &user, nil
}

func (repo *jsonStore) UpdateUser(id string, user *entity.User) error {
	repo.Lock()
	defer repo.Unlock()

	s, err := repo.userStore()
	if err != nil {
		return err
	}

	if _, ok := s.List[id]; !ok {
		return rerror.ErrUserNotFound
	}

	u := s.List[id]
	u.DisplayName = user.DisplayName
	s.List[id] = u

	return repo.commit(s)
}

func (repo *jsonStore) DeleteUser(id string) error {
	repo.Lock()
	defer repo.Unlock()

	s, err := repo.userStore()
	if err != nil {
		return err
	}

	if _, ok := s.List[id]; !ok {
		return rerror.ErrUserNotFound
	}

	delete(s.List, id)

	return repo.commit(s)
}
