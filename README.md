# Terraform Provider for Lenovo Xcloud Jupiter a

The Terraform Jupiter provider is a plugin for Terraform that allows for the management of Lenovo Xcloud Jupiter resources.
This provider is maintained internally by the Lenovo Xcloud Jupiter Provider team.

Please note: We take Terraform's security and our users' trust very seriously. If you believe you have found a security issue in the Terraform Jupiter provider, please responsibly disclose by contacting us at yuhan2@lenovo.com,Thanks

Terraform Jupiter provider是一个用于管理Lenovo Xcloud Jupiter资源的Terraform插件。

该Terraform provider由Lenovo Xcloud Jupiter provider团队进行内部维护。

请注意:我们非常重视Terraform的安全和用户的信任。如果您认为您在Terraform Jupiter provider中发现了安全问题，请通过联系我们(yuhan2@lenovo.com)负责任地披露,十分感谢。

## Quick Starts

here are some TF code to create an jupiter vm and a jupiter volume on your localhost jupiter:
下面是一些TF代码，可以在你本地运行的jupiter上创建一个jupiter vm和一个jupiter volume:
```terraform
terraform {
  required_providers {
    jupiter = {
      source = "lenovoxcloud/jupiter"
      version = "1.0.5"
    }
  }
}

provider "jupiter" {
    jupiter_url = "http://127.0.0.1:8000"  // 要访问的Jupiter的地址
    auth_token = "your auth token"         // 要访问的Jupiter的authToken
}

resource "Jupiter_VM" "testvm1" {
  vm_lookup_key = ""
  items{
    vm{
      instance_name = "tftest1"
      project_global_id = "73e44e46daa5438ead682c7ebd0f9f77"
      cloud_name = "DEV-O1"
      flavor_id = "3e01e507-c0fb-4198-a207-9863843d5b17"
      image_id = "940a8133-903b-4a06-b533-e8beec3e7d2c"
      vpc_id = "d16dba99-86fd-445f-a459-4fe9d21b71ab"
      network_id = "f39503e9-42a4-478e-bd48-e3793e9637d2"
      password_type = "input"
      password = "iE2)iS1&yC"
      power_state = "active"
      
    }
    quantity = 1
  }
}

resource "Jupiter_Volume" "testvolume1" {
  volume_lookup_key = ""
  items{
    volume{
      name = "tfvolume1"
      project_global_id = "73e44e46daa5438ead682c7ebd0f9f77"
      cloud_name = "DEV-O1"
      volume_feature = "blank"
      is_thin_provisioning = "true
      size = "1"
      user = "yuhan2@lenovo.com"
      volume_type = "56015969-8916-4e6b-8049-3a8955509910"
    }
    quantity = 1
  }
}

``` 

# 字段解释 fields

|字段名 Field|含义|explain|
|-----|----|----|
|jupiter_url|你想通过tf管理的Jupiter的地址|Jupiter's address that you want to administer via TF|
|auth_token|访问jupiter的auth_token|Access Jupiter's auth_token|

|字段名 Field|含义|explain|
|-----|----|----|
|instance_name|要创建的云主机的名称|Jupiter VM's name |
|project_global_id|云主机所属的project_global_id|Project_global_id To which the VM belongs|
|cloud_name|云主机所属的云环境名称|Cloud Name which the VM belongs|
|flavor_id|云主机使用的规格的id|VM's flavor uuid|
|image_id|云主机使用的镜像的id|VM's image uuid|
|vpc_id|云主机使用的vpcid|VM's VPC uuid|
|network_id|云主机使用的云网络的id|VM's network id|
|password_type|云主机的密码类型|VM's password type|
|password|云主机的密码值|VM's password|
|power_state|云主机的开机状态|VM's power status|

|字段名 Field|含义|explain|
|-----|----|----|
|name|要创建的云硬盘的名称|Jupiter volume's name |
|project_global_id|云硬盘所属的project_global_id|Project_global_id To which the volume belongs|
|cloud_name|云硬盘所属的云环境名称|Cloud Name which the volume belongs|
|volume_feature|云硬盘使来源|volume's feature|
|is_thin_provisioning|是否精简置备|is this a thin provisioning volume|
|size|云硬盘大小|volume size|
|user|创建的用户|who create the volume|
|volume_type|云硬盘类型|volume type|

--------

Have any other questions,submit a issue or mailing yuhan2@lenovo.com