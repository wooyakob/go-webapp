# Couchbase document metadata

Source:
https://docs.couchbase.com/server/current/learn/data/data.html#metadata

Auto-generated and stored for each document.

Metadata attributes (reference):

- meta: Standard metadata for the saved document (example: `airport_1306`).
- id: The key of the saved document (example: `airport_1306`).
- rev: Revision/sequence number (internal server use). Used for resolving conflicts when replicated documents are updated concurrently on different servers (XDCR conflict resolution).
- expiration: Expiration time (TTL). If non-zero, determines when Couchbase Server removes the document. Can be set explicitly on a document or via `maxTTL` on the collection/bucket.
- flags: SDK-specific values that may identify the type/formatting of the stored value.
- type: The type of the stored value (example: `json`).
- xattrs: Extended Attributes. Special metadata; some system-internal, some optionally written/read by user applications.
