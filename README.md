# My Minecraft Server controller

EC2 instance launch control

## provisioning: ansible

### setup

```
poetory install
```

### apply

```
poetory run ansible-playbook -i ansible/hosts ansible/site.yml
```

## DevOps: lambda

using [Serverless Framework](https://serverless.com/)

### setup deploy tool

https://serverless.com/framework/docs/getting-started/


```
$ npm install -g serverless
```

Or, update the serverless cli from a previous version

```
$ npm update -g serverless
```

### prepare env file

```
$ cp env.yml{.sample,}
```

and edit `env.yml`

### deploy

```
$ make deploy
```

if serverless deploy only is `sls deploy --vervose`

# TODO
- Consider national holiday in stop instance logic
