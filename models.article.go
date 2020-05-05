package main

type article struct {
	Id int         `json:"id"`
	Title string   `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{
		Id:      0,
		Title:   "Tom",
		Content: "Tom is a boy",
	},
	{
		Id:      1,
		Title:   "Tina",
		Content: "Tina is a girl",
	},
}

func getAllArticles() []article {
	return articleList
}
