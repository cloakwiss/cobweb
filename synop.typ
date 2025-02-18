#let guide = "Prof Kaushik Roy"
#let members = (
  ("Ayush Paranjale", 35),
  ("Himanshu Pawar", 43),
  ("Gunangi Bhagat", 5),
  ("Prabhu Kalantri", 57),
)

#import "./extra.typ": header, footer

#header(members, guide)

= Problem Definition

With the rapid evolution of the internet, online content is frequently modified, moved, or deleted, leading to broken links and loss of valuable information.
Users who rely on web pages for research, reference, or personal archiving often face challenges in preserving content for future access.
Existing solutions may be complex, require subscriptions, or fail to provide a convenient offline storage format.
There is a need for a minimalistic and efficient tool that allows users to archive web pages seamlessly and store them in a structured, portable format.

= Project Objective

A minimal tool for downloading and archiving a website which is aimed to be used on personal devices where the webpage/website will be stored as ePUB.
It is aware of network constraints namely bandwidth and total data available to work with, and has interface to ensure user does not download more than they are prepared to, which is not the focus of most other similar tools.

= Proposed Plan of Work

== Literature Review
- _Archiving techniques_: Research existing web archiving methodologies, standards, and best practices
- _Convert for file formats ie html to xhtml_: Study conversion processes and standards for transforming HTML to XHTML while preserving structure and content
- _Authoring ePUB file_: Understand the ePUB specification, structure requirements, and creation process

== Requirement Analysis
- _Site scraping_: Define requirements for fetching web content, including handling of various media types, following links, and respecting robots.txt
- _Show page content size_: Develop mechanisms to analyze and display content size before and during downloading
- _Converting to epub_: Determine requirements for proper ePUB conversion, including metadata handling, content organization, and navigation

== Testing
- _Testing bounded site scraping_: Verify the tool's ability to limit crawling depth and respect site boundaries
- _Conversion to epub for small sites_: Ensure proper conversion of small websites with various content types

== Evaluation
- _Evaluating the scraper on large site and profiling:_ Measure performance, memory usage, and download efficiency on larger websites
- _Mapping web routes to epub chapters/section:_ Assess the effectiveness of the algorithm that converts website structure to ePUB navigation

== Documentation & Reporting
- _Creating a man page_: Develop comprehensive command-line documentation
- _Guides for installation_: Create user-friendly installation instructions for different operating systems

= Methodology

- _Fetching specified webpages and its assets_: Develop algorithms to retrieve the main webpage content along with associated resources like images, CSS, and scripts
- _Recursively fetch all subsequent page if needed_: Implement depth-limited crawling to retrieve linked pages based on user preferences
- _Remap all the routes to point pages stored in file system_: Convert absolute URLs to relative paths that work within the ePUB structure
- _Sanitize and convert HTML to XHTML_: Clean malformed HTML and ensure XHTML compliance for ePUB compatibility
- _Creating manifest file and table of content, etc for ePUB_: Generate the required metadata files according to ePUB specifications
- Archive all the file finally in ePUB: Package all converted content into a valid ePUB container

= Technology

- *Go*
  - Has good concurrency pattern: Utilize Go's goroutines and channels for efficient parallel processing
  - Has feature rich standard library particularly for http client and server: Leverage built-in HTTP handling capabilities for reliable web scraping

- *XML*
  - It is foundation of epub: Use XML processing libraries to handle the ePUB structural requirements

- *HTML-Tidy*
  - HTML-Tidy corrects and clean up HTML content by fixing markup errors such as mismatched, mis-nested, and missing tags; missing end "/" tags; missing quotations; and many, many more discrepant conditions, and serves as an HTML pretty printer.

- *Pandoc*
  - A popular inter-document conversion tool: Potentially integrate with Pandoc for advanced document transformation

= Functional Specification (Deliverables)

- _Fetching and storing webpage/website recursively_: The tool will download specified web content and follow links according to user-defined parameters
- _Conversion of HTML to XHTML_: The tool will sanitize and transform HTML content to valid XHTML for ePUB compatibility
- _Archiving as ePUB_: The tool will package all retrieved content into a standard ePUB format for portability and accessibility

= Project Scope

A minimal web archiving tool that enables users to fetch and store webpages or entire websites recursively.
The tool will ensure content preservation by converting HTML to well-structured XHTML, maintaining compatibility and consistency across different devices.
Additionally, the archived content will be bundled into an EPUB format, providing a portable and easily accessible offline reading experience.
The project will focus on efficiency, simplicity, and usability, ensuring a seamless process for users to save and organize web content for future reference.

#footer(members, guide)
