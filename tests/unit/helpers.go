package data_parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

type HTTP interface {
	Get(url string) (*http.Response, error)
}

type FileWriter interface {
	WriteFile(filename string, data []byte) error
}

type FileHelper struct {
	HTTPClient HTTP
	FileWriter FileWriter
}

func NewFileHelper(httpClient HTTP, fileWriter FileWriter) *FileHelper {
	return &FileHelper{
		HTTPClient: httpClient,
		FileWriter: fileWriter,
	}
}

func (f *FileHelper) DownloadFile(url string, filename string) error {
	resp, err := f.HTTPClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status code %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return f.FileWriter.WriteFile(filename, data)
}

func (f *FileHelper) MkdirRecursive(path string) error {
	return os.MkdirAll(path, 0755)
}

func (f *FileHelper) GetGoVersion() string {
	return runtime.Version()
}

func (f *FileHelper) GetExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(ex)
}

func (f *FileHelper) LoadJSON(filename string, data interface{}) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, data)
}

func (f *FileHelper) WriteJSON(filename string, data interface{}) error {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return f.FileWriter.WriteFile(filename, b)
}

type Files struct {
	sync.RWMutex
	Files map[string][]byte
}

func NewFiles() *Files {
	return &Files{
		Files: make(map[string][]byte),
	}
}

func (f *Files) Get(filename string) ([]byte, bool) {
	f.RLock()
	defer f.RUnlock()

	data, ok := f.Files[filename]
	return data, ok
}

func (f *Files) Set(filename string, data []byte) {
	f.Lock()
	defer f.Unlock()

	f.Files[filename] = data
}