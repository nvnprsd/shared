package entities

import (
	"fmt"
	"os"
	"time"
)

func (f *File) String() string {
	return fmt.Sprintf("File Name : %s File Guid : %s", f.FileName, f.FileGuid)
}
func GetFileInfo(name string, size int64, mode os.FileMode, modTime time.Time, isDir bool, sys any) os.FileInfo {
	return &FileInfo{name: name, size: size, mode: mode, modTime: modTime, isDir: isDir, sys: sys}
}
func (f *FileInfo) Name() string {
	return f.name
}

func (f *FileInfo) Size() int64 {
	return f.size
}

func (f *FileInfo) Mode() os.FileMode {
	return f.mode
}

func (f *FileInfo) ModTime() time.Time {
	return f.modTime
}

func (f *FileInfo) IsDir() bool {
	return f.isDir
}

func (f *FileInfo) Sys() interface{} {
	return f.sys
}
