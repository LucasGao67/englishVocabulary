package main

import "github.com/LucasGao67/englishVocabulary/capture"

func main() {
	//word := &entity.Word{
	//	Name:        "challenge",
	//	Forms:       []string{"change", "changing"},
	//	ShowExample: true,
	//	Contents: []*entity.Content{
	//		{
	//			Type: "adjective",
	//			Explain: []*entity.Slice{
	//				{
	//					Content:  "A",
	//					Imported: false,
	//				},
	//				{
	//					Content:  "challenge",
	//					Imported: true,
	//				},
	//				{
	//					Content:  "is something new and difficult which requires great effort and determination.",
	//					Imported: false,
	//				},
	//			},
	//			ExampleSentences: []string{
	//				"The challenge is to make it taste good.",
	//				"This may feel tricky because it could also challenge your relationship with your husband.",
	//			},
	//		},
	//	},
	//	ExampleSentences: []string{
	//		"The challenge is to make it taste good.",
	//		"This may feel tricky because it could also challenge your relationship with your husband.",
	//	},
	//}
	//f, _ := os.Create(word.Name + ".txt")
	////w:= bufio.NewWriter(f)
	//book.Generate(f, word)

	capture.Convert("challenge")
	//capture.Convert("test")
}
