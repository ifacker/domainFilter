package options

import (
	"domainFilter/public"
	"fmt"

	"github.com/projectdiscovery/goflags"
)

func Options() {
	public.Options = new(public.FlagOptions)
	flagset := goflags.NewFlagSet()
	flagset.SetDescription(fmt.Sprintf(public.Logo, public.Version))

	// 区分大小写
	flagset.CaseSensitive = false

	flagset.CreateGroup("base", "必选参数",
		flagset.StringVarP(&public.Options.Inputs, "input", "i", "", "输入的 domain，可以是多个，但需要以 , 分隔，如：www.baidu.com,www.bing.com"),
		flagset.StringVarP(&public.Options.InputFile, "input-file", "if", "", "输入的文件名"),
	)
	flagset.CreateGroup("other", "可选参数",
		flagset.StringVarP(&public.Options.OutFile, "output", "o", fmt.Sprintf("%s.txt", public.FileName), "输出的文件名，可为空"),
		flagset.BoolVarP(&public.Options.OutDomain, "out-domain", "od", false, "输出txt内容改为domain，原输出为IP"),
		flagset.BoolVarP(&public.Options.JsonPrint, "json-print", "jp", false, "在终端打印执行结果的json串"),
		flagset.BoolVarP(&public.Options.NoLogo, "no-logo", "nl", false, "禁止输出logo等相关信息"),
		flagset.BoolVarP(&public.Options.NoErrDomain, "no-err", "ne", false, "禁止输出探测不到的域名信息"),
		flagset.BoolVarP(&public.Options.NoOutput, "no-txt", "nt", false, "禁止输出txt文件"),
		flagset.BoolVarP(&public.Options.NoOutputCSV, "no-csv", "nc", false, "禁止输出csv文件。（默认会输出csv关系对照表）"),
		flagset.BoolVarP(&public.Options.ViewPrint, "view", "v", false, "-v 打印所有的扫描信息\n🔔温馨提示：终端打印出来的 [+] IP -> Doamin 属于扫描信息实时打印，会有重复，去重后的结果在输出的文件中，使用去重后的数据请在输出的文件中获取"),
	)
	flagset.CreateGroup("custom", "定制参数",
		flagset.BoolVarP(&public.Options.ServerTools, "server-tools", "st", false, "开启后，将使用ServerTools模式"),
	)

	flagset.Parse()
}
