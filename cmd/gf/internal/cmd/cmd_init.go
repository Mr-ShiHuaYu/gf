// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Mr-ShiHuaYu/gf/v2/frame/g"
	"github.com/Mr-ShiHuaYu/gf/v2/os/gcmd"
	"github.com/Mr-ShiHuaYu/gf/v2/os/gfile"
	"github.com/Mr-ShiHuaYu/gf/v2/os/gproc"
	"github.com/Mr-ShiHuaYu/gf/v2/os/gres"
	"github.com/Mr-ShiHuaYu/gf/v2/text/gstr"
	"github.com/Mr-ShiHuaYu/gf/v2/util/gtag"

	"github.com/Mr-ShiHuaYu/gf/cmd/gf/v2/internal/utility/allyes"
	"github.com/Mr-ShiHuaYu/gf/cmd/gf/v2/internal/utility/mlog"
	"github.com/Mr-ShiHuaYu/gf/cmd/gf/v2/internal/utility/utils"
)

var (
	// Init .
	Init = cInit{}
)

type cInit struct {
	g.Meta `name:"init" brief:"{cInitBrief}" eg:"{cInitEg}"`
}

const (
	cInitRepoPrefix  = `github.com/gogf/`
	cInitMonoRepo    = `template-mono`
	cInitMonoRepoApp = `template-mono-app`
	cInitSingleRepo  = `template-single`
	// cInitForkRepo is the fork repository prefix used in generated projects.
	cInitForkRepo = `github.com/Mr-ShiHuaYu/gf`
	// cInitForkVersion is the fork version used in generated projects.
	cInitForkVersion = `v2.9.4-go111`
	// cInitGoVersion is the go version used in generated projects.
	cInitGoVersion = `1.11`
	cInitBrief     = `create and initialize an empty GoFrame project`
	cInitEg        = `
gf init my-project
gf init my-mono-repo -m
gf init my-mono-repo -a
`
	cInitNameBrief = `
name for the project. It will create a folder with NAME in current directory.
The NAME will also be the module name for the project.
`
	// cInitGitDir the git directory
	cInitGitDir = ".git"
	// cInitGitignore the gitignore file
	cInitGitignore = ".gitignore"
)

func init() {
	gtag.Sets(g.MapStrStr{
		`cInitBrief`:     cInitBrief,
		`cInitEg`:        cInitEg,
		`cInitNameBrief`: cInitNameBrief,
	})
}

type cInitInput struct {
	g.Meta  `name:"init"`
	Name    string `name:"NAME" arg:"true" v:"required" brief:"{cInitNameBrief}"`
	Mono    bool   `name:"mono" short:"m" brief:"initialize a mono-repo instead a single-repo" orphan:"true"`
	MonoApp bool   `name:"monoApp" short:"a" brief:"initialize a mono-repo-app instead a single-repo" orphan:"true"`
	Update  bool   `name:"update" short:"u" brief:"update to the latest goframe version" orphan:"true"`
	Module  string `name:"module" short:"g" brief:"custom go module"`
}

type cInitOutput struct{}

