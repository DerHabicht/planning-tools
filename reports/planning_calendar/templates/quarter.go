package templates

// QuarterTemplate is the LaTeX code for the quarter pages of the planning_calendar.
// Template placeholders are:
//
//	+CYQ		the planning_calendar year quarter of this page
//	+FYQ		the fiscal year quarter of this page
//	+M1..3CMD	placeholders for the minimonth macros that make up the content of the quarter page
const QuarterTemplate = `\chapter*{\textbf{+CYQ/+FYQ}}

\vspace{1.5in}

\begin{center}
+M1CMD+M2CMD+M3CMD
\end{center}
`
