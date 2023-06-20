package app

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func (a *App) DownloadDictionaryFiles(repoURL *url.URL, paths ...string) error {
	repoURL, _ = url.Parse("https://raw.githubusercontent.com/nvdajp/nvdajp/alphajp")
	for _, path := range paths {
		if err := a.DownloadDictionaryFile(repoURL, path); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) DownloadDictionaryFile(repoURL *url.URL, path string) error {
	dir := filepath.Join(a.HomeDir, "default", filepath.Dir(path))
	base := filepath.Base(path)

	os.MkdirAll(dir, 0755)

	targetURL := repoURL.JoinPath(path)

	res, err := http.Get(targetURL.String())

	if err != nil {
		return fmt.Errorf("failed to download dictionary: %s: %w", targetURL, err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected HTTP status: %s: %s", res.Status, targetURL)
	}

	file, err := os.Create(filepath.Join(dir, base))

	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := io.Copy(file, res.Body); err != nil {
		return fmt.Errorf("failed to read response: %s: %w", targetURL, err)
	}

	return nil
}
