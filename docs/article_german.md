# Technical Debt Records (TDRs) und das Werkzeug zu ihrer Erstellung

## Einleitung

In der heutigen schnelllebigen Softwareentwicklung stehen Teams vor der Herausforderung, kontinuierlich neue Funktionen bereitzustellen und gleichzeitig die Codequalität zu erhalten. Dabei entstehen oft Kompromisse, die als **Technical Debt** (Technische Schulden) bezeichnet werden. Um diese systematisch zu dokumentieren und zu verwalten, haben sich **Technical Debt Records (TDRs)** etabliert. In diesem Artikel erläutern wir die Bedeutung von TDRs, wie sie Entwicklern, Architekten und Testern helfen, und stellen ein Werkzeug vor, das die Erstellung von TDRs vereinfacht.

## Was sind Technical Debt Records (TDRs)?

Ein **Technical Debt Record (TDR)** ist ein strukturiertes Dokument, das Details über technische Schulden in einem Softwareprojekt festhält. Technische Schulden entstehen, wenn kurzfristige Lösungen gewählt werden, die zwar schnell implementiert werden können, aber langfristig zu erhöhtem Wartungsaufwand, schlechterer Performance oder anderen Nachteilen führen. TDRs bieten eine klare Übersicht über bestehende technische Schulden, deren Auswirkungen und die Maßnahmen zu ihrer Behebung.

## Motivation für TDRs

Technische Schulden können, wenn sie unkontrolliert bleiben, erhebliche negative Auswirkungen auf ein Projekt haben:

- **Codequalität:** Erhöhter Wartungsaufwand und sinkende Codequalität.
- **Skalierbarkeit:** Schwierigkeiten bei der Erweiterung und Anpassung der Software.
- **Performance:** Mögliche Leistungseinbußen durch suboptimale Implementierungen.
- **Risikomanagement:** Erhöhtes Risiko von Systemausfällen oder Sicherheitslücken.

Durch die systematische Dokumentation mittels TDRs können Teams diese Schulden frühzeitig erkennen, priorisieren und gezielt angehen, bevor sie unkontrollierbar werden.

## Vorteile von TDRs für Entwickler, Architekten und Tester

### Für Entwickler

- **Transparenz:** Klare Dokumentation der bestehenden technischen Schulden erleichtert das Verständnis der Codebasis.
- **Priorisierung:** Ermöglicht die Fokussierung auf kritische Bereiche, die sofortige Aufmerksamkeit erfordern.
- **Wiederverwendbarkeit:** Vermeidung von doppeltem Aufwand durch das Bewusstsein über bereits bekannte Probleme.

### Für Architekten

- **Strategische Planung:** Unterstützung bei der Planung von Refactoring-Maßnahmen und Architekturverbesserungen.
- **Risikobewertung:** Einschätzung der Auswirkungen technischer Schulden auf die Gesamtarchitektur.
- **Entscheidungsgrundlage:** Datenbasierte Entscheidungen zur Weiterentwicklung der Systemarchitektur.

### Für Tester

- **Fokus auf kritische Bereiche:** Kenntnis über problematische Bereiche ermöglicht gezieltere Tests und höhere Testabdeckung.
- **Verbesserte Teststrategien:** Anpassung der Testpläne basierend auf den identifizierten technischen Schulden.
- **Qualitätssicherung:** Sicherstellung, dass behobene Schulden die gewünschte Qualitätssteigerung bringen.

## Das TDR-Template und seine Felder

Ein gut strukturiertes TDR-Template ist entscheidend für die effektive Dokumentation technischer Schulden. Unser Werkzeug generiert TDRs mit folgenden Feldern:

1. **Titel:** Eine prägnante Bezeichnung der technischen Schuld.
2. **Autor:** Die Person, die die Schuld identifiziert oder dokumentiert hat.
3. **Version:** Die Version des Projekts oder der Komponente, in der die Schuld existiert.
4. **Datum:** Das Datum der Identifikation oder Dokumentation der Schuld.
5. **State:** Der aktuelle Status der technischen Schuld (z.B. Identified, Analyzed, Approved, In Progress, Resolved, Closed, Rejected).
6. **Relations:** Verweise auf andere verwandte TDRs, um Zusammenhänge zu verdeutlichen.
7. **Zusammenfassung:** Eine kurze Übersicht über die technische Schuld und deren Bedeutung.
8. **Kontext:** Detaillierte Hintergrundinformationen, warum die Schuld entstanden ist (z.B. Zeitdruck, veraltete Technologien).
9. **Auswirkungen:**
   - **Technische Auswirkungen:** Wie die Schuld die Systemleistung, Skalierbarkeit oder Wartbarkeit beeinflusst.
   - **Geschäftliche Auswirkungen:** Die Auswirkungen auf Geschäftsprozesse, Kundenzufriedenheit oder Risikoebenen.
