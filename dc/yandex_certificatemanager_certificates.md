# Table: yandex_certificatemanager_certificates


The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|id (PK)|String|
|folder_id|String|
|created_at|Timestamp|
|name|String|
|description|String|
|labels|JSON|
|type|Int|
|domains|StringArray|
|status|Int|
|issuer|String|
|subject|String|
|serial|String|
|updated_at|Timestamp|
|issued_at|Timestamp|
|not_after|Timestamp|
|not_before|Timestamp|
|challenges|JSON|
|deletion_protection|Bool|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|