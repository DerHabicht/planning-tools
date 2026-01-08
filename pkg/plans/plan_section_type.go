package plans

import (
	"fmt"

	"github.com/pkg/errors"
)

type PST int

const (
	DocChapter PST = iota
	DocSection
	DocSubSection
	DocSubSubSection
	DocParagraph
	DocSubParagraph
)

func ParsePST(s string) (PST, error) {
	switch s {
	case "#":
		return DocChapter, nil
	case "##":
		return DocSection, nil
	case "###":
		return DocSubSection, nil
	case "####":
		return DocSubSubSection, nil
	case "#####":
		return DocParagraph, nil
	case "######":
		return DocSubParagraph, nil
	default:
		return -1, errors.Errorf("unrecognized sectioning level: %s", s)
	}
}

func (pst PST) Markdown() string {
	switch pst {
	case DocChapter:
		return "#"
	case DocSection:
		return "##"
	case DocSubSection:
		return "###"
	case DocSubSubSection:
		return "####"
	case DocParagraph:
		return "#####"
	case DocSubParagraph:
		return "######"
	default:
		panic(errors.Errorf("invalid PST value: %d", pst))
	}
}

func (pst PST) LaTeX(title string) string {
	var latex string
	switch pst {
	case DocChapter:
		latex = `\chapter{%s}`
	case DocSection:
		latex = `\section{%s}`
	case DocSubSection:
		latex = `\subsection{%s}`
	case DocSubSubSection:
		latex = `\subsubsection{%s}`
	case DocParagraph:
		latex = `\paragraph{%s}`
	case DocSubParagraph:
		latex = `\subparagraph{%s}`
	default:
		panic(errors.Errorf("invalid PST value: %d", pst))
	}

	return fmt.Sprintf(latex, title)
}
