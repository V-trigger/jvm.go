package classpath
import(
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface{
	readClass(classname string) ([]byte, Entry, error)
    String() string
}

func newEntry(path string) Entry {
	//路径
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	//通配符
	if strings.HasSuffix(path, "*") {
        return newWildcardEntry(path)
	}

	//jar包
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
	   strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP"){
       return newZipEntry(path)
	}

	return newDirEntry(path)

}

