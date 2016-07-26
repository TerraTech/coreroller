// Code generated by go-bindata.
// sources:
// db/drop_all_tables.sql
// db/migrations/0001_initial.sql
// db/migrations/0002_use_tz_in_timestamps.sql
// db/migrations/0003_package_channels_blacklist.sql
// db/migrations/0004_coreos_action_default_values.sql
// DO NOT EDIT!

package api

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

var _dbDrop_all_tablesSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x91\xcd\x8a\x83\x30\x10\xc7\xef\x3e\x45\xde\xc3\x93\xab\x2e\x08\x0b\xbb\xac\x1e\x7a\x1b\xc6\x74\xd0\xd0\x98\x84\x64\x14\x7c\xfb\x06\x69\x2f\xa5\x74\x3c\x05\xc2\x6f\x66\xfe\x1f\xcd\xff\xef\x9f\x1a\xaa\xaf\x9f\x56\x75\xdf\xaa\xbd\x74\xfd\xd0\x2b\x26\x5c\x54\x5d\xf5\x75\xd5\xb4\x65\xf1\x16\x59\x13\xc5\x24\x30\x18\x82\x35\x1a\xd9\x78\x27\x90\x01\xf5\x0d\x27\x12\x28\xed\x23\xf9\x04\xa8\x4f\x6c\xd4\x33\x3a\x47\x56\xa0\xa6\xe8\xd7\x20\xd9\x30\x2e\x31\x3a\x2d\xa9\x7b\x62\x90\x1f\x5e\xcf\x2e\x85\xf3\x21\xbd\x1c\x80\xd9\x24\xf6\x71\x17\xa6\x68\x23\xc7\xc0\x7b\x90\xf4\x1f\xa0\xd4\x68\x8e\x7e\x33\x2c\xdd\xbc\x22\xe3\x88\x89\x60\x31\x53\x3c\xac\xa5\xcf\xc5\xc3\xa3\x2e\x18\x6d\xfe\xb0\xd9\x58\x59\xdc\x03\x00\x00\xff\xff\x1c\x61\x2a\xd3\x9a\x02\x00\x00")

func dbDrop_all_tablesSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbDrop_all_tablesSql,
		"db/drop_all_tables.sql",
	)
}

