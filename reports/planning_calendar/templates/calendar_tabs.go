package templates

const CalendarTabsTemplate = `\documentclass{minimal}

\usepackage[
    letterpaper,
    landscape,
    top=1in,
    bottom=0.5in,
    right=0.5in,
    left=0.5in
]{geometry}
\usepackage{graphicx}

\setlength{\parindent}{0pt}
\newlength{\hintertab}\setlength{\hintertab}{0.3125in}
\newlength{\vintertab}\setlength{\vintertab}{1in}
\newlength{\vintratab}\setlength{\vintratab}{13pt}

\newcommand{\tab}[1]{%
    \begin{minipage}[c][1in][c]{1.7in}
        \centering
        \parbox[c][0.4in][c]{1.5in}{\centering#1}

        \parbox[c][0.4in][c]{1.5in}{\centering#1}
    \end{minipage}
}

\begin{document}
    +TAB_PAGES
\end{document}
`

const CalendarTabsPage = `\tab{+TAB01}\hspace{\hintertab}
    \tab{+TAB02}\hspace{\hintertab}
    \tab{+TAB03}\hspace{\hintertab}
    \tab{+TAB04}\hspace{\hintertab}
    \tab{+TAB05}

    \vspace{\vintertab}

    \tab{+TAB06}\hspace{\hintertab}
    \tab{+TAB07}\hspace{\hintertab}
    \tab{+TAB08}\hspace{\hintertab}
    \tab{+TAB09}\hspace{\hintertab}
    \tab{+TAB10}

    \vspace{\vintertab}

    \tab{+TAB11}\hspace{\hintertab}
    \tab{+TAB12}\hspace{\hintertab}
    \tab{+TAB13}\hspace{\hintertab}
    \tab{+TAB14}\hspace{\hintertab}
    \tab{+TAB15}

    \vspace{\vintertab}
    
    \tab{+TAB16}\hspace{\hintertab}
    \tab{+TAB17}\hspace{\hintertab}
    \tab{+TAB18}\hspace{\hintertab}
    \tab{+TAB19}\hspace{\hintertab}
    \tab{+TAB20}`
