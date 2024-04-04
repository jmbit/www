---
title: "Gedanken zu XZ"
date: 2024-04-04
draft: false
---

Dieser Blogpost ist dazu da, meine Gedanken zu der XZ-Situation zusammenzufassen.

## 1. Die Entdeckung

Die [Entdeckung der Backdoor in liblzma](https://www.openwall.com/lists/oss-security/2024/03/29/4) 
war einerseits ein purer Glücksgriff, andererseits auch ein Beispiel dafür, warum
Open Source so wichtig für die Entdeckung und Analyse von Schwachstellen ist. In diesem Fall war es dem Entdecker -
trotz der Tatsache, dass er eigentlich Datenbank-Entwickler ist - in verhältnismäßig kurzer Zeit möglich, die Ursache
der "langsamen" SSH-Logins (800 statt 300 ms) zu ermitteln und zu analysieren. Im Vergleich dazu hatte ich deutlich mehr
Zeit investiert in die letztlich erfolglose Analyse eines Memory leaks unter Microsoft Windows.

## 2. Die potenziellen Auswirkungen

Wäre die Backdoor erfolgreich verbreitet worden, wäre früher oder später - 
[abhängig vom Patchzyklus](https://infosec.exchange/@SecureOwl/112181879860988729) - ein großer Teil der Server im 
Internet verwundbar gewesen:
![Shodan zeigt 19 698 358 OpenSSH Server](/img/blog/shodan-openssh-2024-04-04.png)
Der Angreifer hätte ähnlichen Zugriff auf diese Server gehabt wie damals mit Ethernal Blue auf Windows-Server... Nur
dass laut Shodan nur etwa 420 000 Windows Server mit offenem SMB-Port im Internet stehen. 

## 3. Analyse

Die Analyse der Backdoor (abgesehen von dem nicht gerade hilfreichen Takedown durch Github) hat einige interessante
Informationen zu Tage befördert. Einmal die Tatsache, dass diese Backdoor sehr gut versteckt war, enthalten in einer
Testfile, mit dem "Aktivierungscode" nicht in der Versionskontrolle enthalten[^1]. Ebenso auffällig ist, dass die
Backdoor/RCE einen bestimmtes SSH-Zertifikat benötigt, um aktiviert zu werden, und die Payload innerhalb dieses
Zertifikats (genauer, dem CA-Teil) enthalten ist, wo binäre Daten kaum auffallen [^2] . Später ist ebenfalls aufgefallen, dass
der Name "Jia Cheong Tan" zwar chinesisch klingt, aber eine Mischung aus Kantonesisch und Mandarin ist[^3]. Aus der
Analyse der Git commit logs geht ebenfalls heraus, dass die Commits zu UrzeitenSch gemacht wurden, die für China eher
unüblich sind[^4]. Die Vermutung liegt also nahe, dass die tatsächliche Person (oder Personen?) hinter der Backdoor
vermutlich eher in Osteuropa oder dem Nahen Osten (UTC+2/3) zu finden ist.

## 4. Lessons Learned

- Die XZ-Backdoor ist eine seltene Möglichkeit, in die Vorgehensweise und den Code von möglichen staatlichen Akteuren oder
  anderen APTs[^5] hineinsehen zu können und sein Sicherheitsvorgehen dahingehend zu optimieren.  
- Interessant ist, dass die Backdoor nicht im eigentlichen Source Code des Git trees enthalten war, 
  sondern in binären Testdateien (für deren
  Generierung kein Sourcecode verfügbar war) versteckt war und der Code zur Aktivierung nur im Release-Tarball enthalten
  war, nicht jedoch im Git tree, was in Kombination dafür spricht, dass der Urheber der Sicherheitslücke versucht hat, so
  wenig "permanente" Spuren wie möglich zu hinterlassen, die von (automatischer) Code-Analyse einfach gefunden worden
  wären.  
- Ebenfalls stellt sich die Frage, in wie vielen anderen Projekten (Open Source wie proprietär), ähnliche Angriffe
  versucht wurden und vielleicht sogar erfolgreich waren. In Zukunft werde zumindest ich deutlich mehr Vorsicht gegenüber
  binären Dateien in Source Code walten lassen. Möglich wäre z.B., dass eine Backdoor in einem binary blob Treiber
  enthalten ist, der von einem Hardware-Hersteller veröffentlicht wird. Sollte dieser Treiber statisch (also nicht als
  Kernel Modul) in den Kernel compiliert werden, wäre es möglich, diese Schwachstelle auf Millionen von Geräten zu
  installieren. 
- Ähnliches Vorgehen in proprietärer Software (z.B. Microsoft Exchange) dürfte in ähnlicher Weise umsetzbar sein. Das
  Plazieren eines Maulwurfs in einem solchen Entwicklerteam könnte sogar einfacher sein, da das Vertrauen einem anderen
  Mitarbeiter gegenüber deutlich schneller anwächst als einem Fremden gegenüber.[^6]
- Mobbing und Burnout können ein echtes Sicherheitsrisiko sein. "Jia Tan" wäre vermutlich nie in diese Position gekommen,
  wäre der eigentliche Maintainer und Entwickler von xz nicht von Burnout betroffen und effektiv dazu gemobbt worden,
  "Jia Tan" zum Co-Maintainer und sogar primären Maintainer zu machen. 



[^1]: https://gynvael.coldwind.pl/?id=782  
[^2]: https://bsky.app/profile/filippo.abyssdomain.expert/post/3kowjkx2njy2b  
[^3]: https://boehs.org/node/everything-i-know-about-the-xz-backdoor  
[^4]: https://rheaeve.substack.com/p/xz-backdoor-times-damned-times-and
[^5]: https://en.wikipedia.org/wiki/Advanced_persistent_threat
[^6]: https://infosec.exchange/@SecureOwl/112181879860988729

