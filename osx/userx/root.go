package userx

func IAmPrivileged() bool {
	u, err := Current()
	if err != nil {
		return false
	}
	if u.Uid == "0" {
		return true
	}
	return false
}
