SELECT
  product_id,
  resource_id,
  resource_description,
  priority,
  multihash,
  hash_source_type,
  resource_urls,
  content_type
FROM remote_resource
WHERE product_id=ANY($1);
