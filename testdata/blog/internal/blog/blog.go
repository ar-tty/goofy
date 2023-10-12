package blog

type Post struct {
	ID      int    `goofy:"gt=0" json:"ellodisco" json:"id"`
	Title   string `goofy:"required,lte=150,gte=10" json:"title"`
	Content string `goofy:"required,lte=600w,gte=15w" json:"content"`
}

type Post2 struct {
	ID      int    `goofy:"gt=0" json:"ellodisco" json:"id"`
	Title   string `goofy:"required,lte=150,gte=10" json:"title"`
	Content string `goofy:"required,lte=600w,gte=15w" json:"content"`
}
