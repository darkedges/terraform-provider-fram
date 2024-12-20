---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fram_baseurlsource Data Source - terraform-provider-internal"
subcategory: ""
description: |-
  Returns details about the Base URL Source Service https://backstage.forgerock.com/docs/am/6.5/oidc1-guide/index.html#configure-base-url-source
---

# fram_baseurlsource (Data Source)

Returns details about the [Base URL Source Service](https://backstage.forgerock.com/docs/am/6.5/oidc1-guide/index.html#configure-base-url-source)

## Example Usage

```terraform
data "fram_baseurlsource" "all" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **context_path** (String) Specifies the context path for the base URL.
- **extension_class_name** (String) If Extension class is selected as the Base URL source, the Extension class name field.
- **fixed_value** (String) If Fixed value is selected as the Base URL source, the base URL in the Fixed value base URL field.
- **source** (String) Specifies the source of the base URL. Choose from the following:

	- Extension class. `EXTENSION_CLASS`

		Specifies that the extension class returns a base URL from a provided `HttpServletRequest`. In the Extension class name field, enter org.forgerock.openam.services.baseurl.BaseURLProvider.
	- Fixed value. `FIXED_VALUE`

		Specifies that the base URL is retrieved from a specific base URL value. In the Fixed value base URL field, enter the base URL value.
	- Forwarded header. `FORWARDED_HEADER`

		Specifies that the base URL is retrieved from a forwarded header field in the HTTP request. The Forwarded HTTP header field is standardized and specified in [RFC7239](https://tools.ietf.org/html/rfc7239).
	- Host/protocol from incoming request. `REQUEST_VALUES`

		Specifies that the hostname, server name, and port are retrieved from the incoming HTTP request.
	- X-Forwarded-* headers. `X_FORWARDED_HEADERS`

		Specifies that the base URL is retrieved from non-standard header fields, such as `X-Forwarded-For`, `X-Forwarded-By`, and `X-Forwarded-Proto`.


