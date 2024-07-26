package service

import "fmt"

func (s *Service) GetStats() (map[string]interface{}, error) {
	stats, err := s.repo.GetStats()
	if err != nil {
		return nil, fmt.Errorf("failed to get message stats: %w", err)
	}

	return stats, nil
}
