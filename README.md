# AWSCRED

![awscred](./docs/awscred.jpg)

[![Build Status](https://cloud.drone.io/api/badges/hanjunlee/awscred/status.svg)](https://cloud.drone.io/hanjunlee/awscred) [![Go Report Card](https://goreportcard.com/badge/github.com/hanjunlee/awscred)](https://goreportcard.com/report/github.com/hanjunlee/awscred)

---

## Index

- [Concept](#concept)
- [Install](#install)
- [Usage](#usage)
- [License](#license)
- [Changelog](#changelog)

## Concept 

The main concept of Awscred is to handle session token by creating a new AWS `credentials` file. **It helps you by abstracting the process which is to generate a new session token and to share it**.  
<br/>

Suppose we need a session token and we want to store it. The first step is to generate a session token with `aws` command, when you run the command it returns json-format response like below ([aws doc](https://aws.amazon.com/premiumsupport/knowledge-center/authenticate-mfa-cli/)). 

```bash
$ aws sts get-session-token --serial-number arn-of-the-mfa-device --token-code code-from-token 

{
    "Credentials": {
        "SecretAccessKey": "secret-access-key",
        "SessionToken": "temporary-session-token",
        "Expiration": "expiration-date-time",
        "AccessKeyId": "access-key-id"
    }
}
```

After generation, you have to set session token on your AWS `credentials` file if you need to sharing it, or you have to export values as environment variables.

```bash
# credentials file
[defuault-mfa]
aws_access_key_id = example-access-key-as-in-returned-output
aws_secret_access_key = example-secret-access-key-as-in-returned-output
aws_session_token = example-session-Token-as-in-returned-output
```

It is very complicated and also it is a toil because you have to do same process when session token is expired.  
<br/>

Awscred makes you can handle session token without these complicated steps. What is you have to prepare is setting the serial number of IAM user, and after this setting you don’t need to put the serial number anymore because it’s stored at the `config` file of Awscred.

```bash
$ awscred set --on --serial SERIAL 
```

After configuration, let’s generate session token. 

```bash
$ awscred gen --code CODE

$ $(awscred export)
```

Awscred will set session token on the `credentials` file of Awscred (not AWS) automatically.  
<br/>

You can get some benefits by using Awscred. **The best thing is it doesn’t intrude your AWS `credentials`**. In above example, you have to set session token with new profile(`default-mfa`) on AWS `credentials` to share it, but Awscred set session token with the same profile so you don’t need to change your profile 🙂.  And Awscred copies access keys of other profiles on the Awscred `credentials` file so that there’s no side effect to replace `credentials` file.

## Install

### Linux

Download and Install in Linux

```shell
$ curl -L https://github.com/hanjunlee/awscred/releases/latest/download/awscred_linux_amd64.tar.gz | tar zx
$ sudo install -t /usr/local/bin awscred
```

### OSX

Download and Install in OSX

```shell
$ curl -L https://github.com/hanjunlee/awscred/releases/latest/download/awscred_darwin_amd64.tar.gz | tar zx
$ sudo cp awscred /usr/local/bin
```

Or Install in OSX by using Brew

```shell
$ brew tap hanjunlee/awscred
$ brew install awscred
```

### Window

TBD

## Usage

1. Run a new daemon - It creates a new daemon which reflect a session token on new credentials.

```shell
$ awscred run
```

2. Configure the profile - It configures the serial number and the duration. These values are used as options to generate a session token.

```shell
# set up the configuration
$ awscred set --on --serial arn:aws:iam::XXXXXXXX:mfa/USER PROFILE

$ awscred info
NAME         ON       SERIAL                                    DURATION    EXPIRED
PROFILE      true     arn:aws:iam::XXXXXXXX:mfa/USER PROFILE    43200       0001-01-01T00:00:00Z
...
```

3. Generate a new session token 

```shell
# generate a new session token
$ awscred gen --code XXXXXX PROFILE

$ awscred info
NAME         ON       SERIAL                                    DURATION    EXPIRED
PROFILE      true     arn:aws:iam::XXXXXXXX:mfa/USER PROFILE    43200       2020-08-22T23:43:50Z (10.9h)
...
```

4. Set `AWS_SHARED_CREDENTIALS_FILE` environment variable - By set [the environment of shared credentials file](https://docs.aws.amazon.com/cli/latest/topic/config-vars.html#the-shared-credentials-file) Awscred `credentials` file, `aws` command use the new credentials file. 

```shell
$(awscred export)
```

5. Terminate 

```shell
# terminate the daemon
$ awscred terminate 

$ unset AWS_SHARED_CREDENTIALS_FILE
```

## LICENSE

[MIT License](./LICENSE)

## CHANGELOG

[Changelogs](./CHANGELOG.md)
