.PHONY:	build
build:
	platex    -kanji=utf8 main
	pbibtex   -kanji=utf8 main
	platex    -kanji=utf8 main
	platex    -kanji=utf8 main
	dvipdfmx  -p a4 main

.PHONY:	clean
clean:
	/bin/rm -f *~ *.log *.dvi *.blg *.aux *.out *.bbl *.lot *.toc *.lof *.fdb_latexmk *.fls
	/bin/rm -f 00-abstract/*.aux
	/bin/rm -f 01-introduction/*.aux
	/bin/rm -f 02-theory/*.aux
	/bin/rm -f 03-hardware/*.aux
	/bin/rm -f 04-software/*.aux
	/bin/rm -f 05-preliminary-experiment/*.aux
	/bin/rm -f 06-adf/*.aux
	/bin/rm -f 07-conclusion/*.aux
	/bin/rm -f 99-misc/*.aux

.PHONY:	dvi
dvi:
	dvipdfmx  -p a4 main

all:
	build
	clean