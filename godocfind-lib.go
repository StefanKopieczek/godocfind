package godocfind

import "os/exec"
import "strings"

func ByPackage(packageName string) (docs string, err error) {
    cmd := exec.Command("godoc", packageName)
    output, err := cmd.Output()
    docs = string(output)
    return
}

func ByPackageAndKeyword(packageName string, keyword string) (
        docs string, err error) {
    cmd := exec.Command("godoc", packageName, keyword)
    output, err := cmd.Output()
    docs = string(output)
    return
}

func ByKeyword(keyword string) (results []string, err error) {
    for _, packageName := range(STDLIB) {
        var keywordDocs string
        keywordDocs, err = ByPackageAndKeyword(packageName, keyword)
        if hasMatch(keywordDocs) {
            results = append(results, keywordDocs)
        }
    }

    return
}

func hasMatch(docs string) (bool) {
    doclines := strings.Split(docs, "\n")

    for _, line := range(doclines) {
        line = strings.TrimSpace(line)

        if line == "" {
            continue
        }

        if (strings.Contains(line, "DOCUMENTATION") ||
            strings.Contains(line, "package") ||
            strings.Contains(line, "import")) {
            continue
        }

        if (strings.Contains(line, "BUGS") ||
            strings.Contains(line, "SUBDIRECTORIES")) {
            return false
        }

        return true
    }

    return false
}
