// Code generated for package deploy by go-bindata DO NOT EDIT. (@generated)
// sources:
// deploy/staticresources/aro.openshift.io_clusters.yaml
// deploy/staticresources/master/deployment.yaml
// deploy/staticresources/master/rolebinding.yaml
// deploy/staticresources/master/service.yaml
// deploy/staticresources/master/serviceaccount.yaml
// deploy/staticresources/namespace.yaml
// deploy/staticresources/preview.aro.openshift.io_previewfeatures.yaml
// deploy/staticresources/worker/deployment.yaml
// deploy/staticresources/worker/role.yaml
// deploy/staticresources/worker/rolebinding.yaml
// deploy/staticresources/worker/serviceaccount.yaml
package deploy

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _aroOpenshiftIo_clustersYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x18\x4d\x73\xda\xc8\xf2\xce\xaf\xe8\xca\x3b\xf8\xf0\x1e\x02\xe2\x3c\xef\x86\x9b\x17\x27\x59\x2a\xd9\x98\xb2\x5d\xb9\xa4\x72\x68\x46\x2d\x31\xb1\x34\xa3\x9d\x6e\x91\x38\x5b\xfb\xdf\xb7\x66\x24\x81\xc0\x02\xa3\xa4\x72\x59\x5d\x40\x3d\xfd\xfd\x3d\x1a\x0c\x87\xc3\x01\x16\xfa\x03\x39\xd6\xd6\x4c\x01\x0b\x4d\x5f\x85\x8c\x7f\xe3\xe8\xfe\x57\x8e\xb4\x1d\xad\x27\x83\x7b\x6d\xe2\x29\xcc\x4a\x16\x9b\xdf\x10\xdb\xd2\x29\xba\xa2\x44\x1b\x2d\xda\x9a\x41\x4e\x82\x31\x0a\x4e\x07\x00\x68\x8c\x15\xf4\x60\xf6\xaf\x00\xca\x1a\x71\x36\xcb\xc8\x0d\x53\x32\xd1\x7d\xb9\xa4\x65\xa9\xb3\x98\x5c\x60\xde\x88\x5e\x8f\xa3\x8b\xe8\x7c\x38\x8e\x9e\x8f\x9f\x4f\xc6\x2f\x27\x17\x93\xf3\xf1\x2f\x2f\x2e\x86\x2f\x5f\xbc\x18\x4f\x2e\xfe\x3f\xc1\x0b\x75\x3e\x00\x50\x8e\x02\xf3\x3b\x9d\x13\x0b\xe6\xc5\x14\x4c\x99\x65\x03\x00\x83\x39\x4d\x41\x65\x25\x0b\x39\x8e\xd0\xd9\xc8\x16\x64\x78\xa5\x13\x89\xb4\x1d\x70\x41\xca\x6b\x94\x3a\x5b\x16\x53\x78\x74\x5e\x71\xa8\x95\xae\x0d\xae\x98\x05\x48\xa6\x59\xde\xb6\xa1\xef\x34\x4b\x38\x29\xb2\xd2\x61\xb6\x15\x1d\x80\xac\x4d\x5a\x66\xe8\x36\xe0\x01\x00\x2b\x5b\x50\x9b\x6b\x6d\x7c\x90\x39\xac\x0d\x58\x4f\x30\x2b\x56\x38\xa9\xb8\xa8\x15\xe5\x58\xa9\x04\xe0\xd5\xbd\x5c\xcc\x3f\x9c\xdf\xee\x80\x01\x62\x62\xe5\x74\x21\xc1\x91\x35\x7b\xd0\x0c\xb2\x22\xa8\x70\x21\xb1\x2e\xbc\x36\x4a\xc2\xe5\x62\xbe\xa1\x2f\x9c\x2d\xc8\x89\x6e\xac\xaf\x9e\x56\x62\xb4\xa0\x7b\xd2\xce\xbc\x42\x15\x16\xc4\x3e\x23\xa8\x12\x5b\x9b\x46\x71\x6d\x03\xd8\x04\x64\xa5\x19\x1c\x15\x8e\x98\x4c\x95\x23\x3b\x8c\xc1\x23\xa1\x01\xbb\xfc\x4c\x4a\x22\xb8\x25\xe7\xd9\x00\xaf\x6c\x99\xc5\x3e\x91\xd6\xe4\x04\x1c\x29\x9b\x1a\xfd\x6d\xc3\x9b\x41\x6c\x10\x9a\xa1\x50\x1d\x94\xed\xa3\x8d\x90\x33\x98\xc1\x1a\xb3\x92\xfe\x07\x68\x62\xc8\xf1\x01\x1c\x79\x29\x50\x9a\x16\xbf\x80\xc2\x11\xfc\x61\x1d\x81\x36\x89\x9d\xc2\x4a\xa4\xe0\xe9\x68\x94\x6a\x69\x0a\x42\xd9\x3c\x2f\x8d\x96\x87\x51\xc8\x6d\xbd\x2c\xc5\x3a\x1e\xc5\xb4\xa6\x6c\xc4\x3a\x1d\xa2\x53\x2b\x2d\xa4\xa4\x74\x34\xc2\x42\x0f\x83\xea\x26\x14\x45\x94\xc7\xff\x71\x75\x09\xf1\xd9\x8e\xae\xf2\xe0\xd3\x83\xc5\x69\x93\xb6\x0e\x42\x2e\x1e\x89\x80\xcf\x4a\x1f\x6d\xac\x49\x2b\x2b\xb6\x8e\xf6\x20\xef\x9d\x9b\x57\xb7\x77\xd0\x88\x0e\xc1\xd8\xf7\x7e\xf0\xfb\x96\x90\xb7\x21\xf0\x0e\xd3\x26\x21\x57\x05\x31\x71\x36\x0f\x3c\xc9\xc4\x85\xd5\x46\xea\xdc\xd2\x64\xf6\xdd\xcf\xe5\x32\xd7\xe2\xe3\xfe\x67\x49\x2c\x3e\x56\x11\xcc\x42\x97\x80\x25\x41\x59\xc4\x28\x14\x47\x30\x37\x30\xc3\x9c\xb2\x19\x32\xfd\xf4\x00\x78\x4f\xf3\xd0\x3b\xf6\xb4\x10\xb4\x1b\xdc\x3e\x72\xe5\xb5\xd6\x41\xd3\x68\x0e\xc4\xab\xae\xcf\xdb\x82\xd4\x4e\xc5\xc4\xc4\xda\xf9\x9c\x16\x14\xf2\x95\xd0\xee\x3e\xcd\xd3\x5d\xa9\xfe\x41\xe5\xae\x6c\x8e\xda\xec\x1f\x1c\x34\x0a\xaa\x1a\x9f\x1b\x99\x2f\xfa\x11\xb5\xbc\xdb\xd9\x21\xb6\xf4\xbe\xf8\xd2\x3d\x1b\x00\xf0\xdb\x2b\xb3\xd6\xce\x9a\x9c\x8c\xf4\x12\xbd\x44\x63\xc8\x3d\x26\xd9\xf1\xf0\x6f\x01\x69\xe3\x5c\x9d\x00\x36\xb0\xba\x95\x2c\xc9\xff\xfb\x62\x9a\xc6\xa1\xc2\x64\x7b\xa4\xe7\x31\x7f\x43\x3d\xda\x3a\x2d\x78\xc2\x8a\x83\xa9\x13\x98\x56\x61\x6f\xc6\xec\x1b\x3f\xae\xe6\x71\x2f\x2f\xc5\xfd\x13\x21\x45\xa1\x2f\xf8\x50\xa5\x50\x87\xb1\x5a\x28\xef\xf4\xc1\x09\x66\xa2\x73\xf8\xd0\x2d\x6f\xe1\xf4\x1a\x85\x5e\xd5\x6d\xa4\x67\x22\xa6\x64\x68\x8d\xef\x6c\x9a\x6a\x93\x3e\xa6\x7c\x32\x78\x89\x4e\x0f\xe6\x6f\x60\x80\xe2\x67\xc7\x14\xce\x3e\x8e\x87\x2f\x3f\xfd\x37\xaa\x7e\xce\xfa\xc7\x1b\x20\xb7\x46\x8b\xf5\x87\x6f\x66\xb7\x97\x4a\xd9\xf2\x50\xe2\x90\x29\xf3\xee\x93\x21\x5c\xde\x5c\x37\xeb\x87\x4d\x79\xfe\xfe\xee\x24\xbc\xc5\xcd\xf5\xd5\x49\x88\x3f\x6c\xd8\xd1\xba\x7e\xca\xb8\x2b\x8d\xa9\xb1\x2c\x5a\xf1\xc2\xd9\xf8\x00\xd6\xdd\xe3\x11\xdf\x1c\xcd\xf0\x35\x6a\x97\xe0\xd7\x1f\xb6\xe3\xbd\x5f\x05\x0b\x54\xf4\x2f\x08\xd1\x91\x5e\xa3\x4d\xe2\xb0\x67\x73\xd1\x26\x75\xc4\xdc\xb3\x54\xab\x2d\x8c\x64\xb6\x22\x75\xdf\xd5\xc1\x8f\x17\x6b\xe9\xb2\x4e\xf8\x91\xc6\xf4\x84\x42\x6d\x84\xae\x06\x75\xd4\x6f\x99\x55\x61\x7d\xed\xe5\x02\x6f\x1f\x8a\x75\x49\x86\x69\x87\xc6\x18\xc7\xe1\x36\x85\xd9\xe2\xa8\x2b\x8e\xda\xb4\x33\x07\xaf\x6b\x81\xaf\xbd\xc0\xcd\x38\x4c\x08\xfd\xe0\x0e\x1d\x98\x37\xf7\x82\xcb\x9b\xeb\x0d\x7e\x1f\x4f\x34\xfb\x64\x57\x12\xed\x28\xd3\xcc\xb3\xf9\x55\x73\x33\xb9\xfc\xe6\xd5\xd8\x32\xa8\xae\x08\xd4\xba\x30\x9d\x6c\xf8\xda\x90\xf4\x4a\xe3\x43\xab\x9b\xa0\x94\x7c\xc2\xf2\x16\xf0\x76\xd6\x37\xbb\x64\xbf\x2b\x7f\xf7\xfe\xa6\xac\xa9\xe2\xdf\x67\xf8\x76\x86\x7b\xd6\x70\xf2\x8e\xfe\x5c\x72\xb5\x9c\xb3\xa0\x89\xd1\xc5\x5b\x41\x90\x68\xca\x62\x8e\x3a\xf8\x1e\x2f\x46\x80\x0c\x59\xee\x1c\x1a\xd6\xcd\x5d\xfc\x50\x05\x26\xd6\xe5\x28\x53\xf0\x5b\xfe\x50\x74\x4e\xdf\x5b\xa9\x39\x31\x63\x7a\x50\xce\x93\xf4\x8e\x90\x0f\xcd\xf9\x13\xc8\xbb\x32\xa3\x07\x79\x40\xf8\x3e\xe2\x23\xc5\x77\xac\x7f\x35\xdd\xe6\x89\x05\xbd\x53\xac\xa3\xf8\x77\x94\xb7\xf4\xc0\x8b\xea\x26\xf8\xb3\xf7\xc1\x4e\x1b\x1f\x01\xab\x02\x9b\x82\xb8\xb2\x4a\x23\x16\xeb\x7c\x4e\xb4\x20\xe5\x72\x73\xb9\x6e\xb4\xab\x43\x07\x7f\xfd\x3d\xd8\x46\x11\x95\xa2\x42\x28\x7e\xbf\xff\xcd\xe7\xd9\xb3\x9d\x8f\x3a\xe1\xb5\x55\x9b\xf0\xf1\xd3\xa0\x12\x4c\xf1\x87\xe6\xf3\x8d\x07\xfe\x13\x00\x00\xff\xff\xec\xb4\xc0\xad\x4c\x13\x00\x00")

