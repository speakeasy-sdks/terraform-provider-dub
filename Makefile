.PHONY: run
run:
	speakeasy run

overlay.yaml: base.yaml
	speakeasy overlay compare -s base.yaml -s changed.yaml > overlay.yaml

changed.yaml: base.yaml
	speakeasy overlay apply --overlay overlay.yaml -s base.yaml --out changed.yaml

.PHONY: base.yaml
base.yaml:
	curl https://api.dub.co | yq '.'  | yq e '.' - > base.yaml

