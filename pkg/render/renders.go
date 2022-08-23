package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate Renders template using html template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}
	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

const layoutPath = "./templates/*.layout.tmpl"

func createTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range thru all files ending w *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob(layoutPath)
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(layoutPath)
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}

//var tc = make(map[string]*template.Template)
//
//func RenderTemplateOld(w http.ResponseWriter, t string) {
//	// instead of reading from the disk every time store parsed template into a cache
//	var tmpl *template.Template
//	var err error
//
//	// check to see if we already the template in the cache
//
//	_, inMap := tc[t]
//	if !inMap {
//		log.Println("creating template and adding to cache")
//		// need to create the template
//		err = createTemplateCache(t)
//		if err != nil {
//			log.Println(err)
//		}
//	} else {
//		// we have the template in the cache
//		log.Println("using cached template")
//	}
//
//	tmpl = tc[t]
//	err = tmpl.Execute(w, nil)
//	if err != nil {
//		log.Println(err)
//	}
//}
//
//func createTemplateCache(t string) error {
//	templates := []string{
//		fmt.Sprintf("./templates/%s", t),
//		"./templates/base.layout.tmpl",
//	}
//	// parse the template by taking each indv entry
//	// from templates and put them in as indv strings
//	tmpl, err := template.ParseFiles(templates...)
//	if err != nil {
//		return err
//	}
//	// add template to cache (map)
//	tc[t] = tmpl
//	return nil
//}