func aroOpenshiftIo_clustersYamlBytes() ([]byte, error) {
	return bindataRead(
		_aroOpenshiftIo_clustersYaml,
		"aro.openshift.io_clusters.yaml",
	)
}

func aroOpenshiftIo_clustersYaml() (*asset, error) {
	bytes, err := aroOpenshiftIo_clustersYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "aro.openshift.io_clusters.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\xc1\x6e\xdb\x30\x0c\xbd\xfb\x2b\x88\xde\xdd\xa4\xb7\x42\xb7\x62\x0d\x7a\x19\x82\x62\x59\x77\x67\x64\x26\x16\x22\x8b\x02\x49\x07\x75\xbf\x7e\x10\x92\x28\xce\x0a\x64\x3a\x19\x7c\x8f\xef\x3d\xd2\xc4\x1c\xfe\x90\x68\xe0\xe4\x00\x73\xd6\xc5\xf1\xa9\x39\x84\xd4\x39\x78\xa5\x1c\x79\x1a\x28\x59\x33\x90\x61\x87\x86\xae\x01\x88\xb8\xa5\xa8\xe5\x0b\x4a\x83\x03\x14\x6e\x39\x93\xa0\xb1\xb4\x03\xaa\x91\x34\x00\x09\x07\xba\x87\x69\x46\x4f\x0e\x38\x53\xd2\x3e\xec\xac\xc5\xaf\x51\xa8\x92\x1b\xcd\xe4\x8b\x89\x50\x8e\xc1\xa3\x3a\x78\x6a\x00\x94\x22\x79\x63\x39\xd9\x0f\x68\xbe\xff\x39\xcb\x73\x37\x91\x9a\xa0\xd1\x7e\x3a\x51\x85\x63\x0c\x69\xff\x91\x3b\x34\xba\x74\x0f\xf8\xb9\x19\x65\x4f\x27\xb3\x73\xe5\x23\xe1\x11\x43\xc4\x6d\x24\x07\xcb\x06\xc0\x68\xc8\xb1\x76\xcd\x77\x53\x5e\xbc\xc9\x73\x37\x11\xc0\x65\xca\xf2\x3c\x27\xc3\x90\x48\x6a\x73\x0b\x9e\x87\x01\x53\x77\x55\x6b\x8b\xd4\x55\x5b\xf6\x3a\xc7\xea\xf6\xae\xa5\x99\x59\x79\x61\xc0\x32\xde\xdb\x6a\xbd\xfa\xf5\xf2\x7b\xf5\x5a\x81\xef\xff\xab\x42\x99\xc5\x6e\x6c\x6a\xd2\x77\x16\x73\xf0\xbc\x7c\x5e\x56\xf4\xa2\xd4\x9b\xe5\x5a\x8c\xe1\x48\x89\x54\xdf\x85\xb7\xe4\x66\xdc\xc2\x7a\x23\x9b\x97\x00\x32\x5a\xef\x60\xd1\x13\x46\xeb\xbf\x16\x42\xd8\x4d\xb7\x84\x7f\x6d\x13\x77\xb4\xb9\x39\x8d\x4b\xb5\x15\x8e\xf4\x78\x18\xb7\x24\x89\x8c\xf4\x31\xf0\xe2\xb4\x12\x07\x0f\x0f\x67\xaa\x92\x1c\x83\xa7\x17\xef\x79\x4c\xb6\xbe\x73\xb9\xdf\xd9\xf7\x98\x59\x02\x4b\xb0\xe9\x47\x44\xd5\x93\xac\x4e\x6a\x34\xb4\x3e\x8e\x85\xd7\x7a\x09\x16\x3c\xc6\x73\x83\x71\x2c\x3a\x81\xd3\xec\x06\x0e\x34\xb9\xff\xcc\x52\x47\xbe\xe4\x70\xb0\xfa\x0c\x6a\x5a\x01\xda\xed\xc8\x9b\x83\x35\x6f\x7c\x4f\xdd\x18\xa9\xf9\x1b\x00\x00\xff\xff\x57\x5c\x5d\xa2\xfa\x03\x00\x00")

func masterDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterDeploymentYaml,
		"master/deployment.yaml",
	)
}

func masterDeploymentYaml() (*asset, error) {
	bytes, err := masterDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterRolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\xb1\x4e\x03\x31\x0c\x40\x77\x7f\x85\x7f\x20\x87\xd8\x50\x36\x60\x60\x2f\x12\xbb\x9b\xb8\xd4\xf4\x62\x47\x8e\xd3\xa1\x5f\x8f\xaa\xa2\x5b\x90\x6e\xb5\xdf\xf3\x33\x75\xf9\x62\x1f\x62\x9a\xd1\x8f\x54\x16\x9a\x71\x36\x97\x1b\x85\x98\x2e\x97\x97\xb1\x88\x3d\x5d\x9f\xe1\x22\x5a\x33\xbe\xaf\x73\x04\xfb\xc1\x56\x7e\x13\xad\xa2\xdf\xd0\x38\xa8\x52\x50\x06\x44\xa5\xc6\x19\xc9\x2d\x59\x67\xa7\x30\x4f\x8d\xee\x02\xb8\xad\x7c\xe0\xd3\x1d\xa2\x2e\x1f\x6e\xb3\xef\x04\x01\xf1\x5f\x6f\x3b\x5f\x1e\xb3\x44\xb5\x89\xc2\x98\xc7\x1f\x2e\x31\x32\xa4\x3f\xe7\x93\xfd\x2a\x85\x5f\x4b\xb1\xa9\xb1\xfb\xd5\x63\x37\x3a\x15\xce\x68\x9d\x75\x9c\xe5\x14\x89\x6e\xd3\x79\x83\xe1\x37\x00\x00\xff\xff\x4f\x98\xa4\x7c\x24\x01\x00\x00")

func masterRolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterRolebindingYaml,
		"master/rolebinding.yaml",
	)
}

