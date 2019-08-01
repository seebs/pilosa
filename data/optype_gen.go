package data

// GENERATED CODE, DO NOT EDIT
// Generated things from the OpTypes list (see gen/main.go)

var standardOpNames = [OpTypeMax]string{
	OpTypeView: "View",
	OpTypeUpdate: "Update",
	OpTypeViewRange: "ViewRange",
	OpTypeViewBit: "ViewBit",
	OpTypeUpdateBit: "UpdateBit",
	OpTypeViewOther: "ViewOther",
	OpTypeUpdateOther: "UpdateOther",
	OpTypeViewBits: "ViewBits",
	OpTypeUpdateBits: "UpdateBits",
	OpTypeViewOthers: "ViewOthers",
	OpTypeUpdateOthers: "UpdateOthers",
	OpTypeUpdateBytes: "UpdateBytes",
	OpTypeViewGivesBytes: "ViewGivesBytes",
}

const (
	OpTypeView = OpType(iota)
	OpTypeUpdate
	OpTypeViewRange
	OpTypeViewBit
	OpTypeUpdateBit
	OpTypeViewOther
	OpTypeUpdateOther
	OpTypeViewBits
	OpTypeUpdateBits
	OpTypeViewOthers
	OpTypeUpdateOthers
	OpTypeUpdateBytes
	OpTypeViewGivesBytes

	OpTypeMax
)
