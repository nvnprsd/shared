package entities

import (
	"os"
	"time"
)

type File struct {
	FileName string
	FileGuid string
	Type     string
	FileSize *int64
}
type FileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool
	sys     interface{}
}
type RequestFiles struct {
	Files []File
}
type ExzeoIDResponse struct {
	Active     bool
	Sub        string
	Username   string
	Kind       string
	Client_id  string
	Exp        int64
	Iat        int64
	Iss        string
	Jti        string
	Scope      string
	Token_type string
}
