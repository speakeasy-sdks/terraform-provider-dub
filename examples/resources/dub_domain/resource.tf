resource "dub_domain" "my_domain" {
  archived    = false
  expired_url = "https://acme.com/expired"
  placeholder = "https://dub.co/help/article/what-is-dub"
  slug        = "acme.com"
}