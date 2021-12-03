module github.com/hashicorp/terraform-provider-corner

go 1.12

require (
	github.com/hashicorp/go-memdb v1.3.2
	github.com/hashicorp/terraform-plugin-framework v0.4.2
	github.com/hashicorp/terraform-plugin-go v0.4.1-0.20211111190408-1ef303e79b8c
	github.com/hashicorp/terraform-plugin-log v0.1.1-0.20210921194200-489e81e39386
	github.com/hashicorp/terraform-plugin-mux v0.3.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.8.1-0.20211111190908-59c9fb6551c5
)

replace github.com/hashicorp/terraform-plugin-log => ../terraform-plugin-log

replace github.com/hashicorp/terraform-plugin-sdk/v2 => ../terraform-plugin-sdk

replace github.com/hashicorp/terraform-exec => ../terraform-exec

replace github.com/hashicorp/terraform-plugin-go => ../terraform-plugin-go