func (c cInit) Index(ctx context.Context, in cInitInput) (out *cInitOutput, err error) {
	var overwrote = false
	if !gfile.IsEmpty(in.Name) && !allyes.Check() {
		s := gcmd.Scanf(`the folder "%s" is not empty, files might be overwrote, continue? [y/n]: `, in.Name)
		if strings.EqualFold(s, "n") {
			return
		}
		overwrote = true
	}
	mlog.Print("initializing...")

	// Create project folder and files.
	var (
		templateRepoName string
		gitignoreFile    = in.Name + "/" + cInitGitignore
	)

	if in.Mono {
		templateRepoName = cInitMonoRepo
	} else if in.MonoApp {
		templateRepoName = cInitMonoRepoApp
	} else {
		templateRepoName = cInitSingleRepo
	}

	err = gres.Export(templateRepoName, in.Name, gres.ExportOption{
		RemovePrefix: templateRepoName,
	})
	if err != nil {
		return
	}

	// build ignoreFiles from the .gitignore file
	ignoreFiles := make([]string, 0, 10)
	ignoreFiles = append(ignoreFiles, cInitGitDir)
	// in.MonoApp is a mono-repo-app, it should ignore the .gitignore file
	if overwrote && !in.MonoApp {
		err = gfile.ReadLines(gitignoreFile, func(line string) error {
			// Add only hidden files or directories
			// If other directories are added, it may cause the entire directory to be ignored
			// such as 'main' in the .gitignore file, but the path is ' D:\main\my-project '
			if line != "" && strings.HasPrefix(line, ".") {
				ignoreFiles = append(ignoreFiles, line)
			}
			return nil
		})

		// if not found the .gitignore file will skip os.ErrNotExist error
		if err != nil && !os.IsNotExist(err) {
			return
		}
	}

	// Get template name and module name.
	if in.Module == "" {
		in.Module = gfile.Basename(gfile.RealPath(in.Name))
	}
	if in.MonoApp {
		pwd := gfile.Pwd() + string(os.PathSeparator) + in.Name
		in.Module = utils.GetImportPath(pwd)
	}

	// Replace template name to project name and switch to fork repository.
	err = gfile.ReplaceDirFunc(func(path, content string) string {
		for _, ignoreFile := range ignoreFiles {
			if strings.Contains(path, ignoreFile) {
				return content
			}
		}
		c := gfile.GetContents(path)
		// Replace template module name with user's module name.
		c = gstr.Replace(c, cInitRepoPrefix+templateRepoName, in.Module)
		// Replace gogf references with fork repository.
		c = gstr.Replace(c, `github.com/gogf/gf/v2`, cInitForkRepo+`/v2`)
		c = gstr.Replace(c, `github.com/gogf/gf/`, cInitForkRepo+`/`)
		// Replace gogf version with fork version in go.mod.
		c = gstr.Replace(c, cInitForkRepo+`/v2 v2.7.1`, cInitForkRepo+`/v2 `+cInitForkVersion)
		// Replace go version in go.mod.
		c = gstr.Replace(c, "\ngo 1.18\n", "\ngo "+cInitGoVersion+"\n")
		// Add replace directives for compatibility with go 1.11
		if gfile.Ext(path) == ".mod" {
			replaceBlock := `

replace (
	github.com/BurntSushi/toml => github.com/BurntSushi/toml v0.3.1
	github.com/clbanning/mxj/v2 => github.com/clbanning/mxj/v2 v2.5.0
	github.com/fsnotify/fsnotify => github.com/fsnotify/fsnotify v1.4.7
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2
	github.com/grokify/html-strip-tags-go => github.com/grokify/html-strip-tags-go v0.0.1
	github.com/magiconair/properties => github.com/magiconair/properties v1.8.0
	golang.org/x/net => golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	golang.org/x/sys => golang.org/x/sys v0.0.0-20200202164722-d101bd2416d5
	golang.org/x/text => golang.org/x/text v0.3.2
)`
			c += replaceBlock
		}
		return c
	}, in.Name, "*", true)
	if err != nil {
		return
	}

	// Remove go.sum file since it contains references to the original gogf/gf module
	// which will cause dependency mismatches after module path replacement.
	// The go.sum will be regenerated automatically when the user runs go run/go build.
	goSumPath := in.Name + "/go.sum"
	if gfile.Exists(goSumPath) {
		if err = gfile.Remove(goSumPath); err != nil {
			return
		}
	}

	// Update the GoFrame version.
	if in.Update {
		mlog.Print("update goframe...")
		// go get -u github.com/Mr-ShiHuaYu/gf/v2@v2.9.4-go111
		updateCommand := `go get -u github.com/Mr-ShiHuaYu/gf/v2@` + cInitForkVersion
		if in.Name != "." {
			updateCommand = fmt.Sprintf(`cd %s && %s`, in.Name, updateCommand)
		}
		if err = gproc.ShellRun(ctx, updateCommand); err != nil {
			mlog.Fatal(err)
		}
		// go mod tidy
		gomModTidyCommand := `go mod tidy`
		if in.Name != "." {
			gomModTidyCommand = fmt.Sprintf(`cd %s && %s`, in.Name, gomModTidyCommand)
		}
		if err = gproc.ShellRun(ctx, gomModTidyCommand); err != nil {
			mlog.Fatal(err)
		}
	}

	mlog.Print("initialization done! ")
	if !in.Mono {
		enjoyCommand := `gf run main.go`
		if in.Name != "." {
			enjoyCommand = fmt.Sprintf(`cd %s && %s`, in.Name, enjoyCommand)
		}
		mlog.Printf(`you can now run "%s" to start your journey, enjoy!`, enjoyCommand)
	}
	return
}
