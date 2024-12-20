---
title: "Nullbulge-Ranner - Compiled Python Malware"
date: 2024-06-03
draft: false
---

Ich bin in den letzten Tagen über [Eric Parker](https://www.youtube.com/@EricParker) auf Malware gestößen, über BeamNG-Mods verteilt wurde.  

Das interessante an dieser Malware ist, dass sie zwar in Python geschrieben ist, aber mithilfe von Pyinstaller[^1] kompiliert wurde.
Dadurch, in Verbindung mit der gewählten Python-Version, ist die Malware effektiv nicht dekompilierbar und nur schwer zu reverse engineeren.
Die meisten Informationen über die Malware mussten daher über dynamische Analyse in Sandboxen und VMs gewonnen werden. 

## Oberflächliche statische Analyse
Der Erste Schritt war, die Datei "statisch" zu analysieren, das heißt, sich die Eigenschaften der Malware anzuschauen ohne die Malware selbst auszuführen.
Dafür habe ich initial die Tools "capa"[^2], "detect-it-easy"[^3] und "yara"[^4] verwendet.  

Detect-It-Easy gibt uns genauere Informationen über die Datei, vor allem, dass sie mit PyInstaller "gepackt" wurde.
```
PE64
    Linker: Microsoft Linker(14.36.33135)
    Compiler: Microsoft Visual C/C++(19.36.33135)[C]
    Tool: Visual Studio(2022 version 17.6)
    Packer: PyInstaller
```

Capa findet schon sehr viele Informationen:
{{< rawhtml >}}
<pre>
┍━━━━━━━━━━━━━━━━━━━━━━━━┯━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┑
│ ATT&amp;CK Tactic          │ ATT&amp;CK Technique                                                                   │
┝━━━━━━━━━━━━━━━━━━━━━━━━┿━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┥
│ <font color="#06989A">DEFENSE EVASION</font>        │ <font color="#06989A">Obfuscated Files or Information</font> T1027                                              │
│                        │ <font color="#06989A">Virtualization/Sandbox Evasion</font>::System Checks T1497.001                            │
├────────────────────────┼────────────────────────────────────────────────────────────────────────────────────┤
│ <font color="#06989A">DISCOVERY</font>              │ <font color="#06989A">File and Directory Discovery</font> T1083                                                 │
│                        │ <font color="#06989A">System Information Discovery</font> T1082                                                 │
├────────────────────────┼────────────────────────────────────────────────────────────────────────────────────┤
│ <font color="#06989A">EXECUTION</font>              │ <font color="#06989A">Command and Scripting Interpreter</font> T1059                                            │
│                        │ <font color="#06989A">Shared Modules</font> T1129                                                               │
┕━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┙

┍━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┯━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┑
│ MBC Objective               │ MBC Behavior                                                                  │
┝━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┿━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┥
│ <font color="#06989A">ANTI-BEHAVIORAL ANALYSIS</font>    │ <font color="#06989A">Virtual Machine Detection</font> [B0009]                                             │
├─────────────────────────────┼───────────────────────────────────────────────────────────────────────────────┤
│ <font color="#06989A">DATA</font>                        │ <font color="#06989A">Checksum</font>::Adler [C0032.005]                                                   │
│                             │ <font color="#06989A">Compress Data</font> [C0024]                                                         │
│                             │ <font color="#06989A">Compression Library</font> [C0060]                                                   │
│                             │ <font color="#06989A">Encode Data</font>::XOR [C0026.002]                                                  │
├─────────────────────────────┼───────────────────────────────────────────────────────────────────────────────┤
│ <font color="#06989A">DEFENSE EVASION</font>             │ <font color="#06989A">Obfuscated Files or Information</font>::Encoding-Standard Algorithm [E1027.m02]      │
├─────────────────────────────┼───────────────────────────────────────────────────────────────────────────────┤
│ <font color="#06989A">DISCOVERY</font>                   │ <font color="#06989A">Analysis Tool Discovery</font>::Process detection [B0013.001]                        │
│                             │ <font color="#06989A">File and Directory Discovery</font> [E1083]                                          │
│                             │ <font color="#06989A">System Information Discovery</font> [E1082]                                          │
├─────────────────────────────┼───────────────────────────────────────────────────────────────────────────────┤
│ <font color="#06989A">EXECUTION</font>                   │ <font color="#06989A">Command and Scripting Interpreter</font> [E1059]                                     │
├─────────────────────────────┼───────────────────────────────────────────────────────────────────────────────┤
│ <font color="#06989A">FILE SYSTEM</font>                 │ <font color="#06989A">Create Directory</font> [C0046]                                                      │
│                             │ <font color="#06989A">Delete Directory</font> [C0048]                                                      │
│                             │ <font color="#06989A">Delete File</font> [C0047]                                                           │
│                             │ <font color="#06989A">Read File</font> [C0051]                                                             │
│                             │ <font color="#06989A">Writes File</font> [C0052]                                                           │
├─────────────────────────────┼───────────────────────────────────────────────────────────────────────────────┤
│ <font color="#06989A">OPERATING SYSTEM</font>            │ <font color="#06989A">Environment Variable</font>::Set Variable [C0034.001]                                │
├─────────────────────────────┼───────────────────────────────────────────────────────────────────────────────┤
│ <font color="#06989A">PROCESS</font>                     │ <font color="#06989A">Create Process</font> [C0017]                                                        │
│                             │ <font color="#06989A">Terminate Process</font> [C0018]                                                     │
┕━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┙

┍━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┯━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┑
│ Capability                                           │ Namespace                                            │
┝━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┿━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┥
│ <font color="#06989A">reference analysis tools strings</font>                     │ anti-analysis                                        │
│ <font color="#06989A">reference anti-VM strings targeting Qemu</font>             │ anti-analysis/anti-vm/vm-detection                   │
│ <font color="#06989A">reference anti-VM strings targeting Xen</font>              │ anti-analysis/anti-vm/vm-detection                   │
│ <font color="#06989A">compute adler32 checksum</font>                             │ data-manipulation/checksum/adler32                   │
│ <font color="#06989A">compress data via ZLIB inflate or deflate</font>            │ data-manipulation/compression                        │
│ <font color="#06989A">encode data using XOR</font> (6 matches)                    │ data-manipulation/encoding/xor                       │
│ <font color="#06989A">accept command line arguments</font>                        │ host-interaction/cli                                 │
│ <font color="#06989A">query environment variable</font> (4 matches)               │ host-interaction/environment-variable                │
│ <font color="#06989A">set environment variable</font> (4 matches)                 │ host-interaction/environment-variable                │
│ <font color="#06989A">get common file path</font>                                 │ host-interaction/file-system                         │
│ <font color="#06989A">create directory</font> (2 matches)                         │ host-interaction/file-system/create                  │
│ <font color="#06989A">delete directory</font>                                     │ host-interaction/file-system/delete                  │
│ <font color="#06989A">delete file</font>                                          │ host-interaction/file-system/delete                  │
│ <font color="#06989A">enumerate files on Windows</font>                           │ host-interaction/file-system/files/list              │
│ <font color="#06989A">get file size</font>                                        │ host-interaction/file-system/meta                    │
│ <font color="#06989A">read file on Windows</font> (10 matches)                    │ host-interaction/file-system/read                    │
│ <font color="#06989A">write file on Windows</font> (2 matches)                    │ host-interaction/file-system/write                   │
│ <font color="#06989A">get disk information</font> (2 matches)                     │ host-interaction/hardware/storage                    │
│ <font color="#06989A">create process on Windows</font>                            │ host-interaction/process/create                      │
│ <font color="#06989A">terminate process</font>                                    │ host-interaction/process/terminate                   │
│ <font color="#06989A">link many functions at runtime</font> (2 matches)           │ linking/runtime-linking                              │
│ <font color="#06989A">linked against ZLIB</font>                                  │ linking/static/zlib                                  │
│ <font color="#06989A">parse PE header</font> (3 matches)                          │ load-code/pe                                         │
│ <font color="#06989A">resolve function by parsing PE exports</font>               │ load-code/pe                                         │
┕━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┙
</pre>
{{</ rawhtml >}}
Auch Yara findet einige Pattern in der Datei:
```
DebuggerException__SetConsoleCtrl
anti_dbg
screenshot
win_token
win_files_operation
Big_Numbers1 
CRC32_poly_Constant 
CRC32_table 
MachO_File_pyinstaller 
IsPE64 
IsWindowsGUI 
IsPacked 
HasOverlay 
HasDigitalSignature 
HasDebugData 
HasRichSignature 
Microsoft_Visual_Cpp_80 
Microsoft_Visual_Cpp_80_DLL 

```
Zusammenfassend können wir aus dieser statischen Analyse also sagen, dass die Malware es versucht, einem die Analyse schwerer zu machen und möglicherweise ein
RAT ist (kann laut YARA screenshots machen).

## Dynamische Analyse mit CAPE

Eine [CAPE-Sandbox-Umgebung](https://www.jmbit.de/blog/2023-11-30-capev2-sandbox/) ist eine gute Möglichkeit, lokal dynamische Analysen von Malware oder potenzieller
Malware durchzuführen. In diesem Fall enthält die zu analysierende Datei zwar keine sensiblen Daten, sollte das aber der Fall sein, ist es ggf. nicht sinnvoll,
die Datei auf eine Cloud-Platform hochzuladen.

Hier fielen vor allem die Anti-Analyse-Funktionen sowie die Requests zu Discord auf. Da die Malware jedoch ohne Verbindung zu Discord relativ inaktiv ist und keine user-interaktion bietet, ist die 
dynamische analyse hier abgesehen von den Informationen über die installation der Software und ihr Verhalten gegenüber Virtualisierung und Debuggern nicht besonders ergibig.

[REPORT](https://cloud.jmbit.de/s/qLtzrFyjY3TfQmX)

## Analyse des Webtraffic mithilfe von MITMProxy

Um den Webtraffic genauer anzusehen und mitlesen zu können und eventuelle API-Keys abfangen zu können, habe ich die Malware in einer VM gestartet, die über mitmproxy[^5] ins Internet gegangen ist.
Hier wurde aber lediglich versucht, sich als ein Bot in Discord einzuloggen, der Zugriff wurde jedoch abgelehnt:
```http
GET /api/v10/users/@me HTTP/1.1
Host: discord.com
User-Agent: DiscordBot (https://github.com/Rapptz/discord.py 2.3.2) Python/3.10 aiohttp/3.9.5
Authorization: Bot <REDACTED>

HTTP/1.1 401 Unauthorized

{"message": "401: Unauthorized", "code": 0}
```

## Analyse des Programmcodes

Nachdem ich der Malware in der Dynamischen Analyse kaum Informationen entlocken konnte, ging es an die statische Analyse. Da Python 3.10 leider zu neu für die meisten Decompiler ist[^6], musste ich mich mit dem 
Ergebnis eines Disassemblers zufrieden geben. Aus dem Output dieses (Immerhin 21355 Zeilen!) lies sich der Verdacht, dass es sich bei der Malware um eine Version von PySilon handelt, bestätigen.
Es wurden jedoch einige zusätzliche Funktionen eingefügt, so zum Beispiel Ransomware und die Möglichkeit des Nachladens eines Cryptominers. Ebenfalls bin ich über während der Analyse auf andere Malware der
gleichen Gruppe gestoßen.

Die Ransomware-Mitteilung `NULLBULGE-RANSOMWARE-NOTE.TXT` würde diesen Text enthalten:
```
READ THIS ENTIRELY BEFORE ACTING, IF YOU SHUT DOWN NOW, ALL WILL BE LOST.
Your computer is now infected with ransomware. Your file are encrypted with a secure algorithm that is impossible to crack.
By now it has moved from your main drive, and has started encrypting other drives.
DONT PANIC! IF YOU START LOOKING FOR THE VIRUS, YOU WILL NOT GET ANY FILES ALREADY ENCRYPTED BACK! BUT THEY CAN BE SAVED!
To recover your files you need a key. This key is generated once your file have been encrypted. To obtain the key, you must purchase it.
You can do this by sending 100 USD to this monero address:
45i7kjWZuzJ4PdSbandaaE8S6mQATmneTYEpgsaaCqDmc7foEJDXwxd3ABR8bn6YE4c7hZ2dYEEr1CwG48gAknPL6zUpYyV
Don't know how to get monero? Here are some websites:
https://www.kraken.com/learn/buy-monero-xmr
https://localmonero.co/?language=en
https://www.bestchange.com/visa-mastercard-usd-to-monero.html
Cant get monero and want to pay via giftcards instead? contact the email below.
Once you have sent the ransom to the monero address you must write an email this this email address: ZeCBMail@proton.me
In this email you will include your personal ID so we know who you are. Your personal ID is: 
<ID>
Payment is flexable, if you want to discuss pricing send an email with your discord username and I will contact you.
Be warned... pricing can go up too!
Once you have completeted all of the steps, you will be provided with the key to decrypt your files.
Don't know how ransomware works? Read up here:
https://www.trellix.com/en-us/security-awareness/ransomware/what-is-ransomware.html
https://www.checkpoint.com/cyber-hub/threat-prevention/ransomware/
https://www.trendmicro.com/vinfo/us/security/definition/Ransomware
Note: Messing with the ransomware will simply make your files harder to decrypt. Deleting the virus will make it impossible, as the key can not be generated.
Good luck
-NullBulge
```

Der Crypto-Miner ist XMRIG und mined für den Pool "xmrpool.eu".

Der Bot wird sowohl über Discord-Reactions als auch über direkte Befehle gesteuert. Die Reactions können einfachere Funktionen abbilden, wie z.B. das wipen der Malware vom System.
Interessanter sind jedoch die "richtigen" Befehle, z.B.:
 * '.ss' -> Macht einen Screenshot vom Bildschirm des Opfers
 * '.screenrec' -> zeichnet 15 Sekunden den Bildschirm auf
 * '.join' -> Tritt einem Voice Channel bei und überträgt die Aufnahmen des Mikrofons
 * '.show' -> Zeigt informationen, z.B. alle laufenden Prozesse
 * '.kill' -> Beendet einen Prozesse
 * '.block-input' -> Deaktiviert Tastatur- und Mauseingaben
 * '.break-windows' -> Bennent den Boot-Manager um und verhindert damit, dass Windows rebooten kann
 * '.xmrig' -> Führt den Cryptominer aus
 * '.webcam photo' -> Macht ein Bild mit der Webcam
 * '.encrypt' -> Verschlüsselt alle Dateien im angegebenen Verzeichnis
 * '.cmd' -> Führt beliebigen Befehl aus
 * '.admin' -> Kann Funktionen ausführen wie Defender, Taskmanager oder UAC deaktiveren.







[^1]: Pyinstaller ist Software, die es ermöglicht, in Python geschriebene Programme so zu compilen und bundlen, dass sie auf Systemen ohne Python-Interpreter
      benutzt werden können. Das kann z.B. als PE-File (Windows .exe) oder ELF (Linux/Unix) sein.
[^2]: [mandiant/capa](https://github.com/mandiant/capa), ein Tool, das eine Datei analysiert und versucht, deren Fähigkeiten (capabilities) zu erfassen.
[^3]: [horsicq/Detect-It-Easy](https://github.com/horsicq/Detect-It-Easy) gibt einem die Eigenschaften einer Datei zurück.
[^4]: [Yara](https://virustotal.github.io/yara/) ist eine Bibliothek zur Analyse von Malware anhand von Pattern/Eigenschaften der Dateien.
[^5]: [MITMProxy](https://mitmproxy.org/) ist ein Forward-Proxy, der alle Verbindungen die über ihn laufen aufzeichnet
[^6]: Weder [pycdc](https://github.com/zrax/pycdc) noch [unpyc37-3.10](https://github.com/greyblue9/unpyc37-3.10) konnten den Bytecode decompilen.