func masterRolebindingYaml() (*asset, error) {
	bytes, err := masterRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8d\x41\xca\xc2\x40\x0c\x46\xf7\x73\x8a\x5c\x60\xa0\xff\xae\xcc\x29\x7e\x10\xdc\x87\xe9\xa7\x1d\xb4\x93\x90\xc4\x2e\x3c\xbd\xd4\x16\x5d\xb9\x0b\xef\x7b\xbc\xb0\xb6\x33\xcc\x9b\xf4\x42\xeb\x5f\xba\xb5\x3e\x15\x3a\xc1\xd6\x56\x91\x16\x04\x4f\x1c\x5c\x12\x51\xe7\x05\x85\xd8\x24\x8b\xc2\x38\xc4\xf2\xc2\x1e\xb0\x63\x73\xe5\x8a\x42\xa2\xe8\x3e\xb7\x4b\x64\x7e\x3e\x0c\x1f\x39\xb9\xa2\x6e\x1d\xc7\x1d\x35\xc4\xb6\x9b\x88\x55\x7f\x45\x55\x2c\x7c\xb7\xf2\xf1\x7d\x8e\xd0\x37\xd8\xd7\x42\xe3\x30\x0e\x07\x08\xb6\x2b\xe2\xff\x8b\x5f\x01\x00\x00\xff\xff\x10\x70\xf6\x36\xda\x00\x00\x00")

func masterServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterServiceYaml,
		"master/service.yaml",
	)
}

func masterServiceYaml() (*asset, error) {
	bytes, err := masterServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterServiceaccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xca\x31\x8e\x02\x31\x0c\x05\xd0\x3e\xa7\xf0\x05\x52\x6c\xeb\x6e\xcf\x80\x44\xff\x95\xf9\x08\x0b\xc5\x8e\x1c\xcf\x14\x9c\x9e\x06\x51\xbf\x87\x65\x77\xe6\xb6\x70\x95\xeb\xaf\xbd\xcc\x0f\x95\x1b\xf3\xb2\xc1\xff\x31\xe2\xf4\x6a\x93\x85\x03\x05\x6d\x22\x8e\x49\x15\x64\xf4\x58\x4c\x54\x64\x9f\xd8\xc5\xfc\xda\x5e\x18\x54\x89\x45\xdf\x4f\x7b\x54\xc7\xfb\x4c\xfe\x72\xfb\x04\x00\x00\xff\xff\xe4\xf5\x04\x25\x70\x00\x00\x00")

func masterServiceaccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterServiceaccountYaml,
		"master/serviceaccount.yaml",
	)
}

func masterServiceaccountYaml() (*asset, error) {
	bytes, err := masterServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _namespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xca\xb1\x0d\x02\x31\x0c\x05\xd0\x3e\x53\x58\xd7\x07\x44\x9b\x21\x28\xe9\xbf\x2e\x1f\x61\x41\xec\x28\x36\x14\x4c\x8f\xa8\xae\x7f\x98\x7a\xe3\x0a\x75\x6b\xf2\xb9\x94\xa7\x5a\x6f\x72\xc5\x60\x4c\xec\x2c\x83\x89\x8e\x44\x2b\x22\x86\xc1\x26\x3e\x69\xf1\xd0\x7b\x56\x7c\xdf\x8b\xd5\x27\x17\xd2\x57\x11\x81\x99\x27\x52\xdd\xe2\xef\xe5\xb0\x27\xf5\xb3\x79\x67\x0d\xbe\xb8\xa7\xaf\x26\xdb\x56\x7e\x01\x00\x00\xff\xff\xc1\xaf\xa6\x4c\x7c\x00\x00\x00")

func namespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_namespaceYaml,
		"namespace.yaml",
	)
}

