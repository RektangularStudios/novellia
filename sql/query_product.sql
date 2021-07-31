SELECT
  -- product
  product.product_id, product_name, 
  price_currency_id,
  COALESCE(price_unit_amount, 0),
  COALESCE(max_order_size, 0), 
  date_listed,
  date_available,
  -- organization
  product.organization_id, organization_name, organization_description,
  -- market
  product.market_id, market_name, market_description,
  -- native token
  COALESCE(native_token.policy_id, '') AS native_token_policy_id,
  COALESCE(native_token.asset_id, '') AS native_token_asset_id,
  -- product detail
  copyright,
  publisher,
  product_version,
  COALESCE(id, 0),
  tags,
  description_short,
  description_long,
  COALESCE(stock_available, 0),
  COALESCE(total_supply, 0)
FROM product
INNER JOIN organization ON product.organization_id = organization.organization_id
INNER JOIN market ON product.market_id = market.market_id
INNER JOIN product_detail ON product.product_id = product_detail.product_id
LEFT OUTER JOIN native_token ON product.native_token_id = native_token.native_token_id
WHERE
  product.deleted IS NULL AND
  (
    product.product_id=ANY($1) OR
    (product.native_token_id IS NOT NULL AND product.native_token_id=ANY($2))
  ) AND
  (date_listed IS NOT NULL AND date_listed <= NOW())
ORDER BY id ASC;
