package domain

type DictRequest struct {
	// 翻译类型
	TransType string `json:"trans_type"`
	// 待翻译的单词
	Source string `json:"source"`
}
