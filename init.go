package main

import(
    "net/http"
    "io/ioutil"
    "github.com/hoisie/mustache"
    "github.com/ricallinson/stackr"
    "github.com/russross/blackfriday"
)

var markdownCache map[string]string = map[string]string{}

func MarkdownFileRender(filepath string) (string) {

    c, ok := markdownCache[filepath]

    if ok == false {
        b, e := ioutil.ReadFile(filepath)
        if e != nil {
            return e.Error()
        }
        o := blackfriday.MarkdownBasic(b)
        markdownCache[filepath] = string(o)
        c = markdownCache[filepath]
    }

    return c
}

var templateCache map[string]*mustache.Template = map[string]*mustache.Template{}

func MustacheFileRender(filepath string, context ...interface{}) string {

    t, ok := templateCache[filepath]

    if ok == false {
        tmpl, err := mustache.ParseFile(filepath)
        if err != nil {
            return err.Error()
        }
        templateCache[filepath] = tmpl
        t = templateCache[filepath]
    }

    return t.Render(context...)
}

func init() {
    app := stackr.CreateServer()
    app.Use(stackr.Static())
    app.Use("/", func(req *stackr.Request, res *stackr.Response, next func()) {        
        res.SetHeader("content-type", "text/html");
        res.End(MustacheFileRender("./tmpls/index.html", map[string]string{
            "title": "Stackr",
            "body": MarkdownFileRender("./files/home.md"),
        }))
    })
    http.Handle("/", app)
}