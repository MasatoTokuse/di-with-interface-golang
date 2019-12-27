package main

import (
	"fmt"
)

func main() {
	const email = "whitebox@sample.com"
	const password = "pass"

	service := InjectSignUpService()
	service.SignUp(email, password)

	// testable using mock
	serviceMock := InjectSignUpServiceMock()
	serviceMock.SignUp(email, password)
}

// -----------------Injection------------------
func InjectSignUpService() SignUpService {
	return SignUpService {
		InjectUserRepository(),
		InjectUserMailer(),
	}
}

func InjectUserRepository() UserRepository {
	return &ConcreteUserRepository{}
}

func InjectUserMailer() Mailer {
	return &ConcreteMailer{}
}
// -----------------Injection------------------

// ---------------MockInjection----------------
func InjectSignUpServiceMock() SignUpService {
	return SignUpService {
		InjectUserRepositoryMock(),
		InjectUserMailerMock(),
	}
}

func InjectUserRepositoryMock() UserRepository {
	return &MockUserRepository{}
}

func InjectUserMailerMock() Mailer {
	return &MockMailer{}
}
// ---------------MockInjection----------------

// ---------------UserRepository---------------
type UserRepository interface {
    Save(email, password string)
}

type ConcreteUserRepository struct {}

func (cur *ConcreteUserRepository) Save(email, password string) {
	fmt.Println("Use ConcreteUserRepository")
	fmt.Printf("Saved user(%s:%s)\n", email, password)
}

type MockUserRepository struct {}

func (cur *MockUserRepository) Save(email, password string) {
	fmt.Println("Use MockUserRepository")
	fmt.Printf("Saved user(%s:%s) @Mock\n", email, password)
}
// ---------------UserRepository---------------

// -------------------Mailer-------------------
type Mailer interface {
    SendEmail(email string)
}

type ConcreteMailer struct {}

func(cm *ConcreteMailer) SendEmail(email string) {
	fmt.Println("Use ConcreteMailer")
	fmt.Printf("Send email to %s\n", email)
}

type MockMailer struct {}

func(cm *MockMailer) SendEmail(email string) {
	fmt.Println("Use MockMailer")
	fmt.Printf("Send email to %s @Mock\n", email)
}
// -------------------Mailer-------------------

// ----------------SignUpService---------------
type SignUpService struct {
    repo   UserRepository
    mailer Mailer
}

func (s SignUpService) SignUp(email, password string) {
	fmt.Println("")
    s.repo.Save(email, password)
	s.mailer.SendEmail(email)
	fmt.Println("")
}

func NewSignUpService(repo UserRepository, mailer Mailer) SignUpService {
    return SignUpService{
        repo,
        mailer,
    }
}
// ----------------SignUpService---------------