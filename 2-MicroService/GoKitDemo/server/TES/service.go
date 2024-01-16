package TES

import (
	"context"
	"errors"
	"server/TES/structure"
	"sync"
)

// 服务接口=================================
type Service interface {
	PostProfile(ctx context.Context, p structure.Profile) error
	GetProfile(ctx context.Context, id string) (structure.Profile, error)
	PutProfile(ctx context.Context, id string, p structure.Profile) error
	PatchProfile(ctx context.Context, id string, p structure.Profile) error
	DeleteProfile(ctx context.Context, id string) error
}

// 错误定义=================================
var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

// 内部服务结构体=================================
type inmemService struct {
	mtx sync.RWMutex
	m   map[string]structure.Profile
}

func NewInmemService() Service {
	return &inmemService{
		m: map[string]structure.Profile{},
	}
}

// CRUD 操作=================================
func (s *inmemService) PostProfile(ctx context.Context, p structure.Profile) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.m[p.ID]; ok {
		return ErrAlreadyExists // POST = create, don't overwrite
	}
	s.m[p.ID] = p
	return nil
}

func (s *inmemService) GetProfile(ctx context.Context, id string) (structure.Profile, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	p, ok := s.m[id]
	if !ok {
		return structure.Profile{}, ErrNotFound
	}
	return p, nil
}

func (s *inmemService) PutProfile(ctx context.Context, id string, p structure.Profile) error {
	if id != p.ID {
		return ErrInconsistentIDs
	}
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.m[id] = p // PUT = create or update
	return nil
}

func (s *inmemService) PatchProfile(ctx context.Context, id string, p structure.Profile) error {
	if p.ID != "" && id != p.ID {
		return ErrInconsistentIDs
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	existing, ok := s.m[id]
	if !ok {
		return ErrNotFound // PATCH = update existing, don't create
	}

	if p.Name != "" {
		existing.Name = p.Name
	}
	s.m[id] = existing
	return nil
}

func (s *inmemService) DeleteProfile(ctx context.Context, id string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.m[id]; !ok {
		return ErrNotFound
	}
	delete(s.m, id)
	return nil
}
