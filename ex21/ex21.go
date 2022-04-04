package main

import "fmt"

type pet interface { // интерфейс надевания на питомца маленького ошейника
	PutOnASmallCollar()
}

type cat struct { // кошка
}

func (c *cat) PutOnASmallCollar() { // метод надевания ошейника на кошку
	fmt.Println("Маленький ошейник надет на кошку")
}

type smallDog struct { // собака
}

func (d *smallDog) PutOnABiggerCollar() { // метод надевания растянутого ошейника на собаку
	fmt.Println("Ошейник побольше надет на маленького собакена")
}

type collarAdapter struct { // адаптер
	dog *smallDog
}

func (co *collarAdapter) PutOnASmallCollar() { // метода "растягивания ошеника"
	co.dog.PutOnABiggerCollar()
}

type owner struct { // хозяин
}

func (o *owner) PutASmallCollarOnPet(p pet) { // метод надевания хозяином ошеника на питомца
	p.PutOnASmallCollar()
}

func main() {
	owner := owner{}
	cat := cat{}
	owner.PutASmallCollarOnPet(&cat)

	dogge := smallDog{}
	doggeAdapter := collarAdapter{
		dog: &dogge,
	}

	owner.PutASmallCollarOnPet(&doggeAdapter)
}
