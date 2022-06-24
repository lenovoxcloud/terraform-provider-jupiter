resource "Jupiter_VM" "testvm2" {
  vm_lookup_key = "b1e168eb-e34a-46b7-b823-32eadb907201"
  items{
    vm{
      instance_name = "tftest62"
      project_global_id = "73e44e46daa5438ead682c7ebd0f9f77"
      cloud_name = "DEV-O1"
      flavor_id = "3e01e507-c0fb-4198-a207-9863843d5b17"
      image_id = "940a8133-903b-4a06-b533-e8beec3e7d2c"
      vpc_id = "d16dba99-86fd-445f-a459-4fe9d21b71ab"
      network_id = "f39503e9-42a4-478e-bd48-e3793e9637d2"
      password_type = "input"
      password = "iE2)iS1&yC"
      power_state = "shutoff"
      
    }
    quantity = 1
  }
}