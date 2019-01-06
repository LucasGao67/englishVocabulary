package book

import "text/template"


const section = `\section*{ {{- .Name -}} }
{\large \color{blue} {{range .Forms }} {{.}} {{end}} }
\subsection*{Explain}
\begin{enumerate}
{{ range .Contents -}}
\item {{.Type}} \\
{{ range .Explain -}}
{{if .Imported}}\textbf{ {{- .Content -}} } {{else}}{{.Content}} {{end}}
{{- end -}}
\textit{
	\begin{itemize}
	{{ range .ExampleSentences}}\item {{.}}
	{{end -}}
	\end{itemize}
}
{{end -}}
\end{enumerate}
{{ if .ShowExample }}
\subsection*{Example}
\begin{enumerate}
{{ range .ExampleSentences }}\item {{.}}
{{end -}}
\end{enumerate}
{{end}}
`

var tpl = &template.Template{}

func init() {
	tpl = template.Must(tpl.New("section").Parse(section))
}
