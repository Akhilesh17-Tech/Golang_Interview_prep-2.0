package main

import "fmt"

type Product interface {
	GetName() string
	GetPrice() int
}

type Car struct {
	ID    string
	Name  string
	Price int
}

func (c Car) GetName() string {
	return c.Name

}
func (c Car) GetPrice() int {
	return c.Price

}

type Bike struct {
	ID    string
	Name  string
	Price int
}

func (b Bike) GetName() string {
	return b.Name
}
func (b Bike) GetPrice() int {
	return b.Price
}

type ToyMall struct {
	Products []Product
}

func (t *ToyMall) AddProduct(p Product) {
	t.Products = append(t.Products, p)
}

func (t ToyMall) GetTotalPrice() int {
	var totalPrice int
	for _, product := range t.Products {
		totalPrice += product.GetPrice()
	}
	return totalPrice
}

func (t *ToyMall) SellToy(name string) {
	for i, product := range t.Products {
		if product.GetName() == name {
			fmt.Println(i)
			t.Products = append(t.Products[:i], t.Products[i+1:]...)
			break
		}
	}
}

func (t ToyMall) ShowItems() {
	for _, product := range t.Products {
		fmt.Println(product)
	}
}

func main() {

	car1 := Car{ID: "car123", Name: "BMWToy", Price: 5000}
	car2 := Car{ID: "car123", Name: "AudiToy", Price: 5000}
	car3 := Car{ID: "car123", Name: "SkodaToy", Price: 5000}
	car4 := Car{ID: "car123", Name: "TATAToy", Price: 5000}
	car5 := Car{ID: "car123", Name: "HondaToy", Price: 5000}
	car6 := Car{ID: "car123", Name: "Toy", Price: 5000}
	bike := Bike{ID: "bike123", Name: "KTM", Price: 2000}

	toyMall := ToyMall{}
	toyMall.AddProduct(car1)
	toyMall.AddProduct(car2)
	toyMall.AddProduct(car3)
	toyMall.AddProduct(car4)
	toyMall.AddProduct(car5)
	toyMall.AddProduct(car6)
	toyMall.AddProduct(bike)
	fmt.Println(toyMall.GetTotalPrice())
	toyMall.SellToy(car4.Name)
	fmt.Println(toyMall.GetTotalPrice())
	toyMall.ShowItems()

}
