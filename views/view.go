package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type Blog struct {
	Title string
	Para string
	Author string
}

type Review struct {
	Review string
	Client string
	Stylist string
}

var (
	LayoutDir   string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {

	//r1 := Review{
	//	"Wonderful experience, love my hair",
	//	"Anna Alexander",
	//	"Adam",
	//}

	b1 := Blog{
		"Blog Post One",
		"This is the first Blog Post",
		"Adam",
	}
	b2 := Blog{
		"Blog Post 2",
		"This is Blog post two",
		"Nat",
	}
	b3 := Blog{
		"Blog Post 3",
		"This is the third Blog Post",
		"Matt",
	}

	blogs := []Blog{b1, b2, b3}


	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, blogs)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}

	return files
}
