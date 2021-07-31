SELECT
  product_id,
  author_name,
  author_urls,
  work_attributed
FROM product_attribution
WHERE product_id=ANY($1);
