# go-serverless-aaapi-example

Go Serverless Twitter Account Activity API example

- Go 1.x
- AWS Lambda
- Twitter Accout Activity API
- Serverless Framework

## Usage

### 1. Edit `env.yml`

```bash
$ cp env.yml.sample env.yml
$ vi env.yml
# edit env.yml
```

### 2. Configure AWS profile

```bash
$ aws configure --profile aaapi
AWS Access Key ID [None]: XXXXXXXXXXXXXXXX
AWS Secret Access Key [None]: XXXXXXXXXXXXXXXX
Default region name [None]: XXXXXXXXXX
Default output format [None]:
```

### 3. Deploy

```bash
$ make deploy 
```

### 4. Register URL as Webhook

[Getting started with webhooks — Twitter Developers](https://developer.twitter.com/en/docs/accounts-and-users/subscribe-account-activity/guides/getting-started-with-webhooks)
