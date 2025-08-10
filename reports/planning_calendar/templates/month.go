package templates

// MonthWeekdayHeaderTemplate is the LaTeX code that generates the header row of the month. The placeholders are the
// days of the week, allowing the processor to highlight the doomsday of the year.
const MonthWeekdayHeaderTemplate = `KFM & & & +MON & +TUE & +WED & +THU & +FRI & +SAT & +SUN \\`

// MonthDayTemplate is the LaTeX code for each individual day in the month.
// The placeholders are:
//
//	+DY					the day of the month represented by this cell; the abbreviated month should be included if
//						crossing month boundaries
//	+HD					the abbreviation of any holiday that falls on this day
//	+FD					the full date---in ISO 8601 format---for the TikZ moon phase calculation
//	+YD					the ordinal day number
//	+SR					the sunrise time for the configured location
//	+MJD				the Modified Julian Day number
//	+SS					the sunset time for the configured location
const MonthDayTemplate = `+DY\\+HD\moon{+FD}\\+SUN\hspace{1em}+YD\hfill{}+SR\\+MJD\hfill{}+SS`

const MonthWeekTemplate = `
\multirow{5}{0.5in}{} & 
	AG7IF &
	+AQ &
	\multirow[t]{5}{1in}{+D1} & 
	\multirow[t]{5}{1in}{+D2} & 
	\multirow[t]{5}{1in}{+D3} & 
	\multirow[t]{5}{1in}{+D4} & 
	\multirow[t]{5}{1in}{+D5} & 
	\multirow[t]{5}{1in}{+D6} & 
	\multirow[t]{5}{1in}{+D7} 
\\
	&       
	&      
	+AS 
	&                            
	&                            
	&                            
	&                            
	&                            
	&                            
	&                            
\\
	& 
	ISO &      
	+IW &                            
	&                            
	&                            
	&                            
	&                            
	&                            
	&                            
\\
	&       
	FY &      
	+FT/+FQ &                            
	&                            
	&                            
	&                            
	&                            
	&                            
	&                            
\\
	& 
	&      
	+FW &                            
	&
	&                           
	&                            
	&                            
	&                            
	&                            
\\
`

// MonthTemplate is the LaTeX code for generating a month page.
// Template placeholders are:
//
//	+M					the full name and planning_calendar year of the current month
//	+PREV_CMD			the command which invokes the processed MinimonthTemplate for the previous month
//	+NEXT_CMD			the command which invokes the processed MinimonthTemplate for the next month
//	+WEEKDAYS			the macro for the processed MonthWeekdayHeaderTemplate
//	+FT1..6				the fiscal year trimester that this week (row) belongs to
//	+FQ1..6				the fiscal year quarter that this week (row) belongs to
//	+FW1..6				this week's (row's) fiscal year week number
//	+AQ1..6				the planning_calendar year quarter that this week (row) belongs to
//	+AS1..6				the sprint that this week (row) belongs to
//	+IW1..6				this week's (row's) playing card and ISO week number
//	+D01..42			the day data for the cells in the planning_calendar
const MonthTemplate = `\chapter*{\textbf{+M}\hfill{\footnotesize+PREV_CMD+NEXT_CMD}}
\begin{center}
{\ttfamily%
\begin{tabular}{|c|lr|c|c|c|c|c|c|c|}
\hline
    +WEEKDAYS
\hline
	+W1
\hline
	+W2
\hline
	+W3
\hline
	+W4
\hline
	+W5
\hline
	+W6
\hline
\end{tabular}
}
\end{center}
`
