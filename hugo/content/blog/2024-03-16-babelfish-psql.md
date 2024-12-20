---
title: "Babelfish für Postgresql - MSSQL-Kompatibilität für Postgres"
date: 2024-03-16
draft: false
---

[Babelfish](https://babelfishpg.org/) ist ein Projekt, das die Verwendung von Postgresql für Software erlaubt, die
eigentlich für MSSQL (Microsoft SQL Server) geschrieben wurden. Dies erlaubt zum einen das einsparen von Lizenzkosten,
die bei MSSQL, insbesondere für hochverfügbare Umgebungen, leicht sehr hoch werden können, zum anderen aber auch, das
größere und modernere Postgres-Ökosystem verwenden zu können.

## Was ist Babelfish?
Ähnlich wie sein [Namensvorbild](https://de.wikipedia.org/wiki/Babelfisch), ist Babelfish kein migrationstool o.ä.,
sondern eine Kompatibilitäts/Übersetzungsschicht, vergleichbar etwa mit [Wine](https://www.winehq.org/). Im Optimalfall
kann man also eine Applikation, die eine MSSQL-Datenbank benötigt, mit einem Postgres-DB-Server betreiben, ohne die
Applikation umbauen zu müssen. Das kann einerseits für eine vereinfachte Migration sein, andererseits aber auch für die
Ablösung bestehender MSSQL-Server durch Postgres-Server.

## Was ist Babelfish nicht?
Babelfish ist aktuell NICHT 100% MSSQL-Kompatibel. Es fehlen vor allem Funktionalitäten, die mit dem DB-Server selbst zu
tun haben (bestimmte tabellen in Systemdatenbanken, manche Stored Procedures...) oder mit dem darunterliegenden OS
(CLR/Assebly Module). Dadurch ist Babelfish nicht plug&play, sollte die Applikation eher exotische Funktionen verwenden.
Ebenfalls sollte es nicht als einfaches Drop-In-Replacement gesehen werden. 

## Wie finde ich heraus, ob Babelfish für meine Anwendung funktioniert?
Es gibt mehrere Ressourcen, mit deren Hilfe es möglich ist, die Kompatibilität mit Babelfish zu testen:
  - Dokumentation: Es gibt sowohl von AWS für ihre SAAS-Version 
    [Informationen](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/babelfish-compatibility.tsql.limitations-unsupported.html), 
    als auch vom [Babelfish Projekt selbst](https://babelfishpg.org/docs/limitations/limitations-of-babelfish/). 
    Letztere tendieren leider dazu, etwas veraltet zu sein.
  - [Babelfish-Compass](https://github.com/babelfish-for-postgresql/babelfish_compass): Compass ist ein Tool, mit dessen
    Hilfe es möglich ist, die Kompatibilität sowohl von Skripten, als auch Software und bestehender Datenbanken zu
    messen. 

## Welche Alternativen gibt es?
Sollte es nicht möglich sein, eine Applikation von MSSQL auf Babelfish & Postgres zu migrieren, gibt es noch ein paar
andere Möglichkeiten, die Datenbank zu modernisieren. Seit Version 2019 kann MSSQL auch unter Linux betrieben werden,
was abgesehen von der Möglichkeit, Windows Lizenzen einzusparen, die Möglichkeit bietet, die Server mit den üblichen
Verwaltungstools zu verwalten, sowie Dateisysteme und Storage-Backends zu verwenden, die Windows nicht unterstützt.
Andererseits erlaubt es auch den Betrieb der Datenbank auf einem Kubernetes-Cluster, das dann Skalierung, Resilienz und
Verwaltung der DB-Engine übernimmt.

Sollten Sie nach dieser Einleitung weitere Fragen zu Babelfish haben, können Sie mich gerne kontaktieren.
