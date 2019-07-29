package openfigi

//go:generate wget http://central.maven.org/maven2/org/openapitools/openapi-generator-cli/4.0.3/openapi-generator-cli-4.0.3.jar -O openapi-generator-cli.jar
//go:generate java -jar openapi-generator-cli.jar generate -i https://api.openfigi.com/schema -g go -o . --additional-properties withGoCodegenCommen=true,packageName=openfigi

