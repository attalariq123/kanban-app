package web

import (
	"a21hc3NpZ25tZW50/client"
	"embed"
	"log"
	"net/http"
	"path"
	"text/template"
)

type DashboardWeb interface {
	Dashboard(w http.ResponseWriter, r *http.Request)
}

type dashboardWeb struct {
	categoryClient client.CategoryClient
	userClient     client.UserClient
	embed          embed.FS
}

func NewDashboardWeb(catClient client.CategoryClient, uClient client.UserClient, embed embed.FS) *dashboardWeb {
	return &dashboardWeb{catClient, uClient, embed}
}

func (d *dashboardWeb) Dashboard(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id").(string)

	categories, err := d.categoryClient.GetCategories(userId)
	if err != nil {
		log.Println("error get category data: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users, err := d.userClient.GetUserById(userId)
	if err != nil {
		log.Println("error get user data: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var dataTemplate = map[string]interface{}{
		"categories": categories,
		"users":      users,
	}

	var getIndexByCategoryId = func(catId int) int {
		for i := 0; i < len(categories); i++ {
			if categories[i].ID == catId {
				return i
			}
		}

		return -1
	}

	var funcMap = template.FuncMap{
		"categoryInc": func(categoryId int) int {
			idx := getIndexByCategoryId(categoryId)

			if idx == len(categories)-1 {
				return categoryId
			} else {
				return categories[idx+1].ID
			}
		},
		"categoryDec": func(categoryId int) int {
			idx := getIndexByCategoryId(categoryId)

			if idx == 0 {
				return categoryId
			} else {
				return categories[idx-1].ID
			}
		},
	}

	// ignore this
	_ = dataTemplate
	_ = funcMap
	//

	filePath := path.Join("views", "main", "dashboard.html")
	header := path.Join("views", "general", "header.html")

	tmpl := template.Must(template.New("").Funcs(funcMap).ParseFS(d.embed, filePath, header))

	err = tmpl.ExecuteTemplate(w, "dashboard.html", dataTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