func dbDrop_all_tablesSql() (*asset, error) {
	bytes, err := dbDrop_all_tablesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/drop_all_tables.sql", size: 666, mode: os.FileMode(420), modTime: time.Unix(1461422972, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations0001_initialSql = []byte("\x1f\x8b\x08\x00\x61\xcd\x97\x57\x00\x03\xdd\x5c\x5b\x57\xdb\xba\xb6\x7e\xde\xfc\x0a\xad\x71\x1e\x02\xa3\x18\x2c\x5b\xf2\x85\x7d\xf6\x1e\x83\x86\x50\x58\xa5\x50\x6e\xed\x61\xbd\x64\xc8\x92\x9c\xb8\x38\x76\x96\xed\x40\xe1\xd7\x9f\x29\xc7\x36\x0e\x49\xc0\x09\xa4\xb0\x3a\xda\xd2\x60\xeb\x32\xbf\x6f\x5e\x34\xa7\x2c\x47\xd3\xd0\x87\x41\xd0\x4b\x58\x26\xd1\xe5\x70\x6d\x4d\xd3\x50\x3b\x4e\xe4\x59\x1c\x86\x32\x41\x29\xef\xcb\x01\x5b\x5b\x6b\x9f\x75\x76\x2f\x3a\xe8\x62\xf7\xe3\x51\x07\x65\x92\x0d\xd0\xfa\xda\xbf\x02\x81\x2e\x2f\x0f\xf7\xd0\xd7\xb3\xc3\x2f\xbb\x67\x57\xe8\x73\xe7\x0a\xed\x75\xf6\x77\x2f\x8f\x2e\xd0\x68\x14\x88\x6e\x4f\x46\x52\x0d\xdc\xbd\x21\xeb\x1b\x9b\x6b\xff\x8a\xd8\x40\xa2\x1b\x96\xf0\x3e\x4b\xd6\x0d\xba\x81\x8e\x4f\x2e\xd0\xf1\xe5\xd1\x11\x6a\x1f\x74\xda\x9f\xd1\x7a\xde\xe0\x7f\xff\x8b\x5a\xad\x0d\x74\x79\x7c\x78\x7a\xd9\x81\x5e\x3c\x91\x30\x86\xe8\x66\x29\xca\x82\x81\x4c\x33\x36\x18\xa2\xef\x87\x17\x07\x27\x97\x17\xe8\xe2\xf0\x4b\x07\xfd\x75\x72\xdc\xa9\x66\x5e\x8f\xe2\xdb\xf5\x0d\xc4\xb2\xbc\x35\xba\x8f\x23\x89\x5a\xa3\x8c\xb7\x1e\xa6\x5b\xdb\xf8\xf7\x23\x48\xa3\x54\x26\xe9\x52\x98\x54\xcf\x67\x71\x55\x8d\x1e\x63\x4b\x25\xa0\xcb\xaa\xbe\x54\x9f\xee\x5b\x34\x19\xf7\x5c\x09\x1d\x30\xa8\x52\x69\xb7\xc4\x5e\x49\x70\xd6\xd9\xef\x9c\x75\x8e\xdb\x9d\xf3\x42\xe7\x81\xd8\x98\xe6\x8e\x0d\x87\x61\xc0\x59\x16\xc4\xd1\xcb\xad\x62\x16\x03\x35\xe6\xa0\xb9\x90\x29\x4f\x82\x61\x3e\x5d\x26\x7f\x66\x6f\x4f\x09\x34\x1e\x2b\x74\x2c\xea\x26\x2a\xba\xce\xe0\x6a\xc8\xf8\x35\xeb\xc9\xa5\x78\xca\xee\x86\x12\x05\x51\x36\xc5\x4f\x7e\xe3\xbf\x48\x57\x8d\x6e\xc0\x8e\x15\x33\xe7\x9d\x2f\xdf\x3a\x67\x75\x3c\xa3\x24\xac\xd9\xa8\x35\xc3\x48\xa1\x41\xc5\xb2\x1f\x84\x72\x42\x31\x58\xd7\xe7\xb0\x9f\x06\xf7\x35\xeb\xcf\x5b\xf5\x59\xda\xaf\x2e\x59\x64\x65\x66\x5b\x33\xbd\x27\x55\x35\x61\xa2\xa0\x18\x74\x72\x0c\x33\x1d\x75\x40\x2f\xed\xdd\xf3\xf6\xee\x5e\xa7\xd2\xe1\x7a\xc1\xe0\x26\x9a\x1c\x7c\xac\x4c\x88\x8e\x17\x27\x7b\x27\x3b\x28\x91\x37\x81\xbc\x45\x7e\x20\x43\x01\x88\x40\x03\x29\x62\x91\x40\xc0\x5a\x2f\xeb\xa7\xa0\x27\x94\xf5\x03\xb8\xc3\xbc\x50\xa2\x3f\xfe\xf8\x63\xd2\x0e\x38\x84\xd8\x38\xed\x32\xbe\xb4\xd7\xc8\x1b\x19\x65\x13\xb4\x23\x20\xb9\x9f\xc4\x03\x35\x70\x69\x06\x8f\xee\xa7\x7d\x06\xaa\x7f\xa4\x99\x48\x4a\x01\xa2\x88\x01\xc8\xec\xc5\x71\x28\x59\x04\x57\x83\xb4\x2b\x64\x98\xb1\xda\x25\x11\xa4\x0a\x4d\x77\xc8\xee\xc2\x98\x89\xae\x07\xc6\x1c\xfb\x7e\xad\xc5\x40\x66\x4c\xb0\x8c\x75\xd3\xa0\x17\xb1\x6c\x94\xc8\x6e\x92\xb2\x09\xb3\x9b\x6c\x75\x3f\xc3\xc0\x98\x08\x83\x68\xea\xfa\x4a\xec\xa7\x70\xc7\x27\x6d\xa7\x72\xd9\x99\x76\x33\xed\xe2\x20\x74\x14\xc9\x70\xe5\x0b\xa4\xe2\x24\x0e\xe3\x64\x66\xfb\x7f\x84\xc3\x3d\x66\xff\x79\xd2\xcf\x3b\x95\x24\x93\x11\x77\x96\xaf\x4e\x68\xa5\x97\xc4\xa3\xe1\x72\x2b\xfc\x0b\xd7\xa7\x3a\x7d\x09\x64\x55\xf1\x28\xeb\x06\x51\x77\x98\xc4\xbd\x44\xa6\x69\xe9\x3d\x95\x0c\x3e\x0b\x53\x39\x61\xa4\x31\x40\xbb\xeb\x8e\x86\xe0\x33\x32\xed\x42\x58\x06\x27\x14\x53\xfd\xb2\x64\x34\xab\x5b\xca\x7c\xd9\x1d\xc4\x42\x36\xed\x00\x1e\x1d\x70\xd9\xed\xc7\xa3\xa4\xb9\x70\xca\x66\x72\x93\x29\x79\x22\xb9\xd3\x16\x77\x87\x32\x09\x62\x01\xa8\x33\x99\xdc\xb0\xb0\x1e\x95\xa6\xc8\x9c\xd3\xa5\xa2\xb7\xb8\x3f\x60\x3f\x2b\x46\xa0\x6d\xd1\x5e\x2d\x8f\xb2\x07\x79\xeb\x9c\x51\xe7\xf4\x2a\xd6\xce\x09\xa2\x73\x48\xa0\xab\x26\xc2\x3e\xea\xb1\xda\x54\xed\x55\x5d\xb0\x08\x56\xb3\x5c\xb0\x8a\x63\xaf\xe5\x82\x41\x04\x04\x44\xbc\x4c\x7e\xea\x1e\x55\xf7\xc5\x82\x5a\x68\x51\xf1\x18\x0c\xa1\xb3\xcc\x56\x1b\xdd\xe6\x0b\xdc\x85\xff\xb2\x51\x19\x3c\x20\x95\x0f\xc0\x20\x6b\x22\x4f\x05\x6e\x7d\x32\x10\x3f\x15\xa5\x03\x3e\xb9\x4a\x37\x11\xe7\x51\xc2\x3d\x3f\xe5\x5b\x89\xfd\x15\x64\x14\x9e\x06\x17\x42\x96\x66\x5d\xa8\x13\xf9\x75\xd7\x8f\x93\xd2\xbf\x5e\x71\xc6\x7c\x82\xc2\xc7\xa0\x50\x8d\x9e\x87\xf4\xa8\xd3\x24\x43\x2a\x17\x1e\x5f\x5f\x28\x0c\x57\xf4\x3f\xb2\xdd\x59\xbe\xf7\x60\xea\xf3\x1c\xef\x55\xbd\x38\x5f\xdc\x66\xf9\x70\xb9\xea\xcd\x75\xe1\xba\xe3\xad\xd7\x10\x2e\xe0\xce\x85\x77\x74\x21\xdb\xcd\xe2\xe4\xee\x29\x2f\x99\x32\x9d\x29\xc5\xac\xc4\x62\x7f\x1f\xcd\xcd\x4d\x3a\xf3\x6a\xa0\x9b\x17\x82\x4f\xf0\x5f\x56\x90\x13\x4b\xa4\x4a\x4b\x64\x3a\x0a\xb3\x59\x77\xea\xc9\x4c\x3d\x21\x7f\x22\x4c\x8d\x0b\x93\x27\xa4\x58\x4d\x26\xaf\x0a\xb2\x78\x34\x5d\xf9\xe0\x7c\x69\x97\x49\x02\xa1\x89\xab\x34\xe8\x51\x5d\xf1\x8e\x8d\xe3\x41\xa9\x6a\xac\xa9\xd4\xa6\x36\x5c\x5d\xfd\x6a\xb4\x19\x9b\x34\x50\x69\xde\x04\xd9\x93\xee\xb9\x12\xc5\x70\x08\xc4\xe9\x2c\xd3\x4a\x41\xe8\x44\x49\x34\xe3\xde\xfc\x25\xed\x1d\x78\xdf\x52\x69\xd3\x43\xdf\x79\x06\xd7\xdc\xce\xca\x9d\x88\xc3\x28\xc8\x94\x16\x55\x2d\x9d\x5f\xd8\x93\x3e\x53\x8e\x9c\xef\x4c\xa9\xdd\x08\xb5\xef\x88\xd6\xf3\xd2\x7e\x3b\xff\xb9\xb1\x76\x78\x7c\xde\x39\xbb\x40\x87\xc7\x17\x27\xd5\x0e\xd6\x26\x52\xf9\xcb\x06\xfa\xb6\x7b\x74\x09\xd3\xaf\xb7\x84\xe3\x9a\xc4\x10\x5c\x73\x0d\x4c\x34\x42\xb0\xd0\x18\x61\xbe\xe6\x09\xe1\x98\x36\x33\x3d\xc3\x74\x5b\x9b\xa8\x25\xc6\x13\xb6\x40\xa0\xfa\xb8\xc5\x6e\x6a\xb9\xeb\xb9\x89\xc6\x7b\x98\x0f\x1b\x63\x0f\x33\xe5\x52\xa9\xa1\x1c\xcf\xc4\x86\x6b\x08\x62\xdb\x0e\x75\x0c\xae\x4b\xea\x33\xd7\x62\x52\x52\x8a\x4d\x1f\xe7\xd3\x35\x11\xab\x20\xa7\x93\x87\xa1\x7c\x63\x66\x42\xb6\xba\xab\xa8\x9f\x9b\x68\x1c\xfd\x36\x51\x2d\xd6\x3d\x08\x68\x6e\x22\x1d\x66\x3e\x2c\x35\x92\xc8\x61\x9c\x80\x93\x00\xbd\x28\x0f\x2a\x48\x8c\x92\x20\xea\xa9\xdf\xc7\x59\x05\x4a\x33\x39\xdc\x7a\x4c\xc9\xe2\xd3\x62\x98\xf6\x32\x1f\x31\x41\x7d\x96\x22\xc8\x54\x38\x24\x2a\xf9\xd4\x62\x6c\xdf\xf0\xb9\x28\x94\x5f\x61\x3e\xa3\x06\x33\x05\x2c\x90\x69\x09\x98\x20\x8b\x11\x1f\x25\x89\x22\xb3\x34\xed\xc2\x3b\x5f\x3a\x27\x2e\x40\xee\xc5\xb7\x91\xda\x4f\x52\x2c\x86\x2a\x75\xcc\x5e\x6d\x06\x52\xa7\xb1\xda\x54\x60\x49\x12\xdc\x00\xb4\x74\xc4\x15\xa3\xfe\x28\x0c\xef\x5e\x3a\x95\xa3\xeb\xe3\xb9\x72\x0a\xc3\xb0\x1c\x7c\x0b\x15\x93\xf3\x78\x30\x0c\x65\x1e\x8f\xd4\x5a\x25\x55\x16\x8b\xbc\xbb\xca\xd7\xb7\x4a\xc3\x55\x4f\x5f\x4e\x26\x02\xd8\x84\x5c\x8f\x02\xdb\xd8\x75\x27\x84\x9a\xe5\x66\xd2\xb5\x0c\x07\x33\x4b\x13\x18\xfc\x85\x78\x42\x6a\x2e\xd3\x99\xe6\xda\x9e\x6d\x49\x6a\x09\x4e\x6d\xe5\x61\xe3\xc9\xd5\xa7\xa3\x20\x1a\xfd\x44\x90\xd0\xa3\x01\x04\x70\xe0\x4b\xad\x19\xa0\x17\x98\x69\x18\xc6\x77\x03\x00\x90\x2e\xe2\x93\x75\x0c\xa5\x22\x2a\xf1\x0c\x8f\x11\xee\x3a\x44\xa3\xd2\xf5\x34\x82\xb1\xd4\x3c\x9b\x9b\x9a\x67\x4a\x0f\x13\xdf\x66\x86\xa5\x02\x80\xe2\xd7\xb6\xac\x2d\x73\x4b\x57\x53\xf7\xb3\x6c\x98\xee\x6c\x6f\x03\xb5\x83\x38\x52\x51\x50\x65\xa0\xca\x1b\x7a\x71\xdc\x83\x6c\x7e\x18\xa4\x5b\x70\x73\x7b\xec\x98\x5a\x79\x57\x6d\xbe\x6a\x71\xba\x05\x25\xe5\x36\x1b\x08\x8b\x68\xa3\x34\xd9\x2e\x06\xde\x56\x23\x8f\x3b\x6c\xf5\xee\xe1\x97\x7c\xe1\x41\x2d\x4c\x89\x6b\xd9\x84\x3a\xea\x7e\x48\x3e\xdf\xda\xbb\xf2\xe3\x51\xf2\xed\x70\xcf\xfd\xd3\xf3\xef\xbe\xc4\x7f\xca\x8f\xf4\xee\x73\xef\x3f\xea\xbe\xa1\x63\xaa\xe9\xae\x66\xe8\x48\xd7\x77\xb0\xb1\x63\xda\x5b\x14\x88\x30\xf3\xde\x8d\x94\xf1\x1c\x65\xa6\x69\x7b\xa6\x6f\x4b\xcd\xf7\x0d\x57\x23\xb6\x74\x34\xa6\x53\x43\xf3\x75\xc7\x24\xc2\xa0\x9e\xf0\x68\x8d\x32\xb2\x2a\xca\xc8\x53\x94\x51\x1d\x3b\x2e\x36\xd4\x7d\x3f\xb9\xbe\x66\x1f\x3e\x6e\xdf\x5f\xdc\xd8\x27\x5f\xbf\xf7\x02\x71\xf5\xe1\x9a\x9c\x47\x7b\xe7\xd3\x94\x59\x3b\x98\xee\x18\xee\x16\xd6\x1d\xc3\xb2\x5e\x8d\x32\x6e\x30\xd3\x82\x55\x46\xf3\x5c\x07\x28\xd3\x4d\xa9\x31\x8f\xda\x9a\x6e\x71\x9d\x52\x66\x4b\xc6\x8d\x82\x32\x47\x77\xb6\xf4\x55\x50\x56\x0c\x3c\x97\x32\xdb\xb6\xb1\x4d\x30\x51\xf7\xbd\xbf\x4d\xff\xf4\xec\xe0\xab\xf3\xf3\xa3\x79\xb6\x7f\xf9\xa3\x2d\x76\x7d\xf3\xf6\xf4\xaa\x6d\x74\x66\x58\x99\xee\xee\xe8\xd6\x96\x63\xba\x00\xee\xd5\x28\x23\x26\x05\x89\x5d\x43\xe3\x4c\x38\x1a\xb1\x1c\x06\x56\xe6\x79\x9a\xf4\x5c\xa1\x4b\xdd\x95\x9c\x91\x92\x32\x4c\x57\x44\xd9\x78\xe0\xf9\x94\x39\x16\x88\x69\xe7\xa0\xaf\x8f\x09\x1b\xc4\x9f\xaf\xbe\xfd\x75\xf9\xc9\xf8\x1e\x9f\x8b\xd3\x03\xfc\xf5\xe0\xeb\x7d\x42\x77\x27\x29\xa3\x08\x9b\x3b\x14\xac\x4c\xdf\x52\x74\xe3\xd7\xa3\xcc\x70\xc0\xfd\x5c\xea\x69\x14\x3b\x30\x88\xb0\x6c\xcd\x75\x5d\x18\x09\xe2\x87\x03\xa4\x49\xe1\xea\x25\x65\x60\xe3\xab\xa1\x6c\x3c\xf0\x5c\xca\x1c\xcb\x20\x90\x56\xe5\x56\x66\x84\xfd\xf8\xf2\xe6\x26\x8a\xaf\x4c\xea\x0e\x03\x63\x3f\x62\xe7\xdb\x3f\xd3\x5e\x16\xd4\x1c\x13\xeb\xf0\x17\x19\xe6\x0e\xc6\x3b\x58\xdf\x72\x0c\xea\x3a\x74\x59\xca\xca\x3c\xe2\x61\x75\xd2\x2d\xdd\x22\x4c\xa8\xe5\x03\xd6\x10\x57\x27\x30\x88\x05\xa1\x4d\x10\x8b\x52\xd7\x14\xd8\xcb\x63\x48\x9a\x3f\x4d\x53\x9f\xfe\x07\x13\xb0\x40\xab\xae\x52\xec\x22\x9d\x2a\x2f\x30\xc9\x16\x2c\x17\x06\xc1\x8d\xe5\x83\x76\x8d\xc2\xe9\x73\x38\xb0\xe1\x78\x0e\x87\xfe\x54\xa7\xca\x5b\x88\xa9\x39\x12\xf4\xcf\x30\xd3\xa5\xc9\x2d\x82\x79\x6e\x67\x9e\xcc\x58\x8e\xc2\xe7\xb6\x6f\x9a\xf3\x51\x10\xd3\x24\xbf\x1c\x05\x73\x6c\xa6\x9b\x4a\x1b\x6a\x45\x26\x36\xc3\x9a\x23\x38\xd1\x4c\xaa\xdb\x1e\x93\x2e\x96\x32\xa7\x96\x85\xc3\xfe\x18\x06\xf6\x3d\xcf\x79\x42\x19\xd4\xa6\x0b\xc1\x68\xe4\x42\xcf\xc1\x30\x84\x43\x0d\xa1\xeb\x60\x45\x54\x6a\x84\x72\x0e\x39\x85\x02\x84\x0d\x1f\x26\xe3\x60\x57\x6c\x5c\xc3\xdc\xc8\x30\x87\xc1\x20\xd7\x70\xe5\xdb\xc2\x28\x2a\xcf\x0a\x85\xcb\x0c\x21\x3d\x5b\xd7\x40\xb3\x80\x42\x37\x2c\xcd\xa1\x26\xa4\x57\xbe\x10\x96\x67\x12\xdb\xf3\x72\x89\xcf\x2b\xd7\xd8\x87\x94\x0d\x6a\x05\x31\x1a\x3f\x5b\xe6\xe1\x08\x0a\x92\x44\xe5\x6a\xf9\x96\xe6\x66\xfe\xdc\xa7\xfc\x59\x5c\x6a\xed\x42\xa3\x84\x85\x01\xdb\x3e\xbf\x13\x91\xbc\x6b\xe5\xcb\x37\x82\xda\x6c\x04\x09\x79\x6b\x5c\x1d\x58\x7a\xed\xc2\x1c\x8e\x5c\xdd\x32\x16\xe1\xa8\x91\xeb\x3f\xc3\x91\xe9\x4b\xac\x43\x57\x4d\x08\xdb\x84\xb9\x88\x0b\x9a\x36\x98\x66\x38\x8c\x63\x97\x71\x9f\x0b\xae\xe6\xfa\x58\xb8\xdd\xd7\x24\x1e\xc4\x79\x0d\xa7\x2c\x18\xf2\x7a\x08\xaf\xa9\x4c\x37\xf3\x5a\x87\x65\xbc\x8f\xbc\x51\x2f\x45\xe9\x50\xf2\xc0\x0f\xb8\xba\x7c\x17\x8f\x12\x48\xe1\x23\x3f\xe8\x8d\x92\x3c\xf9\x6e\x4d\x12\xb9\x1a\x3a\x6d\xd3\x20\x0b\x99\x5c\xa3\x08\xf4\x0c\x9d\xd4\x73\xb0\x6e\x39\xba\x26\x4d\x8b\x69\xc4\xb1\x5d\x95\x30\xc1\x27\x1f\x32\x0b\xe9\xe8\x9e\xe3\xe6\x51\x6c\xb7\xf4\xff\x8b\x04\xd6\xc0\xb4\x2a\x13\x73\x87\x8a\x87\xaa\x42\x40\xb7\x71\x72\x9d\xd7\xab\x41\x5a\xf2\x2c\x90\x9f\xc8\xbf\x47\x70\x37\xbc\x7b\xa1\x51\xaa\x35\xd4\x6c\xc0\x22\x71\xf1\x42\x8b\x41\xa3\x08\xf8\x0c\x8b\x4d\xc3\xcf\x5e\x19\x7e\x2a\x16\x61\x69\x1f\xe4\xe5\x62\xa0\x1e\xf3\x16\xe5\xe0\x4d\xc0\xd0\xf9\xde\xe7\x07\xc2\x26\x78\x7b\x73\xc6\x1a\xa1\x7d\x1c\xb1\x27\xce\xc0\x54\xc4\x79\x86\x87\x2d\x69\x48\x8d\xda\x3e\x98\xb0\x6d\x53\xcd\x31\x6c\x5f\x73\x7c\xdd\xc3\x18\xd8\xf7\x44\x1e\x61\x86\x71\x9a\x05\xe3\x62\x5b\xfd\x9a\x67\x7f\x4e\xfb\xa3\x93\x7d\x91\xfa\x17\x67\xef\xee\x2f\x7a\xf6\xd7\xed\xfd\xde\xd1\xdd\x85\xb8\x3e\xf8\x71\xb2\x7d\xd5\xf3\x3f\x7f\x8b\x8c\xb3\xde\xe5\x97\xf8\x9a\xff\x67\x0e\x93\xad\x56\xed\xdf\xcc\x32\xce\x34\x80\x91\xfc\x6e\x93\xa2\xb5\x19\x66\x41\x99\xc1\x3d\x1f\xfa\x82\xeb\x69\x44\x3a\x5c\x73\x1c\x58\xbc\x2d\xe1\xbb\xd8\xe7\x90\x08\x70\x6b\x0e\xe6\xd3\xcb\x4f\xd1\xe0\x2b\xc5\xfd\xa1\x7d\x7f\xf7\xe1\xc3\x87\x98\xfa\x1f\x0f\x6f\x3b\xe1\x61\x74\xb1\x3b\x48\xed\xed\xe8\x47\x74\xfd\x73\x94\x45\xdb\xa7\x87\x0b\x63\xae\xea\x30\x6c\x8d\xb3\xbd\xa5\xf2\x8b\xd9\x90\x0d\xd7\xe5\x94\x08\xac\x41\x8a\x03\x76\x02\x7a\xd5\x1c\x26\x0c\x8d\x72\xcb\x14\x60\x58\x16\x1f\xa7\xa7\x33\x20\x9f\xb7\x6f\x1c\xf7\xd3\xd5\xfd\x4f\xfb\xf0\xe7\x87\x8b\xf0\xc7\xdf\xde\x71\x2a\xec\x38\xb2\x68\x1c\x7f\xff\xfb\xe3\x3d\x6f\x27\x47\xfb\x47\xe4\xb6\xdd\x3f\x5d\x42\xcd\xe3\x3a\xca\x05\x1f\xc8\xf7\x1c\x9a\x55\x8d\xcd\x30\xdb\xc4\x11\x3e\xf5\xb9\x86\x0d\x46\x21\x21\x80\x08\xc3\x6c\x30\x1b\x58\xea\x38\xc7\x16\xc4\x5c\xe2\xcc\xc1\xec\x1e\x5c\xa6\xe4\xb6\x1f\xdc\xfb\x77\x37\x1e\x19\xf4\xc2\x0f\xdf\xd9\x31\xfb\x6e\x7e\x3b\x3a\xbd\xba\x4d\xbf\xe3\x4f\x07\xc7\x07\x7f\x1e\x0f\xf9\x7e\x8f\x2c\x86\xb9\x56\x08\x41\x56\x6f\x90\xdc\xb7\x1a\x95\x7d\xcd\x30\xbb\x5c\x10\x9b\x70\xaa\x49\x9f\x99\x2a\xa0\xba\x2a\x09\x02\x3d\x3b\x42\x78\x10\x34\xa4\x70\xe6\xe9\x19\x9f\xbb\xf7\xa7\xed\xa3\x4f\x3f\x06\xd2\xbf\x8a\x3a\xdb\x6c\x9f\x0f\xdb\x3f\x8e\xf0\x71\xda\x1b\x1d\xf4\x4f\x3f\xfd\xd0\x2f\xdb\x03\xea\xeb\x5f\x5c\xa7\x39\xe6\x47\x95\x8c\x8b\x4d\xdb\x76\x16\xc9\xd6\xd4\x1e\xdd\x39\x53\x9b\x79\x13\x7b\x71\xf8\x15\x76\xe9\x3c\x8b\x00\xef\x3a\x04\x3c\x9f\x40\x18\xb0\x0c\x1b\x82\x89\xe9\x69\x9e\xb4\x75\x30\x14\x87\x13\x67\x9c\xec\x4d\x4d\xaf\xae\xfe\x09\xc1\x5f\x6d\x43\xd7\xe7\x86\xd4\x25\xed\xc7\xb7\x48\xfd\xe3\x71\x1c\xd6\xcf\x76\xc3\x92\xbc\xb3\x91\x67\xc0\x2f\xd8\xc2\xcb\xc1\x8d\xf7\x47\x47\x49\x08\xec\x17\xe7\x46\x37\xd1\xbc\x73\x94\xb5\x44\x03\xbb\x54\x45\x3e\x8d\xfa\x8e\xa1\x41\xfc\xa4\x9a\x2b\x20\xd5\xf5\xa5\xef\x3b\x18\x0c\x87\xbb\x3e\xc8\x47\x26\xca\xe3\x44\x26\xb9\xf8\x5b\x71\xd2\xcb\xeb\x5b\xb5\x61\xdc\xc5\x50\xee\xe6\x96\x5b\x7d\x68\xc4\xe5\x0a\x40\x61\xc8\x81\x6d\x9f\x11\x05\xca\x7c\x39\x28\xb3\x04\x65\xbe\x25\x28\x98\x93\x78\x2a\x14\x50\xdf\xb5\x5f\x0e\x8a\x94\xa0\xc8\xb2\xa0\x6a\x4f\xd9\x4a\xdf\xca\x8f\xd9\x3c\xc6\xb0\x89\x1e\x0e\x19\xd6\x1d\xcd\x97\xa6\xe1\x11\xa6\x94\xc4\xe7\xe3\x69\x7d\x61\xaa\x76\xca\x8b\x43\x5d\x6f\xb7\x75\xbd\xb1\xbc\x6a\x77\xa7\x09\x69\xaf\x8b\x8b\x7b\x50\x2d\x4a\xe6\x3c\x63\x7c\xb5\x6a\x11\x70\xb9\xee\xfe\xfe\x22\xb8\x1a\x59\xf8\x33\xd9\xb1\xc7\x19\xb3\x1c\x8f\x3f\xe3\xfa\xaa\x54\x13\xa8\xd3\x36\xd0\x28\xd5\x6e\xc1\x7c\x34\xa3\xbc\x5a\x94\xb7\xe3\x87\x12\x50\xbb\xa9\xbb\xc0\x15\x28\xec\x17\x57\xba\x4d\x69\x6b\xa4\x9b\x67\x68\xb3\x75\x9b\x40\x32\x08\xa5\xad\x0e\x3d\x89\x07\xc9\xa7\x6b\x49\x01\x99\x13\xc8\x23\x3c\x46\x75\x03\x3f\xa6\x0d\x6a\xae\x4c\xc3\xf3\x68\x53\x77\x7f\x7b\xda\xa0\x6e\xd0\x1d\x6c\xb2\xe7\xac\xed\x94\x69\x50\x8e\xe5\x9f\x76\xf3\xaa\xb5\x5e\xca\x56\x94\x55\x04\xbc\x47\xba\x1a\x45\xb6\x47\x74\xd5\x8f\x1b\x6c\xa2\x60\x58\x0b\x28\xe5\xad\xdc\x7e\x20\x5d\x52\x7f\xa6\x8a\x9a\x06\xfd\x8d\x5a\xff\xa9\x6c\xb9\x41\x7f\xb3\xd6\xdf\x5c\xa2\x3f\xa9\xf5\x9f\xca\x5c\x1b\xf4\xa7\xb5\xfe\x53\x15\x4e\x83\xfe\x56\xad\xbf\xb5\x44\x7f\xbb\xd6\x7f\x6a\x3f\xbd\x41\x7f\xa7\xd6\xdf\x59\xa2\xbf\x5b\xeb\x3f\x95\x0b\x36\xb1\x1f\xbd\x6e\x40\x53\xbb\x9e\x4d\x46\x98\x30\xc1\xb9\x36\x38\x71\x38\xf6\x21\x29\x2b\x12\xa8\xe2\xe8\x52\x9e\xa9\x54\x9f\xab\x0f\xe3\xa7\x23\xd5\xaf\x93\xc6\xdf\xd8\x01\x9b\xac\x6c\xbf\x50\xf8\xc5\xa2\xc7\x4a\x84\x37\x96\x16\xbe\x79\xd2\xfb\x0e\x99\x6f\x9e\xdc\x42\xbb\x46\x2b\xfb\x2f\x14\x9e\xbe\xbd\xf0\xcb\x9b\x8d\xf5\xf6\xc2\xe3\xba\xf0\xe6\x22\xc2\xdb\x6f\x2f\x3c\x59\x9a\x79\x67\x21\x87\x6d\x92\x96\xbd\xcc\xe6\xed\x45\x84\x77\xdf\x5e\xf8\x09\x9b\x37\x16\x5a\xa4\x16\xaa\x4d\x57\x23\xfd\xf2\x46\x8f\x17\x5b\x63\x97\x90\xbe\x3a\x00\x5d\x4a\x5c\x09\x31\xef\x18\x33\xd2\x50\xf5\xfa\x59\xcb\x44\xf9\x0b\x71\xad\xfc\x11\x0a\x59\x78\x4f\xa6\xa9\xab\x36\x2d\x78\x26\xb2\x93\xd7\x45\x6a\x55\x48\x69\xae\xc4\xdf\x17\x29\x36\x2a\xa8\xc5\x51\xd4\xdf\x17\xaa\x53\x41\x25\xbf\xb9\xfd\x1a\xa4\x82\x6a\x8c\xeb\xec\xf7\x0b\x75\xde\x3b\x54\x53\xc8\xeb\x1a\x6b\xc4\xc2\x83\xbe\xdf\x20\xa5\x7b\x0e\x0d\x5d\x12\x0d\xaa\xef\xa0\xbc\x43\x5c\xd6\xa2\xb8\xdc\xf7\xac\x25\x7b\x49\x34\x88\xbc\x73\x35\x19\x2f\x07\x86\xb0\xae\x5e\x6a\x89\x23\xf1\x4e\x31\x96\x01\xc3\x68\x8a\xd1\xb4\xde\xb3\x2d\xd2\x65\xe1\xfc\x53\x42\x46\x73\x60\xf6\x7b\xd6\x93\xbd\x2c\x9c\x7f\x4c\xd0\x78\x09\xb2\x5f\x19\x35\xe6\x1c\x53\x30\x5e\xe1\x98\x82\xed\xe8\xc2\x72\x41\x4a\x97\x11\x67\x7c\x70\xd3\x71\x84\xaf\xb9\xd4\x63\x96\xe9\x79\xd2\xe5\x9e\x12\x75\xd6\xf4\xea\xfa\x6e\x14\x67\x7d\xf5\xe5\x72\x53\xf7\x37\x91\x2f\x65\xa8\x0e\x08\x4a\x75\x74\x21\x91\x83\xf8\x46\xa2\x41\x0e\xfc\xad\x4e\x29\x48\xdf\xc3\x8e\xc5\x5d\x4d\x50\xee\x69\x44\xf8\x86\xe6\x9a\x50\x76\x82\x42\x0c\x2c\x89\xe0\x9c\x78\xad\x87\x67\xdf\x3b\xdb\xdb\x61\xcc\x59\xd8\x8f\xd3\x6c\x07\x14\x3a\x3e\x84\x2f\x00\x49\xf7\x06\x97\xa7\xfd\xab\x0f\x8d\xa8\x5c\x01\x28\x8f\x19\x0e\xf3\x61\x4a\x8f\x7a\x2e\x58\x9b\xae\xce\xb9\x58\x4c\x93\x1e\x73\x75\x8b\x4b\x5b\x30\x7f\x01\x50\xb8\x04\x85\x97\x05\xf5\xc2\x07\xdf\xcc\xe6\x0e\x87\xe9\x34\x61\xa8\x53\x46\x36\x15\xe0\x38\x16\xb8\x06\xc6\x90\x8a\x0b\x4b\xba\x6e\xbe\xaf\x56\x7b\xa0\xff\xf0\x06\x41\x53\x7b\x6e\x64\x09\xcf\x1d\xae\xc5\xcc\x34\x18\x73\x35\x93\x52\x43\x23\xae\xcf\x35\x50\x05\xd7\x2c\x4a\x4c\x8f\x73\x5f\xd7\x2d\xf7\x41\x50\x08\x28\x62\xfc\x40\xf2\xa2\x2f\xcb\x97\x23\xd3\x6c\xe4\xfb\xe8\x36\x08\x43\xe4\x81\xf7\x84\xb7\xec\x2e\x45\xe0\x4f\xca\x4b\x7e\xe5\x23\xc9\xa6\xb4\x35\xd2\x4d\x11\xb1\xaa\xaf\xa2\x54\x2f\x85\xae\xad\xed\x9d\x9d\x7c\x2d\x5e\x60\x3f\xdc\x47\x9d\xff\x3b\x3c\xbf\x28\xbe\x77\xaf\x78\x15\xfa\xdf\xb3\x9b\x8c\x5f\x40\x7e\xba\x4d\x3d\x2c\x3d\xdd\xb2\x74\xb8\xa7\x5b\x4d\x9e\x7f\x7b\xa6\x6d\x61\xed\x4f\xb7\x2a\x6c\xe7\xe9\x46\xd5\xa3\xb3\x66\xcd\xca\xef\xb2\x69\xd8\xba\x39\x49\xf3\x16\xd2\xa7\x7b\xd5\xde\xb2\x6d\xd0\xf0\x39\x8d\x96\x1b\x07\x55\xb3\xff\x07\xa7\xbd\x77\xc4\xdd\x54\x00\x00")

func dbMigrations0001_initialSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations0001_initialSql,
		"db/migrations/0001_initial.sql",
	)
}

