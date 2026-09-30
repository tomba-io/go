# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)

### v1.1.2

Bulk jobs (https://docs.tomba.io/bulks)

- New: `CreateBulkJob` with the canonical parameters (`data` rows, zero-based
  `*_field_index`, `verify_emails`, `find_phones`, `webhook_url`, `launch`...),
  `GetBulkJob`, `ListBulkJobs` (status / name filters), `CancelBulk`,
  `RetryBulk`, `RestoreBulk`, `GetBulkEstimate`, `GetBulkHistory`,
  `GetBulkJobProgress` (live metrics, ETA), `GetBulkTypes`, `GetBulkStats`,
  `GetBulkWebhookSecret`, `RotateBulkWebhookSecret`, `DownloadBulkTo` (streams).
- New: `models.BulkConfig`, `models.BulkMetrics` (result distribution,
  `Results.Records`: rows in the full results file), `BulkTypeTechnology`,
  `BulkTypeExport`, status constants, `BulkCreateResponse.Data.Message` (why a
  requested launch was refused).
- New: `WithBaseURL` and a typed `*APIError` (HTTP status, `errors.type`,
  message, Retry-After) for every call.
- Fixed: file uploads kept `Content-Type: application/json` and sent the CSV as
  `application/octet-stream`; they are now multipart with a `text/csv` part.
- Fixed: `BulkItem.Status` / `BulkProgress.Status` are strings (the API never
  sent booleans).
- Deprecated: `CreateBulk`, `CreateBulkWithFile`, `CreateSearchBulk`,
  `CreateFinderBulk`, `CreatePhoneValidatorBulk`, `GetAllBulks`, `GetBulk`,
  `GetBulkProgress`, `DownloadBulk` and `SaveBulkResults`. They keep their
  signatures but send the current parameter names, as the API ignores the old
  ones: `include_sources`, `notify`, `verify_emails`,
  `boost_score_from_sources`, `max_emails_per_domain`, `email_type` +
  `email_type_mode`, `departments` + `departments_mode`,
  `skip_rows_with_email`, zero-based `<field>_field_index` (the `Column*`
  parameters stay one-based), `archived=true` and `?file=`. An option is only
  sent to the types that take it. `Total` is no longer sent.

**Breaking** (the API no longer returns them):

- `BulkItem` loses the flat fields `Maximum`, `EmailType`, `Department`,
  `Sources`, `Verify`, `Phone`, `Full`, `Total`, `TotalList`, `TotalEmails`,
  `VerifyCost`, `SearchCost`, `PhoneCost`, `TimeTrack`, `Chart`, `Table`,
  `FileName`, `Launched` and `Used`. Read `Config` (options, `Input.Rows`) and
  `Metrics` (results, `Emails.Total`, credits, timing, distributions) instead;
  `Used` is now `Billed`, and `Table` is `Preview` (detail only, once billed).
- `BulkProgress.ProcessedEmail` is removed. A job that was never launched now
  has the status `pending`.

### v1.0.0

Initial commit