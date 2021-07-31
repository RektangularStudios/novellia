SELECT
  DISTINCT ma_tx_mint.policy,
  ma_tx_mint.name
FROM ma_tx_mint TABLESAMPLE BERNOULLI(1)
GROUP BY ma_tx_mint.policy, ma_tx_mint.name
LIMIT 100;
