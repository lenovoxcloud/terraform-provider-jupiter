resource "Jupiter_Volume" "jupiter_Jupiter_Volume" {
  volume_lookup_key = ""
  items{
    volume{
      name = "tfvolume2"
      project_global_id = "73e44e46daa5438ead682c7ebd0f9f77"
      cloud_name = "DEV-O1"
      volume_feature = "blank"
      is_thin_provisioning = "true"
      size = "1"
      user = "yuhan2@lenovo.com"
      volume_type = "56015969-8916-4e6b-8049-3a8955509910"
    }
    quantity = 1
  }
}
