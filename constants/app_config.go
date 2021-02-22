package constants

const (
	DownloadPath = "/opt/data"
	MarkdownPath = "/opt/data"
	PicPattern   = "https?://.+\\.(jpg|gif|png)"
	PbTmpl       = `{{ green "DownLoading:" }} {{string . "fileName" | blue}} {{counters . }} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}} {{speed . | rndcolor }} {{percent .}}`
	BackUp       = "b"
	Replace      = "r"
	IgnorePic    = "i.loli.net"
)

var Model = map[string]string{
	BackUp:  "BackUp",
	Replace: "Replace",
}

var AppModel string