func dbMigrations0001_initialSql() (*asset, error) {
	bytes, err := dbMigrations0001_initialSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/0001_initial.sql", size: 21220, mode: os.FileMode(420), modTime: time.Unix(1460645224, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations0002_use_tz_in_timestampsSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x95\xdf\x4e\xf3\x20\x18\x87\xcf\x7b\x15\xef\xd9\x0e\xbe\xec\x0a\xbe\x78\x50\x37\xa2\x4b\xf6\x2f\x96\xc6\xe8\x09\x21\x0c\x37\xb2\x0e\x08\xd0\x19\xef\x5e\xdc\x34\x69\x8c\x43\x4a\x15\x8f\x96\xb5\xf0\xfc\xa0\x2f\x3c\xef\x78\x0c\xff\x0e\x62\x6b\xa8\xe3\x50\xeb\xa2\x28\xe7\x18\xdd\xc1\xb4\xc4\xe5\x75\x59\x21\x60\xca\x70\xa3\x9a\x86\x1b\xa8\x10\x06\x3c\x5b\xa0\xc7\xd5\x12\xc1\x15\x8c\x6a\x3c\x19\xfd\x7f\x1f\xef\x47\xcf\x11\x38\x4e\x0f\x70\x7e\x30\x59\xcd\xeb\xc5\x12\x98\xe1\x1e\xbc\x21\xce\x02\x7e\x58\xa3\xd3\xfc\x0a\x97\x8b\x35\xdc\xcf\xf0\xed\xe9\x2f\x9c\x78\x75\x35\x5b\xde\x74\x87\x97\xb8\xf3\xf6\x8b\xac\xd6\x72\x63\x73\x85\x51\xad\x1b\xc1\xa8\x13\x4a\xe6\x8a\xd4\x94\xed\xe9\x96\xe7\x8a\x7b\xab\xb3\xb2\x84\xb2\x9c\x7b\x64\x3b\x2a\x25\x6f\x72\xc5\x6d\x8d\x6a\x75\xb6\x33\x23\xa4\x75\x54\xb2\x6c\x15\xfc\xc8\x23\x7f\x70\x58\xbf\xcf\x6e\xa8\x75\x84\xed\x38\xdb\x93\x27\x65\x48\xab\x37\x9e\x1d\xb7\x8e\x0b\x53\x7f\x68\x4d\x67\x1c\xf1\x02\x94\x3d\xbe\xcd\x85\xa9\xd1\x6b\xf2\x3f\xae\xb5\x64\x27\xac\x53\xe6\x25\x57\x99\xf8\x91\x4b\x97\xcd\x99\x5e\x25\x47\xe1\x7e\x75\x73\x45\xb7\x7b\x4d\xd5\xb3\x2c\x06\xf5\xa3\x55\xdd\xc9\x18\xd8\x6e\x02\xac\xc4\x0b\x1a\x20\x26\x34\x8b\x00\x2d\xb9\x17\x84\x98\xfd\x55\x1f\xa0\xf5\x37\x79\x00\x96\x22\xea\x08\x5c\xaa\x87\x07\xa1\xe3\x35\x3b\x3c\x26\xca\x9c\x31\x31\xe9\x32\x0c\xd0\x7b\xbb\x2e\x74\x61\x13\x54\xf6\x09\xf7\x1a\x00\x00\xff\xff\x56\x46\xd6\x50\x67\x0b\x00\x00")

