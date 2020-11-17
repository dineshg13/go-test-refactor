package app

type Application struct {
	Name string
}


func NewApp(name string) *Application {
	return &Application{
		Name: name,
	}
}

func (a Application) String() string {
	return a.Name
}
