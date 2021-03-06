// Code generated by go-bindata.
// sources:
// api/api.proto
// DO NOT EDIT!

package protobuf

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _apiProto = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x19\x4d\x73\xdb\xb8\xf5\xae\x5f\xf1\x46\x97\x3a\x9d\x44\x4c\x9c\xdd\xed\x8e\x5d\xb7\x55\x25\x6f\xa2\xd9\x58\xf2\x58\x4a\x32\x7b\xd2\x40\xe4\x13\x89\x35\x09\xa0\x00\x28\x45\xdd\xf1\x7f\xef\x3c\x00\xa4\x48\x4a\xb2\xb3\x1b\xef\xa1\x3a\x24\x16\xde\x07\xde\xf7\x07\x14\x45\x30\x92\x6a\xa7\x79\x9a\x59\x38\x7f\xfd\xe6\x47\x98\xb3\xc2\x94\x22\x85\xf9\x78\x0e\xa3\x5c\x96\x09\x4c\x99\xe5\x1b\x84\x91\x2c\x54\x69\xb9\x48\x61\x81\xac\x00\x56\xda\x4c\x6a\x33\xe8\x45\x51\x2f\x8a\xe0\x03\x8f\x51\x18\x4c\xa0\x14\x09\x6a\xb0\x19\xc2\x50\xb1\x38\xc3\x0a\xf2\x12\x3e\xa1\x36\x5c\x0a\x38\x1f\xbc\x86\x33\x42\xe8\x07\x50\xff\xc5\x25\xb1\xd8\xc9\x12\x0a\xb6\x03\x21\x2d\x94\x06\xc1\x66\xdc\xc0\x9a\xe7\x08\xf8\x25\x46\x65\x81\x0b\x88\x65\xa1\x72\xce\x44\x8c\xb0\xe5\x36\x73\xf7\x04\x2e\x24\x09\xfc\x12\x78\xc8\x95\x65\x5c\x00\x83\x58\xaa\x1d\xc8\x75\x13\x11\x98\x0d\x42\xd3\x27\xb3\x56\x5d\x44\xd1\x76\xbb\x1d\x30\x27\xf0\x40\xea\x34\xca\x3d\xaa\x89\x3e\x4c\x46\xd7\xd3\xf9\xf5\xab\xf3\xc1\xeb\x40\xf4\x51\xe4\x68\x0c\x68\xfc\x4f\xc9\x35\x26\xb0\xda\x01\x53\x2a\xe7\x31\x5b\xe5\x08\x39\xdb\x82\xd4\xc0\x52\x8d\x98\x80\x95\x24\xf4\x56\x73\xb2\xdb\x4b\x30\x72\x6d\xb7\x4c\x23\xb1\x49\xb8\xb1\x9a\xaf\x4a\xdb\xb2\x59\x25\x22\x37\x2d\x04\x29\x80\x09\xe8\x0f\xe7\x30\x99\xf7\xe1\xdf\xc3\xf9\x64\xfe\x92\x98\x7c\x9e\x2c\xde\xcf\x3e\x2e\xe0\xf3\xf0\xee\x6e\x38\x5d\x4c\xae\xe7\x30\xbb\x83\xd1\x6c\x3a\x9e\x2c\x26\xb3\xe9\x1c\x66\x3f\xc1\x70\xfa\x0b\xfc\x3c\x99\x8e\x5f\x02\x72\x9b\xa1\x06\xfc\xa2\x34\x69\x20\x35\x70\xb2\x26\x26\xce\x74\x73\xc4\x96\x08\x6b\xe9\x45\x32\x0a\x63\xbe\xe6\x31\xe4\x4c\xa4\x25\x4b\x11\x52\xb9\x41\x2d\x28\x12\x14\xea\x82\x1b\xf2\xaa\x01\x26\x12\x62\x93\xf3\x82\x5b\x66\xdd\xd1\x81\x5e\x83\x1e\xa1\xdc\xf0\x38\x63\x98\xc3\x27\x14\xf8\x5f\xce\xe0\xef\xc5\xc6\xff\xf5\xaf\xb4\x60\x3c\x1f\xc4\xb2\xf8\x47\xaf\x67\x76\xc2\xb2\x2f\x70\x05\x7d\xa5\xa5\x95\x6f\xfb\x97\xbd\x9e\x62\xf1\x3d\x49\x10\x17\x6c\x53\x90\x25\x2f\x7b\x3d\x5e\x28\xa9\x2d\xf4\x53\x29\xd3\x1c\x23\xa6\x78\xc4\x84\x90\x41\x86\x81\x23\xee\x5f\xd6\x68\xee\x7b\xfc\x2a\x45\xf1\xca\x6c\x59\x9a\xa2\x8e\xa4\x72\xa8\x47\xc9\x7a\x3d\x0f\x85\xb3\x54\xab\x78\x90\x32\x8b\x5b\xb6\xf3\xe0\x78\x99\xa2\x58\x06\x2e\x83\xc0\x65\x20\x15\x0a\xa6\xf8\xe6\xbc\x82\xbc\x80\x2b\xf8\xad\x07\xc0\xc5\x5a\x5e\xb8\xbf\x00\x2c\xb7\x39\x5e\x40\x7f\x94\x97\xc6\xa2\x86\x1b\x26\x58\x8a\x1a\x3e\xdd\x7c\x66\x1a\xe1\x3d\xe6\x0a\x35\x0c\x6f\x27\xfd\x4b\x87\xbf\xf1\xb9\x73\x01\xfd\xcd\xeb\xc1\x9b\xc1\xeb\x70\x1c\x4b\x61\x59\x6c\x2b\xae\xf4\x11\xac\x20\xc6\x1d\x1b\x07\x7c\xfa\x94\x3a\xbf\x80\x3e\x85\xbd\xb9\x88\xa2\x94\xdb\xac\x5c\x91\xc9\x23\xe3\x33\xff\x55\x2c\x62\x1b\xc5\x05\x7b\xe5\x4d\xdc\x20\x45\x72\xcf\x05\xf4\x0f\xfd\x15\x90\x1e\xe8\x3f\xf7\x0f\x7e\xb1\xa8\x05\xcb\x97\x89\x8c\x4d\x25\xdf\x1f\xbc\x3a\x41\x13\x6b\xee\xcc\x4b\x9a\x49\x8d\xc0\x56\xb2\xb4\xf0\x75\xd6\x7b\xe8\x01\x98\x38\xc3\x02\xcd\x05\xbc\x5f\x2c\x6e\xe7\x97\xdd\x13\x3a\x88\xa5\x30\xa5\x3b\xe9\x87\x7c\xa6\x0b\xa3\x5f\x8d\x14\x8e\x8d\xd2\x32\x29\xe3\x53\xf0\x87\xcb\x5e\xcf\xa0\xde\xf0\x18\x6b\xb1\xbc\xd2\x94\xa6\x3c\xcf\x89\x7e\xc3\x5d\x01\x64\x10\x7b\x0c\x07\xd7\x2a\x86\x91\x46\x66\xb1\xa2\x3b\x6b\x7d\xbd\x31\xe9\x0b\xd0\x68\x4b\x2d\x4c\x07\x74\x87\x2a\xdf\xbd\x68\x38\xbf\x8e\x55\x97\x0b\x03\xa6\xf8\x80\xac\x5d\x45\xe0\xfe\xa3\xa4\xb1\x70\x01\x7d\x97\x2e\x9b\x37\x51\x10\xa8\xdf\x42\x5a\xc9\x64\x47\x48\x7f\xdd\x1f\x3f\x04\x3f\xb7\x34\xd3\x68\x35\xc7\x8d\x2f\x1f\xc6\x32\x5b\x1a\x2a\xb9\xb5\x9a\x54\x1a\x80\x5b\x03\xf7\xe5\x0a\x63\x29\xd6\x3c\x75\xd5\x25\x96\x42\x60\x6c\xf9\x86\xdb\x5d\x6d\x8a\x77\x68\x6b\x3b\xec\xff\x6e\x1b\x61\x7f\xfe\xc7\x2d\x90\xe2\xe3\x06\x38\xaa\x69\x82\x39\x5a\x3c\xe2\xc0\xb1\x03\xd4\x82\xb7\xbe\xb6\x65\x6f\x81\xfe\xb8\xf8\x41\x92\xdf\xad\x41\xed\x2b\x06\x39\x37\x96\xfc\x14\x08\xcd\x11\x17\x7c\x20\x94\xb3\xf6\xf7\x53\xae\x20\xd8\x73\xbb\x23\x22\x19\x9f\xd6\xa8\xd4\xa2\xaa\x90\xae\xca\xea\xc2\xe5\x66\xa8\x12\x4c\x71\xa0\xd4\x6c\xb8\xeb\x1d\xda\x30\x8d\x4c\x1a\xe8\x67\xfb\xe3\x03\x25\xc3\xf9\xb3\x29\x18\xc4\x7d\x42\x37\x96\xfc\x5a\x1a\x0b\xec\xd1\xe2\x31\x74\x48\xc1\x0b\x53\x99\xa0\x81\xb3\xd6\x59\x5b\x99\x16\xe8\x1b\x2a\x48\xf9\xac\x05\x84\x5c\x58\xaa\x54\xb3\x04\x83\x0c\xc6\xd5\x08\x06\x29\xdf\xa0\x38\x50\xfa\x1d\xda\x8f\x1e\x3d\x68\xd2\x75\xe4\x49\xe8\x81\x6b\x4f\x62\x3e\x7b\x34\x07\x05\x9f\x72\xba\xb5\x58\x28\x4b\xc3\x63\x65\x91\x43\xa7\xb7\x85\x86\xb3\xf6\xf7\xb6\x8e\x6d\xd8\x73\xbb\xfc\x50\xab\xa7\x5c\xff\xd0\xeb\xa1\x28\x8b\xaa\x4f\xce\x7d\xc7\xa8\xbb\xe5\x54\x5a\x30\x68\xdd\xd7\xf9\x62\xb8\xf8\x38\x5f\x7e\x9c\xce\x6f\xaf\x47\x93\x9f\x26\xd7\x63\xb8\x82\xd7\x97\x15\xea\x22\x43\xb8\xbd\x9b\x7d\x9a\xcc\x27\xb3\xe9\x64\xfa\xce\x75\x1f\x04\x2e\x12\x6a\xcf\x68\x5c\x47\xaa\xba\x10\x37\xb0\x42\x9a\x5a\x63\xd7\x43\x93\x81\xe3\xd2\x22\xbf\x82\x37\x2d\xde\x77\x1f\xa7\x4f\xb2\xcd\x18\xf1\xa5\x10\xf5\x6c\x7d\xb7\x33\xb0\x2e\xf3\x7c\x07\xa5\xa1\xb5\xc0\x5f\x55\x71\xbb\x82\xf3\xf6\x2d\xd7\xa3\xd9\x74\x34\xf9\x70\xfc\x26\x66\xc1\xc8\x02\x61\x2b\xf5\x3d\xf1\x65\xd4\x31\x31\xdf\x05\x65\x12\x29\x90\xf6\x83\x86\x48\x2f\xc1\x94\x71\x06\xcc\x84\xf8\x21\x34\x02\x17\xcc\x09\x2c\x35\x08\x99\x60\xbd\x8d\x04\xe1\x1a\x42\x5c\xc1\xdb\x96\x80\xf3\xc5\xec\xf6\xf6\xab\xcd\xeb\x5b\x53\x12\xfc\x17\x28\xaf\xe0\xbb\x16\xcb\xeb\xbb\xbb\xd9\xdd\xa3\xfc\x68\x8d\x5b\x21\x94\xc2\x9b\xd0\x11\x7b\xaa\x2b\xf8\xbe\xc5\x6b\x7c\xfd\xee\x6e\x38\xbe\x1e\x3f\xca\x2e\xec\x6b\x86\x56\x4b\xed\x8c\x48\x46\x93\xa0\xd1\x58\x9a\x28\xc9\x5d\xb0\x2e\x85\x03\xb0\xbc\x1a\x49\x6a\xde\x57\xf0\xc3\x25\x45\x6e\x81\xc6\xd0\x0a\xd2\x9d\xd1\x1a\xf1\xcb\x0a\xac\x56\xce\xea\x76\x2b\x49\x97\xba\x8a\x07\xeb\xd0\x82\x27\x52\x37\xb2\x1f\x84\x5e\xd5\xcf\xe4\x1a\x7e\x2e\x57\xa8\x05\x92\x46\x54\x12\x29\x10\xd0\xfb\xd0\x0c\x60\x24\x85\xd5\x32\x07\x95\x33\x51\x53\x19\xa0\x39\x38\x41\x4b\xfb\x99\xf0\x4b\x2a\x89\x73\xc3\xe2\x8c\x0b\x9c\x2b\x8c\x07\x4d\x09\xee\x7f\x34\xcb\xea\xc2\x66\x74\x7e\xce\xd0\xad\x8c\x2e\x64\x6c\xd7\xdd\xef\x87\x3d\x9f\xea\x32\x87\x8c\xa7\xd9\x92\x6d\x18\xcf\xd9\x8a\x93\xf5\x0e\x82\x68\xcd\x56\x9a\xc7\xc1\x12\xa5\xe9\x98\x00\x2d\xa9\xb5\x0c\x48\xcd\x68\x09\x32\x1b\xd8\x66\x3c\xce\xdc\x0b\x80\xe6\x06\x9b\xc2\xf8\xaa\x88\xca\xe7\x9f\xdf\x02\x1a\xaa\xba\x35\x49\xcb\x7c\xe9\x6c\xb4\x74\x86\x6b\x05\xd1\x33\x5c\xe1\x9d\x52\xf3\xfe\xa1\xa1\x3a\x37\x60\x32\x59\xe6\x09\x29\xce\x60\xc3\xf2\x12\x21\xe7\xf7\x08\x5c\x5d\xb8\xa5\xd4\x25\xf9\x96\x6a\xbf\xc7\xe0\xda\x96\x2c\x87\xc9\x6d\x44\xe0\x8a\xd3\x2d\x33\x86\x5c\xc9\xe2\x7b\xb2\x62\xb5\x61\x41\x5c\x1a\x2b\x0b\xd4\x26\xd8\xd6\xbd\x43\x58\x49\x6a\x14\xa5\x70\xa9\x40\x5f\xbb\xca\x04\xcb\x33\xc5\x97\x28\x12\x25\xb9\xb0\x70\x05\x7f\xab\x05\xbf\xd5\x7c\x43\xa4\xf7\xb8\x73\xee\x22\x1e\xc6\x64\xc0\x85\x95\x50\x04\x8b\x35\x39\x29\x4f\xb0\x24\x82\x2b\xf8\xf1\x74\xb6\xb8\x0e\xd4\xd8\x8e\x4e\x07\xd9\x96\x99\x66\xd2\xf8\x30\xe6\xfe\xf1\x05\x8d\xdd\x87\x9f\xbc\x3f\x48\xa0\x04\x2d\xe3\xb9\xe9\x66\x62\x20\xa5\xbc\x57\x52\x18\x5f\x57\xaa\xde\x6f\xb1\xa8\x11\x5d\x1e\x34\x54\x68\x2d\x23\x5f\x93\xed\xb9\x94\xf7\x98\x40\xa9\x8e\xe7\xfa\x51\xd6\x1d\xd3\x4c\x3a\x25\xd6\x97\x79\xb3\x33\x16\x8b\x43\xe5\x9b\xaa\x8c\x9d\xf6\x8f\x2a\xd4\x5d\x52\x9a\x1e\x61\x96\x12\xbc\x71\xf7\x5f\x8c\x17\xdd\x4a\x5a\xc7\xad\x96\xbb\x27\xb5\x3a\xdc\x74\xf6\x37\x8c\x5c\x3e\x34\x75\x5b\x61\xc5\x38\x54\x86\x63\x7e\x9d\xd7\xcb\x25\x91\x36\xa3\x20\x08\x12\xb6\xcf\xd3\xbe\x0b\x1b\x0c\xfc\x76\x1a\xfc\x4d\x3e\x08\x44\x1f\x8e\xee\x56\x55\xf9\x38\x12\x6e\x87\x32\x37\x91\xf6\xc2\x8c\x3b\xb1\xd6\x54\x9e\x27\x2d\x19\x8e\x44\xe6\x11\x9f\xed\x8b\xfd\x30\x49\xb8\x6f\x7e\x47\x96\xa8\xf6\x6a\x7f\x82\xa5\x47\x58\x56\x1a\x74\x3b\xc0\x69\xfa\xf6\x24\x58\x3b\xf1\xbb\x63\x06\x69\x44\xf6\xff\xbf\x59\x9a\x99\xd6\x78\x1d\x71\xd5\xdb\x3d\x8e\x3c\x52\xb9\x1b\xf8\xdd\xe9\xea\x77\x5b\xfa\xfb\x96\xa5\xf7\x03\xc7\x07\xb6\xc2\x7c\x6f\x67\xe2\x2d\x82\xfd\x18\xe4\x04\x7c\x7a\x90\x71\xfd\xee\x38\x81\x87\x55\x91\x5f\x09\x1f\x1e\x9c\xbd\x9d\xfd\x12\x58\x3f\x42\x53\x83\xad\xe5\x3c\x6c\xc3\x2d\x49\x69\xd2\x73\x22\x11\x8f\xf9\xfc\x3d\xb0\x38\x46\xd3\xea\x59\x35\x4a\x57\xea\x4c\x1a\xfb\x08\x9d\x03\x77\xe7\x78\xd7\xcb\x8f\xd0\x70\x61\xdf\x9e\x7b\x68\x37\x25\x14\x33\x66\x2b\x75\xd2\x21\x1b\xf8\xb1\x81\x1b\xd7\x11\x79\xa1\x72\x2c\x50\x50\xe9\xd8\x72\x9b\xf1\xd6\xb4\xcf\x14\xaf\x38\xae\x30\x66\xa5\xf1\x3f\x8d\x50\x74\xde\x0b\xb9\x15\x4b\x27\xab\x29\x95\x13\x80\xc1\xcd\x64\x71\x03\x31\x13\x6e\x49\xb5\x0d\x19\x06\x30\xf4\x40\x6e\x2a\x86\xc6\xba\x85\x94\x7a\xf0\x2a\xc7\xc2\x49\x49\xed\x7d\xc5\x68\x20\x60\xa5\xcd\x50\xd8\xe0\xa9\x4b\x40\x5a\xd4\xb9\x8b\xb9\x1d\x24\xd2\xc9\x1e\x2e\xa9\x18\x12\xb1\x03\x93\x00\x9e\x3b\x2f\x14\x6a\x23\x85\x1b\x53\xdc\x96\xe2\xdc\x39\x80\xc5\x6c\x3c\xbb\xd8\x2b\xdf\xd0\xc6\xb4\x86\xd7\xda\x86\xdd\x2c\x70\xe1\x66\xea\x1f\x32\x5a\x33\x4b\x5d\x8b\xbb\xb1\x1e\x88\xaa\x94\xa0\xa0\x7c\x87\xb6\x39\x8e\x0f\x6f\x27\x30\xf7\xef\x49\x8d\x06\xb2\x7f\x38\xf2\xbd\x25\x8a\xc0\x37\x12\xba\xbe\xa2\xae\x3a\xd6\x21\x5d\xb7\xe9\xac\x41\x2a\xd4\x3e\x05\x68\x0a\x9a\xfd\x7c\xa2\xdf\xd7\xb9\x70\xf8\x9e\xb5\xdf\xde\x83\x41\x2c\x4b\xab\x55\x31\xe5\x34\x02\x29\x69\xb8\x95\x7a\x57\x23\x06\x93\xa6\xdc\x36\xf6\x81\x37\x97\x5d\x46\x19\x33\x59\x55\x5d\x88\x13\x8d\x9a\xdc\x1e\xe3\xe2\x21\xfb\x54\x39\x3d\xf3\x59\x8d\xe8\x54\x8d\x73\x64\x02\xb6\x19\x0a\x58\x95\x3c\x3f\xca\x96\x90\x97\x7e\xdd\xab\x53\x2a\xb0\x1e\xd3\xa1\x5c\x3b\xda\xa4\x4b\xeb\x0e\x97\x89\xa7\xfb\xae\x45\xf7\x69\xef\xe1\x54\xd6\x13\x2f\xed\x02\x3c\x6c\x9f\x4d\x19\x64\xc3\x3e\xdf\xb7\xf8\x8c\x3c\x85\xde\xef\x38\x0d\xba\xb8\x02\xd6\x2b\x42\x35\x6d\xe7\xcc\x92\xe7\x80\x5b\x6f\x04\x8f\xe8\x0b\x43\x04\xba\x14\xee\x07\x37\x29\xba\x1c\x55\x45\x58\xcf\xee\x0f\xbd\x5e\x47\xa5\x46\x50\x38\xd0\x91\x58\x09\xda\x2c\x9b\x2d\xee\xc8\x14\xf5\xd8\xab\xda\xa3\xf3\x63\xd8\x71\xd0\x2d\xa6\xb1\x14\x86\x27\xe8\xe4\x27\xfd\xc2\x0b\xd2\xd7\xcc\xc9\x8f\x3f\xd6\x35\x06\x4c\x26\xba\xe3\x65\xb8\xe5\xf4\x74\xe9\xc4\x6e\x6d\xd0\x4a\x1a\xc3\x69\x9d\xf2\x3f\x90\x0b\xb9\x6d\x17\x8f\xaa\x8d\x55\x34\x5d\x8b\x1d\xbc\xca\xfd\x49\x36\x3a\xa2\x80\x63\xb2\xc5\xe6\x13\x90\xfc\x67\xab\xf7\x36\xf7\xfd\x93\x32\x77\x77\x35\x66\xfc\x06\xc6\xc0\x94\xae\x55\xad\xcb\xfc\xf4\x3a\xd6\x60\xdb\x7d\x91\xfe\x73\x2d\xd1\xd9\xe8\xb7\x54\x59\x84\x1b\xa6\x58\x92\x1c\x9b\xa9\x1e\x59\xed\x59\x92\xd4\x7b\xfd\xf9\x57\xdc\xa0\xb1\x90\x1b\x84\xb5\x96\xc5\x53\xd7\xdc\x39\xd4\xe6\x65\x9e\xb8\xbe\xef\x6d\xb7\xca\x9f\x20\x3b\x28\xf5\xa7\xe6\x97\xc3\x19\xe6\x4d\x5d\x35\x4e\x79\xeb\x5b\x63\xe0\x7f\x01\x00\x00\xff\xff\x2e\xb0\x56\xce\x65\x22\x00\x00")

func apiProtoBytes() ([]byte, error) {
	return bindataRead(
		_apiProto,
		"api.proto",
	)
}

func apiProto() (*asset, error) {
	bytes, err := apiProtoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.proto", size: 8805, mode: os.FileMode(420), modTime: time.Unix(1539062535, 0)}
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
	"api.proto": apiProto,
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
	"api.proto": {apiProto, map[string]*bintree{}},
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

func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
