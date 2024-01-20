package main

import (
	"fmt"
	"net/http"

	"golang.org/x/text/language"
	"lure.sh/lure/pkg/search"
)

type Package struct {
	Name          string              `json:"name"`
	Repository    string              `json:"repository"`
	Version       string              `json:"version"`
	Release       int                 `json:"release"`
	Epoch         uint                `json:"epoch"`
	Description   string              `json:"description"`
	Homepage      string              `json:"homepage"`
	Maintainer    string              `json:"maintainer"`
	Architectures []string            `json:"architectures"`
	Licenses      []string            `json:"licenses"`
	Provides      []string            `json:"provides"`
	Conflicts     []string            `json:"conflicts"`
	Replaces      []string            `json:"replaces"`
	Depends       map[string][]string `json:"depends"`
	BuildDepends  map[string][]string `json:"build_depends"`
}

func (p Package) FullVersion() string {
	ver := fmt.Sprintf("%s-%d", p.Version, p.Release)
	if p.Epoch != 0 {
		ver = fmt.Sprintf("%d:%s", p.Epoch, ver)
	}
	return ver
}

func translatePkgs(langs []string, pkgs []search.Package) []Package {
	out := make([]Package, len(pkgs))
	for i, pkg := range pkgs {
		out[i] = translatePkg(langs, pkg)
	}
	return out
}

func translatePkg(langs []string, pkg search.Package) Package {
	return Package{
		Name:          pkg.Name,
		Repository:    pkg.Repository,
		Version:       pkg.Version,
		Release:       pkg.Release,
		Epoch:         pkg.Epoch,
		Description:   performTranslation(langs, pkg.Description),
		Homepage:      performTranslation(langs, pkg.Homepage),
		Maintainer:    performTranslation(langs, pkg.Maintainer),
		Architectures: pkg.Architectures,
		Licenses:      pkg.Licenses,
		Provides:      pkg.Provides,
		Conflicts:     pkg.Conflicts,
		Depends:       pkg.Depends,
		BuildDepends:  pkg.BuildDepends,
	}
}

func getLanguages(req *http.Request) []string {
	al := req.Header.Get("Accept-Language")
	lang := req.URL.Query().Get("lang")

	if al == "" && lang == "" {
		return nil
	}

	tags, _, _ := language.ParseAcceptLanguage(al)

	if lang != "" {
		langTag, err := language.Parse(lang)
		if err == nil {
			tags = append(tags, langTag)
		}
	}

	return getLangBases(tags)
}

func getLangBases(langs []language.Tag) []string {
	out := make([]string, len(langs)+1)
	for i, lang := range langs {
		base, _ := lang.Base()
		out[i] = base.String()
	}
	return out
}

func performTranslation(langs []string, v map[string]string) string {
	if len(langs) == 0 {
		val, ok := v[""]
		if !ok {
			return "<unknown>"
		}
		return val
	}

	for _, name := range langs {
		val, ok := v[name]
		if !ok {
			continue
		}
		return val
	}

	return "<unknown>"
}