func namespaceYaml() (*asset, error) {
	bytes, err := namespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _previewAroOpenshiftIo_previewfeaturesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\x4b\x73\xdb\x36\x10\xbe\xeb\x57\xec\xa4\x07\x5f\x42\x5a\x8a\x53\xb5\xd1\xcd\x75\x1e\xf5\xe4\x61\x8f\xed\x71\xa7\xd3\xe9\x61\x45\x2c\x29\xc4\x20\xc0\x00\x0b\xb9\x6a\xa7\xff\xbd\x03\x90\xb4\x44\x8a\xb4\x9c\x4c\x8a\x93\x88\xc7\xb7\xaf\x6f\x77\x01\x4d\x92\x24\x99\x60\x25\x6f\xc9\x3a\x69\xf4\x02\xb0\x92\xf4\x17\x93\x0e\x5f\x2e\xbd\xfb\xd9\xa5\xd2\x1c\xaf\x67\x93\x3b\xa9\xc5\x02\xce\xbc\x63\x53\x5e\x91\x33\xde\x66\xf4\x9a\x72\xa9\x25\x4b\xa3\x27\x25\x31\x0a\x64\x5c\x4c\x00\x50\x6b\xc3\x18\xa6\x5d\xf8\x04\xc8\x8c\x66\x6b\x94\x22\x9b\x14\xa4\xd3\x3b\xbf\xa4\xa5\x97\x4a\x90\x8d\xe0\xad\xe8\xf5\x34\x9d\xa7\x27\xc9\x34\x7d\x31\x7d\x31\x9b\xbe\x9a\xcd\x67\x27\xd3\x9f\x5e\xce\x93\x57\x2f\x5f\x4e\x67\xf3\x1f\x67\x38\xcf\x4e\x26\x00\x99\xa5\x08\x7e\x23\x4b\x72\x8c\x65\xb5\x00\xed\x95\x9a\x00\x68\x2c\x69\x01\x95\xa5\xb5\xa4\xfb\x9c\x90\xbd\x25\x97\x36\xdf\x29\x5a\x93\x9a\x8a\xb4\x5b\xc9\x9c\x53\x69\x26\xae\xa2\x2c\x28\x58\x58\xe3\xab\x87\x73\xfb\xfb\x6a\xe0\xc6\x96\xda\x0f\x97\xf5\xde\xb7\xb5\x8c\xb8\xa0\xa4\xe3\xf7\x03\x8b\x1f\xa4\xe3\xb8\xa1\x52\xde\xa2\xda\xd3\x2f\xae\x39\xa9\x0b\xaf\xd0\xf6\x57\x27\x00\x2e\x33\x15\x2d\xe0\x4c\x79\xc7\x64\x27\x00\x8d\xbf\xa2\x3e\x49\x63\xf3\x7a\x86\xaa\x5a\xe1\xac\x06\xcb\x56\x54\x62\xad\x2e\x40\x30\xe5\xf4\xf2\xfc\xf6\xe4\xba\x33\x0d\x20\xc8\x65\x56\x56\x1c\x7d\xdf\xd5\x19\xa4\x03\x5e\x11\xd4\x47\x20\x37\x36\x7e\x36\xba\x41\xa3\x1c\x9c\x5e\x9e\x3f\xa0\x55\xd6\x54\x64\x59\xb6\x7e\xaa\xc7\x0e\xb3\x76\x66\x7b\xb2\x8f\x82\x7a\xf5\x2e\x10\x81\x52\x54\x4b\x6f\x0c\x25\xd1\x58\x04\x26\x07\x5e\x49\x07\x96\x2a\x4b\x8e\x74\x4d\xb2\x0e\x30\x84\x4d\xa8\xc1\x2c\x3f\x53\xc6\x29\x5c\x93\x0d\x30\xe0\x56\xc6\x2b\x11\x98\xb8\x26\xcb\x60\x29\x33\x85\x96\x7f\x3f\x60\x3b\x60\x13\x85\x2a\x64\x6a\x02\xb6\x1d\x52\x33\x59\x8d\x0a\xd6\xa8\x3c\x3d\x07\xd4\x02\x4a\xdc\x80\xa5\x20\x05\xbc\xde\xc1\x8b\x5b\x5c\x0a\x1f\x4d\xf0\xa3\xce\xcd\x02\x56\xcc\x95\x5b\x1c\x1f\x17\x92\xdb\x8c\xca\x4c\x59\x7a\x2d\x79\x73\x1c\x93\x43\x2e\x3d\x1b\xeb\x8e\x05\xad\x49\x1d\x3b\x59\x24\x68\xb3\x95\x64\xca\x82\xa3\x8f\xb1\x92\x49\x54\x5d\xc7\xac\x4a\x4b\xf1\x83\x6d\x72\xd0\x1d\x75\x74\xe5\x4d\x20\x8b\x63\x2b\x75\xb1\xb3\x10\x59\xfb\x48\x04\x02\x71\x43\xd0\xb1\x39\x5a\x5b\xb1\x75\x74\x98\x0a\xde\xb9\x7a\x73\x7d\x03\xad\xe8\x18\x8c\xbe\xf7\xa3\xdf\xb7\x07\xdd\x36\x04\xc1\x61\x52\xe7\x64\xeb\x20\xe6\xd6\x94\x11\x93\xb4\xa8\x8c\xd4\x1c\x3f\x32\x25\x49\xf7\xdd\xef\xfc\xb2\x94\x1c\xe2\xfe\xc5\x93\xe3\x10\xab\x14\xce\x62\x99\x81\x25\x81\xaf\x04\x32\x89\x14\xce\x35\x9c\x61\x49\xea\x0c\x1d\xfd\xef\x01\x08\x9e\x76\x49\x70\xec\xd3\x42\xb0\x5b\x21\xfb\x9b\x6b\xaf\xed\x2c\xb4\xa5\x69\x24\x5e\xdd\x6c\xbd\xae\x28\xeb\x24\x4e\x3f\x4f\x43\xfe\x9e\x5e\x5d\x74\xf0\x86\xf3\x35\x0c\xed\x8a\xb7\xca\xdc\x7f\x30\xc5\xde\x52\x4f\x8d\x4f\xd7\xef\xda\x9d\xb1\xc6\xa3\xd4\xf1\x47\x2e\x0b\x6f\x63\x6e\x46\xd1\x9f\xae\xdf\x41\xae\xcc\x3d\x28\x53\xb8\x74\x0f\x12\xe0\xa2\x94\x4d\xf4\x3b\x67\x65\x0e\x1b\xe3\x41\x18\x7d\xc4\x70\x8f\xfa\x61\x4f\xd3\x4c\x42\xca\x86\xc4\xd3\x99\x54\x34\x00\x7b\x40\xee\xb8\x07\xc2\x20\x8d\x4b\x45\x62\x68\xa9\x9f\x3c\x6f\xea\xad\x9d\x10\x2c\x69\x85\x6b\x69\xbc\xad\x8b\x16\x6d\x15\xb5\xc3\x90\x00\xf7\x2b\xd2\xc0\xd6\x53\xdf\xcc\x7b\xa9\x14\xb0\xdd\xd4\xf6\x62\xb6\x8a\x1b\x04\x39\x69\x49\x74\x7d\xf6\xfc\x31\xec\x1c\x95\x23\x90\xdc\x01\x14\xd2\x05\xed\x23\xe4\xd6\x59\x47\x83\x38\x35\x59\x97\xc6\x28\xc2\x7e\xd9\x0d\xc3\x12\xd7\x09\x32\xe6\xb5\x1c\xbd\xe2\x05\xbc\x98\xcd\xa7\xab\xc3\x8e\xbd\x6a\xe1\xa0\x22\x2b\x8d\x88\x64\x0a\x29\x34\xc4\xa1\xd1\xbc\x6b\x87\x63\x63\xb1\xa0\xd3\x2c\x33\x5e\x73\x7b\x85\x39\x7f\x4a\x84\xaf\x87\x8f\xbe\x6e\xbb\xca\x92\x9a\xd8\x8b\xc0\xd9\x5f\x7e\xbf\x48\x9a\x13\x23\xd1\xc0\x1a\x28\x14\x5c\xef\x42\xe1\xfa\xe8\x5d\xac\x64\x52\xc7\x38\x38\x2c\x03\x61\x8a\x60\xba\xc9\x1f\xc2\xf2\x4d\x66\xb3\xc5\x3c\x97\xd9\xa9\x46\xb5\x61\x99\xb9\xf3\xd0\xca\xd6\xa8\x0e\x84\x68\x3e\x2d\x9f\xc0\xfc\x16\x0c\x90\xe1\x7e\x25\x03\x33\x4d\x20\xa4\xf0\x19\xd7\x6a\x63\x2b\x37\x85\xdb\xd8\x18\xc7\xd8\x3f\x9f\x96\xcf\x61\x36\x2d\x53\x78\xbd\xa3\xc3\x08\x11\x49\xfb\x72\x18\x28\x19\xd5\x3c\x09\xe8\xdf\xc3\x83\x1f\x4c\xf1\xf0\xfb\x37\x63\xef\x5c\x85\x4f\xe4\xd1\x15\x7d\xf1\x31\x67\x03\x91\x6f\x7a\xb0\x83\x2c\x18\x71\xd6\x77\xe2\xc6\x7a\xe8\x6a\xb6\xd5\xbc\xcd\xd6\xc3\x86\xf5\x6f\x6f\x0d\x70\x50\x70\xb7\x08\x0f\x2b\xf9\x58\x30\x67\x23\xf3\xc3\x3a\xd5\xe6\x86\xcb\x5a\x11\x6f\xca\xdd\x61\x1b\xef\xef\xcb\x4a\xda\x6a\xbf\xb7\x32\xd2\x9e\xc7\xfb\x36\x23\x7b\xf7\xf4\xce\x1d\xb7\x77\x1a\x87\x59\xba\x70\x5f\x12\x11\x8a\x82\x07\x07\xde\x1a\xed\x18\x6f\x61\x21\x09\xe5\xce\x13\x6c\x77\x48\xa6\x72\x30\x0f\x3b\x9a\x5e\x54\x64\x91\x8d\x3d\x6b\x91\x42\xb9\xfa\x1c\x28\x1a\xc9\xc9\xa8\x05\x5a\xb1\x15\x04\xb9\x24\x25\x06\xfb\xfb\xe3\x9d\x16\x40\xa1\xe3\x1b\x8b\xda\xc9\xf6\x5d\x37\x56\x25\x72\x63\x4b\xe4\x45\xe8\x02\x94\xb0\x1c\x4d\x90\x03\xc4\x0f\xb7\x31\xe7\xb0\x18\x95\x73\xf0\xbc\x25\x74\x63\x99\xf3\x84\xe3\x43\x3c\xf9\x8a\xe3\x71\xc3\xb7\x1d\x1e\x65\xf4\x76\x11\xad\xc5\x4d\xff\x5e\xdf\xb0\x61\xf0\x35\x77\x40\xac\x25\xf1\x2b\xf2\x7b\xda\xb8\xcb\xfa\x51\xf0\x15\x94\x7c\xd4\x9a\x31\x75\x07\x6d\xdc\x9b\xac\xf3\x6c\x11\xaf\x5c\xf5\x44\xdd\xb4\x77\x67\xfc\xf2\xe1\x9d\xd5\x6a\xd7\x84\x0e\xfe\xf9\x77\xb2\x8d\x22\x66\x19\x55\x4c\xe2\x53\xff\x8f\x82\x67\xcf\x3a\x6f\xff\xf8\xb9\x93\x9b\xf0\xc7\x9f\x93\x5a\x30\x89\xdb\xf6\x5d\x1f\x26\xff\x0b\x00\x00\xff\xff\x79\x2d\xb8\x8a\x98\x11\x00\x00")

func previewAroOpenshiftIo_previewfeaturesYamlBytes() ([]byte, error) {
	return bindataRead(
		_previewAroOpenshiftIo_previewfeaturesYaml,
		"preview.aro.openshift.io_previewfeatures.yaml",
	)
}

func previewAroOpenshiftIo_previewfeaturesYaml() (*asset, error) {
	bytes, err := previewAroOpenshiftIo_previewfeaturesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "preview.aro.openshift.io_previewfeatures.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x52\xcb\x6e\xdb\x40\x0c\xbc\xeb\x2b\x88\xdc\x15\x27\xb7\x60\x6f\x41\x63\xe4\x52\x04\x45\xd3\xf4\x4e\xaf\xa6\xd6\xc2\xfb\x02\x49\xbb\x55\xbe\xbe\x10\x64\xcb\x32\x02\x88\x27\x61\x38\x9c\x19\x2e\xc5\x35\xfc\x86\x68\x28\xd9\x11\xd7\xaa\x9b\xd3\x63\x73\x08\xb9\x73\xf4\x82\x1a\xcb\x90\x90\xad\x49\x30\xee\xd8\xd8\x35\x44\x91\x77\x88\x3a\x7e\xd1\x38\xe0\x88\xa5\xb4\xa5\x42\xd8\x8a\xb4\x7f\x8b\x1c\x20\x0d\x51\xe6\x84\xb5\x9e\x56\xf6\x70\x54\x2a\xb2\xf6\xe1\x8f\xb5\xfc\x79\x14\xcc\xe4\x46\x2b\xfc\x68\x22\xa8\x31\x78\x56\x47\x8f\x0d\x91\x22\xc2\x5b\x91\xc9\x3e\xb1\xf9\xfe\xfb\x22\xcf\x6a\x22\x35\x61\xc3\x7e\x98\xa8\x52\x62\x0c\x79\xff\x51\x3b\x36\x5c\xa6\x13\xff\x7b\x3f\xca\x1e\x93\xd9\x19\xf9\xc8\x7c\xe2\x10\x79\x17\xe1\xe8\xa1\x21\x32\xa4\x1a\xe7\xa9\xe5\xdb\x8c\x15\x6f\xf2\xac\x26\x22\xba\x6c\x39\x96\x2f\xd9\x38\x64\xc8\x3c\xdc\x92\x2f\x29\x71\xee\xae\x6a\xed\x28\x75\xd5\x96\xbd\x2e\x7b\xf3\xeb\x5d\xa1\x85\xd9\x58\x21\xf1\xb8\xde\xeb\xf6\x6d\xfb\xf3\xf9\xd7\xf6\x65\x6e\x7c\xbd\xd7\xdc\x8a\xe1\x84\x0c\xd5\x1f\x52\x76\xb8\xda\x11\xf5\x66\xf5\x15\xb6\x84\x88\x2a\x5b\xef\x68\xd3\x83\xa3\xf5\x9f\x1b\x01\x77\xc3\x2d\xa1\x88\x39\x7a\x7a\x78\x7a\x38\xc3\xb9\x74\x78\xbf\x39\xec\x05\x6d\xa5\x44\xdc\x1f\x8e\x3b\x48\x86\x41\xef\x43\xd9\x4c\x0b\x39\xba\xbb\x3b\x53\x15\x72\x0a\x1e\xcf\xde\x97\x63\xb6\xb7\x95\xff\xee\x2b\x7b\x8d\x59\x25\x14\x09\x36\x7c\x8b\xac\x3a\xc9\xea\xa0\x86\xd4\xfa\x78\x54\x83\xb4\x5e\x82\x05\xcf\xb1\xf9\x1f\x00\x00\xff\xff\x4f\x57\x4a\x02\x45\x03\x00\x00")

func workerDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerDeploymentYaml,
		"worker/deployment.yaml",
	)
}

