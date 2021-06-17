# Ebooker

Downloads HTML files and generates epub files from them.

# Input

Input file:

- data.proto - Describes the book metadata and links to actual chapters.
  - title
  - author
  - Links to individual chapters
    - Selector for content in HTML page - shared or per chapter.
    - Name
    - URL
