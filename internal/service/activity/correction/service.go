package correction

import (
	"errors"
	"fmt"
	model "github.com/ozonmp/omp-bot/internal/model/activity"
	repository "github.com/ozonmp/omp-bot/internal/repository/activity"
	"time"
)

type CorrectionService interface {
	Describe(correctionID uint64) (*model.Correction, error)
	List(cursor uint64, limit uint64) ([]model.Correction, error)
	Create(model.Correction) (uint64, error)
	Update(correctionID uint64, correction model.Correction) error
	Remove(correctionID uint64) (bool, error)
}

type DummyCorrectionService struct { }

func NewDummyCorrectionService() *DummyCorrectionService {
	return &DummyCorrectionService{}
}

// Describe Returns the correction with given ID
func (s *DummyCorrectionService) Describe(correctionID uint64) (*model.Correction, error) {
	idx, correction := s.findCorrectionByID(correctionID)
	if idx != -1 {
		return correction, nil
	}
	return nil, errors.New(fmt.Sprintf("correction with ID = %d was not found", correctionID))
}

// List Returns list of corrections from cursor to limit based on index
func (s *DummyCorrectionService) List(cursor uint64, limit uint64) ([]model.Correction, error) {
	if int(cursor) >= len(repository.Corrections) {
		return nil, errors.New(fmt.Sprintf("Invalid start index %d for number of corrections %d",
			cursor, len(repository.Corrections)))
	}
	if int(cursor + limit) < len(repository.Corrections) {
		return repository.Corrections[cursor:cursor + limit], nil
	}
	return repository.Corrections[cursor:], nil
}

func (s *DummyCorrectionService) Create(correction model.Correction) (uint64, error) {
	err := s.validateRequiredFields(correction)
	if err != nil {
		return 0, err
	}

	correction.ID = s.getCorrectionNextID()
	correction.Timestamp = time.Now()
	repository.Corrections = append(repository.Corrections, correction)
	return correction.ID, nil
}

// Update Edits the correction with given ID, update is only performed on optional fields
func (s *DummyCorrectionService) Update(correctionID uint64, correction model.Correction) error {
	if correctionID > 0 {
		idx, exCorrection := s.findCorrectionByID(correctionID)
		if idx != -1 {
			exCorrection.Comments = correction.Comments
			return nil
		}
		return errors.New(fmt.Sprintf("correction with ID = %d was not found", correctionID))
	}
	return errors.New(fmt.Sprintf("correction with ID = %d was not found", correctionID))
}

// Remove Deletes the correction with given ID
func (s *DummyCorrectionService) Remove(correctionID uint64) (bool, error) {
	idx, _ := s.findCorrectionByID(correctionID)
	if idx != -1 {
		repository.Corrections = append(repository.Corrections[:idx], repository.Corrections[idx+1:]...)
		return true, nil
	}
	return false, errors.New(fmt.Sprintf("correction with ID = %d was not found", correctionID))
}

func (s *DummyCorrectionService) GetCorrectionsCount() int {
	return len(repository.Corrections)
}

// findCorrectionByID Returns index of correction by given correction ID, returns -1 if correction not found
func (s *DummyCorrectionService) findCorrectionByID(correctionID uint64) (int, *model.Correction) {
	for idx, correction := range repository.Corrections {
		if correction.ID == correctionID {
			return idx, &repository.Corrections[idx] // to support modification
		}
	}
	return -1, nil
}

// GetCorrectionNextID Gets next ID, single thread
func (s *DummyCorrectionService) getCorrectionNextID() uint64 {
	if len(repository.Corrections) > 0 {
		max := repository.Corrections[0].ID
		for _, correction := range repository.Corrections {
			if correction.ID > max {
				max = correction.ID
			}
		}
		return max + 1
	}
	return 1
}

func (s *DummyCorrectionService) validateRequiredFields(correction model.Correction) error {
	if correction.UserID == 0 {
		return errors.New("correction must have 'userID' property")
	}
	if correction.Object == "" {
		return errors.New("correction must have 'object' property")
	}
	if correction.Action == "" {
		return errors.New("correction must have 'action' property")
	}
	if correction.Data.OriginalData == "" {
		return errors.New("correction must have 'data.originalData' property")
	}
	if correction.Data.RevisedData == "" {
		return errors.New("correction must have 'data.revisedData' property")
	}
	return nil
}