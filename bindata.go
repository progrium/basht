package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _include_basht_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x53\x5d\x6f\xd3\x30\x14\x7d\x8e\x7f\xc5\x25\x8a\x96\x94\x29\xa8\x85\xb7\x4e\x11\x1b\xa3\x13\x93\x46\x8b\x1a\xf6\x54\x55\x91\x95\x3a\x6b\x44\x66\x57\xb6\xc3\x87\x4a\xfe\x3b\xd7\x1f\xcd\x92\x02\x7d\x88\xeb\xeb\x73\xcf\xb9\xf7\x5c\x9b\x14\x45\x29\xb8\xa6\x35\x57\xc9\x04\x8e\x24\x68\x44\x49\x1b\x60\x57\x50\x09\x09\x0c\x6a\x0e\x61\x74\xbc\x9e\xbf\xed\xc2\x2b\xd8\x09\xd8\x6c\x70\xcf\x42\xc8\x32\x5c\x67\x21\x6c\xb7\x70\x71\x01\x92\xe9\x56\x72\x98\x1a\x08\xc7\x5c\xbf\x9f\x91\x8e\xa0\x40\x45\xeb\xe6\x33\x2d\xa5\x18\x2a\x34\x35\x67\x99\xa1\xb8\x02\xb5\xaf\x2b\x4d\x82\xa2\xd0\x4c\xe9\x42\x69\xaa\x5b\x95\xcd\xfa\xc0\x33\x53\x8a\x3e\x19\xb0\xc9\x99\x43\xf4\x3a\x44\x5e\xc9\x28\x6a\x35\xbf\xe0\xab\xe5\xcf\x62\xf6\x1d\x59\x07\x62\x10\x7d\xb8\xc9\x3f\x15\xf9\xea\x71\x7d\xbb\x98\x47\x0f\xf7\xcb\xc5\x72\x15\x13\xf2\x8c\xbd\x0e\x0b\x91\x2d\x07\x93\xc4\x76\x80\xd2\x52\xe3\x57\x1c\x60\xd7\x4a\xaa\x6b\xc1\x49\xb0\x63\x65\x43\x25\x83\x94\xc2\x41\x8a\x12\x8b\x61\x3b\x12\x60\x56\x36\x25\x81\x4b\xb4\xff\xd0\xae\x0a\x37\xce\xb1\x6b\xeb\x16\x09\x02\x25\x5a\x59\x32\x8c\x98\xb3\x10\x03\x06\xa7\x0d\x28\x4a\x7a\xe6\x3b\xf8\x0d\x4f\x92\x1d\x20\xee\x43\x15\xf6\x15\x63\x98\xfe\xf8\x06\xf1\xf1\x20\x6b\xae\x21\x7a\xd7\xc5\x13\xcf\x1b\xd4\x15\xbc\x82\x97\xe1\xa1\x82\x0e\xcd\xa8\xfa\x1a\x37\xd7\x5b\x33\x33\xbd\x67\xdc\xe0\x83\x96\x2b\xa6\x61\x64\x32\x8c\x1d\xb6\x30\x56\xee\x05\x84\x19\xce\x77\xfd\xb8\x04\x24\xb5\x51\xeb\x0c\x4e\x20\x5f\xdc\xae\x96\x1f\x73\x17\x8c\xb4\x5d\xc6\x73\x8b\x8e\xa3\xfd\x3c\x8d\xde\x77\x9e\x42\x1c\xce\x19\x4e\x2e\x67\x51\x92\x98\xf3\xd4\xea\x4c\x26\xf6\xb0\xef\xe4\x32\x4b\x4c\x77\x2e\x6a\x8c\x47\x34\x2e\x97\x33\x0f\x44\x2b\xec\xb5\x1c\x09\xdb\x2b\x3a\xc5\xfb\x39\xb0\xc0\x37\x97\xa6\x29\x7c\xb9\xc9\x73\xec\x0e\x92\xe8\x78\x2a\xa2\x53\x13\x57\x15\x6b\x94\xf3\xe2\x34\x5e\xd4\x73\xff\x7a\xc9\x01\xd3\xdd\xcd\xfd\xc3\xff\x98\xce\x4b\xf3\x46\x87\x67\x65\x79\x36\xc0\xdf\x39\x72\x00\xf0\x25\xd5\xe4\x65\xb1\x5f\xf3\xe4\x88\xff\x3a\xdc\x49\xd5\x15\xed\x1e\xeb\x74\x24\xea\x04\xd7\x14\xaf\xa1\xb9\xff\x46\x52\xbd\x09\x49\x2f\xe4\xce\x8d\x49\x18\xf4\x7e\xfc\x33\x07\xbc\x88\x7f\x42\x7f\x73\x18\x7b\x6c\xf0\x67\xad\x4f\x60\x62\x0a\xef\xc8\x9f\x00\x00\x00\xff\xff\xf9\xb9\x4e\xae\x7c\x04\x00\x00")

func include_basht_bash_bytes() ([]byte, error) {
	return bindata_read(
		_include_basht_bash,
		"include/basht.bash",
	)
}

func include_basht_bash() (*asset, error) {
	bytes, err := include_basht_bash_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "include/basht.bash", size: 1148, mode: os.FileMode(420), modTime: time.Unix(1426447696, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"include/basht.bash": include_basht_bash,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"include": &_bintree_t{nil, map[string]*_bintree_t{
		"basht.bash": &_bintree_t{include_basht_bash, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

