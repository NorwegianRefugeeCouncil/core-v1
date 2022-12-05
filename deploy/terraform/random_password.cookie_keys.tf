# The key pairs (hash_key_1, block_key_1) and (hash_key_2, block_key_2)
# are used to encrypt the session cookie.
#
# They need to be rotated once in a while. But only one of them should
# be rotated at a time. This is a mechanism to allow for a smooth
# transition between the old and the new keys.
#
# If we did not have a smooth transition of keys, every time the keys would
# be changed, all logged in users would lose their session. By having a "new"
# and an "old" key, the sessions would be gracefully upgraded to the new key.
#
# The lifetime of a key is 2 months. It will be used for 1 month, and then
# it will be rotated to the "old" key. After 1 month, the "old" key will be
# removed.
#
# Terraform provides a useful resource called "time_rotating". This resource has
# a base timestamp, and a rotation interval. Every time the interval is reached,
# the "time_rotating" resource is recreated with a new ID.
#
# We associate the "ID" of this "time_rotating" resource to the keepers of the
# "random_password" representing the key. This way, when the "time_rotating"
# resource is recreated, the "random_password" will be recreated as well.
#
# We have "even" and "odd" time_rotating resources. They are recreated respectively
# on even and odd months. The "even" and "odd" random_password are also respectively
# associated to the "even" and "odd" time_rotating resources. This way,
# the "even" random_password is recreated on even months, and the "odd" random_password
# is recreated on odd months.
#
# To compute which of the "even" and "odd" keys is the "new" and "old" one,
# we simply check if the current month of the year is even or odd. If it is even,
# the "even" key is the "new" one, and the "odd" key is the "old" one. If it is odd,
# the "odd" key is the "new" one, and the "even" key is the "old" one.
#
# So
# - On even months
#     - the even key is regenerated
#     - the "new" key is the even key
#     - the "old" key is the odd key
# - On odd months
#     - the odd key is regenerated
#     - the "new" key is the odd key
#     - the "old" key is the even key

resource "time_rotating" "even" {
  rotation_months = 2
  // Base timestamp is on an even month
  // Will be rotated every even month
  rfc3339 = "2022-02-01T00:00:00Z"
}

resource "time_rotating" "odd" {
  rotation_months = 2
  // Base timestamp is on an odd month
  // Will be rotated every odd month
  rfc3339 = "2022-03-01T00:00:00Z"
}

locals {
  use_even = tonumber(formatdate("MM", timestamp())) % 2 == 0
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
  keepers = {
    date = time_rotating.even.id
  }
}

resource "random_password" "even_block_key" {
  length           = 64
  upper            = false
  lower            = false
  numeric          = false
  special          = true
  override_special = "1234567890abcdef"
  keepers = {
    date = time_rotating.even.id
  }
}

resource "random_password" "odd_hash_key" {
  length           = 64
  upper            = false
  lower            = false
  numeric          = false
  special          = true
  override_special = "1234567890abcdef"
  keepers = {
    date = time_rotating.odd.id
  }
}

resource "random_password" "odd_block_key" {
  length           = 64
  upper            = false
  lower            = false
  numeric          = false
  special          = true
  override_special = "1234567890abcdef"
  keepers = {
    date = time_rotating.odd.id
  }
}

