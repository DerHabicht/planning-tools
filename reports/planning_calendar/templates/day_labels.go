package templates

const WeekLabels = `\documentclass[6pt]{scrarticle}

\usepackage[
    paperwidth=2.125in,
    paperheight=1.0in,
	top=0.2in,
	left=0.0in,
	right=0.0in,
	bottom=0.0in
]{geometry}
\usepackage{tabularx}

\newcolumntype{Y}{>{\centering\arraybackslash}X}

\begin{document}
\noindent
+WK

\pagebreak

\noindent
+D1

\pagebreak

\noindent
+D2

\pagebreak

\noindent
+D3

\pagebreak

\noindent
+D4

\pagebreak

\noindent
+D5

\pagebreak

\noindent
+D6


\pagebreak

\noindent
+D7

\end{document}
`

const WeekLabelTemplate = `\begin{tabularx}{\textwidth}{YY}
	\multicolumn{2}{c}{\LARGE +IW} \\ 
		 & 		  \\
	FY+Y & CY+CY1 \\ 
 	+FT & +AQ   \\
	+FQ & +AS   \\
\end{tabularx}
`

const DayLabelTemplate = `\begin{tabularx}{\textwidth}{lrXrr}
    \multicolumn{5}{c}{+FD}                       \\
    FY+Y   &          +FT/+FQ  & & ORD: &   +YD \\
           &               +FW  & & MJD: &  +MJD \\
    CY+CY1 &          +AQ/+AS  & &  SR: &   +SR \\
           &               +IW  & &  SS: &   +SS
\end{tabularx}
`
