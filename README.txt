Il programma esegue il grep di un file di testo in modo distribuito sfruttando il paradigma map reduce.
Per eseguire il programma bisogna:

1 - Eseguire Worker.go tante volte quanto il numero di nodi che si desidera avere, ognuno su un terminale diverso;
2 - Ogni nodo dovrà stare su una porta diversa, da aggiungere come parametro e dovranno partire sempre dalla porta 4041 ed essere contigue
	(e.g. tre nodi :4041, :4042, :4043)
3 - Eseguire una sola volta Master.go che vorrà come parametri il nome del testo da analizzare e il numero di nodi del sistema

E' già presente nella cartella un file di prova "testo.txt"

Esempio con tre nodi:

Terminale 1:
go run Client.go :4041

Terminale 2:
go run Client.go :4042

Terminale 3:
go run Client.go :4043

Terminale 4:
go run Server.go testo.txt 3