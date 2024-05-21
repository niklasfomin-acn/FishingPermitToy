# Get container for vault
docker pull vault:1.13.3

# Run the vault server with env variables
export VAULT_ADDR='http://127.0.0.1:8200'
docker run --cap-add=IPC_LOCK -e 'VAULT_DEV_ROOT_TOKEN_ID=root' -p 8200:8200 vault:1.13.3

# Set Environment Variables for vault adress and VAULT_TOKEN=root

# Store the clientID and the Client Secret in the vault
vault kv patch secret/FishingPermitClone/ServiceConfig ConnectionString="mongodb://localhost:27017"
vault kv patch secret/FishingPermitClone/ServiceConfig ConnectionString="mongodb://mongodbmanaged:QdzAN0rUTnCjRfDOfFYEYTz6j8MUWgRWv6iOLBFZqAj82K30ONoJi0uRpPpBq0Z1ff2IyY5Kh6WnACDbANQ0qQ%3D%3D@mongodbmanaged.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@mongodbmanaged@"


# Store DocumentAI secrets in vault
vault kv patch secret/FishingPermitClone/ServiceConfig ServiceEndpoints="https://kibundai.cognitiveservices.azure.com/documentintelligence/documentModels/prebuilt-idDocument:analyze?api-version=2024-02-29-preview"
vault kv patch secret/FishingPermitClone/ServiceConfig ServiceKeys="aea95a4fbfbb43a78646f7bdc7049341"

# Store Database connection strings and passwords in vault
vault kv patch secret/FishingPermitClone/ServiceConfig ConnectionStrings="mongodb://localhost:27017"

# Store the clientID, realm, endpoint and the Client Secret in the vault
vault kv patch secret/FishingPermitClone/IdentityConfig ClientID="admin-cli"
vault kv patch secret/FishingPermitClone/IdentityConfig Realm="testrealm"
vault kv patch secret/FishingPermitClone/IdentityConfig Endpoint="http://localhost:8080/auth"
vault kv patch secret/FishingPermitClone/IdentityConfig ClientSecret="f7b3b3b3-7b3b-4b3b-8b3b-3b3b3b3b3b3b"
