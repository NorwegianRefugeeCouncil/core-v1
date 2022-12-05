# The key pairs (hash_key_1, block_key_1) and (hash_key_2, block_key_2)
# are used to encrypt the session cookie.
#
# They need to be rotated once in a while. But only one of them should
# be rotated at a time. This is a mechanism to allow for a smooth
# transition between the old and the new keys.
#
# When declaring the random_password resource, we can specify a list of "keepers".
# Whenever the value of one of the keepers changes, the resource will be
# re-created. To periodically rotate the keys, we use a combination of the
# current year and month to calculate the "odd" and "even" keepers.
#
# With this mechanism, the odd and even keys
# will be alternately rotated every 2 months. A single key will begin as the
# "new" key, then after a month it will become the "old" key, and then after
# another month it will be deleted.
#
# This assumes that terraform is run at least once every month. We could add a
# periodic trigger to the pipeline to ensure it is run e.g. every 2 weeks.
#
# Terraform is not able to generate random bytes out of the box. Instead,
# we generate a random password using only hexadecimal characters.
# The application expects a 32 byte key, so we generate a 64 character hexadecimal string.
# The application will convert this 64-character hexadecimal string to a strong 32-byte key.
#
# This table shows which of the "odd" or "even" keys are used as "current" or "old" keys
# throughout the months.
#
# +------+-------+------------+------------+-------------+----------+---------+------+
# | Year | Month | Num        | Odd Keeper | Even Keeper | Use Even | Current | Old  |
# |      |       | Year+Month | (Num+1)/2  | (Num)/2     | Num%2==0 |         |      |
# +------+-------+------------+------------+-------------+----------+---------+------+
# | 2020 | 10    | 2030       | 1015       | 1015        | True     | even    | odd  |
# +------+-------+------------+------------+-------------+----------+---------+------+
# | 2020 | 11    | 2031       | 1016*      | 1015        | False    | odd     | even |
# +------+-------+------------+------------+-------------+----------+---------+------+
# | 2020 | 12    | 2032       | 1016       | 1016*       | True     | even    | odd  |
# +------+-------+------------+------------+-------------+----------+---------+------+
# | 2021 | 1     | 2033       | 1017*      | 1016        | False    | odd     | even |
# +------+-------+------------+------------+-------------+----------+---------+------+
# | 2021 | 2     | 2034       | 1017       | 1017*       | True     | even    | odd  |
# +------+-------+------------+------------+-------------+----------+---------+------+
# | 2021 | 3     | 2035       | 1018*      | 1017        | False    | odd     | even |
# +------+-------+------------+------------+-------------+----------+---------+------+
#
# An asterix * means that the key is regenerated.
#
locals {
  timestamp   = timestamp()
  year        = tonumber(formatdate("YYYY", local.timestamp))
  month       = tonumber(formatdate("MM", local.timestamp))
  date_num    = local.year + local.month
  odd_keeper  = floor((local.date_num + 1) / 2)
  even_keeper = floor(local.date_num / 2)
  use_even    = local.date_num % 2 == 0
}


locals {
  current_hash_key  = local.use_even ? random_password.even_hash_key.result : random_password.odd_hash_key.result
  current_block_key = local.use_even ? random_password.even_block_key.result : random_password.odd_block_key.result
  old_hash_key      = local.use_even ? random_password.odd_hash_key.result : random_password.even_hash_key.result
  old_block_key     = local.use_even ? random_password.odd_block_key.result : random_password.even_block_key.result
}


resource "random_password" "even_hash_key" {
  length           = 64
  upper            = false
  lower            = false
  numeric          = false
  special          = true
  override_special = "1234567890abcdef"
  keepers          = {
    date = local.odd_keeper
  }
}

resource "random_password" "even_block_key" {
  length           = 64
  upper            = false
  lower            = false
  numeric          = false
  special          = true
  override_special = "1234567890abcdef"
  keepers          = {
    date = local.odd_keeper
  }
}

resource "random_password" "odd_hash_key" {
  length           = 64
  upper            = false
  lower            = false
  numeric          = false
  special          = true
  override_special = "1234567890abcdef"
  keepers          = {
    date = local.even_keeper
  }
}

resource "random_password" "odd_block_key" {
  length           = 64
  upper            = false
  lower            = false
  numeric          = false
  special          = true
  override_special = "1234567890abcdef"
  keepers          = {
    date = local.even_keeper
  }
}

