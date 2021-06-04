package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

/*
	Initializing govalidator
	Properties are set with TAGS inside the struct
*/
func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

/*
	Technical details such as DB persistence are not defined here
	The Service Interface is decoupled from specific implementations
	This is a crucial point on Hexagonal Architecture
*/
type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

// Interface for Reading DB
type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

// Interface for Saving on DB
type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

/*
	Defining a single interface for DB operations with interface composition
	Any DB that is implemented will never communicate with the Service directly
*/
type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"` // Optional because when zero, it returns empty value (bug?)
	Status string  `valid:"required"`
}

// Constructor, to set default values
func NewProduct() *Product {
	product := Product{
		ID:     uuid.NewV4().String(),
		Status: DISABLED,
	}
	return &product
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("The status must be 'enabled' or 'disabled")
	}

	if p.Price < 0 {
		return false, errors.New("The price must be equal or greater than zero")
	}

	// Using govalidator
	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("Price must be greater than zero to enable product")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New("Price must be zero to disable product")
}

func (p *Product) GetId() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
