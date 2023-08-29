---
title: "Umgehen von Content Filtering, Proxies und Antivirus mit HTML Dateien"
date: 2023-08-28
draft: false
---

## Einleitung
Seit einigen Monaten arbeite ich hin und wieder an diesem Verfahren, das es einem ermöglicht, auch an ansonsten 
sehr robust konfigurierten Proxies vorbeizukommen, aber auch Endpoint-Basierte Schutzmechanismen zu umgehen, deren 
primäre verteidigung die Quelle eines Dateidownloads ist. Der grundsätzlich Ablauf des Verfahrens ist relativ
einfach. Am besten lässt es sich an einem Beispiel verstehen: 

## 1. Stufe - Seltsame E-Mail
Der Benutzer hat diese Mail bekommen:
![E-Mail mit HTML anhang](/img/blog/2023-08-28-email.png)
Die meisten Benutzer werden diesen Anhang vermutlich öffnen, da sie nichts verwerfliches daran sehen.

## 2.Stufe - Download der Datei aus dem Browser
Das öffnen der Datei führt dann zu dieser Seite:
!["Secure File" Portal](/img/blog/2023-08-28-securefileportal.png)
In diesem Beispiel ist die HTML-Seite bewusst sehr einfach gehalten und nicht obfuscated oder anderweitig 
modifiziert. Auf dieser Seite könnten ggf. auch Instruktionen zum weiteren Verfahren sein. 

## 3.a Stufe - Entblocken
Wenn das Attachment auf einem Windows-Computer geöffnet wurde und ein Microsoft Office-Dokument ist, muss - je nach 
Zonen-Einstellungen - die Datei "Entblockt" werden. Dies kann allerdings dann mit einem Mark-Of-The-Web-Bypass 
umgangen werden. Hierfür können die Dateien z.B. in einem Archiv oder Filesystem-Container untergebracht werden, 
der kein NTFS beherrscht (etwa ISO 9660), oder über einen weniger offensichtlichen Loader geöffnet werden.
Unter Linux/UNIX-Systemen muss hier in der Regel das Execute-Bit auf der Datei gesetzt werden, sofern die Datei
ausführbar sein soll.

## 3.b - Dateien sind lokal auf dem Computer
Bei Dateitypen, die keiner so starken Absicherung unterliegen, oder bei Systemen, die dieses Prinzip der Absicherung
nicht nutzen, steht die Datei im vollen Umfang zur Verfügung. Wird der Benutzer jetzt dazu angeleitet, die Warnung
gegen das Ausführen von aus dem Internet heruntergeladenen Dateien zu ignorieren (oder diese wird umgangen), kann
Schadcode ausgeführt werden. Da die Datei auf dem Rechner selbst generiert wird und nicht aus dem Internet kommt,
greifen auch viele Antivirus/EDR-Tools nicht unbedingt ein.

## 4. Stufe - Exploitation
Nachdem die Malware an den Netzwerk - und manchen Endpoint - angekommen ist und optimalerweise vom Anwender
ausgeführt wurde, kann der Computer übernommen werden und sich weiter im Netzwerk verbreitet werden.

## Wie funktioniert es?
Web- und Emailfilter (Proxy, Gateway, File inspection etc.) führen in der Regel keine dynamische Analyse von Dateien
(e-Mail-Anhängen, heruntergeladene Dateien o.ä.) durch. Daher ist es möglich, Dateien in einem eigentlich harmlosen 
und auch oft verwendeten Dateiformat (z.B. HTML) an diesen Sicherheitsmaßnahmen vorbeizuschleusen.
Beispieldateien finden sich [Hier](https://cloud.jmbit.de/s/AJ6wW32dc5qfGd6), das Projekt zur Erstellung von solchen
Dateien findet sich [Hier](https://codeberg.org/jmbit/trojantool)

## Was kann man dagegen tun?
Das wichtigste Mittel gegen solche Angriffe ist, sich weniger auf statische Attribute zu verlassen, sondern mehr auf
Verhalten. Es gibt Verhaltensmuster, die im Alltag extrem selten auftreten, aber in fast jedem Angriff in
irgendeiner Form auftauchen. Das klassische Beispiel ist das Command `whoami` unter UNIX. ein alltäglicher Benutzer 
hat sehr selten Bedarf daran, anzuzeigen, welcher User er gerade ist. normalerweise steht dieser als teil des Prompt
in der Konsole, oder ist anderweitig indiziert (z.B. $/# als teil des Prompt etc). 



<!-- vim: set wrap linebreak textwidth=120 cc=120 : -->
