resource "fram_baseurlsource" "example" {
  source = "FIXED_VALUE"
  fixed_value = "https://fram.example.com"
  context_path = "/openam"
}