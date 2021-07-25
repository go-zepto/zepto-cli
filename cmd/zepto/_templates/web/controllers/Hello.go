package controllers

import (
    "github.com/go-zepto/zepto/web"
)

func HelloIndex(ctx web.Context) error {
	return ctx.Render("pages/index")
}