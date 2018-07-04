package main

import (
	"fmt"
	apiAdapter "restaurant/adapters/api"
	dbAdapters "restaurant/adapters/database"
	mocksDb "restaurant/mocks/database"
	"restaurant/models"
	apiPort "restaurant/ports/api"
	dbPorts "restaurant/ports/database"
	//"strconv"
	//"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/justinas/alice"
	"html/template"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

const MiB_UNIT = 1 << 20

var (
	daoDb          dbPorts.IDatabase
	dishRepository dbPorts.IDishRepository
	userRepository dbPorts.IUserRepository
	menuRepository dbPorts.IMenuRepository
	api            apiPort.IAPI

	/*MOCKS*/
	mockDatabase       mocksDb.MockDatabase
	mockDishRepository mocksDb.MockDishRepository
)

var templates map[string]*template.Template

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var dishes []models.Dish
	err, dishes := dishRepository.FindAll()
	if err != nil {
		log.Fatal(err)
		return
	}
	renderTemplates(w, "index", "base", dishes)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "login", "base", nil)
}

func authHandler(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	redirectTarget := "/"
	isAuth := userRepository.Authenticate(username, password)
	if isAuth {
		setSession(username, w)
		redirectTarget = "/"

	} else {
		println("NO")
	}
	http.Redirect(w, r, redirectTarget, 302)
}

func createDishHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "upload", "base", nil)
}

func saveDishHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 * MiB_UNIT)
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		return
	}

	saveFile(file, handler)
}

func saveFile(file multipart.File, handle *multipart.FileHeader) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile("./images/"+handle.Filename, data, 0666)
	if err != nil {
		log.Println(err)
		return
	}

}

