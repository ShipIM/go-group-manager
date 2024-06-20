package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ShipIM/go-group-manager/sender/internal/domain/entity"
	"github.com/ShipIM/go-group-manager/sender/internal/repository"
	amqp "github.com/rabbitmq/amqp091-go"
)

type StudentService struct {
	studentRepository repository.StudentRepository
	conn              *amqp.Connection
	exchange          string
}

func NewStudentService(studentRepository repository.StudentRepository, conn *amqp.Connection, exchange string) *StudentService {
	return &StudentService{studentRepository, conn, exchange}
}

func (s *StudentService) CreateStudent(student entity.Student) (entity.Student, error) {
	student, err := s.studentRepository.CreateStudent(student)
	if err != nil {
		return student, err
	}

	err = s.sendStudent(student)
	if err != nil {
		return student, fmt.Errorf("failed to send student: %w", err)
	}

	return student, err
}

func (s *StudentService) sendStudent(student entity.Student) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(student)
	if err != nil {
		return fmt.Errorf("failed to serialize student: %w", err)
	}

	ch, err := s.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %w", err)
	}
	defer ch.Close()

	err = ch.PublishWithContext(ctx,
		s.exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}

func (s *StudentService) GetStudentById(id int) (entity.Student, error) {
	return s.studentRepository.GetStudentById(id)
}

func (s *StudentService) GetStudents() ([]entity.Student, error) {
	return s.studentRepository.FindAllStudents()
}

func (s *StudentService) UpdateStudent(student entity.Student) error {
	return s.studentRepository.UpdateStudent(student)
}

func (s *StudentService) DeleteStudentById(id int) error {
	return s.studentRepository.DeleteStudentById(id)
}
