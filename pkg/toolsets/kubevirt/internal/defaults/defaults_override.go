package defaults

const (
	toolsetNameOverride        = "openshift-virtualization"
	toolsetDescriptionOverride = "OpenShift Virtualization virtual machine management tools, check the [OpenShift Virtualization documentation](https://github.com/containers/kubernetes-mcp-server/blob/main/docs/kubevirt.md) for more details."
)

func ToolsetNameOverride() string {
	return toolsetNameOverride
}

func ToolsetDescriptionOverride() string {
	return toolsetDescriptionOverride
}
