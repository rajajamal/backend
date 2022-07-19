package main

import (
	"github.com/joho/godotenv"
	"github.com/rajajamal/resto-backend/configs"
	"github.com/rajajamal/resto-backend/models"
	"github.com/rajajamal/resto-backend/repositories"
	"github.com/rajajamal/resto-backend/types"
)

func init() {
	godotenv.Load()
	configs.Load()
}

func main() {
	menuRepository := repositories.Menu{
		Storage: configs.Db,
	}

	menu1 := models.Menu{
		Type:  types.MENU_DRINK,
		Name:  "Susu Jahe",
		Price: 5000,
	}

	menu2 := models.Menu{
		Type:  types.MENU_MAIN_COURSE,
		Name:  "Ayam Geprek",
		Price: 15000,
	}

	menu3 := models.Menu{
		Type:  types.MENU_SNACK,
		Name:  "Mendoan",
		Price: 5000,
	}

	menuRepository.Saves(&menu1, &menu2, &menu3)
}
