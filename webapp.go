//adapted from https://go-macaron.com/
package main

import "gopkg.in/macaron.v1"

func main() {
    m := macaron.Classic()
    m.Use(macaron.Renderer())
    m.Get("/macaron", func() string {
        return "Hello world mac!"
    })

     m.Get("/reverse/:name", func(ctx *macaron.Context) {
    // Adapted from: https://go-macaron.com/docs/middlewares/templating
    ctx.Data["Name"] = Reverse(ctx.Params(":name"))
    ctx.HTML(200, "hello")
    })
    m.Run()

    
}
 func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
 }


