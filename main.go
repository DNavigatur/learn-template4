package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name, Descrip, Meal string
	Price               float64
}

type items []item
type itemzz []items

type resturant struct {
	ResName        string
	ResturantMenus itemzz
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	allResturantMenu := []resturant{
		resturant{
			ResName: "Cappa Resturant",
			ResturantMenus: itemzz{items{
				item{
					Name:    "Oatmeal",
					Descrip: "yum yum",
					Meal:    "Breakfast",
					Price:   4.95,
				},
				item{
					Name:    "Hamburger",
					Descrip: "Delicous good eating for you",
					Meal:    "Lunch",
					Price:   6.95,
				},
				item{
					Name:    "Pasta Bolognese",
					Descrip: "From Italy delicious eating",
					Meal:    "Dinner",
					Price:   7.95,
				},
			}},
		},

		resturant{
			ResName: "Yaba Resturant",
			ResturantMenus: itemzz{items{
				item{
					Name:    "Beans",
					Descrip: "crunchy",
					Meal:    "Lunch",
					Price:   8.12,
				},
				item{
					Name:    "Cornflakes",
					Descrip: "Scumpcious",
					Meal:    "Breakfast",
					Price:   2.85,
				},
				item{
					Name:    "Akara & Bread",
					Descrip: "Sweet combimation",
					Meal:    "Dinner",
					Price:   4.35,
				},
			}},
		},
	}

	err := tpl.Execute(os.Stdout, allResturantMenu)
	if err != nil {
		log.Fatalln(err)
	}
}
