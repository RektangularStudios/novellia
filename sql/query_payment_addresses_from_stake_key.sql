SELECT DISTINCT tx_out.address
FROM tx_out
INNER JOIN stake_address ON tx_out.stake_address_id = stake_address.id
WHERE
  stake_address.view = $1 OR
  hash_raw = decode($1, 'hex');
