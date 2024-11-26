---
title: "Investigating Nullbulge - a BeamNG mods malware Group"
date: 2024-06-05
draft: true
---

Over the last few days, I started Investigating a piece of malware that got downloaded with BeamNG[^1] mods.
It was originally discovered by Eric Parker, who describes both its discovery and major features in [his video](https://youtu.be/k52GwOWGy7o?si=4o5A7UkEpeuYWa20).

## First Sample
The sample originally uncovered by Eric Parker is a modified version of the PySilo[^2] malware that has been compiled using PyInstaller and Python 3.10,
making it fairly difficult to reverse engineer. However, I was still able to analyze it using both static and dynamic analysis tooling.

### General static analysis
First, I used a couple general tools for static file analysis. To be mentioned here are [^3]Capa, Detect-It-Easy[^4] and yara[^5].
All of those gave certain insights on the file and informed the way to go forward.

Detect-It-Easy is essentially a more powerful `file` command that gives you additional information about a file.
In this case, it gave me the following Information:  
```
PE64
    Linker: Microsoft Linker(14.36.33135)
    Compiler: Microsoft Visual C/C++(19.36.33135)[C]
    Tool: Visual Studio(2022 version 17.6)
    Packer: PyInstaller
```  
This tells us the File is a 64Bit PE (Portable Executable) packed with PyInstaller[^6], meaning it's compiled and packed from Python for
64 Bit Windows. In this case, we can likely ignore the Linker, compiler and Tool information, as Pyinstaller likely just uses those under the hood.  

Capa is much more verbose, but has a nicely formatted output. It analyzes the calls the program makes to the OS or other common locations to make informed
guesses about the likely capabilities/uses of the program:
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
If we weren't sure already, this wold certainly greatly increase the likelyhood of this software being malicious. No normal software needs to evade
sandboxing and virtualization. We will later see that these features are implemented in a very simple way and not really relevant for most sandboxing
platforms.

Finally, for a good measure, I ran it through yara:
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
Yara can be a bit noisy, but in this case it also shows us that the application has some anti-debug operations, the ability to take screenshots and some
features indicative of encryption.  

### Dynamic analysis

Dynamic analysis essentially just means watching the Malware do its thing and analyzing that. For this I usually use my own installation of CAPE sandbox,
but hosted services like tria.ge and intezer are also available. 

