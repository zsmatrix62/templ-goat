package view

// import (
// 	"fmt"
// 	"io"
// 	"path"
// 	"strings"
// 	"text/template"
// 	"github.com/zsmatrix62/templ-goat/pkg/embeds"
// 	"github.com/zsmatrix62/templ-goat/pkg/logger"
// 	"github.com/zsmatrix62/templ-goat/pkg/route"

// 	"github.com/Masterminds/sprig/v3"
// 	_ "github.com/k0kubun/pp"
// )

// func Render(w io.Writer, name string, data interface{}) {
// 	name = fmt.Sprintf("%s.html", strings.ReplaceAll(name, ".", "/"))
// 	templatesDir := "assets/templates/"

// 	tmpl, err := template.New("").
// 		Funcs(template.FuncMap{
// 			"RouteName2URL": route.RouteName2URL,
// 		}).
// 		Funcs(sprig.FuncMap()).
// 		ParseFS(embeds.AssetTemplates,
// 			path.Join(templatesDir, "layouts/*.html"),
// 			path.Join(templatesDir, "components/*.html"),
// 			path.Join(templatesDir, name),
// 		)
// 	logger.LogIf(err)

// 	err = tmpl.ExecuteTemplate(w, name, data)
// 	logger.LogIf(err)
// }
