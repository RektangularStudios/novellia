WITH modified_timestamps AS (
  SELECT product_id, modified FROM novellia.product
  UNION
  SELECT product_id, novellia.organization.modified AS modified
    FROM novellia.organization
    INNER JOIN novellia.product ON novellia.product.organization_id = novellia.organization.organization_id
  UNION
  SELECT product_id, novellia.market.modified AS modified
    FROM novellia.market
    INNER JOIN novellia.product ON novellia.product.market_id = novellia.market.market_id
  UNION
  SELECT product_id, modified FROM novellia.product_detail
  UNION
  SELECT product_id, novellia.native_token.modified AS modified
    FROM novellia.native_token
    INNER JOIN novellia.product ON novellia.product.native_token_id = novellia.native_token.native_token_id
  UNION
  SELECT product_id, modified FROM novellia.commission
  UNION
  SELECT product_id, modified FROM novellia.product_attribution
  UNION
  SELECT product_id, modified FROM novellia.remote_resource
)

SELECT
  product_id,
  MAX(modified)
FROM modified_timestamps
WHERE product_id=ANY($1)
GROUP BY product_id;