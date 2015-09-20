package pdiscover

import (
	"os"
	"strconv"
	"io/ioutil"
)

func procFileName(pid int, name string) string {
	return "/proc/" + strconv.Itoa(pid) + "/" + name
}

func GetProcInfo(pid int) map[string]string {
	linkFields := []string{"exe","cwd","root"}
	valFields := []string{"cmdline","environ"}

	fields := map[string]string{}

	for _,name := range linkFields {
		val, err := os.Readlink(procFileName(pid, name))

		if err != nil {
			return nil
		}

		fields[name] = val
	}

	for _,name := range valFields {
		val, err := ioutil.ReadFile(procFileName(pid, name))

		if err != nil {
			return nil
		}

		fields[name] = string(val)
	}

	return fields
}




