---
title: "Container"
draft: true
---

Mit den Container-Verwaltungstools Docker, Podman und Kubernetes konnte ich bereits einige Erfahrung sammeln, und nutze
diese regelmäßig. Container sind mittlerweile fast genauso weit verbreitet wie Virtuelle Maschinen, haben diesen
gegenüber jedoch einige Vor-, aber auch Nachteile. Der größte Vorteil ist, dass im Gegensatz zu einer VM nicht das
gesamte Betriebssystem virtualisiert werden muss, d.h. ein Container enthält jediglich das Userland. Dadurch wird nicht
nur der vergleichsweise kleine Linux-Kernel (~258MB on Disk auf dem Laptop, mit dem ich diesen Artikel schreibe)
eingespart, sondern auch die meisten üblichen Verwaltungs-Tools, die durch extern laufende Tools
(Docker/Podman/Kubernetes) ersetzt werden. Dadurch wird die Skalierbarkeit vereinfacht. Zudem sind Container meist
*stateless*, haben also keine persistenten Daten. 

<!-- vim: set wrap linebreak textwidth=120 cc=120 spell spelllang=de,en : -->

