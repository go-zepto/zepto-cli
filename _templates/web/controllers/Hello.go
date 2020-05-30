package controllers

import (
    "github.com/go-zepto/zepto/web"
)

func HelloIndex(ctx web.Context) {
	ctx.Render("pages/index")
}