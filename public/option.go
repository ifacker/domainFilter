package public

type FlagOptions struct {
	Inputs      string // 可输入多个 domain，以, 分隔
	InputFile   string // 输入 domain 文件
	OutFile     string // 输出的文件名
	OutDomain   bool   // 输出txt内容改为domain，原输出为IP
	JsonPrint   bool   // 在终端打印执行结果的json串
	NoLogo      bool   // 禁止输出logo等相关信息
	NoErrDomain bool   // 禁止输出探测不到的域名信息
	NoOutput    bool   // 不输出txt文件
	NoOutputCSV bool   // 不输出csv文件
	ViewPrint   bool   // 打印所有信息
}
