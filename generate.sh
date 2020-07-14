rm -r ./Messages
mkdir Messages
mkdir Messages/Factory
mkdir Messages/Constants
python3 generate.py config.yaml ./dictionary ./Messages
#find ./Messages/ | grep '\.go$' | xargs gofmt -w