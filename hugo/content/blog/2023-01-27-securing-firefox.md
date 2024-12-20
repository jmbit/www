---
title: "Firefox absichern"
date: 2023-01-28
draft: false
author: Johannes Bülow
---

Webbrowser-Sicherheit an sich hat sich in den letzten Jahrzehnten seit ihrer Erfindung und Verbreitung massiv
verbessert.  Einst waren die größten Gefahren, dass eine Website über einen Exploit in der Browser-Engine direkten
Zugriff auf das lokale Dateisystem erlangt.  Mittlerweile ist diese Gefahr zwar noch vorhanden, aber bei weitem
seltener. Die häufigsten Angriffe auf Browser betreffen nicht mehr den Browser an sich, sondern die Inhalte, die darin
dargestellt werden.  Da mittlerweile ein Großteil unseres Lebens digital ist und sich dieses digitale Leben zumindest am
Desktop und Laptop großteils im Webbrowser abspielt, sind diese Inhalte für Angreifer interessanter den je.  
In diesem Blogbeitrag zeige ich ein paar Mittel, um das alltägliche Browsen im Internet sicherer zu machen.  
Mein Browser ist zwar auf Englisch eingestellt, ich werde jedoch in diesem Blogbeitrag jeweils die deutschsprachigen
Informationen verlinken.

### Warum Firefox und nicht z.B. Chrome?
Viele dieser Einstellungen und Erweiterungen sind zwar auch für Chrome verfügbar, jedoch fehlt die meiner Meinung nach
wichtigste Erweiterung, Multi Account Containers, bis jetzt in Chromium-basierten Browsern.

## Wichtige Erweiterungen
### 1. [Multi Account Containers](https://addons.mozilla.org/de-DE/firefox/addon/multi-account-containers/)
Meiner Meinung nach die wichtigste Browsererweiterung, und auch der Grund, warum Firefox der hierfür am besten geeignete
Browser ist. Hiermit kann man bestimmte Tabs quasi in ihrer eigenen Browsersitzung laufen lassen. Ich habe hier z.B. für
alle Websites, die ich regelmäßig besuche und benutze einen eigenen Container, sodass deren Credentials nicht außerhalb
des Containers verfügbar sind. Diese Erweiterung erschwert auch das Tracking etwas.
### 2. [Facebook Container](https://addons.mozilla.org/de-DE/firefox/addon/facebook-container/)
Quasi ein Addon für das Addon "Multi Account Containers", das einen vorgefertigten Container für Facebook-Dienste
bereitstellt.
### 3. [uBlock Origin](https://addons.mozilla.org/de-DE/firefox/addon/ublock-origin/)
uBlock ist vermutlich der effektivste Ad- und Trackerblocker, den es momentan gibt. Der Hauptgrund einen Adblocker zu
nutzen, ist aber weniger, um Werbung an sich zu blockieren (Die Werbung auf bekannten Websites ist in der Regel
harmlos), sondern um Tracker (Tracking-Pixel u.a.) sowie fragwürdigere Werbung zu blockieren.  
Das Problem mit aktueller Internet-Werbung ist, dass hier effektiv beliebiges Javascript von Höchstbietenden im Browser
ausgeführt wird, mit wenig Überprüfung, ob dieser Code tatsächlich nur das tut, was angegeben wird. Hier sind besonders
Cryptominer beliebt, es werde aber auch immer wieder Credential Stealer auf diesem Weg verteilt, insbesondere auf
Websites, die keine ausreichend restriktiven Cookie-Einstellungen haben. 
### 4. [Privacy Badger](https://addons.mozilla.org/de-DE/firefox/addon/privacy-badger17/)
Privacy Badger ist ein Tracking-Blocker von der EFF. uBlock Origin deckt diese Funktionalität zwar weitestgehend ab, ich
habe jedoch die Erfahrung gemacht, dass Privacy Badger mehr Tracker findet als uBlock.
### 5. Passwort-Manager Ihres Vertrauens
Dies ist in meinem Fall Bitwarden, wofür ich meinen eingenen Vault hoste, kann aber auch z.B. KeePass sein. Der große
Vorteil davon, die Browsererweiterung seines Passwort-Managers zu nutzen anstatt die Zugangsdaten von Hand in das
Formular zu kopieren ist, dass das ausfüllen auf Knopfdruck bei "typosquatted"-Domains nicht funktioniert.