func dbMigrations0002_use_tz_in_timestampsSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations0002_use_tz_in_timestampsSql,
		"db/migrations/0002_use_tz_in_timestamps.sql",
	)
}

func dbMigrations0002_use_tz_in_timestampsSql() (*asset, error) {
	bytes, err := dbMigrations0002_use_tz_in_timestampsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/0002_use_tz_in_timestamps.sql", size: 2919, mode: os.FileMode(420), modTime: time.Unix(1461059494, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations0003_package_channels_blacklistSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\xcf\xcf\x4e\x84\x30\x10\xc7\xf1\xf3\xce\x53\xfc\x8e\x10\x77\x9f\x60\x4f\xb5\x9d\x4d\x1a\x6b\xd9\xf4\x4f\xe2\x9e\x36\x15\x08\x36\x22\x12\x25\xf1\xf5\xe5\x00\xc1\x8b\x5c\xdb\xc9\x67\xe6\x7b\x3a\xe1\xe1\x23\x77\x5f\x69\x6a\x11\x47\x22\xe9\x58\x04\x46\x10\x8f\x86\x31\xa6\xfa\x3d\x75\xed\xbd\x7e\x4b\xc3\xd0\xf6\xf7\xd7\x7e\x7e\xe8\xf3\xf7\x84\x82\x0e\xeb\x67\x6e\x10\xa3\x56\xb0\x55\x80\x8d\xc6\xc0\xf1\x85\x1d\x5b\xc9\x7e\x05\x50\xe4\xa6\x44\x65\xa1\xd8\xf0\xac\x4b\xe1\xa5\x50\x7c\xa4\xc3\x2a\xef\x21\xcb\xcc\xff\xc8\xd5\xe9\x67\xe1\x6e\x78\xe2\x1b\x8a\xed\xac\x23\x36\xbd\xa4\xf2\x4c\xf4\x37\x56\x7d\xfe\x0c\x44\xca\x55\xd7\x25\x56\x5f\xc0\x2f\xda\x07\xbf\x93\xbd\xec\x3c\xd3\x6f\x00\x00\x00\xff\xff\x24\x79\x55\x5d\x37\x01\x00\x00")

