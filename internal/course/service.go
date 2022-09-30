package course

import (
	"context"
	"log"
	"time"

	"github.com/ncostamagna/g_ms_domain_ex/domain"
)

type (
	Filters struct {
		Name string
	}

	Service interface {
		Create(ctx context.Context, name, startDate, endDate string) (*domain.Course, error)
		Get(ctx context.Context, id string) (*domain.Course, error)
		GetAll(ctx context.Context, filters Filters, offset, limit int) ([]domain.Course, error)
		Delete(ctx context.Context, id string) error
		Update(ctx context.Context, id string, name, startDate, endDate *string) error
		Count(ctx context.Context, filters Filters) (int, error)
	}

	service struct {
		log  *log.Logger
		repo Repository
	}
)

//NewService is a service handler
func NewService(l *log.Logger, repo Repository) Service {
	return &service{
		log:  l,
		repo: repo,
	}
}

//Create service
func (s service) Create(ctx context.Context, name, startDate, endDate string) (*domain.Course, error) {
	startDateParsed, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		s.log.Println(err)
		return nil, ErrInvalidStartDate
	}

	endDateParsed, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		s.log.Println(err)
		return nil, ErrInvalidEndDate
	}

	course := &domain.Course{
		Name:      name,
		StartDate: startDateParsed,
		EndDate:   endDateParsed,
	}

	if err := s.repo.Create(ctx, course); err != nil {
		return nil, err
	}

	return course, nil
}

func (s service) GetAll(ctx context.Context, filters Filters, offset, limit int) ([]domain.Course, error) {
	courses, err := s.repo.GetAll(ctx, filters, offset, limit)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (s service) Get(ctx context.Context, id string) (*domain.Course, error) {
	course, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (s service) Delete(ctx context.Context, id string) error {

	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

// 2 formas de validar si exite, mediante el get o mediante el result del repository
// ver que podemos tener problemas al agregar texto mayor a lo que se espera en la base de datos
func (s service) Update(ctx context.Context, id string, name, startDate, endDate *string) error {

	var startDateParsed, endDateParsed *time.Time

	if startDate != nil {
		date, err := time.Parse("2006-01-02", *startDate)
		if err != nil {
			s.log.Println(err)
			return ErrInvalidStartDate
		}
		startDateParsed = &date
	}

	if endDate != nil {
		date, err := time.Parse("2006-01-02", *endDate)
		if err != nil {
			s.log.Println(err)
			return ErrInvalidEndDate
		}
		endDateParsed = &date
	}

	if err := s.repo.Update(ctx, id, name, startDateParsed, endDateParsed); err != nil {
		return err
	}
	return nil
}

func (s service) Count(ctx context.Context, filters Filters) (int, error) {
	return s.repo.Count(ctx, filters)
}