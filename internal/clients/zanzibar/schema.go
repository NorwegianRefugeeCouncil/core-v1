package zanzibar

const schema = `
definition notcore/user {}

definition notcore/global { // org = nrc?
  relation admin: notcore/user
  permission write_all_individuals = admin
  permission view_all_individuals = admin + write_all_individuals
}

definition notcore/region {
  relation global: notcore/global
  relation ro_admin: notcore/user
  permission write_ro_individuals = ro_admin + global->write_all_individuals
  permission view_ro_individuals = ro_admin + global->view_all_individuals + write_ro_individuals
}

definition notcore/country {
  relation region: notcore/region
  relation co_admin: notcore/user
  permission write_co_individuals = co_admin + region->write_ro_individuals
  permission view_co_individuals = co_admin + region->view_ro_individuals + write_co_individuals
}

definition notcore/individual {
  relation ind_org: notcore/global
  relation ind_region: notcore/region
  relation ind_country: notcore/country

  relation writer: notcore/user
  relation reader: notcore/user

  permission view = reader + writer + ind_country->view_co_individuals
}
`
