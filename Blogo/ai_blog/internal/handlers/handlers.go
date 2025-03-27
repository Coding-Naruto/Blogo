package handlers

import (
    "html/template"
    "net/http"
    "path/filepath"
    "strconv"

    "ai_blog/internal/models"
    "ai_blog/internal/ai"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    tmplPath := filepath.Join("templates", tmpl+".html")
    layoutPath := filepath.Join("templates", "layout.html")
    tmplParsed, _ := template.ParseFiles(layoutPath, tmplPath)
    tmplParsed.ExecuteTemplate(w, "layout", data)
}

func Home(w http.ResponseWriter, r *http.Request) {
    var posts []models.Post
    models.DB.Find(&posts)
    renderTemplate(w, "home", struct {
        Title string
        Posts []models.Post
    }{Title: "Home", Posts: posts})
}

func Post(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        r.ParseForm()
        title := r.FormValue("title")
        content := r.FormValue("content")
        post := models.Post{Title: title, Content: content}
        models.DB.Create(&post)
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
    renderTemplate(w, "post", struct{ Title string }{Title: "New Post"})
}

func Edit(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(r.URL.Query().Get("id"))
    var post models.Post
    models.DB.First(&post, id)
    if r.Method == http.MethodPost {
        r.ParseForm()
        post.Title = r.FormValue("title")
        post.Content = r.FormValue("content")
        models.DB.Save(&post)
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
    renderTemplate(w, "edit", struct {
        Title string
        Post  models.Post
    }{Title: "Edit Post", Post: post})
}

func Delete(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(r.URL.Query().Get("id"))
    models.DB.Delete(&models.Post{}, id)
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ViewPost(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(r.URL.Query().Get("id"))
    var post models.Post
    models.DB.First(&post, id)
    summary, err := ai.SummarizeText(post.Content)
    if err != nil {
        summary = "Error summarizing post."
    }
    renderTemplate(w, "view", struct {
        Title   string
        Post    models.Post
        Summary string
    }{Title: post.Title, Post: post, Summary: summary})
}