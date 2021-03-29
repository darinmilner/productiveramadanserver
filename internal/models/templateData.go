package models

//TemplateData sent from handlers to HTML templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	CSRFToken string
	FlashMsg  string
	Warning   string
	Error     string
	Data      map[string]interface{}
	Day       int
	Month     string
}
