package router

type Group struct {
	UserRouter
}

var GroupApp = new(Group)
