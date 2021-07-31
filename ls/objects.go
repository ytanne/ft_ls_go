package ftls

import "time"

type Object struct {
	Name         string
	HardLinksNum uint64
	IsDir        bool
	Perm         string
	User         string
	Group        string
	ModTime      time.Time
	Size         int64
}

func printObjects(allObjects []Object) {

}
