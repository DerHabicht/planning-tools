package templates

const AG7IFTable = `\begin{center}
\begin{tabular}{|c|c|c|c|cr|r|r|}
\hline
    \textbf{QTR} & \textbf{XFR} & \textbf{MO} & \textbf{SP} & \multicolumn{2}{c|}{\textbf{WK}} & \textbf{START} & \textbf{END} \\
\hline
    \multirow{13}{*}{Q1} & \multirow{6}{*}{X1}  & \multirow{4}{*}{Q1-1} & \multirow{2}{*}{S01} & W01 &  A$\heartsuit$ & +WKSD01 & +WKED01 \\
                         &                      &                       &                      & W02 &  2$\heartsuit$ & +WKSD02 & +WKED02 \\
\cline{4-8}                                                                                                                              
                         &                      &                       & \multirow{2}{*}{S02} & W03 &  3$\heartsuit$ & +WKSD03 & +WKED03 \\
                         &                      &                       &                      & W04 &  4$\heartsuit$ & +WKSD04 & +WKED04 \\
\cline{3-8}                                                                                                                              
                         &                      & \multirow{4}{*}{Q1-2} & \multirow{2}{*}{S03} & W05 &  5$\heartsuit$ & +WKSD05 & +WKED05 \\
                         &                      &                       &                      & W06 &  6$\heartsuit$ & +WKSD06 & +WKED06 \\
\cline{2-2}\cline{4-8}                                                                                                                   
                         & \multirow{6}{*}{X2}  &                       & \multirow{2}{*}{S04} & W07 &  7$\heartsuit$ & +WKSD07 & +WKED07 \\
                         &                      &                       &                      & W08 &  8$\heartsuit$ & +WKSD08 & +WKED08 \\
\cline{3-8}                                                                                                                              
                         &                      & \multirow{5}{*}{Q1-3} & \multirow{2}{*}{S05} & W09 &  9$\heartsuit$ & +WKSD09 & +WKED09 \\
                         &                      &                       &                      & W10 & 10$\heartsuit$ & +WKSD10 & +WKED10 \\
\cline{4-8}                                                                                                                              
                         &                      &                       & \multirow{2}{*}{S06} & W11 &  J$\heartsuit$ & +WKSD11 & +WKED11 \\
                         &                      &                       &                      & W12 &  Q$\heartsuit$ & +WKSD12 & +WKED12 \\
\cline{2-2}\cline{4-8}                                                                                                                   
                         & SP1                  &                       & SP1                  & W13 &  K$\heartsuit$ & +WKSD13 & +WKED13 \\
\hline
\end{tabular}
\quad
\begin{tabular}{|c|c|c|c|cr|r|r|}
\hline
    \textbf{QTR} & \textbf{XFR} & \textbf{MO} & \textbf{SP} & \multicolumn{2}{c|}{\textbf{WK}} & \textbf{START} & \textbf{END} \\
\hline
    \multirow{13}{*}{Q2} & \multirow{6}{*}{X3}  & \multirow{4}{*}{Q2-1} & \multirow{2}{*}{S07} & W14 &  A$\diamondsuit$ & +WKSD14 & +WKED14 \\
                         &                      &                       &                      & W15 &  2$\diamondsuit$ & +WKSD15 & +WKED15 \\
\cline{4-8}                                                                                                                             
                         &                      &                       & \multirow{2}{*}{S08} & W16 &  3$\diamondsuit$ & +WKSD16 & +WKED16 \\
                         &                      &                       &                      & W17 &  4$\diamondsuit$ & +WKSD17 & +WKED17 \\
\cline{3-8}                                                                                                                             
                         &                      & \multirow{4}{*}{Q2-2} & \multirow{2}{*}{S09} & W18 &  5$\diamondsuit$ & +WKSD18 & +WKED18 \\
                         &                      &                       &                      & W19 &  6$\diamondsuit$ & +WKSD19 & +WKED19 \\
\cline{2-2}\cline{4-8}                                                                                                                  
                         & \multirow{6}{*}{X4}  &                       & \multirow{2}{*}{S10} & W20 &  7$\diamondsuit$ & +WKSD20 & +WKED20 \\
                         &                      &                       &                      & W21 &  8$\diamondsuit$ & +WKSD21 & +WKED21 \\
\cline{3-8}                                                                                                                             
                         &                      & \multirow{5}{*}{Q2-3} & \multirow{2}{*}{S11} & W22 &  9$\diamondsuit$ & +WKSD22 & +WKED22 \\
                         &                      &                       &                      & W23 & 10$\diamondsuit$ & +WKSD23 & +WKED23 \\
\cline{4-8}                                                                                                                             
                         &                      &                       & \multirow{2}{*}{S12} & W24 &  J$\diamondsuit$ & +WKSD24 & +WKED24 \\
                         &                      &                       &                      & W25 &  Q$\diamondsuit$ & +WKSD25 & +WKED25 \\
\cline{2-2}\cline{4-8}                                                                                                                  
                         & SP2                  &                       & SP2                  & W26 &  K$\diamondsuit$ & +WKSD26 & +WKED26 \\
\hline
\end{tabular}

\vspace{\baselineskip}

\begin{tabular}{|c|c|c|c|cr|r|r|}
\hline
    \textbf{QTR} & \textbf{XFR} & \textbf{MO} & \textbf{SP} & \multicolumn{2}{c|}{\textbf{WK}} & \textbf{START} & \textbf{END} \\
\hline
    \multirow{13}{*}{Q3} & \multirow{6}{*}{X5}  & \multirow{4}{*}{Q3-1} & \multirow{2}{*}{S13} & W27 &  A$\clubsuit$ & +WKSD27 & +WKED27 \\
                         &                      &                       &                      & W28 &  2$\clubsuit$ & +WKSD28 & +WKED28 \\
\cline{4-8}                                                                                                                                
                         &                      &                       & \multirow{2}{*}{S14} & W29 &  3$\clubsuit$ & +WKSD29 & +WKED29 \\
                         &                      &                       &                      & W30 &  4$\clubsuit$ & +WKSD30 & +WKED30 \\
\cline{3-8}                                                                                                                                
                         &                      & \multirow{4}{*}{Q3-2} & \multirow{2}{*}{S15} & W31 &  5$\clubsuit$ & +WKSD31 & +WKED31 \\
                         &                      &                       &                      & W32 &  6$\clubsuit$ & +WKSD32 & +WKED32 \\
\cline{2-2}\cline{4-8}                                                                                                                     
                         & \multirow{6}{*}{X6}  &                       & \multirow{2}{*}{S16} & W33 &  7$\clubsuit$ & +WKSD33 & +WKED33 \\
                         &                      &                       &                      & W34 &  8$\clubsuit$ & +WKSD34 & +WKED34 \\
\cline{3-8}                                                                                                                                
                         &                      & \multirow{5}{*}{Q3-3} & \multirow{2}{*}{S17} & W35 &  9$\clubsuit$ & +WKSD35 & +WKED35 \\
                         &                      &                       &                      & W36 & 10$\clubsuit$ & +WKSD36 & +WKED36 \\
\cline{4-8}                                                                                                                                
                         &                      &                       & \multirow{2}{*}{S18} & W37 &  J$\clubsuit$ & +WKSD37 & +WKED37 \\
                         &                      &                       &                      & W38 &  Q$\clubsuit$ & +WKSD38 & +WKED38 \\
\cline{2-2}\cline{4-8}                                                                                                                     
                         & SP3                  &                       & SP3                  & W39 &  K$\clubsuit$ & +WKSD39 & +WKED39 \\
\hline
\end{tabular}
\quad
\begin{tabular}{|c|c|c|c|cr|r|r|}
\hline
    \textbf{QTR} & \textbf{XFR} & \textbf{MO} & \textbf{SP} & \multicolumn{2}{c|}{\textbf{WK}} & \textbf{START} & \textbf{END} \\
\hline
    \multirow{13}{*}{Q4} & \multirow{6}{*}{X7}  & \multirow{4}{*}{Q4-1} & \multirow{2}{*}{S19} & W40 &  A$\spadesuit$ & +WKSD40 & +WKED40 \\
                         &                      &                       &                      & W41 &  2$\spadesuit$ & +WKSD41 & +WKED41 \\
\cline{4-8}
                         &                      &                       & \multirow{2}{*}{S20} & W42 &  3$\spadesuit$ & +WKSD42 & +WKED42 \\
                         &                      &                       &                      & W43 &  4$\spadesuit$ & +WKSD43 & +WKED43 \\
\cline{3-8}
                         &                      & \multirow{4}{*}{Q4-2} & \multirow{2}{*}{S21} & W44 &  5$\spadesuit$ & +WKSD44 & +WKED44 \\
                         &                      &                       &                      & W45 &  6$\spadesuit$ & +WKSD45 & +WKED45 \\
\cline{2-2}\cline{4-8}
                         & \multirow{6}{*}{X8}  &                       & \multirow{2}{*}{S22} & W46 &  7$\spadesuit$ & +WKSD46 & +WKED46 \\
                         &                      &                       &                      & W47 &  8$\spadesuit$ & +WKSD47 & +WKED47 \\
\cline{3-8}
                         &                      & \multirow{5}{*}{Q4-3} & \multirow{2}{*}{S23} & W48 &  9$\spadesuit$ & +WKSD48 & +WKED48 \\
                         &                      &                       &                      & W49 & 10$\spadesuit$ & +WKSD49 & +WKED49 \\
\cline{4-8}
                         &                      &                       & \multirow{2}{*}{S24} & W50 &  J$\spadesuit$ & +WKSD50 & +WKED50 \\
                         &                      &                       &                      & W51 &  Q$\spadesuit$ & +WKSD51 & +WKED51 \\
\cline{2-2}\cline{4-8}
                         & SP4                  &                       & SP4                  & W52 &  K$\spadesuit$ & +WKSD52 & +WKED52 \\
\hline
\end{tabular}
\end{center}
`
