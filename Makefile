.PHONY: base.yaml
base.yaml:
	curl https://api.dub.co | yq '.'  | yq e '.' - > base.yaml

overlay.yaml: