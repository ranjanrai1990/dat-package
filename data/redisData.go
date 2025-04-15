package data

type PackageComponent struct {
	Name    string
	Version string
	Author  string
}

func GetPredefinedComponents() []PackageComponent {
	return []PackageComponent{
		{Name: "auth", Version: "1.0.0", Author: "dev1"},
		{Name: "billing", Version: "2.1.0", Author: "dev2"},
		{Name: "notifications", Version: "1.5.2", Author: "dev3"},
	}
}
