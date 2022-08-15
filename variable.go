package badger

var curOpt Options

func GetCurOpt() string {
	return curOpt.Dir
}

func SetCurOpt(opt Options) {
	curOpt = opt
}
