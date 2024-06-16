package pets

import (
	"errors"
	"fmt"
	"patterns_go/models"
)

// AnimalInterface is the interface for the types we will return from our abstract
// factory. We'll create two types which implement this interface: DogFromFactory and
// CatFromFactory. In order to implement this interface, both types must have a Show()
// method.
type AnimalInterface interface {
	Show() string
}

// DogFromFactory is a type which implements the AnimalInterface, and embeds a dog.
type DogFromFactory struct {
	Pet *models.Dog
}

// Show is a method on the DogFromFactory type, which makes the type implement AnimalInterface.
func (dff *DogFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", dff.Pet.Breed.Breed)
}

// CatFromFactory is a type which implements the AnimalInterface, and embeds a cat.
type CatFromFactory struct {
	Pet *models.Cat
}

// Show is a method on the CatFromFactory type, which makes the type implement AnimalInterface.
func (cff *CatFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", cff.Pet.Breed.Breed)
}

// PetFactoryInterface is the interface for our DogAbstractFactory and CatAbstractFactory functions.
// Both of these types must implement a newPet method which returns a type that satisfies the
// AnimalInterface.
type PetFactoryInterface interface {
	newPet() AnimalInterface
}

type DogAbstractFactory struct{}

func (df *DogAbstractFactory) newPet() AnimalInterface {
	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}

type CatAbstractFactory struct{}

func (cf *CatAbstractFactory) newPet() AnimalInterface {
	return &CatFromFactory{
		Pet: &models.Cat{},
	}
}

// NewPetFromAbstractFactory is the abstract factory method.
func NewPetFromAbstractFactory(species string) (AnimalInterface, error) {
	switch species {
	case "dog":
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPet()
		return dog, nil
	case "cat":
		var catFactory CatAbstractFactory
		cat := catFactory.newPet()
		return cat, nil
	default:
		return nil, errors.New("invalid species supplied")
	}
}