func workerDeploymentYaml() (*asset, error) {
	bytes, err := workerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x8e\xb1\x6e\x2c\x31\x08\x45\x7b\xbe\x82\x1f\xb0\x57\xaf\x7b\x72\x9b\x22\x7d\x14\xa5\x67\x3d\x24\x83\xc6\x63\x2c\xc0\xbb\x52\xbe\x3e\x9a\xd9\x6d\x53\xa5\xe2\x0a\x1d\x0e\x17\x52\x4a\x40\x43\x3e\xd8\x5c\xb4\x17\xb4\x2b\xd5\x4c\x33\x56\x35\xf9\xa6\x10\xed\x79\xfb\xef\x59\xf4\x72\xfb\x07\x9b\xf4\xa5\xe0\x4b\x9b\x1e\x6c\x6f\xda\x18\x76\x0e\x5a\x28\xa8\x00\x62\x35\x3e\x0f\xde\x65\x67\x0f\xda\x47\xc1\x3e\x5b\x03\xc4\x4e\x3b\x17\x24\xd3\xa4\x83\x8d\x42\x2d\xdd\xd5\x36\x36\xb0\xd9\xd8\x0b\x24\xa4\x21\xaf\xa6\x73\xf8\x61\x4a\x07\x9b\x75\x70\xf7\x55\x3e\x23\x8b\x02\xa2\xb1\xeb\xb4\xca\x4f\xa2\x3e\x5a\x38\x20\xde\xd8\xae\xcf\xed\x17\xc7\x39\x9b\xf8\x23\xdc\x29\xea\xfa\x17\xff\xc5\x83\x62\xfe\xf2\x66\x9c\xf6\x23\xcd\xb1\x50\x30\xfc\x04\x00\x00\xff\xff\x30\x78\x19\x41\x50\x01\x00\x00")

func workerRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerRoleYaml,
		"worker/role.yaml",
	)
}

