SELECT
  COALESCE(SUM("value"), 0)
FROM tx_out AS tx_outer
WHERE NOT EXISTS (
  SELECT 1
  FROM tx_out
  INNER JOIN
    tx_in ON tx_out.tx_id = tx_in.tx_out_id AND
    tx_out.index = tx_in.tx_out_index
  WHERE
    tx_outer.id = tx_out.id
) AND
  tx_outer.address = ANY($1);