package fields

type BitMask struct {
	Length  int `json:"length"`
	Default int `json:"default"`
}

type Boolean struct {
	Default bool `json:"default"`
}

type UInt16 struct {
	Default uint16 `json:"default"`
}

type UInt8 struct {
	Default uint8 `json:"default"`
}

type String struct {
	Default string `json:"default"`
}
