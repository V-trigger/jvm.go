package classpath
import(
    "io/ioutil"
    "path/filepath"
)

//单路径classpath
type DirEntry struct{
    //绝对路径
    absDir string
}

func newDirEntry(path string) *DirEntry {
    absDir, error := filepath.Abs(path)
    if error != nil {
        panic(error)
    }
    return &DirEntry{absDir}
}

func (self *DirEntry) readClass(classname string) ([]byte, Entry, error){
    filename := filepath.Join(self.absDir, classname)
    data, error := ioutil.ReadFile(filename)
    return data, self, error
}

func (self *DirEntry) String() string{
    return self.absDir
}