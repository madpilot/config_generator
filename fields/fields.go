package fields

type BitMask struct {
	Length  uint8 `json:"length"`
	Default uint8 `json:"default"`
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
