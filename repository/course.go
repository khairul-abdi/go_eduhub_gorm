package repository

import (
	"errors"
	"go_eduhub/model"

	"gorm.io/gorm"
)

type CourseRepository interface {
	FetchAll() ([]model.Course, error)
	FetchByID(id int) (*model.Course, error)
	Store(g *model.Course) error
	Update(id int, g *model.Course) error
}

type courseRepoImpl struct {
	db *gorm.DB
}

func NewCourseRepo(db *gorm.DB) *courseRepoImpl {
	return &courseRepoImpl{db}
}

func (g *courseRepoImpl) FetchAll() ([]model.Course, error) {
	rows, err := g.db.Table("courses").Select("*").Rows()
	if err != nil {
		return nil, err
	}

	var listCourse []model.Course

	for rows.Next() {
		g.db.ScanRows(rows, &listCourse)
	}

	return listCourse, nil
}

func (g *courseRepoImpl) FetchByID(id int) (*model.Course, error) {
	var course model.Course
	result := g.db.Find(&course, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return &course, nil
}

func (g *courseRepoImpl) Store(course *model.Course) error {
	err := g.db.Create(course).Error
	if err != nil {
		return err
	}

	return nil
}

func (g *courseRepoImpl) Update(id int, course *model.Course) error {
	// err := g.db.Model(&model.Course{}).Where("id= ?", id).Updates(course).Error
	// if err != nil {
	// 	return err
	// }

	err := g.db.Raw(`UPDATE public.courses SET name=$2, schedule=$3, attendance=$4 WHERE id =$1`, id, course.Name, course.Schedule, course.Attendance).Error

	if err != nil {
		return err
	}

	return nil
}
