# Hugbot
## 🤗 Description
Hugbot is a tool built with Slack in mind to **grow appreciation** inside a company.
It allows employees to give each other appreciation through reacting to their messages or explicitly giving them
hugs.

## 🚀 Installation
### Kubernetes
In case you have access to a Kubernetes Cluster, the repository comes with a fully stable standalone
[Helm Chart](https://helm.sh/). To install:

```sh
helm install hugbot ./chart
```

### Manual Installation
Hugbot requires a Postgres Database to persist users and hugs. Refer to [.env.example](./.env.example) for more
information on configuring it.

An up-to-date docker image can be found in [dockerhub](https://hub.docker.com/r/nicoulbricht/hugbot).

### Slack
Once Hugbot is installed, you'll need to wire up a Slack Bot.
Create a new App for your workspace [here](https://api.slack.com/apps?new_app=1). Then get a Bot User Access Token
for the bot and add it as `SLACK_TOKEN` to the environment variables.

At last **Event Subscriptions** have to be enabled and point to `https://{HUGBOT_HOST}/slack/events`, listening to
`message.channels` and `reaction_added` events.

## 🤝 Contribution
Contributions from all skill-levels welcome. Just shoot me a message or PR. :)
