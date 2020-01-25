all:
	echo -e "\n\n\n************************** platex    -kanji=utf8 main *******************************\n\n\n"
	platex    -kanji=utf8 main
	echo -e "\n\n\n************************** pbibtex   -kanji=utf8 main *******************************\n\n\n"
	pbibtex   -kanji=utf8 main
	echo -e "\n\n\n************************** platex    -kanji=utf8 main *******************************\n\n\n"
	platex    -kanji=utf8 main
	echo -e "\n\n\n************************** platex    -kanji=utf8 main *******************************\n\n\n"
	platex    -kanji=utf8 main
	echo -e "\n\n\n************************** dvipdfmx  -p a4 main *******************************\n\n\n"
	dvipdfmx  -p a4 main
build:
	platex    -kanji=utf8 main
	pbibtex   -kanji=utf8 main
	platex    -kanji=utf8 main
	platex    -kanji=utf8 main
	dvipdfmx  -p a4 main
clean:
	/bin/rm -f *~ *.log *.dvi *.blg *.aux *.out *.bbl *.lot *.toc *.lof *.pdf
	/bin/rm -f 00-abstract/*.aux
	/bin/rm -f 01-introduction/*.aux
	/bin/rm -f 02-theory/*.aux
	/bin/rm -f 03-hardware/*.aux
	/bin/rm -f 04-software/*.aux
	/bin/rm -f 05-preliminary-experiment/*.aux
	/bin/rm -f 06-adf/*.aux
	/bin/rm -f 07-conclusion/*.aux
	/bin/rm -f 99-misc/*.aux
dvi:
	dvipdfmx  -p a4 main


