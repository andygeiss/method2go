package project

import (
	"bytes"
	"io"
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"unicode"

	"github.com/andygeiss/method2go/resources"
)

// TemplateEngine ...
type TemplateEngine interface {
	InitTemplates(p *resources.Project) error
}

// DefaultTemplateEngine ...
type DefaultTemplateEngine struct {
	access fs.FS
	files  []string
	path   string
}

func (a *DefaultTemplateEngine) InitTemplates(p *resources.Project) error {
	for _, file := range a.files {
		src := filepath.ToSlash(filepath.Join(a.path, file))
		raw := safeReadAll(a.access, src)
		p.Templates[file] = raw
		p.Contents[file] = safeExecuteTemplate(raw, p)
	}
	return nil
}

// NewDefaultTemplateEngine ...
func NewDefaultTemplateEngine(access fs.FS, path string, files []string) TemplateEngine {
	return &DefaultTemplateEngine{
		access: access,
		files:  files,
		path:   path,
	}
}

var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchLink = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")

func safeExecuteTemplate(src string, data interface{}) string {
	tmpl, err := template.New("t").Funcs(map[string]interface{}{
		"sc": func(in string) string {
			return toSnakeCase(in)
		},
		"lc": func(in string) string {
			return toCamelCase(toSnakeCase(in))
		},
	}).Parse(src)
	if err != nil {
		return ""
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return ""
	}
	return buf.String()
}

func safeReadAll(access fs.FS, path string) string {
	file, err := access.Open(path)
	if err != nil {
		return ""
	}
	content, err := io.ReadAll(file)
	if err != nil {
		return ""
	}
	return string(content)
}

func toSnakeCase(in string) (out string) {
	out = matchFirstCap.ReplaceAllString(in, "${1}_${2}")
	out = matchAllCap.ReplaceAllString(out, "${1}_${2}")
	return strings.ToLower(out)
}

func toCamelCase(in string) string {
	runes := []rune(in)
	runes[0] = unicode.ToLower(runes[0])
	in = string(runes)
	return matchLink.ReplaceAllStringFunc(in, func(s string) string {
		return strings.ToUpper(strings.Replace(s, "_", "", -1))
	})
}
