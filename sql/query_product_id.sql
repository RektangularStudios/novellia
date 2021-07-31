WITH modified_timestamps AS (
  SELECT product_id, modified FROM product
  UNION
  SELECT product_id, organization.modified AS modified
    FROM organization
    INNER JOIN product ON product.organization_id = organization.organization_id
  UNION
  SELECT product_id, market.modified AS modified
    FROM market
    INNER JOIN product ON product.market_id = market.market_id
  UNION
  SELECT product_id, modified FROM product_detail
  UNION
  SELECT product_id, native_token.modified AS modified
    FROM native_token
    INNER JOIN product ON product.native_token_id = native_token.native_token_id
  UNION
  SELECT product_id, modified FROM commission
  UNION
  SELECT product_id, modified FROM product_attribution
  UNION
  SELECT product_id, modified FROM remote_resource
),
modified_timestamps_max AS (
  SELECT
    product_id,
    MAX(modified_timestamps.modified) AS modified
  FROM modified_timestamps
  GROUP BY product_id
)

SELECT
  product.product_id,
  COALESCE(native_token_id, ''),
  modified_timestamps_max.modified
FROM product
INNER JOIN product_detail ON product.product_id = product_detail.product_id
INNER JOIN modified_timestamps_max ON product.product_id = modified_timestamps_max.product_id
WHERE
  product.deleted IS NULL AND
  (($1::TEXT <> '') IS NOT TRUE OR $1::TEXT = product.organization_id) AND
  (($2::TEXT <> '') IS NOT TRUE OR $2::TEXT = product.market_id) AND
  (date_listed IS NOT NULL AND date_listed <= NOW())
ORDER BY product_detail.id ASC;
