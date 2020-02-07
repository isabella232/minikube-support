package coredns

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/qaware/minikube-support/pkg/apis"
	"github.com/qaware/minikube-support/pkg/github"
	"github.com/qaware/minikube-support/pkg/utils/sudos"
	"github.com/sirupsen/logrus"
)

type installer struct {
	ghClient github.Client
	prefix   string
}

const PluginName = "coredns"

func NewInstaller(prefix string, ghClient github.Client) apis.InstallablePlugin {
	return &installer{
		ghClient: ghClient,
		prefix:   prefix,
	}
}

func (i *installer) String() string {
	return PluginName
}

func (i *installer) Install() {
	var errs *multierror.Error
	errs = multierror.Append(errs, sudos.MkdirAll(path.Join(i.prefix, "bin"), 0755))
	errs = multierror.Append(errs, sudos.Chown(i.prefix, os.Getuid(), os.Getgid(), true))

	errs = multierror.Append(errs, os.MkdirAll(path.Join(i.prefix, "etc"), 0755))
	errs = multierror.Append(errs, os.MkdirAll(path.Join(i.prefix, "var", "run"), 0755))
	errs = multierror.Append(errs, os.MkdirAll(path.Join(i.prefix, "var", "log"), 0755))
	errs = multierror.Append(errs, i.writeConfig())

	errs = multierror.Append(errs, i.downloadCoreDns())

	errs = multierror.Append(errs, i.installSpecific())

	if errs.Len() > 0 {
		logrus.Errorf("Unable to install coredns into %s:\n  Errors: %s", i.prefix, errs)
	}
}

func (i *installer) Update() {
	i.Uninstall(false)
	i.Install()
}

func (i *installer) Uninstall(_ bool) {
	var errs *multierror.Error

	errs = multierror.Append(errs, i.uninstallSpecific())
	errs = multierror.Append(errs, sudos.RemoveAll(i.prefix))
	if errs.Len() > 0 {
		logrus.Errorf("Unable to uninstall coredns from %s:\n  Errors: %s", i.prefix, errs)
	}
}

func (i *installer) Phase() apis.Phase {
	return apis.LOCAL_TOOLS_CONFIG
}

func (i *installer) downloadCoreDns() error {
	tagName, e := i.ghClient.GetLatestReleaseTag("coredns", "coredns")
	if e != nil {
		return fmt.Errorf("can not get latest coredns version: %s", e)
	}
	version := strings.TrimPrefix(tagName, "v")

	assetName := fmt.Sprintf("coredns_%s_%s_%s.tgz", version, runtime.GOOS, runtime.GOARCH)
	bytes, e := i.ghClient.DownloadReleaseAsset("coredns", "coredns", tagName, assetName)
	if e != nil {
		return fmt.Errorf("can not download coredns binary: %s", e)
	}

	gzReader, e := gzip.NewReader(bytes)
	if e != nil {
		return fmt.Errorf("can not open gz reader for downloaded coredns: %s", e)
	}
	tarReader := tar.NewReader(gzReader)

	for {
		header, e := tarReader.Next()
		if e == io.EOF {
			break
		}
		if e != nil {
			return fmt.Errorf("unable to extract next file from tar: %s", e)
		}

		if header.Typeflag == tar.TypeReg {
			name := header.Name

			file, e := os.OpenFile(path.Join(i.prefix, "bin", name), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(header.Mode))
			if e != nil {
				return fmt.Errorf("can not write file %s: %s", name, e)
			}

			_, e = io.Copy(file, tarReader)
			if e != nil {
				return fmt.Errorf("can not write file (%s) content: %s", name, e)
			}
		}
	}
	return nil
}
