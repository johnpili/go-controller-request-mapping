package controllers

//var (
//cookieStore   *sessions.CookieStore
//viewBox       *rice.Box
//staticBox     *rice.Box
//db            *sqlx.DB
//servicesHub   *services.Hub
//)

//New ...
/*func New(vb *rice.Box, sb *rice.Box, store *sessions.CookieStore, config *models.Config, database *sqlx.DB) *Hub {
	viewBox = vb
	staticBox = sb
	cookieStore = store
	configuration = config
	db = database
	servicesHub = services.New(configuration, db, store)
	hub := new(Hub)
	hub.LoadControllers()
	return hub
}*/

//New ...
func New() *Hub {
	hub := new(Hub)

	// Load the controllers we specified in controllers/z_controller_loader.go
	// This is something similar to old days using web.xml in SpringFramework
	hub.LoadControllers()

	return hub
}

// Hub ...
type Hub struct {
	Controllers []interface{}
}