func setSession(username string, w http.ResponseWriter) {
	value := map[string]string{
		"name": username,
	}
	if encoded, err := cookieHandler.Encode("session", value); err != nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func menuHandler(w http.ResponseWriter, r *http.Request) {
	var dishes []models.Dish
	var menu models.Menu

	params := mux.Vars(r)
	date := params["date"]
	err, menu := menuRepository.FindByDate(date)
	if err != nil {
		log.Println(err)
		return
	}

	var aux models.Dish
	for _, v := range menu.Dishes {
		err, aux = dishRepository.FindById(v)
		dishes = append(dishes, aux)

	}

	renderTemplates(w, "menu", "base", dishes)
}

func createMenuHandler(w http.ResponseWriter, r *http.Request) {
	var dishes []models.Dish

	_, dishes = dishRepository.FindAll()

	renderTemplates(w, "menu1", "base", dishes)
}

func renderTemplates(w http.ResponseWriter, name string, template string, viewModel interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist.", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func init() {

	/* 	Conexion a la base datos*/
	url := "127.0.0.1:27017"
	db := "restaurant"
	daoDb = &dbAdapters.MongoDbAdapter{url, db}
	daoDb.Connect()

	// Se tiene una variable compartida Dao en el paquete adapters/database
	// esta variable se pasara a los repositorios para que sean almacenados como miembros
	dishRepository = &dbAdapters.DishRepository{dbAdapters.Dao}
	userRepository = &dbAdapters.UserRepository{dbAdapters.Dao}
	menuRepository = &dbAdapters.MenuRepository{dbAdapters.Dao}

	// 	API //
	api = apiAdapter.API{dishRepository, userRepository, menuRepository}

	// MOCKS //
	/*mockDatabase = mocksDb.MockDatabase{}
	mockDatabase.Connect()
	mockDishRepository = mocksDb.MockDishRepository{mockDatabase}

	api = apiAdapter.Api{mockDishRepository}*/

	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	templates["index"] = template.Must(template.ParseFiles("templates/index.html",
		"templates/base.html"))
	templates["login"] = template.Must(template.ParseFiles("templates/login.html",
		"templates/base.html"))
	templates["upload"] = template.Must(template.ParseFiles("templates/upload.html",
		"templates/base.html"))
	templates["menu"] = template.Must(template.ParseFiles("templates/menu.html",
		"templates/base.html"))
	templates["menu1"] = template.Must(template.ParseFiles("templates/updateMenu.html",
		"templates/base.html"))
}

func loggingHandler(next http.Handler) http.Handler {

	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	return handlers.LoggingHandler(logFile, next)
}

func main() {
	//var productRepo pPort.ProductRepository

	/*var buy models.Buy
	ps := []string{"2", "1", "3"}
	buy.UserId = "lt1235"
	buy.Dishes = ps
	buy.DoneAt = time.Now()
	buy.Total = 12.3

	fmt.Println(strconv.FormatFloat(a, 'f', 4, 64))
	fmt.Println("************COMPRA**********")
	for k, v := range buy.Dishes {
		fmt.Println("Producto " + strconv.Itoa(k))
		err, prodAux := dishRepository.FindByDishId(v)
		fmt.Println(prodAux)
		if err != nil {
			fmt.Println(err)
		}
	}
	*/

	r := mux.NewRouter()
	commonHandlers := alice.New(loggingHandler, handlers.CompressHandler)

	//STATIC
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images/"))))

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/auth", authHandler)
	r.HandleFunc("/crearPlato", createDishHandler)
	r.HandleFunc("/menu/{date}", menuHandler)
	r.HandleFunc("/crearMenu", createMenuHandler)

	r.HandleFunc("/guardarPlato", saveDishHandler).Methods("POST")

	//r.Handle("/auth", commonHandlers.ThenFunc(http.HandlerFunc(authHandler))).Methods("POST")

	//API
	r.Handle("/api/platos", commonHandlers.ThenFunc(http.HandlerFunc(api.GetAllDishesHandler))).Methods("GET")
	r.Handle("/api/platos/{id:[0-9]+}", commonHandlers.ThenFunc(http.HandlerFunc(api.GetDishHandler))).Methods("GET")
	r.Handle("/api/platos", commonHandlers.ThenFunc(http.HandlerFunc(api.PostDishHandler))).Methods("POST")
	r.Handle("/api/platos/{id:[0-9]+}", commonHandlers.ThenFunc(http.HandlerFunc(api.DeleteDishHandler))).Methods("DELETE")
	r.Handle("/api/platos/{id:[0-9]+}", commonHandlers.ThenFunc(http.HandlerFunc(api.PutDishHandler))).Methods("PUT")

	r.Handle("/api/usuarios", commonHandlers.ThenFunc(http.HandlerFunc(api.GetAllUsersHandler))).Methods("GET")
	r.Handle("/api/usuarios/{id:[0-9]+}", commonHandlers.ThenFunc(http.HandlerFunc(api.GetUserHandler))).Methods("GET")
	r.Handle("/api/usuarios", commonHandlers.ThenFunc(http.HandlerFunc(api.PostUserHandler))).Methods("POST")
	r.Handle("/api/usuarios/{id:[0-9]+}", commonHandlers.ThenFunc(http.HandlerFunc(api.DeleteUserHandler))).Methods("DELETE")
	r.Handle("/api/usuarios/{id:[0-9]+}", commonHandlers.ThenFunc(http.HandlerFunc(api.PutUserHandler))).Methods("PUT")

	r.Handle("/api/menu", commonHandlers.ThenFunc(http.HandlerFunc(api.PostMenuHandler))).Methods("POST")
	r.Handle("/api/menu", commonHandlers.ThenFunc(http.HandlerFunc(api.GetAllMenuHandler))).Methods("GET")
	r.Handle("/api/menu/{date}", commonHandlers.ThenFunc(http.HandlerFunc(api.GetDailyMenuHandler))).Methods("GET")

	err := http.ListenAndServe(":9000", r)
	log.Println("...")
	fmt.Println("...")
	if err != nil {
		log.Fatal(err)
	}
}
