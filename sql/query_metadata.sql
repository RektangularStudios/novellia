SELECT DISTINCT ON ("policy",  "name")
  "policy",
  "name",
  "json"
FROM tx_metadata
INNER JOIN ma_tx_mint ON ma_tx_mint.tx_id = tx_metadata.tx_id
WHERE
  key='721' AND
  "name"=ANY($2::bytea[]) AND
  "policy"=ANY($1::bytea[]);
