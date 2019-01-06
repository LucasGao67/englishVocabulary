package book

import (
	"github.com/LucasGao67/englishVocabulary/book/entity"
	"github.com/sirupsen/logrus"
	"io"
)

func Generate(wr io.Writer, word *entity.Word) {
	if err := tpl.Execute(wr, word); err != nil {
		logrus.Error(err)
	}
}