## Wichtige Einstellungen
### 1. Cookies und Daten beim Schließen des Browsers löschen
[Supportseite](https://support.mozilla.org/de/kb/webseitendaten-einige-websites-mochten-dateien-auf)
Cookies und gecachte Daten beim schließen des Browsers zu löschen macht nicht nur Tracking schwieriger, sondern
ist auch ein guter Schritt, um Account Highjacking oder ähnliches zu verhindern. Wenn man nicht automatisch eingeloggt
wird, kann man seinen Account auch nicht in einem einzelnen Klick kompromittieren.  
Wenn man sich auf bestimmten Websites z.B. YouTube, Github, Discord, Facebook... ohnehin in jeder Session einloggt, kann
man auch für diese Ausnahmen erstellen, insbesondere falls man sie immer in Multi Account Containern verwendet.
(hierbei ist es Hilfreich, darauf zu achten, die richtige Domain für Authentication Cookies freizuschalten, 
z.B. für YouTube/Google accounts.google.com)

### 2. Datensammlung von Firefox selbst deaktivieren
Selbst wenn man Mozilla selbst 100% vertraut, ist es trotzdem sinnvoll, die Datensammlung von Firefox im normalen
Betrieb zu deaktivieren. Einerseits verringert es unnötigen Traffic, und andererseits können Daten, die nie gesammelt
wurden, auch nicht geleakt werden.

### 3. Firefox-Internen Autofill (Logins/Kreditkarten) deaktivieren (Falls nicht als PW-Manager verwendet)
Ein Passwort-Manager reicht, und es ist nicht sinnvoll, seine Logindaten (auch verschlüsselt) mehr Parteien
anzuvertrauen als unbedingt notwendig.
Der einzige Haken, der hier gesetzt bleiben kann, ist "Breach Alerts" (Warnungen vor Datenlecks). Diese Funktion kann
einen auf geleakte Zugangsdaten hinweisen. Alternativ kann auch die eigene Email-Addresse bei 
[Have I Been Pwned](https://haveibeenpwned.com/) hinterlegt werden.

## Optionale Erweiterungen
### 1. [EditThisCookie2](https://addons.mozilla.org/de/firefox/addon/etc2/)
Mit dieser Erweiterung kann man sich die Cookies einer Website ansehen, diese bearbeiten oder löschen.
### 2. [uMatrix](https://addons.mozilla.org/de/firefox/addon/umatrix/)
Eine Erweiterte Version von uBlock Origin, die deutlich mächtiger ist, aber auch viele Websites "kaputt" macht. Hiermit
können bestimmte Website-Features blockiert oder erlaubt werden.
### 3. [User-Agent Switcher](https://addons.mozilla.org/de/firefox/addon/user-agent-string-switcher/)
Mit dieser Erweiterung kann man seinen User-Agent auf den für einen anderen Browser und/oder für ein anderes
Betriebssystem ändern. Dies kann nützlich sein, um drive-by Downloads zu entschärfen, vor allem aber um Features, die
von der Website offiziell nur in bestimmten Browsern/Betriebssystemen aktiviert werden, auch in anderen Browsern und
Betriebssystemen zu nutzen. Wenn eine Website nur Chrome auf Windows unterstützt, kann man sie so trotzdem z.B. mit Firefox
auf Linux verwenden.



-------------------
Erklärungen:  
* "typosquatting" ist das Verfahren, eine Domain zu registrieren, die aussieht wie die Domain des Ziels, in der
  Hoffnung, dass Opfer über Vertipper ("typo") auf die Angreiferseite gelangen, oder bei einem Link nicht merken,
  dass er nicht auf die echte Website führt. Beispielsweise "goggle.com" statt "google.com"


<!-- vim: set wrap linebreak tw=120 -->
