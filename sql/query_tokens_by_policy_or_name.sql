SELECT
  ma_tx_mint.policy,
  ma_tx_mint.name
FROM ma_tx_mint
WHERE
  encode(ma_tx_mint.policy, 'hex') = ANY($1) OR
  encode(ma_tx_mint.name, 'escape') = ANY($1)
GROUP BY ma_tx_mint.policy, ma_tx_mint.name;