10. **Symptome:** Beobachtbare Anzeichen, die auf die Präsenz der technischen Schuld hinweisen (z.B. häufige Bugs, langsame Performance).
11. **Schweregrad:** Die Kritikalität der Schuld (Critical, High, Medium, Low).
12. **Potenzielle Risiken:** Mögliche negative Folgen, wenn die Schuld nicht behoben wird (z.B. Sicherheitslücken, erhöhte Kosten).
13. **Vorgeschlagene Lösung:** Empfohlene Maßnahmen zur Behebung der technischen Schuld.
14. **Kosten der Verzögerung:** Die Konsequenzen, wenn die Behebung der Schuld verzögert wird.
15. **Aufwand zur Behebung:** Geschätzter Aufwand in Zeit und Ressourcen, um die Schuld zu beheben.
16. **Abhängigkeiten:** Andere Aufgaben, Komponenten oder externe Faktoren, die die Behebung der Schuld beeinflussen.
17. **Zusätzliche Hinweise:** Weitere relevante Informationen oder Überlegungen zur Schuld.

### Rationale für das `State`-Feld

Das `State`-Feld spiegelt den Workflow wider, wie technische Schulden gehandhabt werden. Es hilft dabei, den Fortschritt bei der Bearbeitung der Schuld zu verfolgen und sicherzustellen, dass keine Schulden unbeachtet bleiben. Die definierten Zustände sind:

- **Identified:** Die technische Schuld wurde erkannt.
- **Analyzed:** Die Auswirkungen und der Aufwand zur Behebung wurden bewertet.
- **Approved:** Die Behebung der Schuld wurde genehmigt.
- **In Progress:** Die Arbeit zur Behebung der Schuld ist im Gange.
- **Resolved:** Die technische Schuld wurde behoben.
- **Closed:** Der TDR wurde abgeschlossen und ist nicht mehr relevant.
- **Rejected:** Die Behebung der Schuld wurde abgelehnt.

Diese Zustände ermöglichen es Teams, den Status jeder technischen Schuld klar zu definieren und entsprechende Maßnahmen zu ergreifen.

### Anpassung der Felder je nach State

Beim ersten Identifizieren einer technischen Schuld können einige Felder noch leer bleiben:

- **Initiale Identifikation (`Identified`):**
  - **Gefüllt:** Titel, Autor, Version, Datum, State, Zusammenfassung, Kontext.
  - **Leer:** Auswirkungen, Symptome, Schweregrad, Potenzielle Risiken, Vorgeschlagene Lösung, Kosten der Verzögerung, Aufwand zur Behebung, Abhängigkeiten, Zusätzliche Hinweise.

- **Analysephase (`Analyzed`):**
  - **Gefüllt:** Alle Felder des `Identified`-Status plus Auswirkungen, Symptome, Schweregrad, Potenzielle Risiken.

- **Genehmigungsphase (`Approved`):**
  - **Gefüllt:** Alle bisherigen Felder plus Vorgeschlagene Lösung, Kosten der Verzögerung.

- **Umsetzungsphase (`In Progress`):**
  - **Gefüllt:** Alle bisherigen Felder plus Aufwand zur Behebung, Abhängigkeiten.

- **Abschlussphase (`Resolved` & `Closed`):**
  - **Gefüllt:** Alle Felder einschließlich Zusätzliche Hinweise.
  
Durch diese schrittweise Ergänzung bleibt die Dokumentation stets aktuell und reflektiert den Fortschritt bei der Behebung der technischen Schulden.

## Das Werkzeug zur Erstellung von TDRs

Unser **TDR-Generator** ist ein Go-basiertes Tool, das die Erstellung von Technical Debt Records in verschiedenen Formaten automatisiert. Es unterstützt **Markdown**, **Plain ASCII**, **PDF** und **Excel** und erleichtert somit die Integration in unterschiedliche Entwicklungs- und Dokumentationsprozesse.

