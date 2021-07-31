SELECT
  product_id,
  recipient_name,
  recipient_address,
  commission_percent
FROM commission
WHERE product_id=ANY($1);
