SELECT
  ma_tx_outer.policy,
  ma_tx_outer.name,
  COALESCE(SUM(ma_tx_outer.quantity), 0)
FROM tx_out AS tx_outer
INNER JOIN ma_tx_out AS ma_tx_outer
  ON tx_outer.id = ma_tx_outer.tx_out_id
WHERE NOT EXISTS (
  SELECT 1
  FROM tx_out
  INNER JOIN
    tx_in ON tx_out.tx_id = tx_in.tx_out_id AND
    tx_out.index = tx_in.tx_out_index
  WHERE
    tx_outer.id = tx_out.id
) AND
  tx_outer.address = ANY($1)
  GROUP BY ma_tx_outer.policy, ma_tx_outer.name;