func workerRoleYaml() (*asset, error) {
	bytes, err := workerRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerRolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8d\x31\x6e\xc3\x30\x0c\x45\x77\x9d\x82\x17\x90\x8b\x6e\x85\xb6\xb6\x43\x77\x17\xe8\x4e\xcb\x74\xcd\xda\x26\x05\x8a\x72\x01\x9f\x3e\x08\x12\x64\x09\xe0\xf9\xbf\xf7\x1f\x16\xfe\x21\xab\xac\x92\xc0\x06\xcc\x1d\x36\x9f\xd5\xf8\x40\x67\x95\x6e\x79\xab\x1d\xeb\xcb\xfe\x1a\x16\x96\x31\xc1\xe7\xda\xaa\x93\xf5\xba\xd2\x07\xcb\xc8\xf2\x1b\x36\x72\x1c\xd1\x31\x05\x00\xc1\x8d\x12\xa0\x69\xd4\x42\x86\xae\x16\xff\xd5\x16\xb2\x60\xba\x52\x4f\xd3\x15\xc2\xc2\x5f\xa6\xad\x9c\x04\x03\xc0\x53\xef\xf4\xbe\xb6\xe1\x8f\xb2\xd7\x14\xe2\xdd\xfc\x26\xdb\x39\xd3\x7b\xce\xda\xc4\x4f\xe5\xdb\x56\x0b\x66\x4a\xa0\x85\xa4\xce\x3c\x79\xc4\xa3\x19\x3d\xe0\x70\x09\x00\x00\xff\xff\x73\xce\x57\x9b\x2a\x01\x00\x00")

func workerRolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerRolebindingYaml,
		"worker/rolebinding.yaml",
	)
}

