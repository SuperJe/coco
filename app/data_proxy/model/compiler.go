package model

const (
	LangCPP = "c_cpp"
	LangPY  = "python"
)

type RunCompilerReq struct {
	Lang  string `json:"lang"`
	Code  string `json:"code"`
	Input string `json:"input"`
}

type RunCompilerRsp struct {
	BaseRsp
	OutPut string `json:"output"`
}
