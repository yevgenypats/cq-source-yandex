# Table: yandex_kms_symmetric_keys


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
|status|Int|
|primary_version|JSON|
|default_algorithm|Int|
|rotated_at|Timestamp|
|rotation_period|JSON|
|deletion_protection|Bool|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|