func workerRolebindingYaml() (*asset, error) {
	bytes, err := workerRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerServiceaccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xca\x31\x8a\xc3\x40\x0c\x05\xd0\x7e\x4e\xa1\x0b\x4c\xb1\xad\xba\x3d\x43\x20\xfd\x67\xfc\x43\x84\xb1\x34\x68\x64\x07\x72\xfa\x34\x21\xf5\x7b\x98\x76\x67\x2e\x0b\x57\xb9\xfe\xda\x6e\xbe\xa9\xdc\x98\x97\x0d\xfe\x8f\x11\xa7\x57\x3b\x58\xd8\x50\xd0\x26\xe2\x38\xa8\x82\x8c\x1e\x93\x89\x8a\xec\xaf\xc8\x9d\xf9\xb5\x35\x31\xa8\x12\x93\xbe\x9e\xf6\xa8\x8e\xf7\x99\xfc\xe5\xf6\x09\x00\x00\xff\xff\xe3\x3c\x43\x66\x70\x00\x00\x00")

func workerServiceaccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerServiceaccountYaml,
		"worker/serviceaccount.yaml",
	)
}

func workerServiceaccountYaml() (*asset, error) {
	bytes, err := workerServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"aro.openshift.io_clusters.yaml":                aroOpenshiftIo_clustersYaml,
	"master/deployment.yaml":                        masterDeploymentYaml,
	"master/rolebinding.yaml":                       masterRolebindingYaml,
	"master/service.yaml":                           masterServiceYaml,
	"master/serviceaccount.yaml":                    masterServiceaccountYaml,
	"namespace.yaml":                                namespaceYaml,
	"preview.aro.openshift.io_previewfeatures.yaml": previewAroOpenshiftIo_previewfeaturesYaml,
	"worker/deployment.yaml":                        workerDeploymentYaml,
	"worker/role.yaml":                              workerRoleYaml,
	"worker/rolebinding.yaml":                       workerRolebindingYaml,
	"worker/serviceaccount.yaml":                    workerServiceaccountYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"aro.openshift.io_clusters.yaml": {aroOpenshiftIo_clustersYaml, map[string]*bintree{}},
	"master": {nil, map[string]*bintree{
		"deployment.yaml":     {masterDeploymentYaml, map[string]*bintree{}},
		"rolebinding.yaml":    {masterRolebindingYaml, map[string]*bintree{}},
		"service.yaml":        {masterServiceYaml, map[string]*bintree{}},
		"serviceaccount.yaml": {masterServiceaccountYaml, map[string]*bintree{}},
	}},
	"namespace.yaml": {namespaceYaml, map[string]*bintree{}},
	"preview.aro.openshift.io_previewfeatures.yaml": {previewAroOpenshiftIo_previewfeaturesYaml, map[string]*bintree{}},
	"worker": {nil, map[string]*bintree{
		"deployment.yaml":     {workerDeploymentYaml, map[string]*bintree{}},
		"role.yaml":           {workerRoleYaml, map[string]*bintree{}},
		"rolebinding.yaml":    {workerRolebindingYaml, map[string]*bintree{}},
		"serviceaccount.yaml": {workerServiceaccountYaml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
