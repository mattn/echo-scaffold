package template

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func LoadTemplate(name string) string {
	return LoadTemplateFromFile(TemplatePath(name))
}

func LoadTemplateFromFile(filePath string) string {
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	return strings.Replace(string(data), "}\\\n", "}", -1)
}

func PackageName() string {
	wd, _ := os.Getwd()
	wd = filepath.ToSlash(wd)
	for _, p := range filepath.SplitList(os.Getenv("GOPATH")) {
		p = filepath.ToSlash(p)
		if strings.HasPrefix(strings.ToLower(wd), strings.ToLower(p)) {
			return wd[len(p+"/src/"):]
		}
	}
	return ""
}

func ImportPath() string {
	paths := filepath.SplitList(os.Getenv("GOPATH"))
	wd, _ := os.Getwd()
	wd = filepath.ToSlash(wd)
	found := ""
	for _, p := range paths {
		p = filepath.ToSlash(p)
		if strings.HasPrefix(strings.ToLower(wd), strings.ToLower(p)) {
			found = p
			break
		}
	}
	if found == "" {
		found = paths[0]
	}
	return filepath.Join(found, "src/github.com/mattn/echo-scaffold")
}

func TemplatePath(name string) string {
	return filepath.Join(ImportPath(), "template/data", name)
}
