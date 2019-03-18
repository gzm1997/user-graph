package controllers

import (
	"github.com/astaxie/beego"
	"relation-graph/graph/modelBase"
	"relation-graph/graph/modelRelation"
)

var Content string
var Users []modelBase.User
var Files []modelBase.File
var Groups []modelBase.Group
var	CreateGroupShareLinks []modelRelation.CreateGroupShareLink
var ClickGroupShareLinks []modelRelation.ClickGroupShareLink
var CreateFileLinks []modelRelation.CreateFileLink
var ClickFileLinks []modelRelation.ClickFileLink
var Start modelBase.User
var MayKnows []modelBase.User

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Content"] = Content
	c.Data["Users"] = Users
	c.Data["Files"] = Files
	c.Data["Groups"] = Groups
	c.Data["CreateGroupShareLinks"] = CreateGroupShareLinks
	c.Data["ClickGroupShareLinks"] = ClickGroupShareLinks
	c.Data["CreateFileLinks"] = CreateFileLinks
	c.Data["ClickFileLinks"] = ClickFileLinks
	c.Data["Start"] = Start
	c.Data["MayKnows"] = MayKnows
	c.TplName = "index.html"
	//fmt.Println("get page data", c.Data)
}