func dbMigrations0003_package_channels_blacklistSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations0003_package_channels_blacklistSql,
		"db/migrations/0003_package_channels_blacklist.sql",
	)
}

func dbMigrations0003_package_channels_blacklistSql() (*asset, error) {
	bytes, err := dbMigrations0003_package_channels_blacklistSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/0003_package_channels_blacklist.sql", size: 311, mode: os.FileMode(420), modTime: time.Unix(1461422972, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations0004_coreos_action_default_valuesSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\xd2\xb1\x4e\xc3\x30\x10\xc6\xf1\x3d\x4f\x71\x5b\x07\xd4\x27\x60\x0a\x24\x4c\xa6\x41\x25\x19\x98\xac\x6b\x7c\x29\x16\x8e\x1d\xd9\x47\x11\x3c\x3d\xad\x50\x51\x02\xea\x60\xbb\xf3\x27\xfd\xa4\x3b\xfd\xd7\x6b\xb8\x19\xf5\xde\x23\x13\x74\x53\x51\x94\xa2\xad\xb7\xd0\x96\x77\xa2\x86\x66\x23\x5e\xa0\x77\x9e\x5c\x90\xd8\xb3\x76\x16\x7e\xe6\xfb\x46\x74\x8f\x1b\xa0\x03\x59\x86\xe7\xba\x85\xaa\x7e\x28\x3b\xd1\xc2\x6a\x72\x81\xb5\x0d\x8c\xc6\xac\x6e\xa3\xb0\xfe\xd5\xbb\xf1\x34\x1e\xc8\x87\xd3\xba\x70\x23\x31\x4b\xa4\x8e\x8b\x1a\xf5\x1f\x67\x88\x84\x74\x90\x8a\x0c\x63\x9e\xa2\x74\xc0\x9d\x21\x39\xe1\xa7\x71\xa8\xe4\x0e\xfb\x37\x37\x0c\x4b\x94\x23\xd1\x91\x18\x15\x32\xca\xa0\xf7\x16\xf9\xdd\x93\xf4\x01\xb3\xde\x36\x23\xbf\x28\x4b\x52\x84\xca\x68\xfb\x1f\x29\xe6\xc1\x55\xee\xc3\xa6\x24\x57\x6d\x9b\xa7\xb3\x9a\x59\x59\x3a\x35\x6f\x2c\x5d\xf9\x0d\x2c\x9d\xb8\x54\x57\xba\x78\x21\xad\xab\x80\xc7\xb0\x32\x4e\x3d\x67\xb5\x24\xbe\x03\x00\x00\xff\xff\x7e\x2d\x5c\xad\xc5\x04\x00\x00")

