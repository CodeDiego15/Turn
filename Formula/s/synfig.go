
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// synfigFormula represents a formula in Go.
type synfigFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg synfigFormula) Print() {
    fmt.Printf("Name: synfig\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := synfigFormula{
        Description:  "Command-line renderer",
        Homepage:     "https://synfig.org/",
        URL:          "https://downloads.sourceforge.net/project/synfig/development/1.5.2/source/synfig-1.5.2.tar.gz",
        Sha256:       "8756ad19dc3c0f2b49a0b07a8e07d4766df49f3aa6cdcc6ae1c95cadab4306b0",
        Dependencies: []string{"autoconf", "automake", "intltool", "libtool", "pkg-config", "cairo", "etl", "ffmpeg", "fftw", "fontconfig", "freetype", "fribidi", "gettext", "glib", "glibmm@2.66", "harfbuzz", "imagemagick", "imath", "libmng", "libpng", "libsigc++@2", "libtool", "libxml++", "mlt", "openexr", "pango", "liblqr", "libomp", "little-cms2", "perl-xml-parser"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installsynfig(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg synfigFormula) Installsynfig() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "synfig-1.5.2.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "synfig-1.5.2.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
