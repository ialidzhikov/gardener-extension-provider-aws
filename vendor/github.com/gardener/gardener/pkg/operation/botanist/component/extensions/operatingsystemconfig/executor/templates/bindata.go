// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package templates generated by go-bindata.// sources:
// scripts/execute-cloud-config.tpl.sh
package templates

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _scriptsExecuteCloudConfigTplSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x59\x6d\x73\xe2\xb6\xf6\x7f\xef\x4f\x71\xd6\x4b\x9b\xd0\xc6\xd0\xed\xf4\xff\x7f\x91\x96\x4e\x29\x78\xdb\x4c\x13\xd8\x09\xec\xde\xdb\x7b\x77\x87\x0a\xfb\x10\x54\x8c\xe4\x4a\x72\x12\x96\xf0\xdd\xef\x1c\x49\x36\x04\x4c\x92\x3b\x73\xfb\x0a\x5b\x3e\xfa\x9d\xe7\x07\x89\xd7\xaf\xda\x53\x2e\xda\x53\xa6\xe7\x10\x61\x11\x04\xef\xba\xe3\x5f\x27\xbd\xcb\xe1\xfb\x7e\x6f\x38\x78\x7b\xf1\xcb\xa4\x3f\xfc\xc7\xe0\x72\xd8\xed\xc7\xd7\x93\x51\x7c\xfd\x21\xbe\xee\x84\xeb\x35\xb4\x72\x66\xe6\x3d\x85\x29\x0a\xc3\x59\xa6\x47\xa8\x6e\x51\xc1\x66\x13\x3e\x89\xd0\xeb\x4e\x7a\xf1\xf5\xb8\x16\xa2\xd7\xed\xa1\x32\xb5\x10\x5b\xfa\xbe\xbc\x13\x99\x64\x29\xa6\xbd\x4c\x16\x69\x4f\x8a\x19\xbf\xa9\x67\x3b\xbc\xec\x77\xc2\xc6\x7a\x7f\x7d\xd3\x92\x59\x5a\x92\xff\x1a\xf7\x7e\x1b\xbd\xbf\xaa\xc5\x9f\x63\xb2\xd0\xc5\x72\x07\xbc\xd7\x9f\x8c\x7a\xd7\x17\xef\x76\xe5\xef\xf5\x47\x89\xe2\xb9\xa9\x23\xab\xc1\xaf\xe8\x9f\x82\xaf\xf6\x3d\xd6\xe1\xf0\xfb\xae\x2e\xf1\x3f\xe3\xde\xfb\xf1\xc5\x70\x30\xe9\xc7\x97\xdd\xdf\x27\xa3\xb8\x37\x1c\xf4\x47\x5b\xd6\xf1\x3d\x26\x85\xe1\x52\xf4\x31\x63\xab\x11\x26\x52\xa4\x7a\xcb\x7e\xbb\xff\xb2\x3b\x1a\x4f\xfa\xdd\x71\x5c\xb3\xf7\x92\x69\xd3\x67\x06\xed\xbe\x60\xb9\x48\xb9\x82\x28\x87\x03\xfb\xe9\x3e\x57\x98\x18\xa9\x56\x44\xb9\xfd\xfe\x5b\x31\xc5\x0c\xcd\xa3\xaf\x41\x30\x2b\x44\x42\xe8\x90\xca\x64\x81\x2a\xca\x15\x12\xc8\x69\x13\xd6\x01\x80\x60\x4b\xec\x84\x8d\x37\x61\x00\xc0\x97\xec\x86\x5e\xbe\xa5\x17\x4c\xe6\x12\x42\x6b\x49\x2e\x6e\xe0\x6e\x8e\x66\x8e\x0a\x8c\x04\x0f\x00\x0d\xda\x0b\x33\x25\x97\xd0\xb0\x5b\x2d\xc6\x0c\xfe\x0d\xd1\x67\x68\x9c\x3a\x76\x0e\x54\x43\xf4\x17\x84\x9e\xaa\x09\x9f\xbe\x07\x33\x47\x11\x00\x94\x7c\xde\x39\x4c\xe2\x54\x0b\x0b\x5e\x7a\xc8\x8b\x2c\xab\x90\x48\xcc\x4c\xe3\x0e\xce\x40\x82\x40\x4c\x9f\x15\x73\xc6\x83\x4d\x10\xac\xd7\x24\x6f\x6b\xe1\xed\xc6\x0c\xfb\x20\xb3\x62\x89\x10\x6d\x36\x5b\xbb\xcd\xa4\x5a\x32\x13\xa5\xcc\xb0\x28\xc5\x5b\x9e\xa0\xb7\xdd\x65\xf7\xe7\xf8\xb2\xf3\xdb\xfb\x9f\xe3\x7e\xfc\xc1\xe9\xfe\x0a\xa6\xd9\x82\xa7\x10\x45\x19\x9b\x62\x06\x0d\x4b\x03\x3f\xb6\x53\xbc\x6d\x8b\x22\xcb\x76\x34\xef\xc7\x1f\x2e\x7a\xf1\xa8\xd3\x38\xcd\xf4\x34\x5b\x40\x94\x4e\xb5\x78\x07\x91\x84\x41\xf7\x2a\x3e\x7b\xd7\xbd\x1e\x8f\x7f\x7f\x17\x9f\xbd\x1d\xd9\x9f\xd1\xc5\xbf\xe2\xa6\xdd\x78\xd5\x1d\xf7\x7e\xbd\x18\xfc\x32\xd9\x22\x38\xed\x1b\x7e\x21\x84\x07\xb8\x51\x98\xc3\x49\x09\xd2\x09\xc3\xd6\x57\x0e\xc8\x3e\x12\x98\x8b\xc0\x03\xe5\x5b\x9a\x7f\xb6\x21\x78\xd2\xdc\x31\xec\x15\x33\xc9\x9c\xdc\xe3\xe9\x81\xcc\x01\xce\x1c\x30\x5d\x81\xdd\x74\x0e\x4f\x23\x5a\xbc\x71\xf7\xfa\x97\x78\xec\x65\x9f\x90\xaa\x5b\xf9\xf7\x35\x23\x45\xe6\xc8\x52\x88\xc4\x1b\x78\x80\xa4\x30\x10\xcd\xbe\x85\x28\xfd\x18\xee\x0a\x97\xa2\xc1\xc4\x60\x5a\x2b\x5c\xe3\x90\xa1\x13\x64\xb9\x98\xe9\x16\xde\x9b\xef\x20\xba\x2c\x1d\x15\x0d\xe1\xaf\x42\x1a\x06\x51\x0c\x19\xfb\xbc\x9a\x70\xc3\xa6\x19\x4e\xb8\xe0\xa6\xf3\xcd\x99\x5d\xfa\x53\x16\x4a\xb0\xac\x5c\xb3\xf4\x66\x95\x63\xa7\xd0\xca\xbe\x9c\xdf\xa8\xdc\x3d\xe4\xea\x4f\x07\x07\x36\x00\x6a\x44\xd9\xd1\xc2\xc5\x19\xa9\xc1\x44\x0a\x36\x7e\x30\x7d\xa9\x2a\x54\x2b\xda\x66\x99\xb7\x6f\x99\xca\xf8\x34\xc9\xdd\xba\x2c\x84\xf1\x71\xea\x35\x3c\x24\x72\xdc\x2d\x29\x65\x0e\x2e\x73\x48\x64\xbe\x02\x82\xa4\xb2\xf1\x12\x01\x92\x1c\x22\x06\x84\xdb\xce\xf8\xb4\xfd\xd5\x63\x36\x6d\x4b\x53\x38\x69\x8e\x09\x90\xc8\x9c\x63\x5a\x61\x50\x0a\xbf\x4c\xf7\x1a\x1d\x4b\x8c\x48\x42\x8a\x33\x56\x64\x46\x9f\x95\xbe\x38\x43\xa5\xa4\xd2\x1d\x85\x76\x67\xa4\x64\x8d\x15\x2a\x84\x97\x19\xc0\x95\x93\xc3\x42\x11\xac\xd7\x11\xa0\x48\x61\xe3\xaa\x8d\x62\xe2\x06\x5d\x45\x3a\xf3\xd5\x08\xce\x3b\xd0\x2a\x8b\xe4\x66\x13\x3c\xae\xd2\xb6\xb8\xbb\x0a\x56\x96\x7a\xbf\x8d\xd2\x69\xbd\xae\xc0\x13\x66\xe0\x87\x1f\xe0\x24\x1e\xbe\x3d\x81\x07\x98\x32\x8d\xff\xff\x1d\x44\x29\xfc\x08\x61\x63\xbf\x4d\xdb\x9d\xad\x64\xdb\xe4\xdf\x6b\x54\x94\xaf\x04\x15\x0f\xdf\x06\x81\x2d\xe3\xaf\x20\x9a\xd5\xec\xa6\xc6\x19\xee\xd4\x70\x23\x8b\x64\x7e\x8c\x2e\x98\xf1\x3a\xb4\xfa\x76\x7c\x1c\xf5\x08\xbd\x45\x5f\xa6\xff\x47\xbd\xfe\xa0\x93\x6f\xe0\xc7\xc3\xc5\x6d\x7b\x77\x62\x95\x72\xed\xf5\x50\xfa\x49\xac\x69\xae\x91\x65\xd6\xf6\x9f\x3e\xc1\xc3\xc3\xd1\x0d\xbb\x4d\xb7\x9d\x2f\x78\xdb\x97\xa2\x28\xc9\x38\x0a\x13\x25\x85\x52\x28\x4c\x2b\xc7\x25\x41\x55\x6a\x3a\xbf\xc5\xc3\xb7\xe4\xa8\xa3\x42\xfc\x2c\xa5\xd1\x46\xb1\xdc\xba\x3d\x8a\xa2\x80\xe5\xfc\x03\x2a\xcd\xa5\x38\x87\xdb\x37\xc1\x82\x8b\xf4\x1c\x9c\x33\x03\xcf\x2b\x4a\xa4\x30\x78\x6f\xce\xcb\xb2\x18\x4d\x4b\x98\x9f\x7c\x56\x04\x49\x56\x68\x83\x4a\x9f\x07\x11\xf8\xe7\x73\x97\xd0\xa8\x0c\x9f\xf1\x84\x19\x8c\x58\x61\xe6\x52\x71\xb3\xb2\xa1\x7d\x0e\x8d\x53\x92\xba\xc6\xe1\x87\xc3\x68\xb8\x8d\xc5\x07\x30\x8a\x02\xf2\xe4\xa3\xf0\x9d\x45\xdb\xb9\xf6\x65\x78\x6e\x3c\xb6\x55\x9f\xb2\xe1\x1c\x2a\x0d\x9c\x92\x4e\x03\xaf\xb0\xd3\xc0\xab\x53\x51\xda\x1a\xa4\x69\xe5\xc0\x1e\x15\xea\x71\x4b\xd1\x4e\xcb\xe4\x08\x61\xe0\xc1\x2d\x1b\xa6\x23\x7a\x89\xf0\xde\x28\x76\x0e\xeb\x8d\x5d\x35\x72\x81\xc2\xf5\xc8\x6a\xdb\x98\xd6\xaa\xb4\xf3\xd3\x8c\x5a\x3e\x1d\x93\x8f\xc3\x81\x72\x60\x30\xec\xc7\xb6\x95\x06\xdd\xc1\x60\x38\xee\xda\x59\xf3\x3a\x1e\x8d\xbb\xd7\xe3\xc9\xe8\xf7\xd1\x38\xbe\xea\x5b\x1b\xda\x79\x21\xbc\x93\x6a\x81\xaa\x75\xc3\x54\x8a\x02\x95\x2b\x05\x6d\x85\xda\x30\x65\x22\xbd\xd2\x06\x97\x69\x44\xfe\xe1\x09\xea\x30\x08\x5e\xc3\x58\xad\xa8\x24\xcf\xb8\x48\x61\x20\x53\x04\x39\xfd\x13\x13\x43\xd3\x11\x98\x39\xd7\xb0\x64\x34\x20\x20\x4d\x41\x2c\x53\xc8\xd2\x15\x28\xbc\xe1\xe4\x03\x37\x90\x99\x39\x96\x4e\x69\xf9\xdc\x7b\x69\xe6\x55\xe9\xb2\x5e\xff\x41\xaa\x76\xc2\xc6\x69\x5b\xe6\xc6\x9e\xaf\xc8\x15\x89\xc9\x20\x8a\x16\xd5\xde\x4e\xf8\xc7\x66\xf3\x1c\xf2\x7a\xfd\x47\x08\x37\x68\x40\x90\x3e\x51\x06\x21\x01\x28\x81\x06\x75\x8b\xcb\xf6\x5c\x6a\x63\xe7\xe3\xc6\x69\xf9\xd8\x0c\xa9\xa7\xdc\xc8\x88\xda\x64\xc6\x0c\xda\x21\x8a\x86\x48\x6e\x70\xa9\x2d\x28\x9c\x72\x91\xe2\x7d\xb9\xf4\x4d\xb3\xb5\x44\xc3\x28\x75\x5a\xbe\x90\xbb\x2d\x9e\xec\x38\x35\x13\x82\x46\x0b\x2e\x85\x86\x8f\x61\xe3\x05\x9e\xfd\x18\x36\x61\xb3\xa1\x08\xfb\x3b\xb0\x9d\xe4\xae\xe3\x54\x0f\x61\x93\x4c\x1d\xf8\xd1\xdf\x56\xc7\xcf\x10\x36\xc8\x4b\x8f\x3c\x07\x50\xc5\x68\x58\xcd\x7b\x8e\xea\x01\xd8\xdd\x02\x4e\xd6\xb9\xe2\xc2\x40\xe3\xcd\xe6\xa4\xe9\x5a\xfb\xbe\x08\x93\xf1\xb0\x14\xef\x39\x90\x6f\x3d\x08\x25\x07\xc0\x6b\xb8\x76\xb1\x0d\x3e\xb6\xa1\x8c\x6d\x92\x5a\xe1\x5f\x05\x6a\x83\x29\xa5\x9e\xa3\x9b\x24\x49\xda\x21\xb1\x29\xbc\x3d\x2d\x70\x01\x15\xd3\x27\x44\x23\x59\x34\xa6\x10\xea\xf6\x59\x1b\xda\x37\x61\xf3\x7b\x48\xdd\x90\xe1\x4c\xd4\x58\x7b\xc4\x0d\x74\x3a\xb0\xd7\x8a\xab\x43\xb2\x1a\xb8\x60\xf9\xea\xb1\x15\x1f\x8b\xb8\xf2\x6b\x54\xf9\xb8\x28\xdc\x51\x68\xc6\x77\x26\x1a\xaf\x38\x8d\xee\x7b\xba\x43\xa3\x7c\x48\x0b\xa4\x04\x7d\x49\x14\xc0\x36\x70\x9c\x8f\x1c\x28\xe5\x9f\x97\x0b\xc2\x12\x37\xa4\x86\x69\x94\x95\x2a\x95\x82\x7e\x9e\x49\xda\xe7\x6b\x81\x67\x8f\x2e\x67\x9d\xf3\xed\x14\x06\x61\x63\xfd\x02\xf9\x37\x51\x18\x6c\xfd\xb0\x63\x4a\xeb\x8b\x70\xb5\x17\xb3\xcf\xda\x90\x24\x2e\x04\x37\xe4\xac\x5e\x9d\x17\x29\x1f\xff\xc7\xf6\x7d\x09\xcf\x5d\xdb\x97\x39\x70\x31\xb3\x35\xd8\xf0\x25\x42\xca\x67\x33\x54\x28\x12\x7f\x34\xa6\x0f\x19\xd3\x06\xb0\xbc\x8d\x00\xc3\xb3\x0c\x84\xbc\x03\xae\x41\x2f\x59\x96\xd1\xf1\x7f\xce\x84\xa5\x25\xf3\x47\x3a\xc7\x84\x86\x04\x48\x31\x63\x2b\x6b\x35\xb8\x43\xc0\x7b\x6e\x00\x99\xca\x56\x96\x2d\x9d\x66\x52\x29\x4e\x0c\xb0\x3c\xcf\x2c\x19\x9c\xe6\xd2\xb8\xab\xa9\x6c\x05\x45\x9e\x32\x83\x69\x13\x6c\x1a\x44\xce\xe3\xb6\x9b\xda\xd9\xbb\x05\x63\xea\x2f\x5c\x93\x0d\x75\x4e\xad\xc5\x62\x78\x7b\x68\x90\x4e\x2f\xef\x1a\xcb\x93\xcc\xa3\x2d\x67\x77\x13\x70\x8b\xc2\x80\x91\x12\x96\x4c\xac\xb6\x3b\xe7\x2c\xcf\x51\x90\x5f\xe5\x56\x2d\x0d\xcc\x80\x92\xc5\xcd\xdc\x0b\xab\x29\x11\xc9\x6a\xad\xdd\x2a\x57\x0d\xb3\x47\x2e\x87\xea\x22\xa9\x71\x7a\x4a\xce\xab\x6c\x7c\xc5\xf7\x2f\x8c\xe0\x6b\x68\x5c\x77\x07\xfd\xe1\x15\x7c\x01\x8f\x69\xd9\xfd\x1e\x6d\xb3\xb9\x9d\xee\x8f\x09\x11\xf8\x72\x50\xc1\x4c\xac\xaf\x26\xda\xa1\x74\x1e\xcd\x5e\xc7\x40\x9a\xdb\xf2\x5e\xa7\x76\x75\xa7\xb5\xaf\x72\xc5\x93\x22\x6b\x42\x4e\x3e\xc6\x6f\x0b\xe1\x46\x43\x21\xef\x4a\x7a\xfa\x81\xaf\xbf\xd0\x56\x88\x2a\x75\x4f\x4f\x4b\x12\x88\xea\x18\x35\x9b\x10\x65\x06\x1a\x47\xf4\xde\xaf\xa9\xbe\xaa\x5b\x66\x4d\xa8\x6e\xe4\x7c\x60\x57\x33\x8e\x2d\x3a\x5c\x1f\x87\xf5\xbf\x67\x2e\xf4\x0e\x73\xea\x8e\xd9\xe8\xf2\x5a\x45\x29\xfc\xd4\xa8\x13\xbe\x05\xf1\x3d\xb7\xf5\x46\xc8\xbb\x56\x58\x0a\x49\x79\xf5\xcd\xb6\xc4\xcf\xb8\x9d\xfd\x5e\x43\xd7\xe6\xd5\x52\x6a\x03\x0a\x13\x0a\xf5\x83\x4c\xb2\xa3\xfb\x19\xf8\x53\xe5\x4e\xba\x40\xca\x70\x29\x85\x15\xb8\xac\x30\xf4\xd9\x65\x90\xa1\x8c\x59\xa0\x9b\xdd\xe6\x74\x7a\xd5\xc1\x6b\xc0\xd9\x0c\x13\xc3\x6f\xd1\xce\x71\xaf\x6c\x31\xa9\x3b\x66\x1e\x3d\x3c\x6e\xef\xc3\xa8\x4c\xed\x01\x1c\x9e\xd5\xc2\x67\xcf\x8d\x35\x17\x6c\xce\xa5\x23\x44\x01\x02\xef\x50\x39\x9b\x80\xb7\x89\xdc\x7b\x4f\xb7\xc5\xf3\xd6\x9d\xab\x7c\x9f\xa0\x1c\x74\x56\x73\x55\xb6\x27\x97\x4b\x66\x27\x9f\x83\x36\x31\x2a\x92\x04\xb5\x9e\x15\x54\xd4\xa8\xd6\x71\x4c\x89\xf7\x63\x4e\x3b\xf0\xbb\xd5\xdd\xb9\x21\x72\xac\xec\xbd\xc1\xce\x65\x81\xbd\x24\x70\x1e\xd9\x6c\xec\x47\x9a\xb2\x45\x0a\xa7\xa2\xa4\x68\x54\x2d\xa1\x6f\x6f\x10\x9a\xb5\xdf\x3e\x30\x75\xc9\xa7\x57\xb2\x10\xa6\x9e\xa0\xb6\x9f\xd0\x50\xb9\x27\x2e\x0a\x36\xcd\x6c\xfb\x2b\x6f\x27\xe0\xcb\x2f\x6b\x9a\x55\x14\x09\x19\x4d\x33\x99\x2c\x76\x69\x77\xef\x45\x76\x1e\x8f\x98\xd2\x63\x61\x0a\x2c\xcb\x7c\x60\x2a\xf4\xed\x2b\xa5\xb1\xcc\x1d\x2d\xb6\x46\x6e\x55\xb7\x52\xff\x45\x54\xba\xdb\xa4\x27\x2e\x13\x9e\x0b\xc2\xd7\xd0\xd5\x56\x94\x19\xcf\x10\x12\x26\x60\x8a\x65\x73\x03\x36\x95\xb7\x78\x66\x8f\x1b\x33\x85\x7a\x0e\x89\xff\x93\xa2\xb5\xcd\x64\xb5\x7c\x52\xde\x9a\xc4\xa0\xec\x1f\xc8\x3b\xea\xcb\x2e\x6d\x9f\x2c\x02\xee\x2a\xcb\xd6\x20\x17\x9d\x67\x90\xa1\x39\xd1\x5e\x48\x9f\xe5\x4e\xac\xf6\xee\x6e\x5b\x3d\x76\x46\x13\x90\x22\x78\x6d\xc9\x77\x8f\x82\x7c\x06\xb9\xd4\x9a\x53\x60\x50\x70\x6a\x23\x95\x07\x75\x37\x12\xc4\x1f\x5b\xc1\xc1\x41\xc1\x0d\x70\x9f\x3e\x51\x08\x3d\xea\x31\xdb\x02\xb0\x53\xaf\xff\xce\x19\xf2\xb8\xf6\x8f\x3b\x57\x25\x18\x1d\x08\x23\x79\x8b\xea\x4e\x71\x83\xe4\xc5\xb2\x5f\xd5\xf5\xe6\x6d\x9b\x0b\xfe\x13\x00\x00\xff\xff\xac\x73\x0a\x4f\x1d\x1c\x00\x00")

func scriptsExecuteCloudConfigTplShBytes() ([]byte, error) {
	return bindataRead(
		_scriptsExecuteCloudConfigTplSh,
		"scripts/execute-cloud-config.tpl.sh",
	)
}

func scriptsExecuteCloudConfigTplSh() (*asset, error) {
	bytes, err := scriptsExecuteCloudConfigTplShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/execute-cloud-config.tpl.sh", size: 7197, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	"scripts/execute-cloud-config.tpl.sh": scriptsExecuteCloudConfigTplSh,
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
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"scripts": &bintree{nil, map[string]*bintree{
		"execute-cloud-config.tpl.sh": &bintree{scriptsExecuteCloudConfigTplSh, map[string]*bintree{}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
