
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// php@8.2Formula represents a formula in Go.
type php@8.2Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg php@8.2Formula) Print() {
    fmt.Printf("Name: php@8.2\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := php@8.2Formula{
        Description:  "General-purpose scripting language",
        Homepage:     "https://www.php.net/",
        URL:          "https://www.php.net/downloads",
        Sha256:       "02755d79b5d09e41f038c63f6ea7be828f83f82224ccaa01c04144e65b8f7953",
        Dependencies: []string{"httpd", "pkg-config", "apr", "apr-util", "argon2", "aspell", "autoconf", "curl", "freetds", "gd", "gettext", "gmp", "icu4c", "krb5", "libpq", "libsodium", "libzip", "oniguruma", "openldap", "openssl@3", "pcre2", "sqlite", "tidy-html5", "unixodbc"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installphp@8.2(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg php@8.2Formula) Installphp@8.2() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "downloads"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "downloads"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
