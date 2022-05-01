package dto

type Payload struct {
	Field1 string          `json:"field1"`
	Field2 string          `json:"field2"`
	Field3 string          `json:"field4"`
	Field4 string          `json:"field3"`
	Field5 string          `json:"field5"`
	Detail []PayloadDetail `json:"detail"`
}

type PayloadDetail struct {
	Field6  float32 `json:"field6" bson:",truncate"`
	Field7  float32 `json:"field7" bson:",truncate"`
	Field8  float32 `json:"field8" bson:",truncate"`
	Field9  float32 `json:"field9" bson:",truncate"`
	Field10 float32 `json:"field10" bson:",truncate"`
}

func NewPayload() *Payload {

	detail := []PayloadDetail{
		{
			Field6:  123.345,
			Field7:  435345.345,
			Field8:  15345.323232345,
			Field9:  3232315345.323232345,
			Field10: 323235.323232345,
		},
		{
			Field6:  123.345,
			Field7:  435345.345,
			Field8:  15345.323232345,
			Field9:  3232315345.323232345,
			Field10: 323235.323232345,
		},
		{
			Field6:  123.345,
			Field7:  435345.345,
			Field8:  15345.323232345,
			Field9:  3232315345.323232345,
			Field10: 323235.323232345,
		},
		{
			Field6:  123.345,
			Field7:  435345.345,
			Field8:  15345.323232345,
			Field9:  3232315345.323232345,
			Field10: 323235.323232345,
		},
	}

	return &Payload{
		Field1: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		Field2: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		Field3: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		Field4: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		Field5: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		Detail: detail,
	}
}
