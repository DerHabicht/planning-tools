package templates

const DayCardData = `\begin{tabularx}{\textwidth}{|lrXrr|}
\hline
    \multicolumn{5}{|c|}{+FD}                       \\
    FY+Y   &          +FT/+FQ  & & ORD: &   +YD \\
           &               +FW  & & MJD: &  +MJD \\
    CY+CY1 &          +AQ/+AS  & &  SR: &   +SR \\
           &               +IW  & &  SS: &   +SS \\
\hline
\end{tabularx}
`
const AG7IF5303 = `\hskip 0.63in
\begin{tabular}[H]{|r|p{1.7in}|}
\hline
Context: &
+CONTEXT
\\ \hline
Date: &
+ISODATE
\\ \hline
\end{tabular}
\vspace{0.1in}

\begin{tabular}[H]{|p{0.3in}|p{0.3in}|p{2.1in}|}
\hline
\multicolumn{2}{|c|}{Pom}& \multicolumn{1}{|c|}{Task} \\ \hline
\multicolumn{1}{|c|}{act}& \multicolumn{1}{|c|}{est}&
\multicolumn{1}{|c|}{Progress} \\
\hline
\hline
\multicolumn{2}{|c|}{$\Box$} & \\ \hline
& & \\ \hline
\hline
\multicolumn{2}{|c|}{$\Box$} & \\ \hline
& & \\ \hline
\hline
\multicolumn{2}{|c|}{$\Box$} & \\ \hline
& & \\ \hline
\hline
\multicolumn{2}{|c|}{$\Box$} & \\ \hline
& & \\ \hline
\hline
\multicolumn{2}{|c|}{$\Box$} & \\ \hline
& & \\ \hline
\hline
\multicolumn{2}{|c|}{$\Box$} & \\ \hline
& & \\ \hline
\hline
\multicolumn{2}{|c|}{$\Box$} & \\ \hline
& & \\ \hline
\hline
\multicolumn{2}{|c|}{$\Box$} & \\ \hline
& & \\ \hline
\hline
\multicolumn{2}{|c|}{$\Box$} & \\ \hline
& & \\ \hline
\hline
\multicolumn{2}{|c|}{$\Box$} & \\ \hline
& & \\ \hline
\hline
\multicolumn{2}{|c|}{$\Box$} & \\ \hline
& & \\ \hline
\hline
\multicolumn{2}{|c|}{$\Box$} & \\ \hline
& & \\
\hline
\end{tabular}

\pagebreak
\normalsize

\begin{tabular}[H]{|l|p{1.7in}|}
\hline
\multicolumn{1}{|c|}{Metric}& \multicolumn{1}{|c|}{Value} \\ \hline
\hline
Pomodoros (\%)& \\ \hline
\quad Completed& \\ \hline
\quad Meetings (min) & \\ \hline
\quad Goal& \\ \hline
\hline
Interruptions& \\ \hline
\quad Internal& \\ \hline
\quad External& \\ \hline
\quad Pom, Aborted& \\ \hline
\quad Pom, Invalidated& \\ \hline
\end{tabular}

\vspace{0.1in}
Notes:

\vfill
\noindent
+DAY_LABEL
\vspace{\baselineskip}
`

const DayCards = `\documentclass[10pt]{extarticle}

\usepackage{fancyhdr}
\usepackage[
paperwidth=4in,
paperheight=6in,
margin=0.2in,
top=0.2in,
bottom=0.4in,
left=0.2in,
right=0.2in
]{geometry}
\usepackage{lastpage}
\usepackage{tabularx}
\usepackage{wasysym}

\renewcommand{\headrulewidth}{0pt}
\setlength{\footskip}{0pt}

\pagestyle{fancy}
\fancyhf{}
\lfoot{\small AG7IF Form 5303, Page \thepage~of~2}
\rfoot{\small Rev. 2022-08-26}

\begin{document}
\setcounter{page}{1}
+DC1
\pagebreak
\setcounter{page}{1}
+DC2
\pagebreak
\setcounter{page}{1}
+DC3
\pagebreak
\setcounter{page}{1}
+DC4
\pagebreak
\setcounter{page}{1}
+DC5
\pagebreak
\setcounter{page}{1}
+DC6
\pagebreak
\setcounter{page}{1}
+DC7
\pagebreak
\end{document}
`
