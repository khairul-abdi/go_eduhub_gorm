package repository

import (
	"errors"
	"go_eduhub/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	rows, err := s.db.Table("students").Select("*").Rows()
	if err != nil {
		return nil, err
	}

	var listStudents []model.Student

	for rows.Next() {
		var student model.Student

		err := rows.Scan(&student.ID, &student.FullName, &student.Address, &student.Class)
		if err != nil {
			return nil, err
		}

		listStudents = append(listStudents, student)
	}

	return listStudents, nil
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	var student model.Student

	result := s.db.Find(&student, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return &student, nil
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	err := s.db.Create(student).Error
	if err != nil {
		return err
	}

	return nil
}
