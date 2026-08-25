package service

import (
	"context"
	"frontend_go/domain"
	"frontend_go/storage"
	"sort"
)

type TreeService struct {
	store *storage.Store
	cache map[string]*domain.TreeNode
}

func NewTreeService(s *storage.Store) *TreeService {
	return &TreeService{store: s, cache: map[string]*domain.TreeNode{}}
}
func (s *TreeService) Import(ctx context.Context, meeting string, people []domain.Participant) error {
	for _, p := range people {
		if err := ctx.Err(); err != nil {
			return err
		}
		p.MeetingID = meeting
		if err := domain.ValidateParticipant(p); err != nil {
			return err
		}
		if err := s.store.PutParticipant(p); err != nil {
			return err
		}
	}
	return nil
}
func (s *TreeService) Build(ctx context.Context, meeting string) (*domain.TreeNode, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if cached := s.cache[meeting]; cached != nil {
		return cached, nil
	}
	people, err := s.store.ListParticipants(meeting)
	if err != nil {
		return nil, err
	}
	nodes := map[string]*domain.TreeNode{}
	for _, p := range people {
		nodes[p.ID] = &domain.TreeNode{Participant: p}
	}
	root := &domain.TreeNode{}
	for _, n := range nodes {
		if n.Participant.ParentID == "" {
			root.Children = append(root.Children, n)
		} else if parent := nodes[n.Participant.ParentID]; parent != nil {
			parent.Children = append(parent.Children, n)
		} else {
			root.Children = append(root.Children, n)
		}
	}
	sort.Slice(root.Children, func(i, j int) bool { return root.Children[i].Participant.Name < root.Children[j].Participant.Name })
	s.cache[meeting] = root
	return root, nil
}
func (s *TreeService) Flatten(root *domain.TreeNode) []domain.Participant {
	out := []domain.Participant{}
	var walk func(*domain.TreeNode)
	walk = func(n *domain.TreeNode) {
		for _, c := range n.Children {
			out = append(out, c.Participant)
			walk(c)
		}
	}
	walk(root)
	return out
}