### Features des TDR-Generators

- **Benutzerfreundlich:** Interaktive Eingabeaufforderungen führen den Benutzer durch das Ausfüllen der TDR-Felder.
- **Flexibel:** Unterstützung mehrerer Ausgabeformate zur Anpassung an verschiedene Anforderungen.
- **Automatische Validierung:** Überprüfung der Eingaben auf Vollständigkeit und Korrektheit.
- **Version Control Integration:** Leichtes Einchecken der generierten TDRs in Systeme wie Git oder SVN.

### Repository und Installation

Der TDR-Generator ist auf GitHub verfügbar. Sie können das Repository [hier](https://github.com/yourusername/technical-debt-generator) finden.

#### Schritte zur Installation:

1. **Clone das Repository:**

   ```bash
   git clone https://github.com/yourusername/technical-debt-generator.git
   cd technical-debt-generator
   ```

2. **Initialisiere das Go-Modul:**

   ```bash
   go mod init technical_debt_generator
   ```

3. **Installiere die Abhängigkeiten:**

   ```bash
   go get github.com/phpdave11/gofpdf
   go get github.com/xuri/excelize/v2
   ```

4. **Build das Programm:**

   ```bash
   go build generate-td.go
   ```

### Nutzung des TDR-Generators

Das Programm kann über die Kommandozeile mit verschiedenen Optionen ausgeführt werden.

#### Verfügbare Optionen:

- `-format`: Gibt das Ausgabeformat an. Mögliche Werte sind `markdown`, `ascii`, `pdf`, `excel`. Standard ist `markdown`.
- `-output`: (Optional) Gibt den Dateinamen für die Ausgabe an. Wenn nicht angegeben, wird ein Standardname verwendet.
- `-empty`: (Optional) Erstellt ein leeres TDR mit Platzhaltern, ohne nach Eingaben zu fragen.
- `--help` oder `-h`: Zeigt die Hilfsnachricht mit Nutzungshinweisen an.

#### Beispiele:

1. **Generiere ein Markdown-TDR:**

   ```bash
   ./generate_td -format markdown
   ```

2. **Generiere ein PDF-TDR mit benutzerdefiniertem Dateinamen:**

   ```bash
   ./generate_td -format pdf -output mein_td_record.pdf
   ```

3. **Generiere ein leeres Excel-TDR:**

   ```bash
   ./generate_td -format excel -empty
   ```

4. **Zeige die Hilfemeldung an:**

   ```bash
   ./generate_td --help
   ```

### Integration in Versionskontrollsysteme

Um die Nachverfolgung und Zusammenarbeit zu erleichtern, sollten TDRs in ein Versionskontrollsystem wie **Git** oder **SVN** eingecheckt werden. Dies ermöglicht:

- **Versionierung:** Nachverfolgung von Änderungen und Historie der technischen Schulden.
- **Zusammenarbeit:** Gemeinsame Bearbeitung und Überprüfung von TDRs durch das Team.
- **Zugänglichkeit:** Einfache Integration in bestehende Entwicklungsprozesse und Pipelines.

#### Beispiel für Git:

1. **Füge das TDR dem Repository hinzu:**

   ```bash
   git add technical_debt_record.md
   ```

2. **Commit die Änderung:**

   ```bash
   git commit -m "Add TDR für veraltete Authentifizierungsbibliothek"
   ```

3. **Push die Änderung:**

   ```bash
   git push origin main
   ```

Durch die Einbindung von TDRs in das Versionskontrollsystem bleibt die Dokumentation stets aktuell und für alle Teammitglieder zugänglich.

## Fazit

**Technical Debt Records (TDRs)** sind ein unverzichtbares Instrument zur Verwaltung technischer Schulden in Softwareprojekten. Sie bieten Transparenz, erleichtern die Priorisierung und unterstützen strategische Entscheidungen zur Verbesserung der Codequalität und Systemarchitektur. Der vorgestellte **TDR-Generator** vereinfacht die Erstellung dieser wichtigen Dokumente und integriert sich nahtlos in bestehende Entwicklungs- und Versionskontrollprozesse.

Indem Teams TDRs konsequent verwenden und in ihre Workflows integrieren, können sie die negativen Auswirkungen technischer Schulden minimieren und die langfristige Gesundheit und Wartbarkeit ihrer Softwareprojekte sicherstellen.


