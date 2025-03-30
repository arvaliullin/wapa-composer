package domain

type Design struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Description  Description `json:"description"`
	JSFilename   string      `json:"js_filename"`
	WASMFilename string      `json:"wasm_filename"`
}

type Description struct {
	Name      string               `json:"name"`
	Lang      string               `json:"lang"`
	Functions []ExperimentFunction `json:"functions"`
}

type ExperimentFunction struct {
	Function string    `json:"function"`
	Args     []float64 `json:"args"`
}
