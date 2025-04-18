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
+D01

\pagebreak

\noindent
+D02

\pagebreak

\noindent
+D03

\pagebreak

\noindent
+D04

\pagebreak

\noindent
+D05

\pagebreak

\noindent
+D06


\pagebreak

\noindent
+D07

\end{document}
`

const WeekLabelTemplate = `\begin{tabularx}{\textwidth}{YY}
	\multicolumn{2}{c}{\LARGE +IW1} \\ 
		 & 		  \\
	FY+Y & CY+CY1 \\ 
 	+FT1 & +AQ1   \\
	+FQ1 & +AS1   \\
\end{tabularx}
`

const DayLabelTemplate = `\begin{tabularx}{\textwidth}{lrXrr}
    \multicolumn{5}{c}{+FD}                       \\
    FY+Y   &          +FT1/+FQ1  & & ORD: &   +YD \\
           &               +FW1  & & MJD: &  +MJD \\
    CY+CY1 &          +AQ1/+AS1  & &  SR: &   +SR \\
           &               +IW1  & &  SS: &   +SS
\end{tabularx}
`
