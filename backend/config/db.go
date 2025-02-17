package config

import (
	"fmt"
	"log"
	"time"

	"example.com/se/entity"

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
		&entity.Users{},
		&entity.Genders{},
		&entity.Tags{},
		&entity.Categories{},
		&entity.Privacies{},
		&entity.Works{},
		&entity.WorkTags{},
		&entity.WorkIllustrations{},
		&entity.Platforms{},
		&entity.Contacts{},
		&entity.Albums{},
		&entity.AlbumWorks{},
		&entity.Events{},
		&entity.Prizes{},
		&entity.Community{},
		&entity.Post{},
		&entity.Member{},
		&entity.WorkLikes{},
		&entity.Comments{},
		&entity.Replies{},
		&entity.ProductFiles{},
		&entity.ProductExamples{},
		&entity.ProductOptions{},
		&entity.Products{},
		&entity.ProductTags{},
		&entity.Purchases{},
		&entity.PurchaseGifts{},
		&entity.PurchaseGalleries{},
		&entity.Request{},
		&entity.Respond{},
		&entity.Faq{},
		&entity.Question{},
		&entity.Wallet{},
		&entity.Paymentmethod{},
		&entity.Transaction{},
		&entity.Type{},
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
	User := []entity.Users{
		{Username: "@A",
			Name:         "A",
			Email:        "A@g.com",
			Password:     hashedPassword,
			Description:  "นายA นามสมมุติ",
			Birthday:     BirthDay,
			ProfileImage: "https://i.pinimg.com/736x/1d/5c/74/1d5c74a0964dc147aca12dae6695ab26.jpg",
			CoverImage:   "https://img.freepik.com/free-photo/abstract-studio-background-texture-light-blue-gray-gradient-wall-flat-floor-product_1258-88339.jpg?t=st=1737629944~exp=1737633544~hmac=097f29b58d3c426647a4db396e4bf4bed9377227b43da9816df44856302d133e&w=996",
			GenderID:     2,
			Status:       "admin",
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
		db.FirstOrCreate(&pkg, entity.Users{Username: pkg.Username})
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

	//Request
	Request := &entity.Request{
		RequesttopicID: 2,
		FirstName:      "Jieun",
		LastName:       "Lee",
		Email:          "iu@gmail.com",
		Subject:        "ปัญหาการชำระเงิน",
		MobilePhone:    "0987654321",
		Message:        "ชำระเงิน...",
	}

	db.FirstOrCreate(Request, &entity.Request{
		RequesttopicID: 2,
		FirstName:      "Jieun",
		LastName:       "Lee",
		Email:          "iu@gmail.com",
		Subject:        "ปัญหาการชำระเงิน",
		MobilePhone:    "0987654321",
		Message:        "ชำระเงิน...",
	})

	//Requesttopic
	RequesttopicGeneral := entity.Requesttopic{Category: "General"}
	RequesttopicPayment := entity.Requesttopic{Category: "Payment"}
	RequesttopicReportingAnIssue := entity.Requesttopic{Category: "Reporting an Issue"}
	db.FirstOrCreate(&RequesttopicGeneral, &entity.Requesttopic{Category: "General"})
	db.FirstOrCreate(&RequesttopicPayment, &entity.Requesttopic{Category: "Payment"})
	db.FirstOrCreate(&RequesttopicReportingAnIssue, &entity.Requesttopic{Category: "Reporting an Issue"})

	//Faq
	Faq := &entity.Faq{
		UserID:         1,
		RequesttopicID: 2,
		Question:       "วิธีการชำระเงิน",
		Respond:        "-ขั้นตอนแรก...",
	}

	db.FirstOrCreate(Faq, &entity.Faq{
		UserID:         1,
		RequesttopicID: 2,
		Question:       "วิธีการชำระเงิน",
		Respond:        "-ขั้นตอนแรก...",
	})

	Faq2 := &entity.Faq{
		UserID:         1,
		RequesttopicID: 3,
		Question:       "วิธีแก้ไขโปรไฟล์",
		Respond:        "-ขั้นตอนแรก...",
	}

	db.FirstOrCreate(Faq2, &entity.Faq{
		UserID:         1,
		RequesttopicID: 3,
		Question:       "วิธีแก้ไขโปรไฟล์",
		Respond:        "-ขั้นตอนแรก...",
	})

	//Question
	Question := &entity.Question{
		UserID:         2,
		RequesttopicID: 2,
		Title:          "ฉันจะชำระเงินได้อย่างไร",
		Content:        "ฉันไม่สามารถชำระเงินได้ผ่านช่องทาง...",
		Status:         "Answered",
	}

	db.FirstOrCreate(Question, &entity.Question{
		UserID:         2,
		RequesttopicID: 2,
		Title:          "ฉันจะชำระเงินได้อย่างไร",
		Content:        "ฉันไม่สามารถชำระเงินได้ผ่านช่องทาง...",
		Status:         "Answered",
	})

	//Respond
	Respond := &entity.Respond{
		UserID:     1,
		QuestionID: 1,
		Respond:    "สวัสดีค่ะ จากปัญหาที่พบ คุณสามารถ...",
	}

	db.FirstOrCreate(Respond, &entity.Respond{
		UserID:     1,
		QuestionID: 1,
		Respond:    "สวัสดีค่ะ จากปัญหาที่พบ คุณสามารถ...",
	})

	//Wallet
	Wallet := &entity.Wallet{
		UserID:  1,
		Balance: 7777.77,
	}

	db.FirstOrCreate(Wallet, &entity.Wallet{
		UserID:  1,
		Balance: 7777.77,
	})

	//Transaction
	PaymentmethodTrueMoney := entity.Paymentmethod{Name: "True Money"}

	PaymentmethodKrungthaiBank := entity.Paymentmethod{Name: "Krungthai Bank"}

	db.FirstOrCreate(&PaymentmethodTrueMoney, &entity.Paymentmethod{Name: "True Money"})

	db.FirstOrCreate(&PaymentmethodKrungthaiBank, &entity.Paymentmethod{Name: "Krungthai Bank"})

	TypeDeposit := entity.Type{Name: "Deposit"}

	TypeTransfer := entity.Type{Name: "Transfer"}

	db.FirstOrCreate(&TypeDeposit, &entity.Type{Name: "Deposit"})

	db.FirstOrCreate(&TypeTransfer, &entity.Type{Name: "Transfer"})

	Transaction := &entity.Transaction{
		WalletID:        1,
		TypeID:          2,
		PaymentmethodID: 1,
		AccountNumber:   "0987654321",
		Amount:          777,
	}

	db.FirstOrCreate(Transaction, &entity.Transaction{
		WalletID:        1,
		TypeID:          2,
		PaymentmethodID: 1,
		AccountNumber:   "0987654321",
		Amount:          777,
	})
}
