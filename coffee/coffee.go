package coffee

import (
	"context"
	"errors"
	"log"
)

type Server struct {
}

func (s *Server) AddCoffeeInCart(ctx context.Context, cart *AddCoffeeToCart) (*Response, error) {
	//TODO implement me
	panic("implement me")
}

type Coffee struct {
	Id    int
	Name  string
	Price float64
}

var coffeestore []*Coffee

func (s *Server) AddCoffee(ctx context.Context, request *AddCoffeeRequest) (response *AddCoffeeResponse, err error) {
	id, err := GenerateId()
	if err != nil {
		err = errors.New("Id not generated")
	}
	request.Id = int32(id)
	coffee := &Coffee{
		Price: request.Price,
		Name:  request.Name,
	}
	response = &AddCoffeeResponse{}
	coffeestore = append(coffeestore, coffee)
	response.Name = request.Name
	response.Price = request.Price
	return response, nil
}

func (s *Server) FindCoffeeByID(ctx context.Context, id *FindCoffeeById) (response *AddCoffeeResponse, err error) {
	coffee := &Coffee{}
	for _, coffee = range coffeestore {
		id.Id = int32(coffee.Id)
		if int(id.Id) == coffee.Id {
			response.Name = coffee.Name
			response.Price = coffee.Price
			err = nil

		}
	}
	return response, errors.New("Coffee does not exist")
}

func (s *Server) UpdateCoffeeInfo(ctx context.Context, coffeeUpdateRequest *UpdateCoffee) (response *AddCoffeeResponse, err error) {
	//TODO implement me
	coffeeObj := &Coffee{}
	id := coffeeUpdateRequest.Request.Id
	for _, coffee := range coffeestore {
		if coffee.Id == int(id) {

			coffeeObj.Name = coffeeUpdateRequest.Request.Name
			coffeeObj.Price = coffeeUpdateRequest.Request.Price
			coffeestore = append(coffeestore, coffeeObj)
		}
	}

	response.Name = coffeeObj.Name
	response.Price = coffeeObj.Price
	return response, errors.New("Coffee info not updated")

}

func (s *Server) DeleteCoffeeInfo(ctx context.Context, id *DeleteCoffeeOrder) (response *DeleteOrderResponse, err error) {
	var newCoffeeStore []*Coffee
	coffeeObj := &Coffee{}
	id.Id = int32(coffeeObj.Id)
	for _, coffee := range coffeestore {
		if int(id.Id) != coffee.Id {
			newCoffeeStore = append(newCoffeeStore, coffee)
		}
	}
	if len(newCoffeeStore) == len(coffeestore) {
		log.Fatal(errors.New("Failed to delete coffee record"))
	}
	coffeestore = newCoffeeStore
	return
}

func (s *Server) FindAllCoffee(ctx context.Context, coffee *GetAllCoffee) (*AddCoffeeResponse, error) {

	panic("implement me")
}

func (s *Server) mustEmbedUnimplementedCoffeeServiceServer() {
	//TODO implement me
	panic("implement me")
}

func GenerateId() (id int, err error) {
	if len(coffeestore) < 1 {
		id = len(coffeestore) + 1
		err = nil
		return
	}
	for index, coffee := range coffeestore {
		if index == len(coffeestore)-1 {
			id = coffee.Id + 1
			err = nil
		}

	}
	return -1, errors.New("Id not generated")

}
