# agent
**agent** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project.


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/agent@latest! | sudo bash
```
`username/agent` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)


# Stuff


Basic setup
```bash
ignite scaffold chain agent
cd agent
ignite chain serve

ignite scaffold vue
ignite generate composables
```

For frontend
```bash
cd vue
npm install
npm run dev
```

Generate type
```bash
ignite scaffold type prompt text:string id:uint
```

Generate write endpoint
```bash
ignite scaffold message add-prompt text --response id:uint
```

Generate read endpoint
```bash
ignite scaffold query get-prompt id:uint --response text:string
```

Demo
```bash
agentd tx agent ask-question "whats up?" --from alice --chain-id agent
agentd q agent show-question 0
agentd tx agent answer-question "the sky" 0 --from bob --chain-id agent
agentd q agent show-question 0
```

