SELECT
  ma_tx_mint.policy,
  ma_tx_mint.name
FROM ma_tx_mint
WHERE
  encode(ma_tx_mint.policy, 'hex') = $1 AND
  encode(ma_tx_mint.name, 'escape') = $2
GROUP BY ma_tx_mint.policy, ma_tx_mint.name;
