### api-spec

Defines SCOIR's internal API

To help organization and maintainability of the spec, it is split across multiple
files.

To create a single machine readable json version, install multi-file-swagger,
and mustache.

    npm install -g multi-file-swagger
    sudo apt-get install ruby-mustache

Run:

    spec-me.sh
