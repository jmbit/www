---
title: "CAPEv2 Sandbox Teil 1"
date: 2023-11-30
draft: false
---

Die CAPEv2-Sandbox-Umgebung ist sehr nützlich für die Analyse von Malware, da man hier die Malware in einer
kontrollierten Umgebung "zünden" kann und eine genaue Protokollierung hat. Die Installation und konfiguration von CAPE
ist jedoch nicht besonders einfach. Dieser Blogeintrag kann Ihnen hoffentlich als Einstieg dienen um sich eine
solche Umgebung aufzubauen. Bei Rückfragen und weiterführenden Themen stehe ich gerne zur Verfügung.

## Basis
Die Installations- und Konfigurationsskripte von CAPEv2 sind auf Ubuntu 22.04 optimiert. Eine Installation
unter Debian ist ebenfalls mit vergleichsweise wenig Änderungen möglich. Für andere Distributionen und MacOS sind
deutlich weiterführendere Änderungen notwendig, eine Installation unter Microsoft Windows ist nicht möglich.
Bei der Installation in einer Virtuellen Maschine muss "Nesting"-Aktiviert sein. Die Hardware-Anforderungen sind
stark von der geplanten Art und Anzahl von VMs ab. Es ist sinnvoll, mindestens 4 GB für OS und Applikation
einzurechnen, plus den Hardwareressourcen die für die maximale Anzahl an gleichzeitig laufenden VMs benötigt werden.


## CAPEv2 Download und Vorbereitungen:
CAPEv2 herunterladen:
```bash
sudo git clone https://github.com/kevoreilly/CAPEv2 /opt/CAPEv2
```
Sobald man CAPEv2 heruntergeladen hat, sollte man das Skript `/opt/CAPEv2/installer/kvm-qemu.sh` anpassen.
Die grundsätzliche Anleitung dafür steht in diesem Skript, für den Austausch der \<WOOT\>-Einträge für die
Hardware IDs ist die Website der UEFI alliance mit der [Liste der ACPI-IDs](https://uefi.org/ACPI_ID_List).
Für Testinstallationen ist es aber auch möglich, zufällige 4 Zeichen zu nutzen.
sobald dieses Skript angepasst wurde, kann es ausgeführt werden.
```bash
cd /opt/CAPEv2/installer/
sudo ./kvm-qemu.sh all yourusername | tee kvm-qemu-all.log
# optional
sudo ./kvm-qemu.sh virt-manager yourusername | tee kvm-qemu-virt-manager.log
```
Hierbei sollte besonders darauf geachtet werden, dass APT auch alle Pakete installieren kann. 
Nachdem die installation erfolgreich war, ist es sinnvoll, das System neu zu starten.

## CAPE installieren
für die Installation von CAPEv2 führt man die folgenden Befehle aus:
```bash
cd /opt/CAPEv2/installer
sudo ./cape.sh all | tee cape-all.log
reboot now
cd /opt/CAPEv2
sudo -u cape poetry install
poetry env list
sudo -u cape poetry run pip install -r extra/optional_dependencies.txt
reboot now
```


## Internet via Cloudflare-VPN
Es ist sinnvoll, den Internet-Zugang der VMs über einen VPN laufen zu lassen und nicht über das eigene (lokale) Netz.
hierfür können diverse VPN-Anbieter verwendet werden, in diesem Fall nutze ich der Einfachheit halber Cloudflare Warp
mit seiner SOCKS-Proxy-Funktion.
Dafür führt man die folgenden Befehle aus:
```sh
curl -fsSL https://pkg.cloudflareclient.com/pubkey.gpg | sudo gpg --yes --dearmor --output /usr/share/keyrings/cloudflare-warp-archive-keyring.gpg
echo "deb [arch=amd64 signed-by=/usr/share/keyrings/cloudflare-warp-archive-keyring.gpg] https://pkg.cloudflareclient.com/ $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/cloudflare-client.list
sudo apt-get update && sudo apt-get install cloudflare-warp
warp-cli register --accept-tos
warp-cli set-families-mode off
warp-cli set-mode proxy
warp-cli connect
warp-cli status
ss -tulpn | grep warp # Hier sollte hoffentlich 127.0.0.1:40000 stehen

```

---
Links:
https://github.com/kevoreilly/CAPEv2 - Github-Seite von CAPEv2
https://capev2.readthedocs.io/en/latest/# - CAPEv2 offizielle Dokumentation
https://capesandbox.com/ - CAPEv2 öffentliche Demoinstanz
