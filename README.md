# Collector Components Template

This repo is meant as a guide for you can build your own components for the collector.

## Getting started

1. Clone the repo
1. Open the devcontainer

Now you should be able to run the solution.


```shell
make build
make run
```

Now you should have a running collector with the `mytest` components installed. You can use the test tool to send a span.

```shell
make test-span
```

The default config uses the the debug exporter, so you should see your span in the console.

## Making your own

You can make a new repo with this as a template, the perform the following steps

1. Rename the directories from `myest*` to the name of your component. Keep in mind the `component` part of the directory is important.
1. Update the `metadata.yaml` in the component with the right name and codeowners
1. Update the module path to you new github repository
1. Update the `ocb-manifest` with the new directory and github repo.
1. Update the `collector-config` with the new component name(s)

## Releasing your own component

The OpenTelemetry Collector Builder (OCB) requires tags in a particular format.

`{componentname}{componenttype}/{version}`

Once you've pushed that, you'll be able to reference it like any other component.

## TODO

Adding the different component types.
Making the defaults more meaningful (i.e. actually do something)