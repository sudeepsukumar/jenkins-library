package http

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestDownloadRequest(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(
		http.HandlerFunc(
			func(rw http.ResponseWriter, req *http.Request) { rw.Write([]byte("my fancy file content")) }))
	// Close the server when test finishes
	defer server.Close()

	client := Client{
		logger: log.Entry().WithField("package", "SAP/jenkins-library/pkg/http"),
	}

	workingDir, err := ioutil.TempDir("", "test detailed results")
	if err != nil {
		t.Fatal("Failed to create temporary directory")
	}
	// clean up tmp dir
	defer os.RemoveAll(workingDir)
	targetFile := filepath.Join(workingDir, "abc/123/abc.xml")

	// function under test
	err = client.DownloadFile(server.URL, targetFile, nil, nil)
	// asserts
	assert.NoError(t, err, "Error occured but none expected")
	assert.FileExists(t, targetFile, "File not found")
	bytes, err := ioutil.ReadFile(targetFile)
	assert.NoError(t, err, "Error occured but none expected")
	assert.Equal(t, "my fancy file content", string(bytes))
}

func TestDownloadFileIfURL(t *testing.T) {
	t.Run("Should download file in case path begins with http", func(t *testing.T) {
		// Start a local HTTP server
		server := httptest.NewServer(
			http.HandlerFunc(
				func(rw http.ResponseWriter, req *http.Request) { rw.Write([]byte("my fancy file content")) }))
		// Close the server when test finishes
		defer server.Close()

		client := Client{
			logger: log.Entry().WithField("package", "SAP/jenkins-library/pkg/http"),
		}

		workingDir, err := ioutil.TempDir("", "test detailed results")
		if err != nil {
			t.Fatal("Failed to create temporary directory")
		}
		// clean up tmp dir
		defer os.RemoveAll(workingDir)
		targetFile := filepath.Join(workingDir, "abc/123/abc.xml")

		// function under test
		fileName, err := client.DownloadFileIfURL(server.URL, targetFile)
		// asserts
		assert.NoError(t, err, "Error occured but none expected")
		assert.FileExists(t, fileName, "File not found")
		bytes, err := ioutil.ReadFile(fileName)
		assert.NoError(t, err, "Error occured but none expected")
		assert.Equal(t, "my fancy file content", string(bytes))
	})
	t.Run("Should return path to file if path does not begin with http", func(t *testing.T) {
		client := Client{
			logger: log.Entry().WithField("package", "SAP/jenkins-library/pkg/http"),
		}
		workingDir, err := ioutil.TempDir("", "test detailed results")
		if err != nil {
			t.Fatal("Failed to create temporary directory")
		}
		// clean up tmp dir
		defer os.RemoveAll(workingDir)
		targetFile := filepath.Join(workingDir, "abc/123/abc.xml")
		targetFileName := "myFile.xml"

		// function under test
		fileName, err := client.DownloadFileIfURL(targetFile, targetFileName)
		// asserts
		assert.NoError(t, err, "Error occured but none expected")
		assert.Equal(t, targetFile, fileName, "Different file name than expected")
	})
}
