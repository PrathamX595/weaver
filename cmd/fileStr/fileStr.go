package filestr

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/fiber
var templatesFiber embed.FS

//go:embed templates/http
var templatesHttp embed.FS

//go:embed templates/echo
var templatesEcho embed.FS

//go:embed templates/chi
var templatesChi embed.FS

//go:embed templates/common/envexample.tmpl
var envTmpl string

//go:embed templates/common/gitignore.tmpl
var gitTmpl string

func FileStr(projName string, frameWork string, auth []string) error {
	funcMap := template.FuncMap{
		"name": func() string { return projName },
		"contains": func(val string, slice []string) bool {
			for _, item := range slice {
				if strings.EqualFold(item, val) {
					return true
				}
			}
			return false
		},
		"hasAuth": func(slice []string) bool {
            return len(slice) > 0
        },
	}

	type data struct {
		Name string
		FrameWork string
		Auth []string
	}

	projData := data{
		Name: projName,
		FrameWork: frameWork,
		Auth: auth,
	}

	dirs := []string{"routes", "controller", "utils", "config", "models"}
	
	if len(auth) != 0{
		dirs = append(dirs, "auth")
	}
	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	var templatesFrame embed.FS

	switch frameWork {
	case "fiber":
		templatesFrame = templatesFiber
	case "http":
		templatesFrame = templatesHttp
	case "echo":
		templatesFrame = templatesEcho
	case "chi":
		templatesFrame = templatesChi
	}

	err := fs.WalkDir(templatesFrame, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		paths := "templates/" + frameWork

		if d.IsDir() || !strings.HasSuffix(path, ".tmpl") || !strings.Contains(path, paths) {
			return nil
		}

        if strings.Contains(path, "/auth/") && len(auth) == 0 {
            return nil
        }
        if strings.Contains(path, "authentication.go.tmpl") && len(auth) == 0 {
            return nil
        }
        if strings.Contains(path, "authModel.go.tmpl") && len(auth) == 0 {
            return nil
        }

		content, err := templatesFrame.ReadFile(path)
		if err != nil {
			return err
		}

		tmpl, err := template.New(filepath.Base(path)).Funcs(funcMap).Parse(string(content))
		if err != nil {
			return err
		}

		outPathdes := paths + "/"

		outPath := strings.TrimPrefix(path, outPathdes)
		outPath = strings.TrimSuffix(outPath, ".tmpl")

		if dir := filepath.Dir(outPath); dir != "." {
			if err := os.MkdirAll(dir, 0755); err != nil && !os.IsExist(err) {
				return err
			}
		}

		f, err := os.Create(outPath)
		if err != nil {
			return err
		}
		defer f.Close()

		err = tmpl.Execute(f, projData)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	tmpl, err := template.New("envTmpl").Funcs(funcMap).Parse(envTmpl)
    if err != nil {
        return err
    }
	var f *os.File
	f, err = os.Create(".env.example")
	if err != nil {
		return err
	}

	defer f.Close()

	err = tmpl.Execute(f, projData)
	if err != nil {
		return err
	}

	tmpl2, err := template.New("gitTmpl").Funcs(funcMap).Parse(gitTmpl)
    if err != nil {
        return err
    }

	f1, err := os.Create(".gitignore")
	if err != nil {
		return err
	}

	defer f1.Close()

	err = tmpl2.Execute(f1, projData)
	if err != nil {
		return err
	}
	
	return nil
}