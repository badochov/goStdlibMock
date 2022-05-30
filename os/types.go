package os

import (
	"os"
	"stdlibMock/io/fs"
)

type FileMode = fs.FileMode
type FileInfo = fs.FileInfo
type File struct{ *os.File }
type Process struct{ *os.Process }
type ProcAttr struct{ *os.ProcAttr }

func toFile(file *os.File) *File {
	return &File{file}
}

func toProcess(process *os.Process) *Process {
	return &Process{process}
}

func fromProcAttr(attr *ProcAttr) *os.ProcAttr {
	return attr.ProcAttr
}
