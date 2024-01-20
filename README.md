# lure-web

`lure-web` is a website for LURE and a web interface that lets users search and get information about LURE's packages.

LURE Web is powered by [Salix](https://gitea.elara.ws/Elara6331/salix) and [PicoCSS](https://picocss.com/).

## API

### Search Packages

- **Endpoint:** `GET /api/search`
- **Description:** Search for packages based on specified search options.
- **Parameters:**
  - None
- **Query Parameters:**
  - `sort` (optional) - Sort the results by a specific attribute.
    - Possible values: "name", "version", "repo"
  - `filter` (optional) - Apply a filter to narrow down the search results.
    - Possible values: "inrepo", "arch"
  - `fv` (optional) - Filter value to be used based on the selected filter.
  - `q` (optional) - The search query string.
- **Response:**
  - Content-Type: application/json
  - Body: List of translated packages based on search criteria.

### Get Package Information

- **Endpoint:** `GET /api/pkg/:repo/:pkg`
- **Description:** Get detailed information about a specific package.
- **Parameters:**
  - `:repo` - Repository name
  - `:pkg` - Package name
- **Response:**
  - Content-Type: application/json
  - Body: Translated information about the specified package.

### Get Package Script

- **Endpoint:** `GET /api/pkg/:repo/:pkg/script`
- **Description:** Get the script associated with a specific package.
- **Parameters:**
  - `:repo` - Repository name
  - `:pkg` - Package name
- **Query Parameters:**
  - `highlight` (optional) - If set to `true`, the script will be returned with syntax highlighting.
- **Response:**
  - If `highlight` is `true`:
    - Content-Type: text/html
    - Body: Script with syntax highlighting in HTML format.
  - If `highlight` is not provided or set to "false":
    - Content-Type: text/x-shellscript
    - Body: Raw script content.

### Language Translation

- **Header:** `Accept-Language`
  - Use the Accept-Language header to specify the preferred language for response translations.
- **Query Parameter:** `lang`
  - You can also include a `lang` query parameter instead of `Accept-Language` to specify the translation language.

### Error Handling

- **Response:**
  - Content-Type: application/json
  - Body: A JSON object with a single `error` string field containing an explanation of the error.

---


