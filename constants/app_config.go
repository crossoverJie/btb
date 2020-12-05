package constants

const (
	DownloadPath = "/opt/data"
	MarkdownPath = "/opt/data"
	PicPattern   = "https?://.+\\.(jpg|gif|png)"
	PbTmpl       = `{{ red "DownLoading:" }} {{string . "fileName" | blue}} {{counters . }} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}} {{speed . | rndcolor }} {{percent .}}`
)
