package main

import (
	"html/template"
	"log"
	"net/http"
)

type indexPage struct {
	Title         string
	FeaturedPosts []featuredPostData
	RecentPosts   []recentPostData
}

type featuredPostData struct {
	Title         string
	Subtitle      string
	AuthorPhoto   string
	AuthorName    string
	Date          string
	BackgroundImg string
}

type recentPostData struct {
	PostImg     string
	Title       string
	Subtitle    string
	AuthorPhoto string
	AuthorName  string
	Date        string
}

func index(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := indexPage{
		Title:         "Escape",
		FeaturedPosts: featuredPosts(),
		RecentPosts:   recentPosts(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/post.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	err = ts.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func featuredPosts() []featuredPostData {
	return []featuredPostData{
		{
			Title:         "The Road Ahead",
			Subtitle:      "The road ahead might be paved - it might not be.",
			AuthorPhoto:   "../static/img/Mat.png",
			AuthorName:    "Mat Vogels",
			Date:          "September 25, 2015",
			BackgroundImg: "card-big_background-1",
		},
		{
			Title:         "From Top Down",
			Subtitle:      "Once a year, go someplace you’ve never been before.",
			AuthorPhoto:   "../static/img/William.png",
			AuthorName:    "William Wong",
			Date:          "September 25, 2015",
			BackgroundImg: "card-big_background-2",
		},
	}
}

func recentPosts() []recentPostData {
	return []recentPostData{
		{
			PostImg:     "../static/img/balloons.jpg",
			Title:       "Still Standing Tall",
			Subtitle:    "Life begins at the end of your comfort zone.",
			AuthorPhoto: "../static/img/William.png",
			AuthorName:  "William Wong",
			Date:        "9/25/2015",
		},
		{
			PostImg:     "../static/img/bridge.jpg",
			Title:       "Sunny Side Up",
			Subtitle:    "No place is ever as bad as they tell you it’s going to be.",
			AuthorPhoto: "../static/img/Mat.png",
			AuthorName:  "Mat Vogels",
			Date:        "9/25/2015",
		},
		{
			PostImg:     "../static/img/sunset.jpg",
			Title:       "Water Falls",
			Subtitle:    "We travel not to escape life, but for life not to escape us.",
			AuthorPhoto: "../static/img/Mat.png",
			AuthorName:  "Mat Vogels",
			Date:        "9/25/2015",
		},
		{
			PostImg:     "../static/img/ocean.jpg",
			Title:       "Through the Mist",
			Subtitle:    "Travel makes you see what a tiny place you occupy in the world.",
			AuthorPhoto: "../static/img/William.png",
			AuthorName:  "William Wong",
			Date:        "9/25/2015",
		},
		{
			PostImg:     "../static/img/funicular.jpg",
			Title:       "Awaken Early",
			Subtitle:    "Not all those who wander are lost.",
			AuthorPhoto: "../static/img/Mat.png",
			AuthorName:  "Mat Vogels",
			Date:        "9/25/2015",
		},
		{
			PostImg:     "../static/img/waterfall.jpg",
			Title:       "Try it Always",
			Subtitle:    "The world is a book, and those who do not travel read only one page.",
			AuthorPhoto: "../static/img/Mat.png",
			AuthorName:  "Mat Vogels",
			Date:        "9/25/2015",
		},
	}
}
