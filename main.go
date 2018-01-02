package main

import (
	"html/template"
	"os"
	"fmt"
)

func main() {
	tmpl := template.Must(template.ParseGlob("./src/tmpl/*.tmpl"));

	err := tmpl.Execute(os.Stdout, nil);
	if err != nil {
		fmt.Errorf("in template execution: %s", err);
		os.Exit(1);
	}

}
