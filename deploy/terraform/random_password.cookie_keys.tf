# The key pairs (hash_key_1, block_key_1) and (hash_key_2, block_key_2)
# are used to encrypt the session cookie.
#
# They need to be rotated once in a while. But only one of them should
# be rotated at a time. This is a mechanism to allow for a smooth
# transition between the old and the new keys.
#
# Terraform is not able to generate random bytes out of the box. Instead,
# we generate a random password using only hexadecimal characters.
# The application expects a 32 byte key, so we generate a 64 character hexadecimal string.
# The application will convert this 64-character hexadecimal string to a strong 32-byte key.
#
# 2020-01
# odd_keeper = 101001
# even_keeper = 101000
# use_even = false
# current = odd
# old = even

# 2020-02
# odd_keeper = 101001
# even_keeper = 101001 # triggers re-creation of the even key since it is different from the previous value
# use_even = true
# current = even
# old = odd

# 2020-03
# odd_keeper = 101002 # triggers re-creation of the odd key since it is different from the previous value
# even_keeper = 101001
# use_even = false
# current = odd
# old = even
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
  length  = 64
  upper = false
  lower = false
  numeric = false
  special = true
  override_special = "1234567890abcdef"
  keepers = {
    date = local.odd_keeper
  }
}

resource "random_password" "even_block_key" {
  length  = 64
  upper = false
  lower = false
  numeric = false
  special = true
  override_special = "1234567890abcdef"
  keepers = {
    date = local.odd_keeper
  }
}

resource "random_password" "odd_hash_key" {
  length  = 64
  upper = false
  lower = false
  numeric = false
  special = true
  override_special = "1234567890abcdef"
  keepers = {
    date = local.even_keeper
  }
}

resource "random_password" "odd_block_key" {
  length  = 64
  upper = false
  lower = false
  numeric = false
  special = true
  override_special = "1234567890abcdef"
  keepers = {
    date = local.even_keeper
  }
}

