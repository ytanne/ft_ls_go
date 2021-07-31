package ftls

import (
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func getDotsInfo() []Object {
	fstat, err := os.Stat(".")
	if err != nil {
		panic(err)
	}
	ugstat := fstat.Sys().(*syscall.Stat_t)
	uid := ugstat.Uid
	gid := ugstat.Gid
	nlinks := ugstat.Nlink

	u := strconv.FormatUint(uint64(uid), 10)
	g := strconv.FormatUint(uint64(gid), 10)

	userInfo, _ := user.LookupId(u)
	groupInfo, _ := user.LookupGroupId(g)

	dot := Object{
		Name:         ".",
		HardLinksNum: nlinks,
		IsDir:        fstat.IsDir(),
		Perm:         fstat.Mode().Perm().String(),
		User:         userInfo.Username,
		Group:        groupInfo.Name,
		ModTime:      fstat.ModTime(),
		Size:         fstat.Size(),
	}

	fstat, err = os.Stat("..")
	if err != nil {
		panic(err)
	}
	dotdot := Object{
		Name: "..",
		// IsDir:   fstat.IsDir(),
		// Mode:    uint32(fstat.Mode()),
		// ModTime: fstat.ModTime(),
		// Size:    fstat.Size(),
	}
	return []Object{dot, dotdot}
}
