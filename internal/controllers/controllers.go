package controllers

type MainController struct {
}

// For the case of this app, since its simple, I'm manually supplying the dependancies of MainController. However, in larger projects
// I'd want to use a proper dependancy injection library, something like Container (https://github.com/golobby/container).
// You register several all your services during app start up, and do a "resolve" during in the constructor. It will automatically
// load all the dependancies for you.
func NewMainController() *MainController {
	return &MainController{}
}
