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

func GetOwnProcPath() string {
	return os.Readlink("/proc/self/exe")
}

func GetProcPath(pid int) (string,error) {
	procInfo, err := GetProcInfo(pid)
	if err != nil {
		return "",nil
	}

	if procInfo == nil {
		return "",ErrInvalidProcInfo
	}

	return procInfo["exe"],nil
}