func dbMigrations0004_coreos_action_default_valuesSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations0004_coreos_action_default_valuesSql,
		"db/migrations/0004_coreos_action_default_values.sql",
	)
}

func dbMigrations0004_coreos_action_default_valuesSql() (*asset, error) {
	bytes, err := dbMigrations0004_coreos_action_default_valuesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/0004_coreos_action_default_values.sql", size: 1221, mode: os.FileMode(420), modTime: time.Unix(1465219150, 0)}
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
	"db/drop_all_tables.sql": dbDrop_all_tablesSql,
	"db/migrations/0001_initial.sql": dbMigrations0001_initialSql,
	"db/migrations/0002_use_tz_in_timestamps.sql": dbMigrations0002_use_tz_in_timestampsSql,
	"db/migrations/0003_package_channels_blacklist.sql": dbMigrations0003_package_channels_blacklistSql,
	"db/migrations/0004_coreos_action_default_values.sql": dbMigrations0004_coreos_action_default_valuesSql,
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
	"db": &bintree{nil, map[string]*bintree{
		"drop_all_tables.sql": &bintree{dbDrop_all_tablesSql, map[string]*bintree{}},
		"migrations": &bintree{nil, map[string]*bintree{
			"0001_initial.sql": &bintree{dbMigrations0001_initialSql, map[string]*bintree{}},
			"0002_use_tz_in_timestamps.sql": &bintree{dbMigrations0002_use_tz_in_timestampsSql, map[string]*bintree{}},
			"0003_package_channels_blacklist.sql": &bintree{dbMigrations0003_package_channels_blacklistSql, map[string]*bintree{}},
			"0004_coreos_action_default_values.sql": &bintree{dbMigrations0004_coreos_action_default_valuesSql, map[string]*bintree{}},
		}},
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

