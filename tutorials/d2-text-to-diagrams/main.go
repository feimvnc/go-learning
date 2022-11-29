package main

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"strings"

	"oss.terrastruct.com/d2/d2compiler"
	"oss.terrastruct.com/d2/d2exporter"
	"oss.terrastruct.com/d2/d2layouts/d2dagrelayout"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2renderers/textmeasure"
	"oss.terrastruct.com/d2/d2themes/d2themescatalog"
)

func main() {
	graph, _ := d2compiler.Compile("path", strings.NewReader("x -> y: hello world"), &d2compiler.CompileOptions{UTF16: true})
	ruler, _ := textmeasure.NewRuler()
	graph.SetDimensions(nil, ruler)
	d2dagrelayout.Layout(context.Background(), graph)
	diagram, _ := d2exporter.Export(context.Background(), graph, d2themescatalog.NeutralDefault.ID)
	out, _ := d2svg.Render(diagram)
	ioutil.WriteFile(filepath.Join("out.svg"), out, 0600)

}
