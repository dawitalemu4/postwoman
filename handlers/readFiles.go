package handlers

import (
    "strings"
    "os"
)

func ReadFiles() map[string]string {

    fileMap := make(map[string]string)
    fileNames := []string{"script.sh"}

    for i := 0; i < len(fileNames); i++ {

        fileExt := strings.Split(fileNames[i], ".")

        fileContents, err := os.ReadFile("scripts/" +  fileNames[i])

        if err != nil {
            panic(err)
        }

        fileMap[fileExt[0] + fileExt[1]] = string(fileContents)
    }

    return fileMap
}
