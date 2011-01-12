package filestuts

import "os"
import "syscall"
type File struct {
    fd int
    name string
}
func newFile(fd int, name string) *File{
    if fd < 0 {
        return nil
    }
    return &File{fd, name}
}
func open(name string, mode int, perm uint32)(file *File, err os.Errno) {
    r,e := syscall.Open(name, mode, perm)
    if e != 0 {
        err = os.Errno(e)
    }
    return newFile(r, name) , err
}

func (file *File) Close() os.Error{
    if file == nil {
        return os.EINVAL
    }
    e := syscall.Close(file.fd)
    file.fd = -1
    if e != 0 {
        return os.Errno(e)
    }
    return nil
}
func (file *File) Read(b []byte) (ret int, err os.Error) {
    if file == nil {
        return -1, os.EINVAL
    }
    r, e := syscall.Read(file.fd, b)
    if e != 0 {
        err = os.Errno(e)
    }
    return int(r), err
}
func (file *File) Write(b []byte) (ret int, err os.Error) {
    if file == nil {
        return -1, os.EINVAL
    }
    r, e := syscall.Write(file.fd, b)
    if e != 0 {
        err = os.Errno(e)
    }
    return int(r), err
}
func (file *File) String() string {
    return file.name
}
