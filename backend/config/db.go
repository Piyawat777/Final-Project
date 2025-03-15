package config

import (
	"fmt"
	"log"
	"time"

	"github.com/Piyawat777/Final-Project/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("Hotel.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

func SetupDatabase() {

	err := db.AutoMigrate(
		&entity.Amenity{},
		&entity.Booking{},
		&entity.User{},
		&entity.Hotel{},
		&entity.Invoice{},
		&entity.Maintenance{},
		&entity.Payment{},
		&entity.Promotion{},
		&entity.Review{},
		&entity.Room{},
		&entity.RoomImage{},
		&entity.Staff{},
	)

	if err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}

	Gender := []entity.Genders{
		{Gender: "Male"},
		{Gender: "Female"},
		{Gender: "Other"},
	}
	for _, pkg := range Gender {
		// Ensure that gender is created if not already existing
		if err := db.FirstOrCreate(&pkg, entity.Genders{Gender: pkg.Gender}).Error; err != nil {
			fmt.Printf("Error creating gender: %v\n", err)
		}
	}
	// Hash password and handle error
	hashedPassword, err := HashPassword("123456")
	if err != nil {
		log.Fatalf("failed to hash password: %v", err)
	}

	// Parse birth date and handle error
	BirthDay, err := time.Parse("2006-01-02", "1988-11-12")
	if err != nil {
		log.Fatalf("failed to parse birthdate: %v", err)
	}

	// Insert default user data
	User := []entity.User{
		{Username: "@A",
			Name:         "A",
			Email:        "A@g.com",
			Password:     hashedPassword,
			Birthday:     BirthDay,
			ProfileImage: "https://i.pinimg.com/736x/1d/5c/74/1d5c74a0964dc147aca12dae6695ab26.jpg",
			CoverImage:   "https://img.freepik.com/free-photo/abstract-studio-background-texture-light-blue-gray-gradient-wall-flat-floor-product_1258-88339.jpg?t=st=1737629944~exp=1737633544~hmac=097f29b58d3c426647a4db396e4bf4bed9377227b43da9816df44856302d133e&w=996",
			GenderID:     2,
			Role:       "admin",
		},
		{Username: "@fourarm",
			Name:         "fourarm",
			Email:        "sa@gmail.com",
			Password:     hashedPassword,
			Description:  "นายfourarm",
			Birthday:     BirthDay,
			ProfileImage: "https://i.pinimg.com/736x/45/fd/d3/45fdd38b20a18b1fe3308158b9e8ea70.jpg",
			CoverImage:   "https://img.freepik.com/free-photo/abstract-studio-background-texture-light-blue-gray-gradient-wall-flat-floor-product_1258-88339.jpg?t=st=1737629944~exp=1737633544~hmac=097f29b58d3c426647a4db396e4bf4bed9377227b43da9816df44856302d133e&w=996",
			GenderID:     1,
			Status:       "admin",
		},
	}
	for _, pkg := range User {
		db.FirstOrCreate(&pkg, entity.User{Username: pkg.Username})
	}

	// Insert default tags
	tags := []entity.Tags{
		{Tag: "#Zenless Zone Zero"}, {Tag: "#Genshin Impact"}, {Tag: "#Honkai Impact 3rd"},
		{Tag: "#Arknights"}, {Tag: "#Blue Archive"}, {Tag: "#Project Sekai"},
		{Tag: "#Fate/Grand Order"}, {Tag: "#Azur Lane"}, {Tag: "#Alchemy Stars"},
		{Tag: "#Punishing Gray Raven"}, {Tag: "#Girls' Frontline"}, {Tag: "#Tower of Fantasy"},
		{Tag: "#Eversoul"}, {Tag: "#Guardian Tales"}, {Tag: "#NIKKE"},
		{Tag: "#Epic Seven"}, {Tag: "#Seven Knights"}, {Tag: "#Grand Cross"},
		{Tag: "#Dragon Raja"}, {Tag: "#Shadowverse"},
	}

	for _, tag := range tags {
		db.FirstOrCreate(&tag, entity.Tags{Tag: tag.Tag})
	}

	// Insert default categories
	categories := []entity.Categories{
		{Category: "Others"},
		{Category: "Illustrations"},
		{Category: "Manga"},
	}

	for _, category := range categories {
		db.FirstOrCreate(&category, entity.Categories{Category: category.Category})
	}

	privacies := []entity.Privacies{
		{Privacy: "Public"},
		{Privacy: "Private"},
	}

	for _, privacy := range privacies {
		db.FirstOrCreate(&privacy, entity.Privacies{Privacy: privacy.Privacy})
	}

	// ข้อมูล Platform (Social Media platforms)
	Platform := []entity.Platforms{
		{PlatformName: "Facebook", PlatformIcon: "https://cdn-icons-png.freepik.com/512/2496/2496095.png"},
		{PlatformName: "Instagram", PlatformIcon: "https://img.freepik.com/free-vector/instagram-logo_1199-122.jpg?semt=ais_hybrid"},
		{PlatformName: "Twitter", PlatformIcon: "https://cdn-icons-png.freepik.com/512/2496/2496110.png"},
		{PlatformName: "YouTube", PlatformIcon: "https://cdn-icons-png.freepik.com/512/2496/2496113.png?ga=GA1.1.672280926.1732045333"},
		{PlatformName: "TikTok", PlatformIcon: "https://i.pinimg.com/736x/82/90/df/8290df78623a45b9d4c30079b7487a94.jpg"},
		{PlatformName: "LinkedIn", PlatformIcon: "https://cdn-icons-png.freepik.com/512/2496/2496097.png?ga=GA1.1.672280926.1732045333"},
		{PlatformName: "Other", PlatformIcon: "https://icon-library.com/images/no-profile-picture-icon/no-profile-picture-icon-14.jpg"},
	}
	for _, pkg := range Platform {
		db.FirstOrCreate(&pkg, entity.Platforms{PlatformName: pkg.PlatformName})
	}
}