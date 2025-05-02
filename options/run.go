package options

import (
	"domainFilter/controllor"
	"domainFilter/public"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ifacker/myutil"
)

func Run() {
	if !public.Options.NoLogo {
		fmt.Print(fmt.Sprintf(public.Logo, public.Version))
	}
	if public.Options.Inputs == "" && public.Options.InputFile == "" {
		fmt.Println("请使用 -i 或 -if 参数输入目标。")
		return
	}
	if public.Options.Inputs != "" {
		tmps := strings.Split(public.Options.Inputs, ",")
		public.Targets = append(public.Targets, tmps...)
	}
	if public.Options.InputFile != "" {
		data, err := myutil.ReadTxt(public.Options.InputFile)
		if err != nil {
			fmt.Println(err)
		}
		tmps := strings.Split(data, "\n")
		public.Targets = append(public.Targets, tmps...)
	}
	if len(public.Targets) <= 0 {
		fmt.Println("未检测到目标，请检查输入的文件。")
		return
	}
	controllor.Run()

	if public.Options.JsonPrint {
		jsonData, err := json.Marshal(public.IpDomainMap)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(jsonData))
	}

	if !public.Options.NoOutputCSV {
		// 输出csv关系对照表
		data := [][]string{}
		for key, value := range public.IpDomainMap {
			valueStr := ""
			for _, valueLine := range value {
				valueStr += valueLine + "\n"
			}
			data = append(data, []string{key, strings.TrimSpace(valueStr)})
		}
		myutil.WriteCsv2(public.FileName+".csv", []string{"IP", "Domain"}, data)
	}

	if !public.Options.NoOutput {
		ipStr := ""
		for key, value := range public.IpDomainMap {
			if public.Options.OutDomain {
				if len(value) >= 1 {
					ipStr += value[0] + "\n"
				}
			} else {
				ipStr += key + "\n"
			}
		}
		myutil.WriteTxt(public.FileName+".txt", ipStr)
	}
}
