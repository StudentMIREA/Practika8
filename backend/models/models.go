package main

type Product struct {
	ID          int
	Name        string
	Image       string
	Cost        int
	Describtion string
}

type User struct {
	ID    int
	Image string
	Name  string
	Phone string
	Mail  string
}

type ShopCartProduct struct {
	ID     int
	Count  int
	UserID int
}

type favoriteProduct struct {
	ID     int
	UserID int
}

var products = []Product{
	{
		ID:          1,
		Name:        "Кружка \"Гусь\"",
		Image:       "https://ir.ozone.ru/s3/multimedia-1-3/wc1000/7110855183.jpg",
		Cost:        600,
		Describtion: "Материал: Полистоун\nРазмер: 380мл\nМожно мыть в посудомоечной машине.",
	},
	{
		ID:          2,
		Name:        "Кружка керамическая",
		Image:       "https://ir.ozone.ru/s3/multimedia-1-j/wc1000/7075750735.jpg",
		Cost:        500,
		Describtion: "Материал: Керамика\nРазмер: 400мл\nМожно мыть в посудомоечной машине.",
	},
	{
		ID:          3,
		Name:        "Кружка Доляна Клубника",
		Image:       "https://ir.ozone.ru/s3/multimedia-h/wc1000/6785147933.jpg",
		Cost:        600,
		Describtion: "Материал: Керамика, Силикон\nРазмер: 350мл\nМожно мыть в посудомоечной машине.",
	},
	{
		ID:          4,
		Name:        "Кружка хамелеон Чеширский Кот",
		Image:       "https://ir.ozone.ru/s3/multimedia-0/wc1000/6294798312.jpg",
		Cost:        400,
		Describtion: "Материал: Керамика\nРазмер: 300мл\nОсобенность этой кружки в том, что она меняет свой цвет при нагревании. Можно мыть в посудомоечной машине.",
	},
	{
		ID:          5,
		Name:        "Кружка металлическая \"Череп\"",
		Image:       "https://ir.ozone.ru/s3/multimedia-h/wc1000/6304198541.jpg",
		Cost:        400,
		Describtion: "Материал: Нержавеющая сталь\nРазмер: 400мл\nНельзя мыть в посудомоечной машине.",
	},
	{
		ID:          6,
		Name:        "Кружка керамическая Доляна Twist",
		Image:       "https://ir.ozone.ru/s3/multimedia-1-j/wc1000/7090710499.jpg",
		Cost:        400,
		Describtion: "Материал: Керамика\nРазмер: 350мл\nНельзя мыть в посудомоечной машине.",
	},
	{
		ID:          7,
		Name:        "Кружка фарфоровая \"Шины\"",
		Image:       "https://ir.ozone.ru/s3/multimedia-1-9/wc1000/7092772317.jpg",
		Cost:        500,
		Describtion: "Материал: Керамика\nРазмер: 460мл\nМожно мыть в посудомоечной машине.",
	},
	{
		ID:          8,
		Name:        "Кружка с двойными стенками \"Cердце\"",
		Image:       "https://ir.ozone.ru/s3/multimedia-j/wc1000/6808645387.jpg",
		Cost:        400,
		Describtion: "Материал: Стекло, Боросиликатное стекло\nРазмер: 250мл\nМожно мыть в посудомоечной машине.",
	},
	{
		ID:          9,
		Name:        "Starbucks Кружка",
		Image:       "https://ir.ozone.ru/s3/multimedia-1-8/wc1000/7107265016.jpg",
		Cost:        3600,
		Describtion: "Материал: Керамика\nРазмер: 355мл\nВращающееся, Двойные стенки. Можно мыть в посудомоечной машине.",
	},
	{
		ID:          10,
		Name:        "Кружка керамическая Mushroom",
		Image:       "https://ir.ozone.ru/s3/multimedia-m/wc1000/6632228578.jpg",
		Cost:        1600,
		Describtion: "Материал: Керамика\nРазмер: 300мл\nКружка керамическая Mushroom Гриб из Супер Марио. Можно мыть в посудомоечной машине.",
	},
}
