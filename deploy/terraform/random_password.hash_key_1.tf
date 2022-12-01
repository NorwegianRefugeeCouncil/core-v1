# The key pairs (hash_key_1, block_key_1) and (hash_key_2, block_key_2)
# are used to encrypt the session cookie.
#
# They need to be rotated once in a while. But only one of them should
# be rotated at a time. This is a mechanism to allow for a smooth
# transition between the old and the new keys.
#
# 2020-01
# odd_keeper = 101001
# even_keeper = 101000
# use_even = false
# current_hash_key = odd
# current_block_key = odd
# old_hash_key = even
# old_block_key = even

# 2020-02
# odd_keeper = 101001
# even_keeper = 101001 # triggers re-creation of the even key since it is different from the previous value
# use_even = true
# current_hash_key = even
# current_block_key = even
# old_hash_key = odd
# old_block_key = odd

# 2020-03
# odd_keeper = 101002 # triggers re-creation of the odd key since it is different from the previous value
# even_keeper = 101001
# use_even = false
# current_hash_key = odd
# current_block_key = odd
# old_hash_key = even
# old_block_key = even
#
# and so on...

locals {
  now         = formatdate("YYYYMM", timestamp())
  date        = tonumber(local.now)
  odd_keeper  = floor((local.date + 1) / 2)
  even_keeper = floor(local.date / 2)
  use_even    = local.date % 2 == 0
}


locals {
  current_hash_key  = local.use_even ? random_password.even_hash_key.result : random_password.odd_hash_key.result
  current_block_key = local.use_even ? random_password.even_block_key.result : random_password.odd_block_key.result
  old_hash_key      = local.use_even ? random_password.odd_hash_key.result : random_password.even_hash_key.result
  old_block_key     = local.use_even ? random_password.odd_block_key.result : random_password.even_block_key.result
}


resource "random_password" "even_hash_key" {
  length  = 32
  keepers = {
    date = local.odd_keeper
  }
}

resource "random_password" "even_block_key" {
  length  = 32
  keepers = {
    date = local.odd_keeper
  }
}

resource "random_password" "odd_hash_key" {
  length  = 32
  keepers = {
    date = local.even_keeper
  }
}

resource "random_password" "odd_block_key" {
  length  = 32
  keepers = {
    date = local.even_keeper
  }
}

